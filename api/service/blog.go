package service

import (
	"github.com/Zulhaditya/divetech-blog/api/repository"
	"github.com/Zulhaditya/divetech-blog/models"
)

// post service struct
type PostService struct {
	repository repository.PostRepository
}

// returns the post service struct instance
func NewPostService(r repository.PostRepository) PostService {
	return PostService{
		repository: r,
	}
}

// calls post repository save method
func (p PostService) Save(post models.Post) error {
	return p.repository.Save(post)
}

// call find all method
func (p PostService) FindAll(post models.Post, keyword string) (*[]models.Post, int64, error) {
	return p.repository.FindAll(post, keyword)
}

// call update method
func (p PostService) Update(post models.Post) error {
	return p.repository.Update(post)
}

// call delete method
func (p PostService) Delete(id int64) error {
	var post models.Post
	post.ID = id
	return p.repository.Delete(post)
}

// call find method
func (p PostService) Find(post models.Post) (models.Post, error) {
	return p.repository.Find(post)
}
