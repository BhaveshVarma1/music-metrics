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
	SongId     string `json:"songId"`
	ArtistId   string `json:"artistId"`
}

type RawArtistTime struct {
	Artist   string
	Millis   int
	ArtistId string
}

type RecentlyPlayedObject struct {
	Song      SongBean
	Album     AlbumBean
	Timestamp int64
}

type TopSong struct {
	Song     string `json:"song"`
	Artist   string `json:"artist"`
	Count    int    `json:"count"`
	SongId   string `json:"songId"`
	ArtistId string `json:"artistId"`
	Image    string `json:"image"`
}

type TopAlbum struct {
	Album    string `json:"album"`
	Artist   string `json:"artist"`
	Image    string `json:"image"`
	Count    int    `json:"count"`
	AlbumId  string `json:"albumId"`
	ArtistId string `json:"artistId"`
}

type TopArtist struct {
	Artist   string `json:"artist"`
	Count    int    `json:"count"`
	ArtistId string `json:"artistId"`
}
