package service

import (
	"pccth/portal-blog/internal/entity"
	"pccth/portal-blog/internal/model"
	"pccth/portal-blog/internal/repository"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) CreateUser(createRequest *model.CreateUserRequest) error {
	user := &entity.Users{
		UserId:     createRequest.UserId,
		Name:       createRequest.Name,
		Username:   createRequest.Username,
		GivenName:  createRequest.GivenName,
		FamilyName: createRequest.FamilyName,
		Email:      createRequest.Email,
	}
	return repository.CreateUser(s.db, user)
}

func (s *UserService) UpdateUserInfo(userId string, updateRequest *model.UpdateUserRequest) error {
	user, err := repository.GetUserByUserId(s.db, userId)
	if err != nil {
		return err
	}

	if updateRequest.Name != "" {
		user.Name = updateRequest.Name
	}
	if updateRequest.Username != "" {
		user.Username = updateRequest.Username
	}
	if updateRequest.GivenName != "" {
		user.GivenName = updateRequest.GivenName
	}
	if updateRequest.FamilyName != "" {
		user.FamilyName = updateRequest.FamilyName
	}
	if updateRequest.Email != "" {
		user.Email = updateRequest.Email
	}

	return repository.UpdateUser(s.db, user)
}

func (s *UserService) GetUserInfoByUserId(userId string) (*entity.Users, error) {
	return repository.GetUserByUserId(s.db, userId)
}

func (s *UserService) CheckUser(userInfo *model.CreateUserRequest) (bool, error) {
	_, err := repository.GetUserByUserId(s.db, userInfo.UserId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			if err := s.CreateUser(userInfo); err != nil {
				return false, err
			}
			return false, nil
		}
		return false, err
	}
	return true, nil
}
