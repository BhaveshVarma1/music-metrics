package model

type UpdateCodeRequest struct {
	Code string `json:"code"`
}

type GetAccessTokenRequest struct {
	GrantType   string `json:"grant_type"`
	Code        string `json:"code"`
	RedirectURL string `json:"redirect_uri"`
}

type ExtendedStreamingObject struct {
	Timestamp        string `json:"ts"`
	Username         string `json:"username"`
	Platform         string `json:"platform"`
	MsPlayed         int    `json:"ms_played"`
	ConnCountry      string `json:"conn_country"`
	IPAddr           string `json:"ip_addr_decrypted"`
	UserAgent        string `json:"user_agent_decrypted"`
	TrackName        string `json:"master_metadata_track_name"`
	ArtistName       string `json:"master_metadata_album_artist_name"`
	AlbumName        string `json:"master_metadata_album_album_name"`
	TrackURI         string `json:"spotify_track_uri"`
	EpisodeName      string `json:"episode_name"`
	EpisodeShowName  string `json:"episode_show_name"`
	EpisodeURI       string `json:"spotify_episode_uri"`
	ReasonStart      string `json:"reason_start"`
	ReasonEnd        string `json:"reason_end"`
	Shuffle          bool   `json:"shuffle"`
	Skipped          bool   `json:"skipped"`
	Offline          bool   `json:"offline"`
	OfflineTimestamp int64  `json:"offline_timestamp"`
	IncognitoMode    bool   `json:"incognito_mode"`
}

type ExtendedStreamingHistory struct {
	StreamingHistory []ExtendedStreamingObject `json:"streaming_history"`
}
