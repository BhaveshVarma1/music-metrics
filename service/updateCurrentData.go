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

	// Grab tracks to iterate through
	tracks, err := da.RetrieveAllTracks(tx)
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

	// Update all tracks
	iteration := 0
	for _, track := range tracks {
		fmt.Printf("Iteration: %d\n", iteration)
		spotifyTrack, err := GetTrack(token, track.Id)
		if err != nil {
			fmt.Println("Error getting track data for track: " + track.Id)
			continue
		}

		newTrack := model.TrackBean{
			Id:         spotifyTrack.ID,
			Name:       spotifyTrack.Name,
			Artist:     ArtistsToString(spotifyTrack.Artists),
			ArtistId:   ArtistIdsToString(spotifyTrack.Artists),
			Album:      spotifyTrack.Album.ID,
			Explicit:   spotifyTrack.Explicit,
			Popularity: spotifyTrack.Popularity,
			Duration:   spotifyTrack.DurationMs,
		}

		if da.UpdateTrack(tx, &newTrack) != nil {
			fmt.Println("Error updating track: " + newTrack.Id)
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
			Image:       GetAlbumImage(spotifyAlbum),
			Popularity:  spotifyAlbum.Popularity,
		}

		if da.UpdateAlbum(tx, &newAlbum) != nil {
			fmt.Println("Error updating album: " + newAlbum.Id)
			continue
		}

		// Sleep to avoid rate limiting
		time.Sleep(100 * time.Millisecond)
		iteration++
	}

	// Commit transaction
	if da.CommitAndClose(tx, db, true) != nil {
		fmt.Println("Error committing transaction")
		return
	}

}
