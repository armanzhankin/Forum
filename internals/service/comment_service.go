package service

import (
	"forum/internals/repository"
	"forum/models"
)

type CommentService struct {
	repo repository.Comment
}

type Comment interface {
	CreateComment(*models.Comment) error
	GetCommentsByPostId(int) ([]*models.Comment, error)
	ValidComment(*models.Comment) (*models.Comment, error)
}

func CreateCommentService(repo repository.Comment) *CommentService {
	return &CommentService{
		repo: repo,
	}
}

func (c *CommentService) CreateComment(comment *models.Comment) error {
	return c.repo.CreateComment(comment)
}

func (c *CommentService) GetCommentsByPostId(id int) ([]*models.Comment, error) {
	comments, err := c.repo.GetCommentsByPostId(id)
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (c *CommentService) ValidComment(comment *models.Comment) (*models.Comment, error) {
	com, flag := CheckComment(comment.Content)
	if !flag {
		return comment, models.ErrInvalidComment
	}

	comment.Content = com
	return comment, nil
}
