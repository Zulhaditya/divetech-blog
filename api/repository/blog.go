package repository

import (
	"github.com/Zulhaditya/divetech-blog/initdb"
	"github.com/Zulhaditya/divetech-blog/models"
)

// post repository
type PostRepository struct {
	db initdb.Database
}

// new post repository : fetching database
func NewPostRepository(db initdb.Database) PostRepository {
	return PostRepository{
		db: db,
	}
}

// method for save post to database
func (p PostRepository) Save(post models.Post) error {
	return p.db.DB.Create(&post).Error
}

// method for fetching all post from database
func (p PostRepository) FindAll(post models.Post, keyword string) (*[]models.Post, int64, error) {
	var posts []models.Post
	var totalRows int64 = 0

	queryBuilder := p.db.DB.Order("created_at desc").Model(&models.Post{})
	// search parameter
	if keyword != "" {
		queryKeyword := "%" + keyword + "%"
		queryBuilder = queryBuilder.Where(
			p.db.DB.Where("post.title LIKE ? ", queryKeyword))
	}

	err := queryBuilder.Where(post).Find(&posts).Count(&totalRows).Error
	return &posts, totalRows, err
}

// method for updating post
func (p PostRepository) Update(post models.Post) error {
	return p.db.DB.Save(&post).Error
}

// method for fetching post by id
func (p PostRepository) Find(post models.Post) (models.Post, error) {
	var posts models.Post
	err := p.db.DB.Debug().Model(&models.Post{}).Where(&post).Take(&posts).Error
	return posts, err
}

// method for delete post
func (p PostRepository) Delete(post models.Post) error {
	return p.db.DB.Delete(&post).Error
}
