package model

type AnalyticsResponse struct {
	TotalUsers int    `json:"totalUsers"`
	LastAction string `json:"lastAction"`
}

type GenericResponse struct {
	Message string `json:"message,omitempty"`
	Success bool   `json:"success"`
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
