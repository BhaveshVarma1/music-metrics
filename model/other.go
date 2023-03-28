package model

type SongCount struct {
	Song   string `json:"song"`
	Artist string `json:"artist"`
	Count  int    `json:"count"`
}

type RecentlyPlayedObject struct {
	Song      Song
	Timestamp int64
}
