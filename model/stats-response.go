package model

type StatsResponse interface {
}

type AveragePopularityResponse struct {
	Items []PopularityObject `json:"items,omitempty"`
}

type DecadeBreakdownResponse struct {
	Items []DecadeBreakdown `json:"items,omitempty"`
}

type HourBreakdownResponse struct {
	Items []int `json:"items,omitempty"`
}

type ModeYearResponse struct {
	Items []ModeYear `json:"items,omitempty"`
}

type SingleIntResponse struct {
	Value int `json:"value,omitempty"`
}

type TopAlbumsResponse struct {
	Items []TopAlbum `json:"items,omitempty"`
}

type TopArtistsResponse struct {
	Items []TopArtist `json:"items,omitempty"`
}

type TopSongsResponse struct {
	Items []TopSong `json:"items,omitempty"`
}

type WeekDayBreakdownResponse struct {
	Items []int `json:"items,omitempty"`
}
