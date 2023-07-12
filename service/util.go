package service

import (
	"fmt"
	"math/rand"
	"music-metrics/model"
	"strconv"
	"strings"
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

func ArtistIdsToString(artists []model.Artist) string {
	var arr []string
	for _, artist := range artists {
		arr = append(arr, artist.ID)
	}
	return strings.Join(arr, SEPARATOR)
}

func ArtistsToString(artists []model.Artist) string {
	var arr []string
	for _, artist := range artists {
		arr = append(arr, artist.Name)
	}
	return strings.Join(arr, SEPARATOR)
}

func DatetimeToUnixMilli(datetime string) int64 {
	t, err := time.Parse(time.RFC3339, datetime)
	if err != nil {
		return -1
	}
	return t.UnixMilli()
}

func GenerateID(length int) string {
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

func YearFromReleaseDate(date string) int {
	i, err := strconv.Atoi(date[:4])
	if err != nil {
		return -1
	}
	return i
}
