package model

import "mime/multipart"

type VideoResponse struct {
	Success  bool   `json:"success"`
	Error    string `json:"error,omitempty"`
	FullURL  string `json:"fullUrl,omitempty"`
	Duration string `json:"duration,omitempty"`
	Size     string `json:"size,omitempty"`
}

type CreateVideoRequest struct {
	VdoId       string `json:"vdo_id"`
	VdoName     string `json:"vdo_name"`
	VdoSize     string `json:"vdo_size"`
	VdoDuration string `json:"vdo_duration"`
}

type UploadedFile struct {
	File     *multipart.FileHeader
	Filename string
	Size     int64
	SaveFunc func(string) error
}
