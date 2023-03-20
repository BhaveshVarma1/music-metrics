package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"music-metrics-back/dal"
	"music-metrics-back/model"
	"music-metrics-back/service"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func main() {

	fmt.Println("STARTING TRACKING SCRIPT...")

	for {
		// Instantiate connection
		db := dal.Db()
		if db == nil {
			fmt.Println("Error connecting to database")
			return
		}
		tx, err := db.Begin()
		if err != nil {
			fmt.Println("Error starting transaction")
			return
		}

		users, err := dal.RetrieveAllUsers(tx)
		if err != nil {
			if service.CommitAndClose(tx, db, false) != nil {
				fmt.Println("Error committing transaction")
				return
			}
		}

		for _, user := range users {
			// Get new access token
			newToken, err := refreshToken(user.Refresh)
			if err != nil || newToken == "" {
				fmt.Println("Error refreshing token for username: " + user.Username)
				continue
			}
			fmt.Println("New token: " + newToken)

			// Use new access token to call /recently-played
			recentlyPlayed, err := getRecentlyPlayed(newToken)
			fmt.Println("Recently played: " + strconv.Itoa(len(recentlyPlayed)))
		}

		if service.CommitAndClose(tx, db, true) != nil {
			fmt.Println("Error committing transaction")
			return
		}

		// Sleep for 2 hours
		time.Sleep(2 * time.Hour)
	}

}

func refreshToken(refresh string) (string, error) {

	uri := service.SPOTIFY_BASE_ACCOUNT + "/api/token"
	secret := service.GetSecret()
	if secret == "" {
		return "", fmt.Errorf("secret is empty")
	}
	encodedSecret := base64.StdEncoding.EncodeToString([]byte(service.SPOTIFY_CLIENT_ID + ":" + secret))

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

	uri := service.SPOTIFY_BASE_API + "/me/player/recently-played"

	payload := map[string]int64{
		"limit":  50,
		"before": time.Now().UnixMilli(),
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", uri, bytes.NewBuffer(jsonPayload))
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
		song := model.Song{
			Id:         item.Track.ID,
			Name:       item.Track.Name,
			Artist:     artistsToString(item.Track.Artists),
			Album:      item.Track.Album.Name,
			Genre:      strings.Join(item.Track.Album.Genres, service.SEPARATOR),
			Explicit:   item.Track.Explicit,
			Popularity: item.Track.Popularity,
			Duration:   item.Track.DurationMs,
			Year:       yearFromReleaseDate(item.Track.Album.ReleaseDate),
		}
		returnObj := model.RecentlyPlayedObject{
			Song:      song,
			Timestamp: item.PlayedAt,
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
	return strings.Join(arr, service.SEPARATOR)
}

func yearFromReleaseDate(date string) int {
	i, err := strconv.Atoi(date[:4])
	if err != nil {
		return -1
	}
	return i
}
