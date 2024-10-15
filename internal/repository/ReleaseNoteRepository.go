package repository

import (
	"pccth/portal-blog/internal/entity"
	"gorm.io/gorm"
)

func CreateRelease(db *gorm.DB, release *entity.Release) error {
	result := db.Create(release)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
