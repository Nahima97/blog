package services

import (
	"blog/models"
	"blog/repository"

	"github.com/google/uuid"
)

type PostService struct {
	Repo repository.PostRepository
}

func (s *PostService) CreatePost(req *models.Post) error {
		_, err := s.Repo.GetPostByTitle(req.Title)
	if err == nil {
		return err
	}

	uuid := uuid.New()
	req.ID = uuid

	err = s.Repo.CreatePost(req)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostService) GetAllPosts() ([]models.Post, error) {

	posts, err := s.Repo.GetAllPosts()
	if err != nil {
		return nil, err
	}
	return posts, nil

}

func (s *PostService) GetPostByID(id string) (*models.Post, error) {
	post, err := s.Repo.GetPostByID(id)
	if err != nil {
		return &models.Post{}, err
	}
	return post, nil
}

func (s *PostService) UpdateOwnPost(req *models.Post, id string) error {
	
	err := s.Repo.UpdateOwnPost(req, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostService) DeleteOwnPost(id string) error {

	err := s.Repo.DeleteOwnPost(id)
	if err != nil {
		return err
	}
	return nil
}
