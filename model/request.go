package model

type UpdateCodeRequest struct {
	Code string `json:"code"`
}

type GetAccessTokenRequest struct {
	GrantType   string `json:"grant_type"`
	Code        string `json:"code"`
	RedirectURL string `json:"redirect_uri"`
}

/*type GetRecentlyPlayedRequest struct {
	Before int `json:"before"`
	Limit  int `json:"limit"`
}*/
