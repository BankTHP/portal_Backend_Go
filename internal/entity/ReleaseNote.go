package entity

import "time"


type Release struct {
    ID          uint      `gorm:"primaryKey;autoIncrement" json:"releaseId"`
    Body        string    `gorm:"type:text;not null" json:"releaseBody"`
    Header      string    `gorm:"type:varchar(255);not null" json:"releaseHeader"`
    CreateDate  time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"createDate"`
}


func (Release) TableName() string {
    return "releases"
}
