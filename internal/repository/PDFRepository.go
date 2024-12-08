package repository

import (
	"errors"
	"pccth/portal-blog/internal/entity"

	"gorm.io/gorm"
)

func CreatePDF(db *gorm.DB, pdf *entity.PDFs) error {
	return db.Create(pdf).Error
}

func GetPDFByName(db *gorm.DB, pdfName string) (*entity.PDFs, error) {
	var pdf entity.PDFs
	err := db.Where("pdf_name = ?", pdfName).First(&pdf).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	return &pdf, nil
}

func GetPDFsByPostID(db *gorm.DB, postID uint) ([]entity.PDFs, error) {
	var pdfs []entity.PDFs

	err := db.Where("post_id = ?", postID).
		Order("id asc").
		Find(&pdfs).Error
	if err != nil {
		return nil, err
	}

	return pdfs, nil
}
