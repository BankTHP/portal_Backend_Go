package repository

import (
	"errors"
	"pccth/portal-blog/internal/entity"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, user *entity.Users) error {
	return db.Create(user).Error
}

func GetUserByUserId(db *gorm.DB, userId string) (*entity.Users, error) {
	var user entity.Users
	err := db.Where("user_id = ?", userId).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	return &user, nil
}

func UpdateUser(db *gorm.DB, user *entity.Users) error {
	return db.Save(user).Error
}