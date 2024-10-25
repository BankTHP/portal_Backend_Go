package repository

import (
	"pccth/portal-blog/internal/entity"

	"gorm.io/gorm"
)

func CreateNews(db *gorm.DB, news *entity.News) error {
    result := db.Create(news)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

func GetNewsByID(db *gorm.DB, id uint) (*entity.News, error) {
    var news entity.News
    if err := db.First(&news, id).Error; err != nil {
        return nil, err
    }
    return &news, nil
}

func UpdateNews(db *gorm.DB, news *entity.News) error {
    result := db.Save(news)
    return result.Error
}


func DeleteNews(db *gorm.DB, id uint) error {
    return db.Delete(&entity.News{}, id).Error
}


func GetAllNews(db *gorm.DB) ([]entity.News, error) {
    var news []entity.News
    result := db.Find(&news)
    if result.Error != nil {
        return nil, result.Error
    }
    return news, nil
}
