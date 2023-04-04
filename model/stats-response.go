package model

type StatsResponse interface {
}

type AverageYearResponse struct {
	Message     string `json:"message,omitempty"`
	Success     bool   `json:"success"`
	AverageYear int    `json:"averageYear,omitempty"`
}

type TopSongsResponse struct {
	Message string    `json:"message,omitempty"`
	Success bool      `json:"success"`
	Items   []TopSong `json:"items,omitempty"`
}

type TopAlbumsResponse struct {
	Message string     `json:"message,omitempty"`
	Success bool       `json:"success"`
	Items   []TopAlbum `json:"items,omitempty"`
}

type DecadeBreakdownResponse struct {
	Message string            `json:"message,omitempty"`
	Success bool              `json:"success"`
	Items   []DecadeBreakdown `json:"items,omitempty"`
}

type TopArtistsResponse struct {
	Message string      `json:"message,omitempty"`
	Success bool        `json:"success"`
	Items   []TopArtist `json:"items,omitempty"`
}
