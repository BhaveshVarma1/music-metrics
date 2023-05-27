package service

import (
	"database/sql"
	"fmt"
	"music-metrics/da"
	"music-metrics/model"
	"strconv"
	"time"
)

func StartTracking() {

	fmt.Println("STARTING TRACKING SCRIPT...")

	for {

		startTime := time.Now()

		// Instantiate connection
		tx, db, err := da.BeginTX()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		users, err := da.RetrieveAllUsers(tx)
		if err != nil {
			if da.CommitAndClose(tx, db, false) != nil {
				fmt.Println("Error committing transaction")
				return
			}
		}

		// For every user in the database
		for _, user := range users {
			startTimeUser := time.Now()
			// Get new access token
			newToken, err := RefreshToken(user.Refresh)
			if err != nil || newToken == "" {
				fmt.Println("Error refreshing token for username: " + user.Username)
				continue
			}

			// Use new access token to call /recently-played
			recentlyPlayed, err := GetRecentlyPlayed(newToken)
			if err != nil {
				fmt.Println("Error getting recently played for username: " + user.Username)
				continue
			}

			// Get most recent listen
			mostRecentListen, err := da.GetMostRecentListen(tx, user.Username)
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

		if da.CommitAndClose(tx, db, true) != nil {
			fmt.Println("Error committing transaction")
			return
		}

		// Sleep for 2 hours - time it took to run the script
		time.Sleep((2 * time.Hour) - time.Since(startTime))
	}

}

func loopThroughRecentListens(listens []model.RecentlyPlayedObject, tx *sql.Tx, user model.UserBean, oldTime int64) int {
	newSongsCount := 0
	for _, rpObj := range listens {
		if rpObj.Timestamp > oldTime {
			// Add song to database if it isn't already there
			newSongsCount++
			song, err := da.RetrieveSong(tx, rpObj.Song.Id)
			if err != nil {
				fmt.Println("Error retrieving song for username: " + user.Username)
				fmt.Println(err)
				continue
			}
			if (song == model.SongBean{}) {
				err = da.CreateSong(tx, &rpObj.Song)
				if err != nil {
					fmt.Println("Error creating song for username: " + user.Username)
					continue
				}
			} else {
				// Update song if it is already there
				err = da.UpdateSong(tx, &rpObj.Song)
				if err != nil {
					fmt.Println("Error updating song for username: " + user.Username)
					continue
				}
			}

			// Add album to database if it isn't already there
			album, err := da.RetrieveAlbum(tx, rpObj.Album.Id)
			if err != nil {
				fmt.Println("Error retrieving album for username: " + user.Username)
				fmt.Println(err)
				continue
			}
			if (album == model.AlbumBean{}) {
				err = da.CreateAlbum(tx, &rpObj.Album)
				if err != nil {
					fmt.Println("Error creating album for username: " + user.Username)
					continue
				}
			} else {
				// Update album if it is already there
				err = da.UpdateAlbum(tx, &rpObj.Album)
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
			err = da.CreateListen(tx, newListen)
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
