package model

import "pccth/portal-blog/internal/entity"

type PostResponse struct {
	ID           uint          `json:"id"`
	PostHeader   string        `json:"postHeader"`
	PostBody     string        `json:"postBody"`
	PostCreateBy string        `json:"postCreateBy"`
	Views        uint          `json:"views"`
	PDFs         []entity.PDFs `json:"pdfs"`
}
