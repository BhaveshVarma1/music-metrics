package service

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"music-metrics/model"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

func GetTrack(token string, trackID string) (model.Track, error) {

	uri := SPOTIFY_BASE_API + "/tracks/" + trackID

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return model.Track{}, err
	}

	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return model.Track{}, err
	}

	var respBody model.Track
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		return model.Track{}, err
	}

	err = resp.Body.Close()
	if err != nil {
		return model.Track{}, err
	}

	return respBody, nil

}

func GetSeveralTracks(token string, trackIDs []string) ([]model.Track, error) {

	// Split trackIDs into groups of 50
	var trackIDGroups [][]string
	for i := 0; i < len(trackIDs); i += 50 {
		end := i + 50
		if end > len(trackIDs) {
			end = len(trackIDs)
		}
		trackIDGroups = append(trackIDGroups, trackIDs[i:end])
	}

	// Get data for each group of 50
	var tracks []model.Track
	for _, group := range trackIDGroups {

		uri := SPOTIFY_BASE_API + "/tracks?ids=" + strings.Join(group, ",")

		req, err := http.NewRequest("GET", uri, nil)
		if err != nil {
			return nil, err
		}

		req.Header.Set("Authorization", "Bearer "+token)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}

		var allTracks model.SeveralTracks
		err = json.NewDecoder(resp.Body).Decode(&allTracks)
		if err != nil {
			return nil, err
		}

		err = resp.Body.Close()
		if err != nil {
			return nil, err
		}

		// Add newly fetched data to tracks[]
		for _, track := range allTracks.Tracks {
			tracks = append(tracks, track)
		}
	}

	return tracks, nil
}

func GetAlbum(token string, albumID string) (model.Album, error) {

	uri := SPOTIFY_BASE_API + "/albums/" + albumID

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return model.Album{}, err
	}

	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return model.Album{}, err
	}

	fmt.Println("RESPONSE: " + resp.Status)

	var respBody model.Album
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		return model.Album{}, err
	}

	err = resp.Body.Close()
	if err != nil {
		return model.Album{}, err
	}

	return respBody, nil
}

func GetSeveralAlbums(token string, albumIDs []string) ([]model.Album, error) {

	// Split albumIDs into groups of 20
	var albumIDGroups [][]string
	for i := 0; i < len(albumIDs); i += 20 {
		end := i + 20
		if end > len(albumIDs) {
			end = len(albumIDs)
		}
		albumIDGroups = append(albumIDGroups, albumIDs[i:end])
	}

	// Get data for each group of 20
	var albums []model.Album
	for _, group := range albumIDGroups {

		uri := SPOTIFY_BASE_API + "/albums?ids=" + strings.Join(group, ",")

		req, err := http.NewRequest("GET", uri, nil)
		if err != nil {
			return nil, err
		}

		req.Header.Set("Authorization", "Bearer "+token)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}

		var allAlbums model.SeveralAlbums
		err = json.NewDecoder(resp.Body).Decode(&allAlbums)
		if err != nil {
			return nil, err
		}

		err = resp.Body.Close()
		if err != nil {
			return nil, err
		}

		// Add newly fetched data to albums[]
		for _, album := range allAlbums.Albums {
			albums = append(albums, album)
		}
	}

	return albums, nil
}

func GetRecentlyPlayed(token string) ([]model.RecentlyPlayedObject, error) {

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
		track := model.TrackBean{
			Id:         item.Track.ID,
			Name:       item.Track.Name,
			Artist:     ArtistsToString(item.Track.Artists),
			ArtistId:   ArtistIdsToString(item.Track.Artists),
			Album:      item.Track.Album.ID,
			Explicit:   item.Track.Explicit,
			Popularity: item.Track.Popularity,
			Duration:   item.Track.DurationMs,
		}
		album := model.AlbumBean{
			Id:          item.Track.Album.ID,
			Name:        item.Track.Album.Name,
			Artist:      ArtistsToString(item.Track.Album.Artists),
			ArtistId:    ArtistIdsToString(item.Track.Album.Artists),
			Genre:       strings.Join(item.Track.Album.Genres, SEPARATOR),
			TotalTracks: item.Track.Album.TotalTracks,
			Year:        YearFromReleaseDate(item.Track.Album.ReleaseDate),
			Image:       item.Track.Album.Images[0].URL,
			Popularity:  item.Track.Album.Popularity,
		}
		returnObj := model.RecentlyPlayedObject{
			Track:     track,
			Album:     album,
			Timestamp: DatetimeToUnixMilli(item.PlayedAt),
		}
		toReturn = append(toReturn, returnObj)
	}

	return toReturn, nil
}

func RefreshToken(refresh string) (string, error) {

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

func RequestAccessToken(code string) (string, string, error) {

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

func RequestUserInfo(access string) (model.UserBean, error) {

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
