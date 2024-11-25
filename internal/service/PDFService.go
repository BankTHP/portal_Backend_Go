package service

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
    "pccth/portal-blog/internal/entity"
    "pccth/portal-blog/internal/model"
    "pccth/portal-blog/internal/repository"
    "github.com/google/uuid"
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

func (s *PDFService) ensureDir(dirPath string) error {
    if _, err := os.Stat(dirPath); os.IsNotExist(err) {
        return os.MkdirAll(dirPath, 0755)
    }
    return nil
}

func (s *PDFService) ProcessPDFUpload(file *model.UploadedPDFFile) (*model.PDFResponse, error) {
    // ตรวจสอบขนาดไฟล์
    maxSize := int64(10 * 1024 * 1024) // 10MB
    if file.Size > maxSize {
        return &model.PDFResponse{
            Success: false,
            Error: fmt.Sprintf("ไม่สามารถอัปโหลดได้: ขนาดไฟล์ %.2f MB เกินขนาดที่กำหนด (ไม่เกิน %.2f MB)",
                float64(file.Size)/(1024*1024),
                float64(maxSize)/(1024*1024)),
        }, nil
    }

    // ตรวจสอบนามสกุลไฟล์
    ext := strings.ToLower(filepath.Ext(file.Filename))
    if ext != ".pdf" {
        return &model.PDFResponse{
            Success: false,
            Error:   "นามสกุลไฟล์ไม่ถูกต้อง รองรับเฉพาะ: .pdf",
        }, nil
    }

    // สร้างโฟลเดอร์ถ้ายังไม่มี
    if err := s.ensureDir(s.uploadPath); err != nil {
        return nil, fmt.Errorf("ไม่สามารถสร้างโฟลเดอร์ได้: %v", err)
    }

    // สร้างชื่อไฟล์ใหม่
    filename := uuid.New().String() + ext
    fullPath := filepath.Join(s.uploadPath, filename)

    // บันทึกไฟล์
    if err := file.SaveFunc(fullPath); err != nil {
        return nil, fmt.Errorf("เกิดข้อผิดพลาดในการบันทึกไฟล์: %v", err)
    }

    fileSizeMB := fmt.Sprintf("%.2f MB", float64(file.Size)/1024/1024)
    urlPath := filepath.Join("/pdfs", filename)

    // บันทึกข้อมูลลงฐานข้อมูล
    pdf := &entity.PDFs{
        PDFName: filename,
        PDFSize: fileSizeMB,
        PDFPath: urlPath,
    }

    if err := repository.CreatePDF(s.db, pdf); err != nil {
        os.Remove(fullPath)
        return nil, fmt.Errorf("ไม่สามารถบันทึกข้อมูล PDF ได้: %v", err)
    }

    return &model.PDFResponse{
        Success: true,
        FullURL: urlPath,
        Size:    fileSizeMB,
    }, nil
}

func (s *PDFService) GetPDFByName(pdfName string) (*entity.PDFs, error) {
    pdf, err := repository.GetPDFByName(s.db, pdfName)
    if err != nil {
        return nil, fmt.Errorf("ไม่พบไฟล์ PDF: %v", err)
    }
    return pdf, nil
} 