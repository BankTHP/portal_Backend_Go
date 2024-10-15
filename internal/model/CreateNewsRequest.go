package model

type CreateNewsRequest struct {
	NewsHeader string `json:"newsHeader" validate:"required"`
	NewsBody   string `json:"newsBody" validate:"required"`
}

type UpdateNewsRequest struct {
	NewsHeader string `json:"newsHeader" validate:"required"`
	NewsBody   string `json:"newsBody" validate:"required"`
}
