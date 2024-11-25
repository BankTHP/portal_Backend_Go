package model

import "mime/multipart"

type PDFResponse struct {
    Success  bool   `json:"success"`
    Error    string `json:"error,omitempty"`
    FullURL  string `json:"fullUrl,omitempty"`
    Size     string `json:"size,omitempty"`
}

type UploadedPDFFile struct {
    File     *multipart.FileHeader
    Filename string
    Size     int64
    SaveFunc func(string) error
} 