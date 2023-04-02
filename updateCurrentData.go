package main

import (
	"encoding/json"
	"fmt"
	"music-metrics/dal"
	"music-metrics/model"
	"music-metrics/service"
	"net/http"
	"strings"
)

func main() {

	tx, db, err := dal.BeginTX()
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
		track, err := getSongData(song.Id)
		if err != nil {
			fmt.Println("Error getting song data for song: " + song.Id)
			continue
		}

		newAlbum := model.AlbumBean{
			Id:          track.Album.ID,
			Name:        track.Album.Name,
			Artist:      artistsToString(track.Album.Artists),
			Genre:       strings.Join(track.Album.Genres, service.SEPARATOR),
			TotalTracks: track.Album.TotalTracks,
			Year:        yearFromReleaseDate(track.Album.ReleaseDate),
			Image:       track.Album.Images[0].URL,
			Popularity:  track.Album.Popularity,
		}

		album, err := dal.RetrieveAlbum(tx, track.Album.ID)
		if err != nil {
			fmt.Println("Error retrieving album: " + track.Album.ID)
			continue
		}

		if (album == model.AlbumBean{}) {
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
			Artist:     artistsToString(track.Artists),
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

func getSongData(songID string) (model.Track, error) {

	uri := service.SPOTIFY_BASE_API + "/tracks/" + songID

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return model.Track{}, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return model.Track{}, err
	}

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
