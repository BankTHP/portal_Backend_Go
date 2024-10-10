package entity

import "time"

type News struct {
    ID          uint      `gorm:"primaryKey;autoIncrement" json:"newsId"`      
    NewsHeader      string    `gorm:"type:varchar(255);not null" json:"newsHeader"`
    NewsBody        string    `gorm:"type:text;not null" json:"newsBody"`        
    NewsCreateDate  time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"newsCreateDate"`
}

func (News) TableName() string {
    return "news"
}
