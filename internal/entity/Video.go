package entity

type Videos struct {
	ID          uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	VdoName     string `gorm:"type:varchar(255)" json:"vdoName"`
	VdoSize     string `gorm:"type:varchar(255)" json:"vdoSize"`
	VdoDuration string `gorm:"type:varchar(255)" json:"vdoDuration"`
}

func (Videos) TableName() string {
	return "videos"
}
