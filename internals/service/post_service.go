package service

import (
	"forum/internals/repository"
	"forum/models"
)

type PostService struct {
	repo repository.Post
}

type Post interface {
	CreatePost(*models.Post) (int, error)
	GetPostByID(int) (*models.Post, error)
	AllPosts() ([]*models.Post, error)
	ValidPost(*models.Post) (*models.Post, error)
	GetMyPosts(models.User) ([]*models.Post, error)
	GetPostsByCategories(string) ([]*models.Post, error)
	GetLikedPosts(models.User) ([]*models.Post, error)
}

func CreatePostService(repo repository.Post) *PostService {
	return &PostService{
		repo: repo,
	}
}

func (p *PostService) GetMyPosts(user models.User) ([]*models.Post, error) {
	posts, err := p.repo.GetMyPosts(user)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (p *PostService) ValidPost(post *models.Post) (*models.Post, error) {
	content, flag := CheckPostContent(post.Content)
	title, flag2 := CheckTitle(post.Title)
	category, flag3 := CheckCategory(post.Category)
	if !flag || !flag2 || !flag3 {
		return post, models.ErrInvalidPost
	}

	post.Title = title
	post.Content = content
	post.Category = category
	return post, nil
}

func (p *PostService) CreatePost(post *models.Post) (int, error) {
	return p.repo.CreatePost(post)
}

func (p *PostService) GetPostByID(id int) (*models.Post, error) {
	post, err := p.repo.GetPostByID(id)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (p *PostService) AllPosts() ([]*models.Post, error) {
	posts, err := p.repo.AllPosts()
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (p *PostService) GetPostsByCategories(s string) ([]*models.Post, error) {
	posts, err := p.repo.CategoryPost(s)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (p *PostService) GetLikedPosts(user models.User) ([]*models.Post, error) {
	posts, err := p.repo.GetMylikedPosts(user)
	if err != nil {
		return nil, err
	}
	return posts, nil
}
