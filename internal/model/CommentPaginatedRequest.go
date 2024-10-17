package model

type CommentPaginatedRequest struct {
	PostID uint `json:"postID"`
	Page   uint `json:"page"`
	Size   uint `json:"size"`
}

type CommentByUserIdPaginatedRequest struct {
	UserId uint `json:"userID"`
	Page   uint `json:"page"`
	Size   uint `json:"size"`
}
