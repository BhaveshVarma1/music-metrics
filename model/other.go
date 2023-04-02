package model

type RecentlyPlayedObject struct {
	Song      SongBean
	Album     AlbumBean
	Timestamp int64
}

type SongCount struct {
	Song   string `json:"song"`
	Artist string `json:"artist"`
	Count  int    `json:"count"`
}
