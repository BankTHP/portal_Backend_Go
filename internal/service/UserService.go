package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"pccth/portal-blog/internal/entity"
	"pccth/portal-blog/internal/model"
	"pccth/portal-blog/internal/repository"

	"github.com/spf13/viper"
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
		UniversityName: createRequest.UniversityName,
	}
	return repository.CreateUser(s.db, user)
}

func (s *UserService) UpdateUserInfo(updateRequest *model.UpdateUserRequest, token string) error {
	user, err := repository.GetUserByUserId(s.db, updateRequest.UserId)
	if err != nil {
		return err
	}
	if updateRequest.Name != "" {
		user.Name = updateRequest.Name
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

	err = s.updateUserViaMicroservice(updateRequest, token)
	if err != nil {
		return err
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

func (s *UserService) updateUserViaMicroservice(updateRequest *model.UpdateUserRequest, token string) error {
	url := viper.GetString("keycloak.server") + "edit-user"

	data := model.UserUpdateRequest{
		Email:     updateRequest.Email,
		FirstName: updateRequest.GivenName,
		LastName:  updateRequest.FamilyName,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to update user, status code: %d", resp.StatusCode)
	}

	return nil
}
