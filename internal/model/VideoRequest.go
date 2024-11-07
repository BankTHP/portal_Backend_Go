package model

type VideoResponse struct {
	Success  bool   `json:"success"`
	FullURL  string `json:"fullURL"`
	Error    string `json:"error,omitempty"`
}