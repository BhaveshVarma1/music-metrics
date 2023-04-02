package model

type AverageYearResponse struct {
	Message     string `json:"message,omitempty"`
	Success     bool   `json:"success"`
	AverageYear int    `json:"averageYear,omitempty"`
}

type GenericResponse struct {
	Message string `json:"message,omitempty"`
	Success bool   `json:"success"`
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

type UpdateCodeResponse struct {
	Message     string `json:"message,omitempty"`
	Success     bool   `json:"success"`
	Token       string `json:"token,omitempty"`
	Username    string `json:"username,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	Email       string `json:"email,omitempty"`
	Timestamp   int64  `json:"timestamp,omitempty"`
}
