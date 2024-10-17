package model

type CreatePostRequest struct {
	PostHeader   string `json:"postHeader" validate:"required"`
	PostBody     string `json:"postBody" validate:"required"`
	PostCreateBy string `json:"postCreateBy" validate:"required"`
}

type UpdatePostRequest struct {
	PostHeader string `json:"postHeader" validate:"required"`
	PostBody   string `json:"postBody" validate:"required"`
}

type PostPaginatedRequest struct {
	Page uint `json:"page"`
	Size uint `json:"size"`
}

type PostByUserIdPaginatedRequest struct {
	PostCreateBy string `json:"postCreateBy"`
	Page         uint   `json:"page"`
	Size         uint   `json:"size"`
}
