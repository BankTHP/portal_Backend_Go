package repository

import (
	"pccth/portal-blog/internal/entity"

	"gorm.io/gorm"
)

func CreatePost(db *gorm.DB, post *entity.Post) error {
    result := db.Create(post)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

func GetPostByID(db *gorm.DB, id uint) (*entity.Post, error) {
    var post entity.Post
    if err := db.First(&post, id).Error; err != nil {
        return nil, err
    }
    return &post, nil
}

func UpdatePost(db *gorm.DB, post *entity.Post) error {
    result := db.Save(post)
    return result.Error
}


func DeletePost(db *gorm.DB, id uint) error {
    return db.Delete(&entity.Post{}, id).Error
}


func GetAllPosts(db *gorm.DB) ([]entity.Post, error) {
    var posts []entity.Post
    result := db.Find(&posts)
    if result.Error != nil {
        return nil, result.Error
    }
    return posts, nil
}
