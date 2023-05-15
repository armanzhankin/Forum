package repository

import (
	"database/sql"
	"fmt"

	"forum/models"
)

type CommentRepo struct {
	db *sql.DB
}

type Comment interface {
	CreateComment(*models.Comment) error
	GetCommentsByPostId(int) ([]*models.Comment, error)
}

func CreateCommentRepo(db *sql.DB) *CommentRepo {
	return &CommentRepo{
		db: db,
	}
}

func (c *CommentRepo) CreateComment(comment *models.Comment) error {
	stmt := `INSERT INTO comment(
		username,
		content,
		post_id)
		VALUES(?, ?, ?)`
	_, err := c.db.Exec(stmt, comment.Author, comment.Content, comment.PostId)
	if err != nil {
		return err
	}

	return nil
}

func (c *CommentRepo) GetCommentsByPostId(id int) ([]*models.Comment, error) {
	stmt := `Select 
	* FROM comment WHERE post_id = $1`
	rows, err := c.db.Query(stmt, id)
	if err != nil {
		return []*models.Comment{}, fmt.Errorf("repo: : %s", err)
	}

	defer rows.Close()

	comments := []*models.Comment{}

	for rows.Next() {
		com := &models.Comment{}

		err = rows.Scan(&com.Id, &com.Author, &com.Content, &com.Likes, &com.Dislikes, &com.PostId)

		if err != nil {
			return []*models.Comment{}, fmt.Errorf("repo: : %s", err)
		}

		comments = append(comments, com)
	}

	return comments, nil
}
