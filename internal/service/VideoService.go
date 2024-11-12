package service

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"pccth/portal-blog/internal/entity"
	"pccth/portal-blog/internal/model"
	"pccth/portal-blog/internal/repository"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VideoService struct {
	db         *gorm.DB
	uploadPath string
}

func NewVideoService(db *gorm.DB, uploadPath string) *VideoService {
	return &VideoService{
		db:         db,
		uploadPath: uploadPath,
	}
}

func (s *VideoService) ensureDir(dirPath string) error {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		return os.MkdirAll(dirPath, 0755)
	}
	return nil
}

func (s *VideoService) getVideoDuration(filename string) (string, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return "", fmt.Errorf("ไม่พบไฟล์: %v", err)
	}

	cmd := exec.Command("ffmpeg", "-i", filename, "2>&1")
	output, err := cmd.CombinedOutput()
	if err != nil {
		outputStr := string(output)
		for _, line := range strings.Split(outputStr, "\n") {
			if strings.Contains(line, "Duration:") {
				parts := strings.Split(line, "Duration: ")
				if len(parts) > 1 {
					duration := strings.Split(parts[1], ",")[0]
					return strings.TrimSpace(duration), nil
				}
			}
		}
	}
	return "", fmt.Errorf("ไม่สามารถอ่านความยาววิดีโอได้")
}

func (s *VideoService) ProcessVideoUpload(file *model.UploadedFile) (*model.VideoResponse, error) {
	// ตรวจสอบขนาดไฟล์
	maxSize := int64(500 * 1024 * 1024) // 500MB
	if file.Size > maxSize {
		return &model.VideoResponse{
			Success: false,
			Error: fmt.Sprintf("ไม่สามารถอัปโหลดได้: ขนาดไฟล์ %.2f MB เกินขนาดที่กำหนด (ไม่เกิน %.2f MB)",
				float64(file.Size)/(1024*1024),
				float64(maxSize)/(1024*1024)),
		}, nil
	}

	// ตรวจสอบนามสกุลไฟล์
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if ext != ".mp4" && ext != ".mov" && ext != ".avi" && ext != ".wmv" {
		return &model.VideoResponse{
			Success: false,
			Error:   "นามสกุลไฟล์ไม่ถูกต้อง รองรับเฉพาะ: .mp4, .mov, .avi, .wmv",
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
	duration, err := s.getVideoDuration(fullPath)
    if err != nil {
        duration = "ไม่สามารถอ่านความยาววิดีโอได้"
    } else {
        parts := strings.Split(duration, ":")
        if len(parts) == 3 {
            hours := strings.TrimSpace(parts[0])
            if hours != "00" {
                os.Remove(fullPath)
                return &model.VideoResponse{
                    Success: false,
                    Error:   "วิดีโอต้องมีความยาวไม่เกิน 1 ชั่วโมง",
                }, nil
            }
        }
    }

	fileSizeMB := fmt.Sprintf("%.2f MB", float64(file.Size)/1024/1024)
	urlPath := filepath.Join("/videos", filename)

	// บันทึกข้อมูลลงฐานข้อมูล
	video := &entity.Videos{
		VdoName:     filename,
		VdoSize:     fileSizeMB,
		VdoDuration: duration,
	}

	if err := repository.CreateVideo(s.db, video); err != nil {
		os.Remove(fullPath)
		return nil, fmt.Errorf("ไม่สามารถบันทึกข้อมูลวิดีโอได้: %v", err)
	}

	return &model.VideoResponse{
		Success:  true,
		FullURL:  urlPath,
		Size:     fileSizeMB,
		Duration: duration,
	}, nil
}

func (s *VideoService) GetVideoByName(vdoName string) (*entity.Videos, error) {
	video, err := repository.GetVideoByName(s.db, vdoName)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("ไม่พบวิดีโอที่ต้องการ")
		}
		return nil, fmt.Errorf("เกิดข้อผิดพลาดในการค้นหาวิดีโอ: %v", err)
	}
	return video, nil
}

