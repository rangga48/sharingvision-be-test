package services

import (
	"errors"
	"sharingvision-be-test/dto"
	"sharingvision-be-test/models"
	"sharingvision-be-test/repository"
)

type ArticleService interface {
	CreateArticle(req *dto.ArticleRequest) error
	GetArticles(limit, offset int) ([]dto.ArticleResponse, error)
	GetArticleByID(id int) (*dto.ArticleResponse, error)
	UpdateArticle(id int, req *dto.ArticleRequest) error
	DeleteArticle(id int) error
}

type articleService struct {
	repo repository.ArticleRepository
}

func NewArticleService(repo repository.ArticleRepository) ArticleService {
	return &articleService{repo: repo}
}

func (s *articleService) CreateArticle(req *dto.ArticleRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}

	post := models.Post{
		Title:    req.Title,
		Content:  req.Content,
		Category: req.Category,
		Status:   req.Status,
	}

	return s.repo.Create(&post)
}

func (s *articleService) GetArticles(limit, offset int) ([]dto.ArticleResponse, error) {
	posts, err := s.repo.GetPaged(limit, offset)
	if err != nil {
		return nil, err
	}

	responses := make([]dto.ArticleResponse, 0, len(posts))
	for _, p := range posts {
		responses = append(responses, dto.ArticleResponse{
			ID:          p.ID,
			Title:       p.Title,
			Content:     p.Content,
			Category:    p.Category,
			CreatedDate: p.CreatedDate.Format("2006-01-02 15:04:05"),
			UpdatedDate: p.UpdatedDate.Format("2006-01-02 15:04:05"),
			Status:      p.Status,
		})
	}

	return responses, nil
}

func (s *articleService) GetArticleByID(id int) (*dto.ArticleResponse, error) {
	post, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if post == nil {
		return nil, errors.New("article not found")
	}

	return &dto.ArticleResponse{
		ID:          post.ID,
		Title:       post.Title,
		Content:     post.Content,
		Category:    post.Category,
		CreatedDate: post.CreatedDate.Format("2006-01-02 15:04:05"),
		UpdatedDate: post.UpdatedDate.Format("2006-01-02 15:04:05"),
		Status:      post.Status,
	}, nil
}

func (s *articleService) UpdateArticle(id int, req *dto.ArticleRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}

	post := models.Post{
		Title:    req.Title,
		Content:  req.Content,
		Category: req.Category,
		Status:   req.Status,
	}

	return s.repo.Update(id, &post)
}

func (s *articleService) DeleteArticle(id int) error {
	return s.repo.Delete(id)
}
