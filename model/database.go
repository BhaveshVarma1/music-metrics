package model

type User struct {
	Username    string `json:"username"`
	DisplayName string `json:"displayName"`
	Email       string `json:"email"`
	Refresh     string `json:"refresh,omitempty"`
	Access      string `json:"access,omitempty"`
}

type AuthToken struct {
	Token      string `json:"token"`
	Username   string `json:"username"`
	Expiration string `json:"expiration,omitempty"`
}

type Song struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Artist     string `json:"artist"`
	Album      string `json:"album"`
	Explicit   bool   `json:"explicit"`
	Popularity int    `json:"popularity"`
	Duration   int    `json:"duration"`
	Year       int    `json:"year"`
}

type Listen struct {
	Username  string `json:"username"`
	Timestamp string `json:"timestamp"`
	SongId    string `json:"songID"`
}
