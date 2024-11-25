package entity

type PDFs struct {
    ID          uint   `gorm:"primaryKey;autoIncrement" json:"id"`
    PDFName     string `gorm:"type:varchar(255)" json:"pdfName"`
    PDFSize     string `gorm:"type:varchar(255)" json:"pdfSize"`
    PDFPath     string `gorm:"type:varchar(255)" json:"pdfPath"`
}

func (PDFs) TableName() string {
    return "pdfs"
} 