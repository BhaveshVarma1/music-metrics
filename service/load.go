package service

import (
	"fmt"
	"music-metrics/da"
	"music-metrics/model"
	"strings"
)

func Load(history []model.ExtendedStreamingObject, username string) {

	// Instantiate DB connection
	tx, db, err := da.BeginTX()
	if err != nil {
		fmt.Println("Error beginning transaction in load service: ", err)
		return
	}

	// Get authtoken (for obtaining spotify data) and timestamp of account creation
	user, err := da.RetrieveUser(tx, username)
	if err != nil {
		fmt.Println("Error retrieving user in load service: ", err)
		return
	}
	token, err := RefreshToken(user.Refresh)
	if err != nil {
		fmt.Println("Error refreshing token in load service: ", err)
		return
	}
	endTime := user.Timestamp

	// Loop necessary to fix a weird bug. Due to primary key constraints, listens will not be duplicated.
	for {

		// Filter out bad data and build slice of unique track IDs
		var uniqueTrackIDs []string
		var listensToAdd []model.ListenBean
		for _, item := range history {
			if item.TrackName != "" && item.MsPlayed > 29999 {
				millis := DatetimeToUnixMilli(item.Timestamp)
				if millis != -1 {
					//item.Timestamp = strconv.FormatInt(millis, 10)
					if millis < endTime {
						// At this point, item is a listen of at least 30 seconds that occurred before the account was created
						trackID := item.TrackURI[strings.LastIndex(item.TrackURI, ":")+1:]
						if !SliceContainsString(uniqueTrackIDs, trackID) {
							uniqueTrackIDs = append(uniqueTrackIDs, trackID)
						}
						// Create listen object and add it to the slice that will be added to the DB
						listen := model.ListenBean{
							Username:  username,
							Timestamp: millis,
							TrackId:   trackID,
						}
						listensToAdd = append(listensToAdd, listen)
					}
				}
			}
		}
		if listensToAdd == nil {
			// User has already submitted this file
			fmt.Println("No new listens to add for this file.")
			return
		}

		// The listening history has now been filtered and the listen beans have been created
		// Because of foreign key constraints, the listensToAdd cannot be added to the DB until track/album metadata is added

		// We only want to get data from Spotify for tracks and albums that aren't already in the DB to save time
		var newTrackIDs []string
		var dbTrackIDs []string
		dbTracks, err := da.RetrieveAllTracks(tx)
		if err != nil {
			fmt.Println("Error retrieving tracks in load service: ", err)
			return
		}
		for _, track := range dbTracks {
			dbTrackIDs = append(dbTrackIDs, track.Id)
		}
		for _, track := range uniqueTrackIDs {
			if !SliceContainsString(dbTrackIDs, track) {
				newTrackIDs = append(newTrackIDs, track)
			}
		}

		// Obtain track metadata for new tracks
		tracks, err := getAllTrackData(token, newTrackIDs)
		if err != nil {
			fmt.Println("Error getting track data in load service: ", err)
			return
		}

		// Using the track metadata, obtain album metadata for new albums only
		var dbAlbumIDs []string
		dbAlbums, err := da.RetrieveAllAlbums(tx)
		if err != nil {
			fmt.Println("Error retrieving albums in load service: ", err)
			return
		}
		for _, album := range dbAlbums {
			dbAlbumIDs = append(dbAlbumIDs, album.Id)
		}
		var newUniqueAlbumIDs []string
		for _, track := range tracks {
			if !SliceContainsString(dbAlbumIDs, track.Album) {
				if !SliceContainsString(newUniqueAlbumIDs, track.Album) {
					newUniqueAlbumIDs = append(newUniqueAlbumIDs, track.Album)
				}
			}
		}

		albums, err := getAllAlbumData(token, newUniqueAlbumIDs)
		if err != nil {
			fmt.Println("Error getting album data in load service: ", err)
			return
		}

		// Add track and album metadata to DB
		// Add albums first due to foreign key constraints
		for _, album := range albums {
			if da.CreateAlbum(tx, &album) != nil {
				// There's a good chance the album already exists, so ignore the error
				continue
			}
		}
		for _, track := range tracks {
			if da.CreateTrack(tx, &track) != nil {
				// There's a good chance the track already exists, so ignore the error
				continue
			}
		}

		// Finally, we add the listensToAdd to the DB
		counter := 0
		for _, listen := range listensToAdd {
			counter++
			if da.CreateListen(tx, listen) != nil {
				counter--
				continue
			}
		}

		fmt.Println(counter, "/", len(listensToAdd), "listens added to the database")

		if counter == 0 {
			break
		}

	}

	if da.CommitAndClose(tx, db, true) != nil {
		fmt.Println("Error committing and closing transaction in load service: ", err)
	}
}

func getAllTrackData(token string, trackIDs []string) ([]model.TrackBean, error) {

	var beans []model.TrackBean

	tracks, err := GetSeveralTracks(token, trackIDs)
	if err != nil {
		fmt.Println("Error getting track data in load service: ", err)
		return nil, err
	}

	for _, track := range tracks {
		bean := model.TrackBean{
			Id:         track.ID,
			Name:       track.Name,
			Artist:     ArtistsToString(track.Artists),
			ArtistId:   ArtistIdsToString(track.Artists),
			Album:      track.Album.ID,
			Explicit:   track.Explicit,
			Popularity: track.Popularity,
			Duration:   track.DurationMs,
		}
		beans = append(beans, bean)
	}

	return beans, nil
}

func getAllAlbumData(token string, albumIDs []string) ([]model.AlbumBean, error) {

	var albums []model.AlbumBean

	albumsFull, err := GetSeveralAlbums(token, albumIDs)
	if err != nil {
		fmt.Println("Error getting album data in load service: ", err)
		return nil, err
	}

	for _, album := range albumsFull {
		albumBean := model.AlbumBean{
			Id:          album.ID,
			Name:        album.Name,
			Artist:      ArtistsToString(album.Artists),
			ArtistId:    ArtistIdsToString(album.Artists),
			Genre:       strings.Join(album.Genres, SEPARATOR),
			TotalTracks: album.TotalTracks,
			Year:        YearFromReleaseDate(album.ReleaseDate),
			Image:       GetAlbumImage(album),
			Popularity:  album.Popularity,
		}
		albums = append(albums, albumBean)
	}

	return albums, nil
}
