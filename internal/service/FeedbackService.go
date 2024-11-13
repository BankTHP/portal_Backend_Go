package service

import (
	"fmt"
	"pccth/portal-blog/internal/entity"
	"pccth/portal-blog/internal/model"
	"pccth/portal-blog/internal/repository"
	"time"

	"gorm.io/gorm"
)

type FeedbackService struct {
	db *gorm.DB
}

func NewFeedbackService(db *gorm.DB) *FeedbackService {
	return &FeedbackService{db: db}
}

func (s *FeedbackService) CreateFeedback(createRequest *model.CreateFeedbackRequest) error {
	if len(createRequest.FeedBackPhone) > 10 {
		return fmt.Errorf("เบอร์โทรศัพท์ต้องไม่เกิน 10 หลัก")
	}

	feedback := &entity.Feedback{
		FeedBackName:  createRequest.FeedBackName,
		FeedBackEmail: createRequest.FeedBackEmail,
		FeedBackPhone: createRequest.FeedBackPhone,
		FeedBackText:  createRequest.FeedBackText,
		FeedBackDate:  time.Now(),
	}
	return repository.CreateFeedback(s.db, feedback)
}

func (s *FeedbackService) GetFeedbackById(id uint) (*entity.Feedback, error) {
	return repository.GetFeedbackById(s.db, id)
}

func (s *FeedbackService) DeleteFeedback(id uint) error {
	return repository.DeleteFeedback(s.db, id)
}

func (s *FeedbackService) GetPaginatedFeedbacks(page, limit int) (model.PaginatedResponse, error) {
	feedbacks, totalFeedbacks, err := repository.GetPaginatedFeedbacks(s.db, page, limit)
	if err != nil {
		return model.PaginatedResponse{}, err
	}

	totalPages := int(totalFeedbacks) / limit
	if int(totalFeedbacks)%limit != 0 {
		totalPages++
	}

	response := model.PaginatedResponse{
		Data:       feedbacks,
		TotalCount: int(totalFeedbacks),
		TotalPages: totalPages,
		Page:       page,
		PageSize:   limit,
	}

	return response, nil
}

func (s *FeedbackService) GetAllFeedbacks() ([]entity.Feedback, error) {
	feedbacks, err := repository.GetAllFeedbacks(s.db)
	if err != nil {
		return nil, fmt.Errorf("ไม่สามารถดึงข้อมูล Feedback ได้: %v", err)
	}
	return feedbacks, nil
} 