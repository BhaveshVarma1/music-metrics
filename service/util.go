package service

import (
	"fmt"
	"math/rand"
	"time"
)

var serverErrorStr = "Internal server error"
var SPOTIFY_BASE_ACCOUNT = "https://accounts.spotify.com"
var SPOTIFY_BASE_API = "https://api.spotify.com/v1"
var SPOTIFY_REDIRECT_URL = "https://dev.musicmetrics.app/spotify-landing" // note: this has to be the same as the one on the front end
var DEFAULT_ID_LENGTH = 32
var SPOTIFY_CLIENT_ID = "8b99139c99794d4b9e89b8367b0ac3f4"
var SEPARATOR = ";;"
var verbose = true

func generateID(length int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = chars[r.Intn(len(chars))]
	}
	return string(result)
}

func PrintMessage(message string) {
	if verbose {
		fmt.Println(message)
	}
}

func SliceContainsString(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}
