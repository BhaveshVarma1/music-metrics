package service

import (
	"bufio"
	"database/sql"
	"fmt"
	"math/rand"
	"music-metrics-back/dal"
	"os"
	"time"
)

var serverErrorStr = "Internal server error"
var SPOTIFY_BASE_ACCOUNT = "https://accounts.spotify.com"
var SPOTIFY_BASE_API = "https://api.spotify.com/v1"
var SPOTIFY_REDIRECT_URL = "https://dev.musicmetrics.app/spotify-landing" // note: this has to be the same as the one on the front end
var DEFAULT_ID_LENGTH = 32
var SPOTIFY_CLIENT_ID = "8b99139c99794d4b9e89b8367b0ac3f4"
var SEPARATOR = ";;"
var verbose = false

func GetSecret() string {
	file, err := os.Open("nogit2.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return ""
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		return scanner.Text()
	}
	return ""
}

func generateID(length int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = chars[r.Intn(len(chars))]
	}
	return string(result)
}

func CommitAndClose(tx *sql.Tx, db *sql.DB, commit bool) error {
	if commit {
		err := tx.Commit()
		if err != nil {
			return err
		}
	} else {
		err := tx.Rollback()
		if err != nil {
			return err
		}
	}
	err := dal.DbClose(db)
	if err != nil {
		return err
	}
	return nil
}

func PrintMessage(message string) {
	if verbose {
		fmt.Println(message)
	}
}
