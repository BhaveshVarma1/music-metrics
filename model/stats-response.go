package model

type StatsResponse interface {
}

type AverageYearResponse struct {
	Message     string `json:"message,omitempty"`
	Success     bool   `json:"success"`
	AverageYear int    `json:"averageYear,omitempty"`
}

type TopSongsResponse struct {
	Message  string    `json:"message,omitempty"`
	Success  bool      `json:"success"`
	TopSongs []TopSong `json:"topSongs,omitempty"`
}

type TopAlbumsResponse struct {
	Message   string     `json:"message,omitempty"`
	Success   bool       `json:"success"`
	TopAlbums []TopAlbum `json:"topAlbums,omitempty"`
}

type DecadeBreakdownResponse struct {
	Message          string            `json:"message,omitempty"`
	Success          bool              `json:"success"`
	DecadeBreakdowns []DecadeBreakdown `json:"decadeBreakdown,omitempty"`
}

type TopArtistsResponse struct {
	Message    string      `json:"message,omitempty"`
	Success    bool        `json:"success"`
	TopArtists []TopArtist `json:"topArtists,omitempty"`
}
