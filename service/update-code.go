package service

import (
	"music-metrics/da"
	"music-metrics/model"
	"time"
)

func UpdateCode(code string) model.UpdateCodeResponse {

	tx, db, err := da.BeginTX()
	if err != nil {
		return model.UpdateCodeResponse{Success: false, Message: serverErrorStr}
	}

	var response model.UpdateCodeResponse

	// Request Access/Refresh Token from Spotify
	PrintMessage("\nGetting access token from Spotify...")
	accessToken, refreshToken, err := RequestAccessToken(code)
	if err != nil {
		if da.CommitAndClose(tx, db, false) != nil {
			return model.UpdateCodeResponse{Success: false, Message: serverErrorStr}
		}
		return model.UpdateCodeResponse{Success: false, Message: serverErrorStr}
	}
	PrintMessage("Successfully got access token from Spotify, now getting user info...")

	// Get UserBean Info from Spotify
	currUser, err := RequestUserInfo(accessToken)
	if err != nil {
		if da.CommitAndClose(tx, db, false) != nil {
			return model.UpdateCodeResponse{Success: false, Message: serverErrorStr}
		}
		return model.UpdateCodeResponse{Success: false, Message: err.Error()}
	}
	PrintMessage("Successfully got user info from Spotify")
	currUser.Refresh = refreshToken

	// Determine if user already exists
	PrintMessage("Checking if user already exists...")
	existingUser, err := da.RetrieveUser(tx, currUser.Username)
	if err != nil {
		if da.CommitAndClose(tx, db, false) != nil {
			return model.UpdateCodeResponse{Success: false, Message: serverErrorStr}
		}
		return model.UpdateCodeResponse{Success: false, Message: err.Error()}
	}

	var token model.AuthTokenBean

	// If user does not exist, create user and auth token
	if (existingUser == model.UserBean{}) {

		PrintMessage("UserBean does not exist, creating user and auth token...")
		currUser.Timestamp = time.Now().UnixMilli()
		err = da.CreateUser(tx, currUser)
		if err != nil {
			if da.CommitAndClose(tx, db, false) != nil {
				return model.UpdateCodeResponse{Success: false, Message: serverErrorStr}
			}
			return model.UpdateCodeResponse{Success: false, Message: serverErrorStr}
		}
		token = model.AuthTokenBean{
			Username: currUser.Username,
			Token:    GenerateID(DEFAULT_ID_LENGTH),
		}
		err = da.CreateAuthToken(tx, token)
		if err != nil {
			if da.CommitAndClose(tx, db, false) != nil {
				return model.UpdateCodeResponse{Success: false, Message: serverErrorStr}
			}
			return model.UpdateCodeResponse{Success: false, Message: serverErrorStr}
		}
		PrintMessage("Successfully created user and auth token")

		// Add recent listens to DB for instant access
		PrintMessage("Adding recent listens to DB...")
		recentListens, err := GetRecentlyPlayed(accessToken)
		if err != nil {
			if da.CommitAndClose(tx, db, false) != nil {
				return model.UpdateCodeResponse{Success: false, Message: serverErrorStr}
			}
			return model.UpdateCodeResponse{Success: false, Message: err.Error()}
		}

		// Add all 50 recent listens to DB
		loopThroughRecentListens(recentListens, tx, currUser, 0)
		PrintMessage("Successfully added listens to DB")

	} else {

		// UserBean already exists, update them and get auth token
		PrintMessage("UserBean already exists, updating user and getting auth token...")
		currUser.Timestamp = existingUser.Timestamp
		err = da.UpdateUser(tx, currUser)
		if err != nil {
			if da.CommitAndClose(tx, db, false) != nil {
				return model.UpdateCodeResponse{Success: false, Message: serverErrorStr}
			}
			return model.UpdateCodeResponse{Success: false, Message: err.Error()}
		}
		token, err = da.RetrieveAuthTokenByUsername(tx, currUser.Username)
		if err != nil {
			if da.CommitAndClose(tx, db, false) != nil {
				return model.UpdateCodeResponse{Success: false, Message: serverErrorStr}
			}
			return model.UpdateCodeResponse{Success: false, Message: err.Error()}
		}
		if (token == model.AuthTokenBean{}) {
			PrintMessage("Token is null, returning (this should not happen)")
			if da.CommitAndClose(tx, db, false) != nil {
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
	if da.CommitAndClose(tx, db, true) != nil {
		return model.UpdateCodeResponse{Success: false, Message: serverErrorStr}
	}

	return response
}
