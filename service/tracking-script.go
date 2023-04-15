package service

import (
	"database/sql"
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

func StartTracking() {

	fmt.Println("STARTING TRACKING SCRIPT...")

	for {

		startTime := time.Now()

		// Instantiate connection
		tx, db, err := dal.BeginTX()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		users, err := dal.RetrieveAllUsers(tx)
		if err != nil {
			if dal.CommitAndClose(tx, db, false) != nil {
				fmt.Println("Error committing transaction")
				return
			}
		}

		// For every user in the database
		for _, user := range users {
			startTimeUser := time.Now()
			// Get new access token
			newToken, err := refreshToken(user.Refresh)
			if err != nil || newToken == "" {
				fmt.Println("Error refreshing token for username: " + user.Username)
				continue
			}

			// Use new access token to call /recently-played
			recentlyPlayed, err := getRecentlyPlayed(newToken)
			if err != nil {
				fmt.Println("Error getting recently played for username: " + user.Username)
				continue
			}

			// Get most recent listen
			mostRecentListen, err := dal.GetMostRecentListen(tx, user.Username)
			var oldTime int64
			if err != nil {
				fmt.Println("Error getting most recent listen for username: " + user.Username)
				continue
			}
			if (mostRecentListen == model.ListenBean{}) {
				oldTime = 0
			} else {
				oldTime = mostRecentListen.Timestamp
			}

			// Determine which listens are new and add them if they are
			newSongsCount := loopThroughRecentListens(recentlyPlayed, tx, user, oldTime)

			fmt.Println(user.Username + " listened to " + strconv.Itoa(newSongsCount) + " songs in the last 2 hours. (" + time.Since(startTimeUser).String() + "), (" + time.Now().Format("2006-01-02 15:04:05 -0700 MST") + ")")

			// Sleep for a little bit to avoid rate limiting
			time.Sleep(500 * time.Millisecond)

		}

		if dal.CommitAndClose(tx, db, true) != nil {
			fmt.Println("Error committing transaction")
			return
		}

		// Sleep for 2 hours - time it took to run the script
		time.Sleep(2*time.Hour - time.Since(startTime))
	}

}

func loopThroughRecentListens(listens []model.RecentlyPlayedObject, tx *sql.Tx, user model.UserBean, oldTime int64) int {
	newSongsCount := 0
	for _, rpObj := range listens {
		if rpObj.Timestamp > oldTime {
			// Add song to database if it isn't already there
			newSongsCount++
			song, err := dal.RetrieveSong(tx, rpObj.Song.Id)
			if err != nil {
				fmt.Println("Error retrieving song for username: " + user.Username)
				fmt.Println(err)
				continue
			}
			if (song == model.SongBean{}) {
				err = dal.CreateSong(tx, &rpObj.Song)
				if err != nil {
					fmt.Println("Error creating song for username: " + user.Username)
					continue
				}
			} else {
				// Update song if it is already there
				err = dal.UpdateSong(tx, &rpObj.Song)
				if err != nil {
					fmt.Println("Error updating song for username: " + user.Username)
					continue
				}
			}

			// Add album to database if it isn't already there
			album, err := dal.RetrieveAlbum(tx, rpObj.Album.Id)
			if err != nil {
				fmt.Println("Error retrieving album for username: " + user.Username)
				fmt.Println(err)
				continue
			}
			if (album == model.AlbumBean{}) {
				err = dal.CreateAlbum(tx, &rpObj.Album)
				if err != nil {
					fmt.Println("Error creating album for username: " + user.Username)
					continue
				}
			} else {
				// Update album if it is already there
				err = dal.UpdateAlbum(tx, &rpObj.Album)
				if err != nil {
					fmt.Println("Error updating album for username: " + user.Username)
					continue
				}
			}

			// Add listen to database
			newListen := model.ListenBean{
				Username:  user.Username,
				Timestamp: rpObj.Timestamp,
				SongId:    rpObj.Song.Id,
			}
			err = dal.CreateListen(tx, newListen)
			if err != nil {
				fmt.Println("Error creating listen for username: " + user.Username)
				fmt.Println(err)
				continue
			}
		} else {
			break
		}
	}
	return newSongsCount
}

func refreshToken(refresh string) (string, error) {

	uri := SPOTIFY_BASE_ACCOUNT + "/api/token"
	secret := os.Getenv("SPOTIFY_CLIENT_SECRET")
	if secret == "" {
		return "", fmt.Errorf("secret is empty")
	}
	encodedSecret := base64.StdEncoding.EncodeToString([]byte(SPOTIFY_CLIENT_ID + ":" + secret))

	reqBody := url.Values{}
	reqBody.Set("grant_type", "refresh_token")
	reqBody.Set("refresh_token", refresh)

	encodedRequestBody := reqBody.Encode()

	req, err := http.NewRequest("POST", uri, strings.NewReader(encodedRequestBody))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Basic "+encodedSecret)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	var tokenResp model.GetRefreshTokenResponse
	err = json.NewDecoder(resp.Body).Decode(&tokenResp)
	if err != nil {
		panic(err)
	}

	err = resp.Body.Close()
	if err != nil {
		return "", err
	}
	return tokenResp.AccessToken, nil
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
