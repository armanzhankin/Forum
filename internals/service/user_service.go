package service

import (
	"errors"
	"time"

	"forum/internals/repository"
	"forum/models"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo repository.User
}

type User interface {
	CreateUser(*models.User) error
	ValidUser(*models.User) error
	CheckUser(string, string) (models.User, error)
	DeleteToken(string) error
	GetUserByToken(string) (models.User, error)
	GetUsernameById(int) (string, error)
}

func CreateUserService(repo repository.User) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) ValidUser(user *models.User) error {
	if !CheckEmail(user.Email) || !CheckPassword(user.Password) || !CheckUsername(user.Username) {
		return models.ErrInvalidData
	}

	return nil
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.repo.CreateUser(user)
}

func (s *UserService) CheckUser(username, password string) (models.User, error) {
	user, err := s.repo.GetUserByEmail(username)
	if errors.Is(err, models.ErrInvalidData) {
		user, err = s.repo.GetUserByUsername(username)
		if errors.Is(err, models.ErrInvalidData) {
			return user, err
		} else if err != nil {
			return user, err
		}
	} else if err != nil {
		return user, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return user, err
	} else if err != nil {
		return user, err
	}
	gen := uuid.NewGen()
	token, err := gen.NewV4()
	if err != nil {
		return user, err
	}
	strToken := token.String()
	duration := time.Now().Add(time.Hour * 6)

	err = s.repo.SaveToken(user.Id, strToken, duration)
	if err != nil {
		return user, err
	}

	user.Token = strToken
	user.TokenDuration = duration
	return user, nil
}

func (s *UserService) DeleteToken(token string) error {
	err := s.repo.DeleteToken(token)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) GetUserByToken(token string) (models.User, error) {
	user, err := s.repo.GetUserByToken(token)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *UserService) GetUsernameById(id int) (string, error) {
	name, err := s.repo.GetUsernameById(id)
	if err != nil {
		return "", err
	}

	return name, err
}
