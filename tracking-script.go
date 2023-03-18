package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"music-metrics-back/dal"
	"music-metrics-back/model"
	"music-metrics-back/service"
	"net/http"
	"net/url"
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

			// use new access token to call /recently-played
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
