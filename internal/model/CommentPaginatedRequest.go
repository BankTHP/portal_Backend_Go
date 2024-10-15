package model

type CommentPaginatedRequest struct {
	PostID uint `json:"postID"`
	Page   uint `json:"page"`
	Limit  uint `json:"limit"`
}
