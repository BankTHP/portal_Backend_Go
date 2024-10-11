package repository

import (
	"fmt"
	"pccth/portal-blog/internal/entity"

	"gorm.io/gorm"
)

func CreateComment(db *gorm.DB, comment *entity.Comment) error {
	result := db.Create(comment)
	if result.Error != nil {
		return result.Error
	}
	fmt.Println("Comment created successfully! ID:", comment.ID)
	return nil
}

func GetCommentByID(db *gorm.DB, id uint) (*entity.Comment, error) {
	var comment entity.Comment
	if err := db.First(&comment, id).Error; err != nil {
		return nil, err
	}
	return &comment, nil
}

func GetCommentByPostID(db *gorm.DB, postID uint) ([]entity.Comment, error) {
	var comments []entity.Comment
	if err := db.Where("post_id = ?", postID).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

func DeleteComment(db *gorm.DB, id uint) error {
	return db.Delete(&entity.Comment{}, id).Error
}
