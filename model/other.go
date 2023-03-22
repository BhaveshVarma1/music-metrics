package model

type SongCount struct {
	Song   string
	Artist string
	Count  int
}

type RecentlyPlayedObject struct {
	Song      Song
	Timestamp int64
}
