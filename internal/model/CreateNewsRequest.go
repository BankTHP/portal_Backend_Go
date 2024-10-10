package model

type CreateNewsRequest struct {
	NewsHeader string `json:"postHeader" validate:"required"`
	NewsBody   string `json:"postBody" validate:"required"`
}

type UpdateNewsRequest struct {
	NewsHeader string `json:"postHeader" validate:"required"`
	NewsBody   string `json:"postBody" validate:"required"`
}
