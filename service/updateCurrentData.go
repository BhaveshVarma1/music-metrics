package service

import (
	"fmt"
	"music-metrics/da"
	"music-metrics/model"
	"strings"
	"time"
)

func UpdateCurrentData() {

	// Instantiate DB connection
	tx, db, err := da.BeginTX()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Grab any user's refresh token
	user, err := da.RetrieveUser(tx, "prattnj")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	token, err := RefreshToken(user.Refresh)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Grab songs to iterate through
	songs, err := da.RetrieveAllSongs(tx)
	if err != nil {
		if da.CommitAndClose(tx, db, false) != nil {
			fmt.Println("Error committing transaction")
			return
		}
	}

	// Grab albums to iterate through
	albums, err := da.RetrieveAllAlbums(tx)
	if err != nil {
		if da.CommitAndClose(tx, db, false) != nil {
			fmt.Println("Error committing transaction")
			return
		}
	}

	// Update all songs
	iteration := 0
	for _, song := range songs {
		fmt.Printf("Iteration: %d\n", iteration)
		track, err := GetTrack(token, song.Id)
		if err != nil {
			fmt.Println("Error getting song data for song: " + song.Id)
			continue
		}

		newSong := model.SongBean{
			Id:         track.ID,
			Name:       track.Name,
			Artist:     ArtistsToString(track.Artists),
			ArtistId:   ArtistIdsToString(track.Artists),
			Album:      track.Album.ID,
			Explicit:   track.Explicit,
			Popularity: track.Popularity,
			Duration:   track.DurationMs,
		}

		if da.UpdateSong(tx, &newSong) != nil {
			fmt.Println("Error updating song: " + newSong.Id)
			continue
		}

		time.Sleep(100 * time.Millisecond)
		iteration++
	}

	// Update all albums
	iteration = 0
	for _, album := range albums {
		fmt.Printf("Iteration: %d\n", iteration)
		spotifyAlbum, err := GetAlbum(token, album.Id)
		if err != nil {
			fmt.Println("Error getting album data for album: " + album.Id)
			continue
		}

		newAlbum := model.AlbumBean{
			Id:          spotifyAlbum.ID,
			Name:        spotifyAlbum.Name,
			Artist:      ArtistsToString(spotifyAlbum.Artists),
			ArtistId:    ArtistIdsToString(spotifyAlbum.Artists),
			Genre:       strings.Join(spotifyAlbum.Genres, SEPARATOR),
			TotalTracks: spotifyAlbum.TotalTracks,
			Year:        YearFromReleaseDate(spotifyAlbum.ReleaseDate),
			Image:       spotifyAlbum.Images[0].URL,
			Popularity:  spotifyAlbum.Popularity,
		}

		if da.UpdateAlbum(tx, &newAlbum) != nil {
			fmt.Println("Error updating album: " + newAlbum.Id)
			continue
		}

		time.Sleep(100 * time.Millisecond)
		iteration++
	}

	// Commit transaction
	if da.CommitAndClose(tx, db, true) != nil {
		fmt.Println("Error committing transaction")
		return
	}

}
