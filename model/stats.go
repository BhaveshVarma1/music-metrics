package model

type StatsResponse interface {
}

type AllStatsResponse struct {
	AverageLength     StatsResponse `json:"averageLength,omitempty"`
	AveragePopularity StatsResponse `json:"averagePopularity,omitempty"`
	AverageYear       StatsResponse `json:"averageYear,omitempty"`
	DecadeBreakdown   StatsResponse `json:"decadeBreakdown,omitempty"`
	HourBreakdown     StatsResponse `json:"hourBreakdown,omitempty"`
	MedianYear        StatsResponse `json:"medianYear,omitempty"`
	ModeYear          StatsResponse `json:"modeYear,omitempty"`
	PercentExplicit   StatsResponse `json:"percentExplicit,omitempty"`
	TopAlbums         StatsResponse `json:"topAlbums,omitempty"`
	TopAlbumsTime     StatsResponse `json:"topAlbumsTime,omitempty"`
	TopArtists        StatsResponse `json:"topArtists,omitempty"`
	TopArtistsTime    StatsResponse `json:"topArtistsTime,omitempty"`
	TopTracks         StatsResponse `json:"topTracks,omitempty"`
	TopTracksTime     StatsResponse `json:"topTracksTime,omitempty"`
	TotalMinutes      StatsResponse `json:"totalMinutes,omitempty"`
	TotalTracks       StatsResponse `json:"totalTracks,omitempty"`
	UniqueAlbums      StatsResponse `json:"uniqueAlbums,omitempty"`
	UniqueArtists     StatsResponse `json:"uniqueArtists,omitempty"`
	UniqueTracks      StatsResponse `json:"uniqueTracks,omitempty"`
	WeekDayBreakdown  StatsResponse `json:"weekDayBreakdown,omitempty"`
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

type TopTracksResponse struct {
	Items []TopTrack `json:"items,omitempty"`
}

type WeekDayBreakdownResponse struct {
	Items []int `json:"items,omitempty"`
}
