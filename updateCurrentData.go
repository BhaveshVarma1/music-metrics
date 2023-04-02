package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"music-metrics/dal"
	"music-metrics/model"
	"music-metrics/service"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func main() {

	tx, db, err := dal.BeginTX()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	user, err := dal.RetrieveUser(tx, "prattnj")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	token, err := refreshToken2(user.Refresh)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	songs, err := dal.RetrieveAllSongs(tx)
	if err != nil {
		if dal.CommitAndClose(tx, db, false) != nil {
			fmt.Println("Error committing transaction")
			return
		}
	}

	for _, song := range songs {
		track, err := getSongData(token, song.Id)
		if err != nil {
			fmt.Println("Error getting song data for song: " + song.Id)
			continue
		}

		newAlbum := model.AlbumBean{
			Id:          track.Album.ID,
			Name:        track.Album.Name,
			Artist:      artistsToString2(track.Album.Artists),
			Genre:       strings.Join(track.Album.Genres, service.SEPARATOR),
			TotalTracks: track.Album.TotalTracks,
			Year:        yearFromReleaseDate2(track.Album.ReleaseDate),
			Image:       track.Album.Images[0].URL,
			Popularity:  track.Album.Popularity,
		}

		album, err := dal.RetrieveAlbum(tx, track.Album.ID)
		if err != nil {
			fmt.Println("Error retrieving album: " + track.Album.ID)
			continue
		}

		if (album == model.AlbumBean{}) {
			fmt.Println("Creating album: " + newAlbum.Id)
			if dal.CreateAlbum(tx, &newAlbum) != nil {
				fmt.Println("Error creating album: " + newAlbum.Id)
				continue
			}
		} else {
			if dal.UpdateAlbum(tx, &newAlbum) != nil {
				fmt.Println("Error updating album: " + newAlbum.Id)
				continue
			}
		}

		newSong := model.SongBean{
			Id:         track.ID,
			Name:       track.Name,
			Artist:     artistsToString2(track.Artists),
			Album:      track.Album.ID,
			Explicit:   track.Explicit,
			Popularity: track.Popularity,
			Duration:   track.DurationMs,
		}

		if dal.UpdateSong(tx, &newSong) != nil {
			fmt.Println("Error updating song: " + newSong.Id)
			continue
		}

	}

}

func getSongData(token string, songID string) (model.Track, error) {

	uri := service.SPOTIFY_BASE_API + "/tracks/" + songID

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return model.Track{}, err
	}

	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return model.Track{}, err
	}

	fmt.Println("RESPONSE: " + resp.Status)

	var respBody model.Track
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		return model.Track{}, err
	}

	err = resp.Body.Close()
	if err != nil {
		return model.Track{}, err
	}

	return respBody, nil

}

func refreshToken2(refresh string) (string, error) {

	uri := service.SPOTIFY_BASE_ACCOUNT + "/api/token"
	secret := os.Getenv("SPOTIFY_CLIENT_SECRET")
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

func artistsToString2(artists []model.Artist) string {
	var arr []string
	for _, artist := range artists {
		arr = append(arr, artist.Name)
	}
	return strings.Join(arr, service.SEPARATOR)
}

func yearFromReleaseDate2(date string) int {
	i, err := strconv.Atoi(date[:4])
	if err != nil {
		return -1
	}
	return i
}
