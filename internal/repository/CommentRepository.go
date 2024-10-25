package repository

import (
	"errors"
	"pccth/portal-blog/internal/entity"

	"gorm.io/gorm"
)

func CreateComment(db *gorm.DB, comment *entity.Comment) error {
	result := db.Create(comment)
	if result.Error != nil {
		return result.Error
	}
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

func UpdateComment(db *gorm.DB, comment *entity.Comment) error {
	result := db.Model(comment).Updates(map[string]interface{}{
		"comment_body": comment.CommentBody,
	})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("no comment found with the given ID or no changes made")
	}

	return nil
}

func DeleteComment(db *gorm.DB, id uint) error {
	return db.Delete(&entity.Comment{}, id).Error
}

func DeleteCommentByPostId(db *gorm.DB, postID uint) error {
	result := db.Where("post_id = ?", postID).Delete(&entity.Comment{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no comments found for the given post ID")
	}
	return nil
}
