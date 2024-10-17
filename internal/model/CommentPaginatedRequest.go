package model

type CommentPaginatedRequest struct {
	PostID uint `json:"postID"`
	Page   uint `json:"page"`
	Size   uint `json:"size"`
}

type CommentByUserIdPaginatedRequest struct {
	CommentCreateBy string `json:"commentCreateBy"`
	Page            uint   `json:"page"`
	Size            uint   `json:"size"`
}
