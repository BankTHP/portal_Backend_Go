package repository

import (
	"errors"
	"pccth/portal-blog/internal/entity"

	"gorm.io/gorm"
)

func CreateVideo(db *gorm.DB, video *entity.Videos) error {
	return db.Create(video).Error
}

func GetVideoByVdoId(db *gorm.DB, vdoId string) (*entity.Videos, error) {
	var video entity.Videos
	err := db.Where("vdo_id = ?", vdoId).First(&video).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	return &video, nil
}
