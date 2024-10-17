package service

import (
	"errors"
	"pccth/portal-blog/internal/entity"
	"pccth/portal-blog/internal/model"
	"pccth/portal-blog/internal/repository"

	"gorm.io/gorm"
)

type PostService struct {
    db *gorm.DB
}

func NewPostService(db *gorm.DB) *PostService {
    return &PostService{db: db}
}


func (s *PostService) CreatePost(createRequest *model.CreatePostRequest) error {
    post := &entity.Post{
        PostHeader: createRequest.PostHeader,
        PostBody:   createRequest.PostBody,
        PostCreateBy: createRequest.PostCreateBy,
    }
    return repository.CreatePost(s.db, post)
}

func (s *PostService) GetPostByID(id uint) (*entity.Post, error) {
    return repository.GetPostByID(s.db, id)
}

func (s *PostService) UpdatePost(id uint, updateRequest *model.UpdatePostRequest) error {
    if updateRequest.PostHeader == "" || updateRequest.PostBody == "" {
        return errors.New("PostHeader and PostBody cannot be empty")
    }

    post, err := repository.GetPostByID(s.db, id)
    if err != nil {
        return err
    }


    post.PostHeader = updateRequest.PostHeader
    post.PostBody = updateRequest.PostBody

    return s.db.Save(post).Error
}

func (s *PostService) DeletePost(id uint) error {
    tx := s.db.Begin()

    var post entity.Post
    if err := tx.Where("id = ?", id).First(&post).Error; err != nil {
        tx.Rollback()
        return errors.New("post not found")
    }

    if err := tx.Where("post_id = ?", id).Delete(&entity.Comment{}).Error; err != nil {
        tx.Rollback()
        return errors.New("failed to delete comments")
    }

    if err := tx.Delete(&post).Error; err != nil {
        tx.Rollback()
        return errors.New("failed to delete post")
    }

    return tx.Commit().Error
}


func (s *PostService) GetAllPosts() ([]entity.Post, error) {
    return repository.GetAllPosts(s.db)
}

func (s *PostService) GetPaginatedPosts(page, limit int) (model.PaginatedResponse, error) {
	var posts []entity.Post
	var totalPosts int64

	offset := (page - 1) * limit

	s.db.Model(&entity.Post{}).Count(&totalPosts)

	result := s.db.Limit(limit).Offset(offset).Find(&posts)
	if result.Error != nil {
		return model.PaginatedResponse{}, result.Error
	}

	response := model.PaginatedResponse{
		Data:       posts,
		TotalCount: int(totalPosts),
		Page:       page,
		PageSize:   limit,
	}

	return response, nil
}
