package model

type CreateFeedbackRequest struct {
	FeedBackName  string `json:"feedBackName" validate:"required"`
	FeedBackEmail string `json:"feedBackEmail" validate:"required,email"`
	FeedBackPhone string `json:"feedBackPhone" validate:"required,max=10"`
	FeedBackText  string `json:"feedBackText" validate:"required"`
}

type FeedbackPaginatedRequest struct {
	Page uint `json:"page"`
	Size uint `json:"size"`
} 