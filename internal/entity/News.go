package entity

import "time"

type News struct {
    ID          uint      `gorm:"primaryKey;autoIncrement" json:"newsId"`      
    Header      string    `gorm:"type:varchar(255);not null" json:"newsHeader"`
    Body        string    `gorm:"type:text;not null" json:"newsBody"`        
    CreateDate  time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"newsCreateDate"`
}

func (News) TableName() string {
    return "news"
}
