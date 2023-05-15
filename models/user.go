package models

import "time"

type User struct {
	Id            int
	Username      string
	Email         string
	Password      string
	Token         string
	TokenDuration time.Time
	Auth          int
}
