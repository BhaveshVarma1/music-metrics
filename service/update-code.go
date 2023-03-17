package service

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"music-metrics-back/dal"
	"music-metrics-back/model"
	"net/http"
	"net/url"
	"strings"
)

func UpdateCode(code string) model.UpdateCodeResponse {

	db := dal.Db()
	if db == nil {
		return model.UpdateCodeResponse{Success: false, Message: serverErrorStr}
	}
	tx, err := db.Begin()
	if err != nil {
		return model.UpdateCodeResponse{Success: false, Message: serverErrorStr}
	}

	var response model.UpdateCodeResponse

	/*userWithCode, err := dal.RetrieveUserByCode(tx, code)
	if err != nil {
		if commitAndClose(tx, db, false) != nil {
			return model.UpdateCodeResponse{Success: false, Message: serverErrorStr}
		}
		return model.UpdateCodeResponse{Success: false, Message: serverErrorStr}
	}
	if (userWithCode != model.User{}) {
		// User with code already exists, return associated auth token
		tokenToReturn, err := dal.RetrieveAuthTokenByUsername(tx, userWithCode.Username)
		if err != nil {
			if commitAndClose(tx, db, false) != nil {
				return model.UpdateCodeResponse{Success: false, Message: serverErrorStr}
			}
			return model.UpdateCodeResponse{Success: false, Message: serverErrorStr}
		}
		if tokenToReturn == (model.AuthToken{}) {
			fmt.Print("tokenToReturn is empty, this should not happen")
		}
		response = model.UpdateCodeResponse{
			Success:     true,
			Token:       tokenToReturn.Token,
			Username:    userWithCode.Username,
			DisplayName: userWithCode.DisplayName,
			Email:       userWithCode.Email,
		}
	} else {*/

	// make http request to /api/token with code
	// access token and refresh token are returned, put them aside for now (dynamically, in memory)
	// use access token to make http request to /get/me
	// user info is returned, use it to create new user along with code, refresh token, and access token
	// add the new user to DB
	// create new auth token for user
	// add new auth token to DB
	// return auth token just created

	// Request Access/Refresh Token from Spotify
	fmt.Println("Getting access token from Spotify...")
	accessToken, refreshToken, err := requestAccessToken(code)
	if err != nil {
		if commitAndClose(tx, db, false) != nil {
			return model.UpdateCodeResponse{Success: false, Message: serverErrorStr}
		}
		return model.UpdateCodeResponse{Success: false, Message: serverErrorStr}
	}
	fmt.Println("Successfully got access token from Spotify, now getting user info...")

	// Get User Info from Spotify
	currUser, err := requestUserInfo(accessToken)
	if err != nil {
		if commitAndClose(tx, db, false) != nil {
			return model.UpdateCodeResponse{Success: false, Message: serverErrorStr}
		}
		return model.UpdateCodeResponse{Success: false, Message: err.Error()}
	}
	fmt.Println("Successfully got user info from Spotify")
	currUser.Code = code
	currUser.Refresh = refreshToken
	currUser.Access = accessToken

	// Determine if user already exists
	fmt.Println("Checking if user already exists...")
	existingUser, err := dal.RetrieveUser(tx, currUser.Username)
	if err != nil {
		if commitAndClose(tx, db, false) != nil {
			return model.UpdateCodeResponse{Success: false, Message: serverErrorStr}
		}
		return model.UpdateCodeResponse{Success: false, Message: err.Error()}
	}

	var token model.AuthToken

	// If user does not exist, create user and auth token
	if (existingUser == model.User{}) {
		fmt.Println("User does not exist, creating user and auth token...")
		err = dal.CreateUser(tx, currUser)
		if err != nil {
			if commitAndClose(tx, db, false) != nil {
				return model.UpdateCodeResponse{Success: false, Message: serverErrorStr}
			}
			return model.UpdateCodeResponse{Success: false, Message: serverErrorStr}
		}
		token = model.AuthToken{
			Username: currUser.Username,
			Token:    generateID(DEFAULT_ID_LENGTH),
		}
		err = dal.CreateAuthToken(tx, token)
		if err != nil {
			if commitAndClose(tx, db, false) != nil {
				return model.UpdateCodeResponse{Success: false, Message: serverErrorStr}
			}
			return model.UpdateCodeResponse{Success: false, Message: serverErrorStr}
		}
		fmt.Println("Successfully created user and auth token")
	} else { // User already exists, update them and get auth token
		fmt.Println("User already exists, updating user and getting auth token...")
		err = dal.UpdateUser(tx, currUser)
		if err != nil {
			if commitAndClose(tx, db, false) != nil {
				return model.UpdateCodeResponse{Success: false, Message: serverErrorStr}
			}
			return model.UpdateCodeResponse{Success: false, Message: err.Error()}
		}
		token, err = dal.RetrieveAuthTokenByUsername(tx, currUser.Username)
		if err != nil {
			if commitAndClose(tx, db, false) != nil {
				return model.UpdateCodeResponse{Success: false, Message: serverErrorStr}
			}
			return model.UpdateCodeResponse{Success: false, Message: err.Error()}
		}
		if (token == model.AuthToken{}) {
			fmt.Println("Token is null, returning (this should not happen)")
			if commitAndClose(tx, db, false) != nil {
				return model.UpdateCodeResponse{Success: false, Message: serverErrorStr}
			}
			return model.UpdateCodeResponse{Success: false, Message: serverErrorStr}
		}
		fmt.Println("Successfully updated user and got auth token")
	}

	response = model.UpdateCodeResponse{
		Success:     true,
		Token:       token.Token,
		Username:    currUser.Username,
		DisplayName: currUser.DisplayName,
		Email:       currUser.Email,
	}

	//}

	fmt.Println("Committing to DB...")
	if commitAndClose(tx, db, true) != nil {
		return model.UpdateCodeResponse{Success: false, Message: serverErrorStr}
	}

	return response
}

func requestAccessToken(code string) (string, string, error) {

	uri := SPOTIFY_BASE_ACCOUNT + "/api/token"
	secret := GetSecret()
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

func requestUserInfo(access string) (model.User, error) {

	uri := SPOTIFY_BASE_API + "/me"

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return model.User{}, err
	}

	req.Header.Set("Authorization", "Bearer "+access)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return model.User{}, err
	}

	var getMeResponse model.GetMeResponse
	err = json.NewDecoder(resp.Body).Decode(&getMeResponse)
	if err != nil {
		return model.User{}, err
	}

	err = resp.Body.Close()
	if err != nil {
		return model.User{}, err
	}
	return model.User{Username: getMeResponse.ID, DisplayName: getMeResponse.DisplayName, Email: getMeResponse.Email}, nil
}
