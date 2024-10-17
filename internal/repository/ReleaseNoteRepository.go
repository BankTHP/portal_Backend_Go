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

func GetReleaseByID(db *gorm.DB, id uint) (*entity.Release, error) {
	var release entity.Release
	if err := db.First(&release, id).Error; err != nil {
		return nil, err
	}
	return &release, nil
}

func UpdateRelease(db *gorm.DB, release *entity.Release) error {
	result := db.Save(release)
	return result.Error
}

func DeleteRelease(db *gorm.DB, id uint) error {
	return db.Delete(&entity.Release{}, id).Error
}

func GetAllRelease(db *gorm.DB) ([]entity.Release, error) {
	var release []entity.Release
	result := db.Find(&release)
	if result.Error != nil {
		return nil, result.Error
	}
	return release, nil
}
