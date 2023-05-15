package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"forum/models"

	"github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

type UserRepo struct {
	db *sql.DB
}

type User interface {
	CreateUser(*models.User) error
	GetUserByEmail(string) (models.User, error)
	GetUserByUsername(string) (models.User, error)
	GetUserByToken(string) (models.User, error)
	SaveToken(int, string, time.Time) error
	DeleteToken(string) error
	GetUsernameById(int) (string, error)
}

func CreateUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) CreateUser(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		return err
	}

	stmt := `INSERT INTO users(
		username,
		email,
		password)
		VALUES(?, ?, ?);`

	_, err = r.db.Exec(stmt, user.Username, user.Email, string(hashedPassword))

	if err != nil {
		if sqliteErr, ok := err.(sqlite3.Error); ok {
			if sqliteErr.Code == sqlite3.ErrConstraint && strings.Contains(sqliteErr.Error(), "UNIQUE constraint failed") {
				return models.ErrDuplicate
			}
		} else if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return models.ErrDuplicate
		}
		return err
	}
	return nil
}

func (r *UserRepo) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	stmt := `SELECT id, username, password FROM users WHERE email = ?`
	row := r.db.QueryRow(stmt, email)

	err := row.Scan(&user.Id, &user.Username, &user.Password)
	if err == sql.ErrNoRows {
		log.Println(err)

		return user, models.ErrInvalidData
	} else if err != nil {
		log.Println(err)
		return user, err
	}

	return user, nil
}

func (r *UserRepo) GetUserByUsername(username string) (models.User, error) {
	var user models.User
	stmt := `SELECT id, username, password FROM users WHERE username = ?`
	row := r.db.QueryRow(stmt, username)
	err := row.Scan(&user.Id, &user.Username, &user.Password)
	if errors.Is(err, sql.ErrNoRows) {
		return user, models.ErrInvalidData
	} else if err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserRepo) GetUsernameById(id int) (string, error) {
	var user string
	stmt := `SELECT username FROM users WHERE id = ?`
	row := r.db.QueryRow(stmt, id)
	err := row.Scan(&user)
	if errors.Is(err, sql.ErrNoRows) {
		return "", models.ErrInvalidData
	} else if err != nil {
		return "", err
	}

	return user, nil
}

func (r *UserRepo) GetUserByToken(token string) (models.User, error) {
	var user models.User
	stmt := `SELECT * FROM users WHERE token = ?`
	row := r.db.QueryRow(stmt, token)
	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Token, &user.TokenDuration)
	if err != nil {
		return user, err
	}

	return user, err
}

func (r *UserRepo) SaveToken(id int, token string, duration time.Time) error {
	stmt := `UPDATE users SET token=$1, token_duration=$2 WHERE id=$3`
	_, err := r.db.Exec(stmt, token, duration, id)
	if err != nil {
		return fmt.Errorf("ERROR: /repository save token: %w", err)
	}
	return nil
}

func (r *UserRepo) DeleteToken(token string) error {
	stmt := `UPDATE users SET token=NULL, token_duration=NULL WHERE token=?`
	_, err := r.db.Exec(stmt, token)
	if err != nil {
		return err
	}

	return nil
}
