package entity

type PDFs struct {
	ID      uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	PostID  uint   `gorm:"not null" json:"postId"`
	PDFName string `gorm:"type:varchar(255)" json:"pdfName"`
	PDFSize string `gorm:"type:varchar(255)" json:"pdfSize"`
	PDFPath string `gorm:"type:varchar(255)" json:"pdfPath"`
}

func (PDFs) TableName() string {
	return "pdfs"
}
