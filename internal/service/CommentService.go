package service

import (
	"errors"
	"pccth/portal-blog/internal/entity"
	"pccth/portal-blog/internal/model"
	"pccth/portal-blog/internal/repository"

	"gorm.io/gorm"
)

type CommentService struct {
	db *gorm.DB
}

func NewCommentService(db *gorm.DB) *CommentService {
	return &CommentService{db: db}
}

func (s *CommentService) CreateComment(createRequest *model.CreateCommentRequest) error {
	_, err := repository.GetPostByID(s.db, createRequest.PostID)
	if err != nil {
		return errors.New("post not found")
	}

	comment := &entity.Comment{
		PostID:          createRequest.PostID,
		CommentBody:     createRequest.CommentBody,
		CommentCreateBy: createRequest.CommentCreateBy,
	}
	return repository.CreateComment(s.db, comment)
}

func (s *CommentService) GetCommentByID(id uint) (*entity.Comment, error) {
	return repository.GetCommentByID(s.db, id)
}

func (s *CommentService) GetCommentByPostID(id uint) ([]entity.Comment, error) {
	return repository.GetCommentByPostID(s.db, id)
}

func (s *CommentService) DeleteComment(id uint) error {
	return repository.DeleteComment(s.db, id)
}

func (s *CommentService) GetPaginatedComments(page, limit, postId int) (model.PaginatedResponse, error) {
	var comments []entity.Comment
	var totalComment int64

	offset := (page - 1) * limit

	s.db.Model(&entity.Comment{}).Where("post_id = ?", postId).Count(&totalComment)

	totalPages := int(totalComment) / limit
	if int(totalComment)%limit != 0 {
		totalPages++
	}

	result := s.db.Where("post_id = ?", postId).Limit(limit).Offset(offset).Find(&comments)
	if result.Error != nil {
		return model.PaginatedResponse{}, result.Error
	}

	response := model.PaginatedResponse{
		Data:       comments,
		TotalCount: int(totalComment),
		TotalPages: totalPages,
		Page:       page,
		PageSize:   limit,
	}

	return response, nil
}

func (s *CommentService) UpdateComment(id uint, updateRequest *model.UpdateCommentRequest) error {
	if updateRequest.CommentBody == "" {
		return errors.New("CommentBody cannot be empty")
	}

	comment, err := repository.GetCommentByID(s.db, id)
	if err != nil {
		return err
	}

	comment.CommentBody = updateRequest.CommentBody

	if err := repository.UpdateComment(s.db, comment); err != nil {
		return err
	}

	return nil
}

func (s *CommentService) GetPaginatedCommentsByUserId(page, limit int, CommentCreateBy string) (model.PaginatedResponse, error) {
	var comments []entity.Comment
	var totalComment int64

	offset := (page - 1) * limit

	s.db.Model(&entity.Comment{}).Where("comment_create_by = ?", CommentCreateBy).Count(&totalComment)

	totalPages := int(totalComment) / limit
	if int(totalComment)%limit != 0 {
		totalPages++
	}

	result := s.db.Where("comment_create_by = ?", CommentCreateBy).Limit(limit).Offset(offset).Find(&comments)
	if result.Error != nil {
		return model.PaginatedResponse{}, result.Error
	}

	response := model.PaginatedResponse{
		Data:       comments,
		TotalCount: int(totalComment),
		TotalPages: totalPages,
		Page:       page,
		PageSize:   limit,
	}

	return response, nil
}
