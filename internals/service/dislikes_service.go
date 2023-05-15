package service

import (
	"fmt"

	"forum/internals/repository"
	"forum/models"
)

type DislikesService struct {
	repo repository.Dislikes
}

type Dislikes interface {
	DislikePost(*models.Reaction) error
	DislikeComment(*models.Reaction) error
}

func CreateDislikesService(repo repository.Dislikes) *DislikesService {
	return &DislikesService{
		repo: repo,
	}
}

func (s *DislikesService) DislikePost(dislike *models.Reaction) error {
	flag, err := s.repo.CheckPostDislike(dislike)
	if err != nil {
		fmt.Println("Like Post service: %w", err)
		return err
	}
	if flag {
		err := s.repo.DeletePostDislike(dislike)
		if err != nil {
			return err
		}
	} else {
		flag2, err := s.repo.CheckPostLike(dislike.PostId, dislike.User_Id)
		if flag2 {
			err = s.repo.DeletePostLike(dislike.PostId, dislike.User_Id)
			if err != nil {
				return err
			}
			err = s.repo.DislikePost(dislike)
			if err != nil {
				return err
			}
		} else {
			err = s.repo.DislikePost(dislike)
			if err != nil {
				return err
			}
		}
	}
	err = s.repo.UpdatePostReaction(dislike.PostId)
	if err != nil {
		return err
	}

	return nil
}

func (s *DislikesService) DislikeComment(dislike *models.Reaction) error {
	flag, err := s.repo.CheckCommentDislike(dislike)
	if err != nil {
		fmt.Println("Like Post service: %w", err)
		return err
	}
	if flag {
		err := s.repo.DeleteCommentDislike(dislike)
		if err != nil {
			return err
		}
	} else {
		flag2, err := s.repo.CheckCommentLike(dislike.CommentId, dislike.User_Id)
		if flag2 {
			err = s.repo.DeleteCommentLike(dislike.CommentId, dislike.User_Id)
			if err != nil {
				return err
			}
			err = s.repo.DislikeComment(dislike)
			if err != nil {
				return err
			}
		} else {
			err = s.repo.DislikeComment(dislike)
			if err != nil {
				return err
			}
		}
	}
	err = s.repo.UpdateCommentReaction(dislike.CommentId)
	if err != nil {
		return err
	}

	return nil
}
