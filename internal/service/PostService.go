﻿package service

import (
	"errors"
	"fmt"
	"mime/multipart"
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

func (s *PostService) CreatePost(createRequest *model.CreatePostRequest, files []*multipart.FileHeader) error {
	post := &entity.Post{
		PostHeader:   createRequest.PostHeader,
		PostBody:     createRequest.PostBody,
		PostCreateBy: createRequest.PostCreateBy,
	}

	if err := repository.CreatePost(s.db, post); err != nil {
		return err
	}

	if len(files) == 0 {
		return nil
	}

	for _, file := range files {
		if err := s.SavePDF(post.ID, file); err != nil {
			return err
		}
	}

	return nil
}



func (s *PostService) GetPostByID(id uint) (*model.PostResponse, error) {
	post, err := repository.GetPostByID(s.db, id)
	if err != nil {
		return nil, err
	}

	post.Views++

	user, err := repository.GetUserByUserId(s.db, post.PostCreateBy)
	if err != nil {
		return nil, err
	}

	Res := &model.PostResponse{
		ID:             post.ID,
		PostHeader:     post.PostHeader,
		PostBody:       post.PostBody,
		PostCreateBy:   user.Username,
		PostCreateDate: post.PostCreateDate,
		Views:          post.Views,
	}

	if err := s.db.Model(&entity.Post{}).Where("id = ?", id).Update("views", post.Views).Error; err != nil {
		return nil, fmt.Errorf("ไม่สามารถอัพเดทจำนวนการดูได้: %v", err)
	}

	pdfs, err := repository.GetPDFsByPostID(s.db, id)
	if err != nil {
		return nil, err
	}

	Res.PDFs = pdfs

	return Res, nil
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
	var post entity.Post

	if err := s.db.Where("id = ?", id).First(&post).Error; err != nil {
		return errors.New("post not found")
	}

	if err := s.db.Where("post_id = ?", id).Delete(&entity.Comment{}).Error; err != nil {
		return errors.New("failed to delete comments")
	}

	if err := s.db.Where("post_id = ?", id).Delete(&entity.PDFs{}).Error; err != nil {
		return errors.New("failed to delete pdfs")
	}

	if err := s.db.Delete(&post).Error; err != nil {
		return errors.New("failed to delete post")
	}

	return nil
}

func (s *PostService) GetAllPosts() ([]entity.Post, error) {
	posts, err := repository.GetAllPosts(s.db)
	if err != nil {
		return nil, err
	}

	for i := range posts {
		user, err := repository.GetUserByUserId(s.db, posts[i].PostCreateBy)
		if err != nil {
			return nil, err
		}
		posts[i].PostCreateBy = user.Username
	}

	return posts, nil
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

	for i := range posts {
		user, err := repository.GetUserByUserId(s.db, posts[i].PostCreateBy)
		if err != nil {
			return model.PaginatedResponse{}, err
		}
		posts[i].PostCreateBy = user.Username
	}

	response := model.PaginatedResponse{
		Data:       posts,
		TotalCount: int(totalPosts),
		Page:       page,
		PageSize:   limit,
	}

	return response, nil
}

func (s *PostService) GetPaginatedPostsByUserId(page, limit int, PostCreateBy string) (model.PaginatedResponse, error) {
	var posts []entity.Post
	var totalPost int64

	offset := (page - 1) * limit

	s.db.Model(&entity.Post{}).Where("post_create_by = ?", PostCreateBy).Count(&totalPost)

	totalPages := int(totalPost) / limit
	if int(totalPost)%limit != 0 {
		totalPages++
	}

	result := s.db.Where("post_create_by = ?", PostCreateBy).Limit(limit).Offset(offset).Find(&posts)
	if result.Error != nil {
		return model.PaginatedResponse{}, result.Error
	}

	for i := range posts {
		user, err := repository.GetUserByUserId(s.db, posts[i].PostCreateBy)
		if err != nil {
			return model.PaginatedResponse{}, err
		}
		posts[i].PostCreateBy = user.Username
	}

	response := model.PaginatedResponse{
		Data:       posts,
		TotalCount: int(totalPost),
		TotalPages: totalPages,
		Page:       page,
		PageSize:   limit,
	}

	return response, nil
}
