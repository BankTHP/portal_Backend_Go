package model

type CreateCommentRequest struct {
	PostID          uint   `json:"postID" validate:"required"`
	CommentBody     string `json:"CommentBody" validate:"required"`
	CommentCreateBy string `json:"CommentCreateBy" validate:"required"`
}

type UpdateCommentRequest struct {
	CommentBody string `json:"commentBody" validate:"required"`
}
