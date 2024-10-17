package entity

import "time"

type Post struct {
    ID             uint           `gorm:"primaryKey;autoIncrement" json:"id"`                       
    PostHeader     string         `gorm:"type:varchar(255);not null" json:"postHeader"`              
    PostBody       string         `gorm:"type:text;not null" json:"postBody"`                        
    PostCreateDate time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"postCreateDate"`  
    PostCreateBy   string         `gorm:"type:varchar(100);not null" json:"postCreateBy"`
    
    Comments       []Comment      `gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE" json:"comments,omitempty"`
}

func (Post) TableName() string {
	return "post"
}

