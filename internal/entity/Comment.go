package entity

import "time"

type Comment struct {
    ID              uint      `gorm:"primaryKey;autoIncrement" json:"id"`
    PostID          uint      `gorm:"not null" json:"postId"`
    CommentBody     string    `gorm:"type:text;not null" json:"commentBody"`
    CommentCreateDate time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"commentCreateDate"`
    CommentCreateBy string    `gorm:"type:varchar(100);not null" json:"commentCreateBy"`
}

func (Comment) TableName() string {
    return "comments"
}
