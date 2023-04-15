package service

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"music-metrics/dal"
	"music-metrics/model"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

func UpdateCode(code string) model.UpdateCodeResponse {

	tx, db, err := dal.BeginTX()
	if err != nil {
		return model.UpdateCodeResponse{Success: false, Message: serverErrorStr}
	}

	var response model.UpdateCodeResponse

	// Request Access/Refresh Token from Spotify
	PrintMessage("Getting access token from Spotify...")
	accessToken, refreshToken, err := requestAccessToken(code)
	if err != nil {
		if dal.CommitAndClose(tx, db, false) != nil {
			return model.UpdateCodeResponse{Success: false, Message: serverErrorStr}
		}
		return model.UpdateCodeResponse{Success: false, Message: serverErrorStr}
	}
	PrintMessage("Successfully got access token from Spotify, now getting user info...")

	// Get UserBean Info from Spotify
	currUser, err := requestUserInfo(accessToken)
	if err != nil {
		if dal.CommitAndClose(tx, db, false) != nil {
			return model.UpdateCodeResponse{Success: false, Message: serverErrorStr}
		}
		return model.UpdateCodeResponse{Success: false, Message: err.Error()}
	}
	PrintMessage("Successfully got user info from Spotify")
	currUser.Refresh = refreshToken

	// Determine if user already exists
	PrintMessage("Checking if user already exists...")
	existingUser, err := dal.RetrieveUser(tx, currUser.Username)
	if err != nil {
		if dal.CommitAndClose(tx, db, false) != nil {
			return model.UpdateCodeResponse{Success: false, Message: serverErrorStr}
		}
		return model.UpdateCodeResponse{Success: false, Message: err.Error()}
	}

	var token model.AuthTokenBean

	// If user does not exist, create user and auth token
	if (existingUser == model.UserBean{}) {
		PrintMessage("UserBean does not exist, creating user and auth token...")
		currUser.Timestamp = time.Now().UnixMilli()
		err = dal.CreateUser(tx, currUser)
		if err != nil {
			if dal.CommitAndClose(tx, db, false) != nil {
				return model.UpdateCodeResponse{Success: false, Message: serverErrorStr}
			}
			return model.UpdateCodeResponse{Success: false, Message: serverErrorStr}
		}
		token = model.AuthTokenBean{
			Username: currUser.Username,
			Token:    generateID(DEFAULT_ID_LENGTH),
		}
		err = dal.CreateAuthToken(tx, token)
		if err != nil {
			if dal.CommitAndClose(tx, db, false) != nil {
				return model.UpdateCodeResponse{Success: false, Message: serverErrorStr}
			}
			return model.UpdateCodeResponse{Success: false, Message: serverErrorStr}
		}
		PrintMessage("Successfully created user and auth token")
		// Add recent listens to DB for instant access
		PrintMessage("Adding recent listens to DB...")
		recentListens, err := getRecentlyPlayed(accessToken)
		if err != nil {
			if dal.CommitAndClose(tx, db, false) != nil {
				return model.UpdateCodeResponse{Success: false, Message: serverErrorStr}
			}
			return model.UpdateCodeResponse{Success: false, Message: err.Error()}
		}
		// Determine which listens are new and add them if they are
		// todo: remove this duplication with tracking script
		for _, rpObj := range recentListens {
			// Add song to database if it isn't already there
			song, err := dal.RetrieveSong(tx, rpObj.Song.Id)
			if err != nil {
				fmt.Println("Error retrieving song for username: " + currUser.Username)
				fmt.Println(err)
				continue
			}
			if (song == model.SongBean{}) {
				err = dal.CreateSong(tx, &rpObj.Song)
				if err != nil {
					fmt.Println("Error creating song for username: " + currUser.Username)
					continue
				}
			} else {
				// Update song if it is already there
				err = dal.UpdateSong(tx, &rpObj.Song)
				if err != nil {
					fmt.Println("Error updating song for username: " + currUser.Username)
					continue
				}
			}

			// Add album to database if it isn't already there
			album, err := dal.RetrieveAlbum(tx, rpObj.Album.Id)
			if err != nil {
				fmt.Println("Error retrieving album for username: " + currUser.Username)
				fmt.Println(err)
				continue
			}
			if (album == model.AlbumBean{}) {
				err = dal.CreateAlbum(tx, &rpObj.Album)
				if err != nil {
					fmt.Println("Error creating album for username: " + currUser.Username)
					continue
				}
			} else {
				// Update album if it is already there
				err = dal.UpdateAlbum(tx, &rpObj.Album)
				if err != nil {
					fmt.Println("Error updating album for username: " + currUser.Username)
					continue
				}
			}

			// Add listen to database
			newListen := model.ListenBean{
				Username:  currUser.Username,
				Timestamp: rpObj.Timestamp,
				SongId:    rpObj.Song.Id,
			}
			err = dal.CreateListen(tx, newListen)
			if err != nil {
				fmt.Println("Error creating listen for username: " + currUser.Username)
				fmt.Println(err)
				continue
			}
		}
		PrintMessage("Successfully added listens to DB")
	} else { // UserBean already exists, update them and get auth token
		PrintMessage("UserBean already exists, updating user and getting auth token...")
		currUser.Timestamp = existingUser.Timestamp
		err = dal.UpdateUser(tx, currUser)
		if err != nil {
			if dal.CommitAndClose(tx, db, false) != nil {
				return model.UpdateCodeResponse{Success: false, Message: serverErrorStr}
			}
			return model.UpdateCodeResponse{Success: false, Message: err.Error()}
		}
		token, err = dal.RetrieveAuthTokenByUsername(tx, currUser.Username)
		if err != nil {
			if dal.CommitAndClose(tx, db, false) != nil {
				return model.UpdateCodeResponse{Success: false, Message: serverErrorStr}
			}
			return model.UpdateCodeResponse{Success: false, Message: err.Error()}
		}
		if (token == model.AuthTokenBean{}) {
			PrintMessage("Token is null, returning (this should not happen)")
			if dal.CommitAndClose(tx, db, false) != nil {
				return model.UpdateCodeResponse{Success: false, Message: serverErrorStr}
			}
			return model.UpdateCodeResponse{Success: false, Message: serverErrorStr}
		}
		PrintMessage("Successfully updated user and got auth token")
	}

	response = model.UpdateCodeResponse{
		Success:     true,
		Token:       token.Token,
		Username:    currUser.Username,
		DisplayName: currUser.DisplayName,
		Email:       currUser.Email,
		Timestamp:   currUser.Timestamp,
	}

	PrintMessage("Committing to DB...")
	if dal.CommitAndClose(tx, db, true) != nil {
		return model.UpdateCodeResponse{Success: false, Message: serverErrorStr}
	}

	return response
}

func requestAccessToken(code string) (string, string, error) {

	uri := SPOTIFY_BASE_ACCOUNT + "/api/token"
	secret := os.Getenv("SPOTIFY_CLIENT_SECRET")
	if secret == "" {
		return "", "", fmt.Errorf("secret is empty")
	}
	encodedSecret := base64.StdEncoding.EncodeToString([]byte(SPOTIFY_CLIENT_ID + ":" + secret))

	reqBody := url.Values{}
	reqBody.Set("grant_type", "authorization_code")
	reqBody.Set("code", code)
	reqBody.Set("redirect_uri", SPOTIFY_REDIRECT_URL)

	encodedRequestBody := reqBody.Encode()

	req, err := http.NewRequest("POST", uri, strings.NewReader(encodedRequestBody))
	if err != nil {
		return "", "", err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Basic "+encodedSecret)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", "", err
	}

	var tokenResp model.GetAccessTokenResponse
	err = json.NewDecoder(resp.Body).Decode(&tokenResp)
	if err != nil {
		panic(err)
	}

	err = resp.Body.Close()
	if err != nil {
		return "", "", err
	}
	return tokenResp.AccessToken, tokenResp.Refresh, nil
}

func requestUserInfo(access string) (model.UserBean, error) {

	uri := SPOTIFY_BASE_API + "/me"

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return model.UserBean{}, err
	}

	req.Header.Set("Authorization", "Bearer "+access)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return model.UserBean{}, err
	}

	var getMeResponse model.GetMeResponse
	err = json.NewDecoder(resp.Body).Decode(&getMeResponse)
	if err != nil {
		return model.UserBean{}, err
	}

	err = resp.Body.Close()
	if err != nil {
		return model.UserBean{}, err
	}
	return model.UserBean{Username: getMeResponse.ID, DisplayName: getMeResponse.DisplayName, Email: getMeResponse.Email}, nil
}

func getRecentlyPlayed(token string) ([]model.RecentlyPlayedObject, error) {

	uri := SPOTIFY_BASE_API + "/me/player/recently-played"

	params := url.Values{}
	params.Add("before", strconv.FormatInt(time.Now().UnixMilli(), 10))
	params.Add("limit", "50")
	urlWithParams := fmt.Sprintf("%s?%s", uri, params.Encode())

	req, err := http.NewRequest("GET", urlWithParams, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var respBody model.GetRecentlyPlayedResponse
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		return nil, err
	}

	err = resp.Body.Close()
	if err != nil {
		return nil, err
	}

	var toReturn []model.RecentlyPlayedObject
	for _, item := range respBody.Items {
		song := model.SongBean{
			Id:         item.Track.ID,
			Name:       item.Track.Name,
			Artist:     artistsToString(item.Track.Artists),
			Album:      item.Track.Album.ID,
			Explicit:   item.Track.Explicit,
			Popularity: item.Track.Popularity,
			Duration:   item.Track.DurationMs,
		}
		album := model.AlbumBean{
			Id:          item.Track.Album.ID,
			Name:        item.Track.Album.Name,
			Artist:      artistsToString(item.Track.Album.Artists),
			Genre:       strings.Join(item.Track.Album.Genres, SEPARATOR),
			TotalTracks: item.Track.Album.TotalTracks,
			Year:        yearFromReleaseDate(item.Track.Album.ReleaseDate),
			Image:       item.Track.Album.Images[0].URL,
			Popularity:  item.Track.Album.Popularity,
		}
		returnObj := model.RecentlyPlayedObject{
			Song:      song,
			Album:     album,
			Timestamp: datetimeToUnixMilli(item.PlayedAt),
		}
		toReturn = append(toReturn, returnObj)
	}

	return toReturn, nil
}

func artistsToString(artists []model.Artist) string {
	var arr []string
	for _, artist := range artists {
		arr = append(arr, artist.Name)
	}
	return strings.Join(arr, SEPARATOR)
}

func yearFromReleaseDate(date string) int {
	i, err := strconv.Atoi(date[:4])
	if err != nil {
		return -1
	}
	return i
}

func datetimeToUnixMilli(datetime string) int64 {
	t, err := time.Parse(time.RFC3339, datetime)
	if err != nil {
		return -1
	}
	return t.UnixMilli()
}
