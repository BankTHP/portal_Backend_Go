package entity

import "time"


type Notification struct {
    ID            uint      `gorm:"primaryKey;autoIncrement" json:"id"`        
    PostID        uint      `gorm:"not null" json:"postId"`                    
    CommentID     uint      `gorm:"not null" json:"commentId"`                 
    UserID        uint      `gorm:"not null" json:"userId"`                    
    IsRead        bool      `gorm:"default:false" json:"isRead"`              
    CreateDate    time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"createDate"` 
}


func (Notification) TableName() string {
    return "notifications"
}
