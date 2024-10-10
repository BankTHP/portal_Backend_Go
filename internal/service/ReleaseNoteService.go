package service

import (
	"pccth/portal-blog/internal/entity"
	"pccth/portal-blog/internal/model"
	"pccth/portal-blog/internal/repository"

	"gorm.io/gorm"
)

type ReleaseNoteService struct {
    db *gorm.DB
}

func NewReleaseNoteService(db *gorm.DB) *ReleaseNoteService {
    return &ReleaseNoteService{
        db:          db,
    }
}

func (s *ReleaseNoteService) CreateRelease(createRequest *model.CreateReleaseNoteRequest) error {
    release := &entity.Release{
        Header: createRequest.ReleaseNoteHeader,
        Body:   createRequest.ReleaseNoteBody,
    }
    return repository.CreateRelease(s.db, release)
}