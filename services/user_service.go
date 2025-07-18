package services

import (
	"blog/middleware"
	"blog/models"
	"blog/repository"
	"blog/utils"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserService struct {
	Repo repository.UserRepository
}

func (s *UserService) RegisterUser(req *models.User) error {
	existingUser, err := s.Repo.GetUserByUsername(req.Name)
	if err == nil && existingUser != nil {
    return fmt.Errorf("user already exists")
}

if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
    return err 
} 

	hashedPass, err := utils.HashPassword(req.Password)
	if err != nil {
		return err
	}
	req.Password = hashedPass

	uuid := uuid.New()
	req.ID = uuid

	err = s.Repo.CreateUser(req)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) Login(req *models.User) (string, error) {
	user, err := s.Repo.GetUserByUsername(req.Name)
	if err != nil {
		return "", err
	}

	err = utils.ComparePassword(user.Password, req.Password)
	if err != nil {
		return "", err
	}

	token, err := middleware.GenerateJWT(user.ID.String(), "user")
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *UserService) GetUserInfo(id string) (*models.User, error) {

	user, err := s.Repo.GetUserByID(id)
	if err != nil {
		return &models.User{}, err
	}
	return user, nil
}

