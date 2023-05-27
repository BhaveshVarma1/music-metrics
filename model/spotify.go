package model

type Album struct {
	AlbumType            string       `json:"album_type"`
	TotalTracks          int          `json:"total_tracks"`
	AvailableMarkets     []string     `json:"available_markets"`
	ExternalURLs         ExternalURLs `json:"external_urls"`
	Href                 string       `json:"href"`
	ID                   string       `json:"id"`
	Images               []Image      `json:"images"`
	Name                 string       `json:"name"`
	ReleaseDate          string       `json:"release_date"`
	ReleaseDatePrecision string       `json:"release_date_precision"`
	Restrictions         Restrictions `json:"restrictions"`
	Type                 string       `json:"type"`
	URI                  string       `json:"uri"`
	Copyrights           []Copyright  `json:"copyrights"`
	ExternalIDs          ExternalIDs  `json:"external_ids"`
	Genres               []string     `json:"genres"`
	Label                string       `json:"label"`
	Popularity           int          `json:"popularity"`
	AlbumGroup           string       `json:"album_group"`
	Artists              []Artist     `json:"artists"`
}

type Artist struct {
	ExternalURLs ExternalURLs `json:"external_urls"`
	Followers    Followers    `json:"followers"`
	Genres       []string     `json:"genres"`
	Href         string       `json:"href"`
	ID           string       `json:"id"`
	Images       []Image      `json:"images"`
	Name         string       `json:"name"`
	Popularity   int          `json:"popularity"`
	Type         string       `json:"type"`
	URI          string       `json:"uri"`
}

type Context struct {
	Type         string       `json:"type"`
	Href         string       `json:"href"`
	ExternalURLs ExternalURLs `json:"external_urls"`
	URI          string       `json:"uri"`
}

type Copyright struct {
	Text string `json:"text"`
	Type string `json:"type"`
}

type Cursors struct {
	After  string `json:"after"`
	Before string `json:"before"`
}

type ExplicitContent struct {
	FilterEnabled bool `json:"filter_enabled"`
	FilterLocked  bool `json:"filter_locked"`
}

type ExternalIDs struct {
	Isrc string `json:"isrc"`
	Ean  string `json:"ean"`
	Upc  string `json:"upc"`
}

type ExternalURLs struct {
	Spotify string `json:"spotify"`
}

type Followers struct {
	Href  string `json:"href"`
	Total int    `json:"total"`
}

type GetRecentlyPlayedResponse struct {
	Href    string  `json:"href"`
	Limit   int     `json:"limit"`
	Next    string  `json:"next"`
	Cursors Cursors `json:"cursors"`
	Total   int     `json:"total"`
	Items   []Item  `json:"items"`
}

type Image struct {
	URL    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

type Item struct {
	Track    Track   `json:"track"`
	PlayedAt string  `json:"played_at"`
	Context  Context `json:"context"`
}

type Restrictions struct {
	Reason string `json:"reason"`
}

type Track struct {
	Album            Album        `json:"album"`
	Artists          []Artist     `json:"artists"`
	AvailableMarkets []string     `json:"available_markets"`
	DiscNumber       int          `json:"disc_number"`
	DurationMs       int          `json:"duration_ms"`
	Explicit         bool         `json:"explicit"`
	ExternalIDs      ExternalIDs  `json:"external_ids"`
	ExternalURLs     ExternalURLs `json:"external_urls"`
	Href             string       `json:"href"`
	ID               string       `json:"id"`
	Restrictions     Restrictions `json:"restrictions"`
	Name             string       `json:"name"`
	Popularity       int          `json:"popularity"`
	PreviewURL       string       `json:"preview_url"`
	TrackNumber      int          `json:"track_number"`
	Type             string       `json:"type"`
	URI              string       `json:"uri"`
	IsLocal          bool         `json:"is_local"`
}

type SeveralTracks struct {
	Tracks []Track `json:"tracks"`
}

type SeveralAlbums struct {
	Albums []Album `json:"albums"`
}

type GetAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
	ExpiresIn   int    `json:"expires_in"`
	Refresh     string `json:"refresh_token"`
}

type GetMeResponse struct {
	Country         string          `json:"country"`
	DisplayName     string          `json:"display_name"`
	Email           string          `json:"email"`
	ExplicitContent ExplicitContent `json:"explicit_content"`
	ExternalURLs    ExternalURLs    `json:"external_urls"`
	Followers       Followers       `json:"followers"`
	Href            string          `json:"href"`
	ID              string          `json:"id"`
	Images          []Image         `json:"images"`
	Product         string          `json:"product"`
	Type            string          `json:"type"`
	URI             string          `json:"uri"`
}

type GetRefreshTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
	ExpiresIn   int    `json:"expires_in"`
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
