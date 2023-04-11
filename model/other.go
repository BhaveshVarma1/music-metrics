package model

type DecadeBreakdown struct {
	Decade string `json:"decade"`
	Count  int    `json:"count"`
}

type ModeYear struct {
	Year  int
	Count int
}

type PopularityObject struct {
	Song       string `json:"song"`
	Artist     string `json:"artist"`
	Popularity int    `json:"popularity"`
}

type RawArtistTime struct {
	Artist string
	Millis int
}

type RecentlyPlayedObject struct {
	Song      SongBean
	Album     AlbumBean
	Timestamp int64
}

type TopSong struct {
	Song   string `json:"song"`
	Artist string `json:"artist"`
	Count  int    `json:"count"`
}

type TopAlbum struct {
	Album  string `json:"album"`
	Artist string `json:"artist"`
	Image  string `json:"image"`
	Count  int    `json:"count"`
}

type TopArtist struct {
	Artist string `json:"artist"`
	Count  int    `json:"count"`
}
