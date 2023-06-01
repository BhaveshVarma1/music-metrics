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

	/*dbListens, err := da.RetrieveAllListensForUser(tx, username)
	if err != nil {
		fmt.Println("Error retrieving listensToAdd in load service: ", err)
		return
	}*/
	dbTimestamps, err := da.RetrieveAllTimestampsForUser(tx, username)
	if err != nil {
		fmt.Println("Error retrieving timestamps in load service: ", err)
		return
	}

	// Filter out bad data and build slice of unique track IDs
	var uniqueTrackIDs []string
	var uniqueTimestamps []int64
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
						SongId:    trackID,
					}
					if !SliceContainsInt64(dbTimestamps, millis) && !SliceContainsInt64(uniqueTimestamps, millis) {
						listensToAdd = append(listensToAdd, listen)
					}
				}
			}
		}
	}
	if listensToAdd == nil {
		// User has already submitted this file
		fmt.Println("No new listensToAdd to add.")
		return
	}

	// The listening history has now been filtered and sorted and the listen beans have been created
	// Because of foreign key constraints, the listensToAdd cannot be added to the DB until song/album metadata is added

	// We only want to get data from Spotify for songs and albums that aren't already in the DB to save time
	var newSongIDs []string
	var dbSongIDs []string
	dbSongs, err := da.RetrieveAllSongs(tx)
	if err != nil {
		fmt.Println("Error retrieving songs in load service: ", err)
		return
	}
	for _, song := range dbSongs {
		dbSongIDs = append(dbSongIDs, song.Id)
	}
	for _, song := range uniqueTrackIDs {
		if !SliceContainsString(dbSongIDs, song) {
			newSongIDs = append(newSongIDs, song)
		}
	}

	// Obtain song metadata for new songs
	songs, err := getAllSongData(token, newSongIDs)
	if err != nil {
		fmt.Println("Error getting song data in load service: ", err)
		return
	}

	// Using the song metadata, obtain album metadata for new albums only
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
	for _, song := range songs {
		if !SliceContainsString(dbAlbumIDs, song.Album) {
			if !SliceContainsString(newUniqueAlbumIDs, song.Album) {
				newUniqueAlbumIDs = append(newUniqueAlbumIDs, song.Album)
			}
		}
	}

	albums, err := getAllAlbumData(token, newUniqueAlbumIDs)
	if err != nil {
		fmt.Println("Error getting album data in load service: ", err)
		return
	}

	// Add song and album metadata to DB
	// Add albums first due to foreign key constraints
	for _, album := range albums {
		if da.CreateAlbum(tx, &album) != nil {
			// There's a good chance the album already exists, so ignore the error
			continue
		}
	}
	for _, song := range songs {
		if da.CreateSong(tx, &song) != nil {
			// There's a good chance the song already exists, so ignore the error
			continue
		}
	}

	// Finally, we add the listensToAdd to the DB
	for _, listen := range listensToAdd {
		exists, err := da.HasTimestamp(tx, listen.Username, listen.Timestamp)
		if err != nil {
			fmt.Println("Error checking if timestamp exists in load service: ", err)
			continue
		}
		if !exists {
			if da.CreateListen(tx, listen) != nil {
				continue // This should never be reached
			}
		}
	}

	if da.CommitAndClose(tx, db, true) != nil {
		fmt.Println("Error committing and closing transaction in load service: ", err)
	}

	fmt.Println(len(listensToAdd), "listens added to the database")
}

func getAllSongData(token string, trackIDs []string) ([]model.SongBean, error) {

	var songs []model.SongBean

	tracks, err := GetSeveralTracks(token, trackIDs)
	if err != nil {
		fmt.Println("Error getting track data in load service: ", err)
		return nil, err
	}

	for _, track := range tracks {
		song := model.SongBean{
			Id:         track.ID,
			Name:       track.Name,
			Artist:     ArtistsToString(track.Artists),
			ArtistId:   ArtistIdsToString(track.Artists),
			Album:      track.Album.ID,
			Explicit:   track.Explicit,
			Popularity: track.Popularity,
			Duration:   track.DurationMs,
		}
		songs = append(songs, song)
	}

	return songs, nil
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
			Image:       album.Images[0].URL,
			Popularity:  album.Popularity,
		}
		albums = append(albums, albumBean)
	}

	return albums, nil
}
