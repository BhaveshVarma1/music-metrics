package service

import (
	"fmt"
	"music-metrics/da"
	"music-metrics/model"
	"strconv"
	"strings"
)

func Load(history []model.ExtendedStreamingObject, username string) {

	// Instantiate DB connection
	tx, db, err := da.BeginTX()
	if err != nil {
		fmt.Println("Error beginning transaction in load service: ", err)
		return
	}

	// Get authtoken and timestamp of account creation
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

	// Filter out bad data and build slice of unique track IDs
	var filteredHistory []model.ExtendedStreamingObject
	var uniqueTrackIDs []string
	var listens []model.ListenBean
	var timestamps []int64
	for _, item := range history {
		if item.TrackName != "" && item.MsPlayed > 29999 {
			millis := DatetimeToUnixMilli(item.Timestamp)
			if millis != -1 {
				item.Timestamp = strconv.FormatInt(millis, 10)
				if millis < endTime {
					if !SliceContainsInt64(timestamps, millis) {
						// At this point, item is a unique listen of at least 30 seconds that occurred before the account was created
						timestamps = append(timestamps, millis)
						filteredHistory = append(filteredHistory, item)
						trackID := item.TrackURI[strings.LastIndex(item.TrackURI, ":")+1:]
						if !SliceContainsString(uniqueTrackIDs, trackID) {
							uniqueTrackIDs = append(uniqueTrackIDs, trackID)
						}
						// Create listen object
						listen := model.ListenBean{
							Username:  username,
							Timestamp: millis,
							SongId:    trackID,
						}
						listens = append(listens, listen)
					}
				}
			}
		}
	}

	// The listening history has now been filtered and sorted and the listen beans have been created
	// Because of foreign key constraints, the listens cannot be added to the DB until song/album metadata is added

	// Obtain song and album metadata
	songs, err := getAllSongData(token, uniqueTrackIDs)
	if err != nil {
		fmt.Println("Error getting song data in load service: ", err)
		return
	}
	var uniqueAlbumIDs []string
	for _, song := range songs {
		if !SliceContainsString(uniqueAlbumIDs, song.Album) {
			uniqueAlbumIDs = append(uniqueAlbumIDs, song.Album)
		}
	}
	albums, err := getAllAlbumData(token, uniqueAlbumIDs)
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

	// Finally, we add the listens to the DB
	fmt.Println("Adding listens to DB: ", len(listens))
	for _, listen := range listens {
		if da.CreateListen(tx, listen) != nil {
			// Listen could already exist if the user already submitted this history
			continue
		}
	}

	if da.CommitAndClose(tx, db, true) != nil {
		fmt.Println("Error committing and closing transaction in load service: ", err)
	}
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
