package service

import (
	"fmt"
	"math/rand"
	"music-metrics/model"
	"net/smtp"
	"os"
	"strconv"
	"strings"
	"time"
)

var serverErrorStr = "Internal server error"
var SPOTIFY_BASE_ACCOUNT = "https://accounts.spotify.com"
var SPOTIFY_BASE_API = "https://api.spotify.com/v1"
var SPOTIFY_REDIRECT_URL = "https://musicmetrics.app/spotify-landing" // note: this has to be the same as the one on the front end
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

func GetAlbumImage(album model.Album) string {
	// Donda album cover for non-existent album covers
	defaultImage := "https://i.scdn.co/image/ab67616d0000b273cad190f1a73c024e5a40dddd"
	if album.Images == nil || len(album.Images) == 0 {
		return defaultImage
	} else {
		return album.Images[0].URL
	}
}

func PrintMessage(message string) {
	if verbose {
		fmt.Println(message)
	}
}

func SendEmail(subject string, message string) {

	from := "musicmetricsapp@gmail.com"
	password := os.Getenv("MM_GMAIL")

	auth := smtp.PlainAuth("", from, password, "smtp.gmail.com")

	to := []string{"prattnj@gmail.com"}

	header := make(map[string]string)
	header["From"] = "Music Metrics <" + from + ">"
	body := "To: " + to[0] + "\r\nSubject: " + subject + "\r\n\r\n" + message

	err := smtp.SendMail("smtp.gmail.com:587", auth, from, to, []byte(body))
	if err != nil {
		fmt.Println("Error sending email: ", err)
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
