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
    return repository.DeletePost(s.db, id)
}

func (s *PostService) GetAllPosts() ([]entity.Post, error) {
    return repository.GetAllPosts(s.db)
}
