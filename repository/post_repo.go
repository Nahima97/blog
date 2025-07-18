package repository

import (
	"blog/models"

	"gorm.io/gorm"
)

type PostRepository interface {
	GetPostByTitle(title string) (*models.Post, error)
	CreatePost(req *models.Post) error
	GetAllPosts() ([]models.Post, error)
	GetPostByID(id string) (*models.Post, error)
	UpdateOwnPost(req *models.Post, id string) error
	DeleteOwnPost(id string) error 
}

type PostRepo struct {
Db *gorm.DB
}

func (r *PostRepo) GetPostByTitle(title string) (*models.Post, error) {
	var post models.Post
	err := r.Db.Where("title =?", title).First(&post).Error
	if err != nil {
		return &models.Post{}, err
	}
	return &post, nil
}


func (r *PostRepo) CreatePost(req *models.Post) error {
	err := r.Db.Create(&req).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *PostRepo) GetAllPosts() ([]models.Post, error) {
	var posts []models.Post
	err := r.Db.Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil

}

func (r *PostRepo) GetPostByID(id string) (*models.Post, error) {
	var post models.Post

	err := r.Db.Where("id =? ", id).First(&post).Error
	if err != nil {
		return &models.Post{}, err
	}
	return &post, nil
}

func (r *PostRepo) UpdateOwnPost(req *models.Post, id string) error {
var post models.Post

err := r.Db.Model(&post).Where("id =?", id).Save(&req).Error
if err != nil {
	return err
}
return nil 
}

func (r *PostRepo) DeleteOwnPost(id string) error {
	var post models.Post

	err := r.Db.Where("id = ?", id).Delete(&post).Error
	if err != nil {
		return err
	}
	return nil
}
