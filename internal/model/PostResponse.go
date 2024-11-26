package model

import (
	"pccth/portal-blog/internal/entity"
	"time"
)

type PostResponse struct {
	ID             uint          `json:"id"`
	PostHeader     string        `json:"postHeader"`
	PostBody       string        `json:"postBody"`
	PostCreateBy   string        `json:"postCreateBy"`
	PostCreateDate time.Time     `json:"postCreateDate"`
	Views          uint          `json:"views"`
	PDFs           []entity.PDFs `json:"pdfs"`
}
