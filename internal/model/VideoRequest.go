package model

type VideoResponse struct {
	Success  bool   `json:"success"`
	Error    string `json:"error,omitempty"`
	FullURL  string `json:"fullUrl,omitempty"`
	Duration string `json:"duration,omitempty"`
	Size     string `json:"size,omitempty"`
}
