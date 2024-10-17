package service

import (
	"errors"
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
		db: db,
	}
}

func (s *ReleaseNoteService) CreateRelease(createRequest *model.CreateReleaseNoteRequest) error {
	release := &entity.Release{
		Header: createRequest.ReleaseNoteHeader,
		Body:   createRequest.ReleaseNoteBody,
	}
	return repository.CreateRelease(s.db, release)
}

func (s *ReleaseNoteService) GetReleaseByID(id uint) (*entity.Release, error) {
	return repository.GetReleaseByID(s.db, id)
}

func (s *ReleaseNoteService) UpdateRelease(id uint, updateRequest *model.UpdateReleaseNoteRequest) error {
	if updateRequest.ReleaseNoteHeader == "" || updateRequest.ReleaseNoteBody == "" {
		return errors.New("ReleaseHeader and ReleaseBody cannot be empty")
	}

	release, err := repository.GetReleaseByID(s.db, id)
	if err != nil {
		return err
	}

	release.Header = updateRequest.ReleaseNoteHeader
	release.Body = updateRequest.ReleaseNoteBody

	return s.db.Save(release).Error
}

func (s *ReleaseNoteService) DeleteRelease(id uint) error {
	return repository.DeleteRelease(s.db, id)
}

func (s *ReleaseNoteService) GetAllRelease() ([]entity.Release, error) {
	return repository.GetAllRelease(s.db)
}
