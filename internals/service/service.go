package service

import "forum/internals/repository"

type Service struct {
	Post
	Comment
	Likes
	Dislikes
	User
}

func NewService(repo repository.Repository) *Service {
	return &Service{
		Post:     CreatePostService(repo.Post),
		Comment:  CreateCommentService(repo.Comment),
		Likes:    CreateLikesService(repo.Likes),
		Dislikes: CreateDislikesService(repo.Dislikes),
		User:     CreateUserService(repo.User),
	}
}
