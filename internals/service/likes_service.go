package service

import (
	"fmt"

	"forum/internals/repository"
	"forum/models"
)

type LikesService struct {
	repo repository.Likes
}

type Likes interface {
	LikePost(*models.Reaction) error
	LikeComment(*models.Reaction) error
}

func CreateLikesService(repo repository.Likes) *LikesService {
	return &LikesService{
		repo: repo,
	}
}

func (s *LikesService) LikePost(like *models.Reaction) error {
	flag, err := s.repo.CheckPostLike(like)
	if err != nil {
		fmt.Println("Like Post service: %w", err)
		return err
	}
	if flag {
		err := s.repo.DeletePostLike(like)
		if err != nil {
			return err
		}
	} else {
		flag2, err := s.repo.CheckPostDislike(like.PostId, like.User_Id)
		if flag2 {
			err = s.repo.DeletePostDislike(like.PostId, like.User_Id)
			if err != nil {
				return err
			}
			err = s.repo.LikePost(like)
			if err != nil {
				return err
			}
		} else {
			err = s.repo.LikePost(like)
			if err != nil {
				return err
			}
		}
	}
	err = s.repo.UpdatePostReaction(like.PostId)
	if err != nil {
		return err
	}

	return nil
}

func (s *LikesService) LikeComment(like *models.Reaction) error {
	flag, err := s.repo.CheckCommentLike(like)
	if err != nil {
		fmt.Println("Like Post service: %w", err)
		return err
	}

	if flag {
		err := s.repo.DeleteCommentLike(like)
		if err != nil {
			return err
		}
	} else {
		flag2, err := s.repo.CheckCommentDislike(like.CommentId, like.User_Id)
		if flag2 {
			fmt.Println(like.CommentId)

			err = s.repo.DeleteCommentDislike(like.CommentId, like.User_Id)
			if err != nil {
				return err
			}
			err = s.repo.LikeComment(like)
			if err != nil {
				return err
			}
		} else {

			err = s.repo.LikeComment(like)
			if err != nil {
				return err
			}
		}
	}
	err = s.repo.UpdateCommentReaction(like.CommentId)
	if err != nil {
		return err
	}

	return nil
}
