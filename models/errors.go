package models

import "errors"

var (
	ErrNoRecord       = errors.New("models: no matching record found")
	ErrDuplicate      = errors.New("username/email already registered")
	ErrPasswordConf   = errors.New("password not confirmed")
	ErrInvalidData    = errors.New("Invalid username/email or password")
	ErrInvalidPost    = errors.New("Invalid Post: too long OR empty title/content OR unexisting category")
	ErrInvalidComment = errors.New("Too long commentary, you can write maximum 280 characters")
	ErrNoPost         = errors.New("No posts with such category")
)
