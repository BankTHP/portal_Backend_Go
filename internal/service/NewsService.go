package service

import (
	"errors"
	"pccth/portal-blog/internal/entity"
	"pccth/portal-blog/internal/model"
	"pccth/portal-blog/internal/repository"

	"gorm.io/gorm"
)

type NewsService struct {
    db *gorm.DB
}

func NewsPService(db *gorm.DB) *NewsService {
    return &NewsService{db: db}
}


func (s *NewsService) CreateNews(createRequest *model.CreateNewsRequest) error {
    news := &entity.News{
        NewsHeader: createRequest.NewsHeader,
        NewsBody:   createRequest.NewsBody,
    }
    return repository.CreateNews(s.db, news)
}

func (s *NewsService) GetNewsByID(id uint) (*entity.News, error) {
    return repository.GetNewsByID(s.db, id)
}

func (s *NewsService) UpdateNews(id uint, updateRequest *model.UpdateNewsRequest) error {
    if updateRequest.NewsHeader == "" || updateRequest.NewsBody == "" {
        return errors.New("NewsHeader and NewsBody cannot be empty")
    }

    news, err := repository.GetNewsByID(s.db, id)
    if err != nil {
        return err
    }


    news.NewsHeader = updateRequest.NewsHeader
    news.NewsBody = updateRequest.NewsBody

    return s.db.Save(news).Error
}

func (s *NewsService) DeleteNews(id uint) error {
    return repository.DeleteNews(s.db, id)
}

func (s *NewsService) GetAllNews() ([]entity.News, error) {
    return repository.GetAllNews(s.db)
}
