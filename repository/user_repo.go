package repository

import (
	"blog/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByUsername(username string) (*models.User, error)
	CreateUser(user *models.User) error
	GetUserByID(id string) (*models.User, error)
}

type UserRepo struct {
	Db *gorm.DB
}

func (r *UserRepo) GetUserByUsername(username string) (*models.User, error) {
	var user models.User

	err := r.Db.Where("name =?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) GetUserByID(id string) (*models.User, error) {
	var user models.User
	err := r.Db.Where("id =?", id).First(&user).Error
	if err != nil {
		return &models.User{}, err
	}
	return &user, nil
}
func (r *UserRepo) CreateUser(user *models.User) error {
	err := r.Db.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}
