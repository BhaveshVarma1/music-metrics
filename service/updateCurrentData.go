package service

import (
	"encoding/json"
	"fmt"
	"music-metrics/dal"
	"music-metrics/model"
	"net/http"
	"strings"
	"time"
)

func UpdateCurrentData() {

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

	token, err := refreshToken(user.Refresh)
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

	albums, err := dal.RetrieveAllAlbums(tx)
	if err != nil {
		if dal.CommitAndClose(tx, db, false) != nil {
			fmt.Println("Error committing transaction")
			return
		}
	}

	iteration := 0
	for _, song := range songs {
		fmt.Printf("Iteration: %d\n", iteration)
		track, err := getSongData(token, song.Id)
		if err != nil {
			fmt.Println("Error getting song data for song: " + song.Id)
			continue
		}

		newSong := model.SongBean{
			Id:         track.ID,
			Name:       track.Name,
			Artist:     artistsToString(track.Artists),
			ArtistId:   artistsToString(track.Artists),
			Album:      track.Album.ID,
			Explicit:   track.Explicit,
			Popularity: track.Popularity,
			Duration:   track.DurationMs,
		}

		if dal.UpdateSong(tx, &newSong) != nil {
			fmt.Println("Error updating song: " + newSong.Id)
			continue
		}

		time.Sleep(100 * time.Millisecond)
		iteration++
	}

	iteration = 0
	for _, album := range albums {
		fmt.Printf("Iteration: %d\n", iteration)
		spotifyAlbum, err := getAlbumData(token, album.Id)
		if err != nil {
			fmt.Println("Error getting album data for album: " + album.Id)
			continue
		}

		newAlbum := model.AlbumBean{
			Id:          spotifyAlbum.ID,
			Name:        spotifyAlbum.Name,
			Artist:      artistsToString(spotifyAlbum.Artists),
			ArtistId:    artistIdsToString(spotifyAlbum.Artists),
			Genre:       strings.Join(spotifyAlbum.Genres, SEPARATOR),
			TotalTracks: spotifyAlbum.TotalTracks,
			Year:        yearFromReleaseDate(spotifyAlbum.ReleaseDate),
			Image:       spotifyAlbum.Images[0].URL,
			Popularity:  spotifyAlbum.Popularity,
		}

		if dal.UpdateAlbum(tx, &newAlbum) != nil {
			fmt.Println("Error updating album: " + newAlbum.Id)
			continue
		}

		time.Sleep(100 * time.Millisecond)
		iteration++
	}

	// Commit transaction
	if dal.CommitAndClose(tx, db, true) != nil {
		fmt.Println("Error committing transaction")
		return
	}

}

func getSongData(token string, songID string) (model.Track, error) {

	uri := SPOTIFY_BASE_API + "/tracks/" + songID

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

func getAlbumData(token string, albumID string) (model.Album, error) {

	uri := SPOTIFY_BASE_API + "/albums/" + albumID

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return model.Album{}, err
	}

	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return model.Album{}, err
	}

	fmt.Println("RESPONSE: " + resp.Status)

	var respBody model.Album
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		return model.Album{}, err
	}

	err = resp.Body.Close()
	if err != nil {
		return model.Album{}, err
	}

	return respBody, nil

}
