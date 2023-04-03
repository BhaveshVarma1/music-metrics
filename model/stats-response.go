package model

type StatsResponse interface {
}

type AverageYearResponse struct {
	Message     string `json:"message,omitempty"`
	Success     bool   `json:"success"`
	AverageYear int    `json:"averageYear,omitempty"`
}

type SongCountsResponse struct {
	Message    string      `json:"message,omitempty"`
	Success    bool        `json:"success"`
	SongCounts []SongCount `json:"songCounts,omitempty"`
}

type TopAlbumsResponse struct {
	Message   string     `json:"message,omitempty"`
	Success   bool       `json:"success"`
	TopAlbums []TopAlbum `json:"topAlbums,omitempty"`
}

type DecadeBreakdownResponse struct {
	Message          string            `json:"message,omitempty"`
	Success          bool              `json:"success"`
	DecadeBreakdowns []DecadeBreakdown `json:"decadeBreakdowns,omitempty"`
}
