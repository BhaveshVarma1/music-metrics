package service

import (
	"fmt"
	"music-metrics/da"
	"music-metrics/model"
	"strconv"
	"time"
)

func Analytics() model.AnalyticsResponse {

	// Instantiate DB connection
	tx, db, err := da.BeginTX()
	if err != nil {
		fmt.Println("Error beginning transaction in analytics service: ", err)
		return model.AnalyticsResponse{}
	}

	// Get all users
	users, err := da.RetrieveAllUsers(tx)
	if err != nil {
		fmt.Println("Error retrieving users in analytics service: ", err)
		return model.AnalyticsResponse{}
	}

	// Get last action
	username, timestamp, err := da.RetrieveLastAction(tx)
	if err != nil {
		fmt.Println("Error retrieving last action in analytics service: ", err)
		return model.AnalyticsResponse{}
	}

	// Commit changes and close the connection
	if da.CommitAndClose(tx, db, true) != nil {
		fmt.Println("Error committing and closing in analytics service: ", err)
		return model.AnalyticsResponse{}
	}

	// Get display name from username
	displayName := ""
	for _, user := range users {
		if user.Username == username {
			displayName = user.DisplayName
			break
		}
	}

	// Convert timestamp into readable format
	readableTime, err := formatTimestamp(timestamp)
	if err != nil {
		fmt.Println("Error formatting timestamp in analytics service: ", err)
		return model.AnalyticsResponse{}
	}

	return model.AnalyticsResponse{
		TotalUsers: len(users),
		LastAction: displayName + " on the " + readableTime,
	}
}

func formatTimestamp(timestampMillis int64) (string, error) {
	loc, err := time.LoadLocation("America/Denver")
	if err != nil {
		return "", err
	}

	// Convert milliseconds to seconds and adjust to the local time zone
	timestamp := time.Unix(timestampMillis/1000, 0).In(loc)

	// Extract day
	day := timestamp.Day()

	// Determine the appropriate suffix for the day (e.g., "23rd")
	daySuffix := "th"
	if day == 1 || day == 21 || day == 31 {
		daySuffix = "st"
	} else if day == 2 || day == 22 {
		daySuffix = "nd"
	} else if day == 3 || day == 23 {
		daySuffix = "rd"
	}

	// Format the time
	formattedTime := timestamp.Format("3:04 PM")

	// Combine day, suffix, and time into the desired format
	formattedTimestamp := strconv.Itoa(day) + daySuffix + " at " + formattedTime

	return formattedTimestamp, nil
}
