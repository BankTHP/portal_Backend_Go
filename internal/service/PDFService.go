package service

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"pccth/portal-blog/internal/entity"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type PDFService struct {
    db         *gorm.DB
    uploadPath string
}

func NewPDFService(db *gorm.DB, uploadPath string) *PDFService {
    return &PDFService{
        db:         db,
        uploadPath: uploadPath,
    }
    
}
func (s *PostService) SavePDF(postID uint, file *multipart.FileHeader) error {
	filename := file.Filename
	uploadPath := fmt.Sprintf("uploads/pdfs/%s", filename)

	port := viper.GetString("app.port")

	dir := "uploads/pdfs"
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(uploadPath)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return err
	}

	fullURL := fmt.Sprintf("http://localhost:%s/pdfs/%s", port, filename)

	pdf := &entity.PDFs{
		PostID:  postID,
		PDFName: filename,
		PDFSize: fmt.Sprintf("%d", file.Size),
		PDFPath: fullURL,
	}

	return s.db.Create(pdf).Error
}