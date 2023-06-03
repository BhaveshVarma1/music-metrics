package model

type AlbumBean struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Artist      string `json:"artist"`
	ArtistId    string `json:"artistId"`
	Genre       string `json:"genre"`
	TotalTracks int    `json:"totalTracks"`
	Year        int    `json:"year"`
	Image       string `json:"image"`
	Popularity  int    `json:"popularity"`
}

type AuthTokenBean struct {
	Token      string `json:"token"`
	Username   string `json:"username"`
	Expiration string `json:"expiration,omitempty"`
}

type ListenBean struct {
	Username  string `json:"username"`
	Timestamp int64  `json:"timestamp"`
	SongId    string `json:"songID"`
}

type SongBean struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Artist     string `json:"artist"`
	ArtistId   string `json:"artistId"`
	Album      string `json:"album"`
	Explicit   bool   `json:"explicit"`
	Popularity int    `json:"popularity"`
	Duration   int    `json:"duration"`
}

type UserBean struct {
	Username    string `json:"username"`
	DisplayName string `json:"displayName"`
	Email       string `json:"email"`
	Refresh     string `json:"refresh,omitempty"`
	Timestamp   int64  `json:"timestamp,omitempty"`
}

type LogBean struct {
	Username  string `json:"username"`
	Timestamp int64  `json:"timestamp"`
	Action    string `json:"action"`
	IP        string `json:"ip"`
}
