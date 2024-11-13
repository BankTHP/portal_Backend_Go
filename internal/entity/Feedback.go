package entity

import "time"

type Feedback struct {
	FeedBackId    uint      `gorm:"primaryKey;autoIncrement" json:"feedBackId"`
	FeedBackName  string    `gorm:"type:varchar(255)" json:"feedBackName"`
	FeedBackEmail string    `gorm:"type:varchar(255)" json:"feedBackEmail"`
	FeedBackPhone string    `gorm:"type:varchar(10)" json:"feedBackPhone"`
	FeedBackDate  time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"feedBackDate"`
	FeedBackText  string    `gorm:"type:text" json:"feedBackText"`
}

func (Feedback) TableName() string {
	return "feedbacks"
} 