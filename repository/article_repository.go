package repository

import (
	"errors"
	"sharingvision-be-test/models"

	"gorm.io/gorm"
)

type ArticleRepository interface {
	Create(post *models.Post) error
	GetPaged(limit, offset int) ([]models.Post, error)
	GetByID(id int) (*models.Post, error)
	Update(id int, post *models.Post) error
	Delete(id int) error
}

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) ArticleRepository {
	return &articleRepository{db: db}
}

func (r *articleRepository) Create(post *models.Post) error {
	return r.db.Create(post).Error
}

func (r *articleRepository) GetPaged(limit, offset int) ([]models.Post, error) {
	var posts []models.Post
	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}
	err := r.db.Limit(limit).Offset(offset).Order("id DESC").Find(&posts).Error
	return posts, err
}

func (r *articleRepository) GetByID(id int) (*models.Post, error) {
	var post models.Post
	err := r.db.First(&post, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &post, nil
}

func (r *articleRepository) Update(id int, post *models.Post) error {
	var existing models.Post
	err := r.db.First(&existing, id).Error
	if err != nil {
		return err
	}

	existing.Title = post.Title
	existing.Content = post.Content
	existing.Category = post.Category
	existing.Status = post.Status

	return r.db.Save(&existing).Error
}

func (r *articleRepository) Delete(id int) error {
	var post models.Post
	err := r.db.First(&post, id).Error
	if err != nil {
		return err
	}
	return r.db.Delete(&post).Error
}
