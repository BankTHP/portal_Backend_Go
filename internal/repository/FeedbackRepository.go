package repository

import (
	"pccth/portal-blog/internal/entity"

	"gorm.io/gorm"
)

func CreateFeedback(db *gorm.DB, feedback *entity.Feedback) error {
	return db.Create(feedback).Error
}

func GetFeedbackById(db *gorm.DB, id uint) (*entity.Feedback, error) {
	var feedback entity.Feedback
	if err := db.First(&feedback, id).Error; err != nil {
		return nil, err
	}
	return &feedback, nil
}

func DeleteFeedback(db *gorm.DB, id uint) error {
	return db.Delete(&entity.Feedback{}, id).Error
}

func GetAllFeedbacks(db *gorm.DB) ([]entity.Feedback, error) {
	var feedbacks []entity.Feedback
	if err := db.Find(&feedbacks).Error; err != nil {
		return nil, err
	}
	return feedbacks, nil
}

func GetPaginatedFeedbacks(db *gorm.DB, page, limit int) ([]entity.Feedback, int64, error) {
	var feedbacks []entity.Feedback
	var totalFeedbacks int64

	if err := db.Model(&entity.Feedback{}).Count(&totalFeedbacks).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit

	if err := db.Limit(limit).Offset(offset).Find(&feedbacks).Error; err != nil {
		return nil, 0, err
	}

	return feedbacks, totalFeedbacks, nil
} 