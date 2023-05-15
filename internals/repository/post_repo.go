package repository

import (
	"database/sql"
	"fmt"
	"regexp"
	"strings"

	"forum/models"
)

type PostRepo struct {
	db *sql.DB
}

type Post interface {
	CreatePost(*models.Post) (int, error)
	GetPostByID(int) (*models.Post, error)
	AllPosts() ([]*models.Post, error)
	GetMylikedPosts(models.User) ([]*models.Post, error)
	GetMyPosts(models.User) ([]*models.Post, error)
	CategoryPost(string) ([]*models.Post, error)
}

func CreatePostRepo(db *sql.DB) *PostRepo {
	return &PostRepo{
		db: db,
	}
}

func (p *PostRepo) CreatePost(post *models.Post) (int, error) {
	stmt := `INSERT INTO post (
		user_id,
		category,
		title,
		content)
		VALUES(?, ?, ?, ?)`

	var categoryStr string

	if len(post.Category) == 1 {
		categoryStr = post.Category[0]
	} else {
		categoryStr = strings.Join(post.Category, ", ")
	}
	res, err := p.db.Exec(stmt, post.User_Id, categoryStr, post.Title, post.Content)
	fmt.Println(post.Content)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, nil
	}
	return int(id), nil
}

func (p *PostRepo) GetPostByID(id int) (*models.Post, error) {
	stmt := `SELECT id, user_id, category, title, content, like, dislike FROM post WHERE id = ?`

	row := p.db.QueryRow(stmt, id)

	m := &models.Post{}
	var categoryStr string
	err := row.Scan(&m.Id, &m.User_Id, &categoryStr, &m.Title, &m.Content, &m.Like, &m.Dislike)
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}

	m.Category = strings.Split(categoryStr, ", ")

	return m, nil
}

func (p *PostRepo) AllPosts() ([]*models.Post, error) {
	stmt := `SELECT 
	p.id, 
	u.username, 
	p.category, 
	p.title, 
	p.content,
	p.like,
	p.dislike
FROM 
	post p  
JOIN 
	users u
ON 
	p.user_id = u.id
	`

	rows, err := p.db.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	posts := []*models.Post{}
	var categoryStr string
	for rows.Next() {
		m := &models.Post{}
		err = rows.Scan(&m.Id, &m.Author, &categoryStr, &m.Title, &m.Content, &m.Like, &m.Dislike)
		if err != nil {
			return nil, err
		}
		m.Category = strings.Split(categoryStr, ", ")
		posts = append(posts, m)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return posts, nil
}

func (p *PostRepo) CategoryPost(s string) ([]*models.Post, error) {
	stmt := `SELECT 
	p.id, 
	u.username, 
	p.category, 
	p.title, 
	p.content,
	p.like,
	p.dislike
FROM 
	post p  
JOIN 
	users u
ON 
	p.user_id = u.id
	WHERE p.category LIKE '%' || ? || '%'`
	rows, err := p.db.Query(stmt, s)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	posts := []*models.Post{}
	var categoryStr string

	for rows.Next() {
		m := &models.Post{}

		err = rows.Scan(&m.Id, &m.Author, &categoryStr, &m.Title, &m.Content, &m.Like, &m.Dislike)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		fmt.Println(categoryStr)
		match, err := regexp.MatchString(".*"+s+".*", categoryStr)
		if err != nil {
			fmt.Println(err)
			return nil, models.ErrNoPost
		}

		if match {
			m.Category = strings.Split(categoryStr, ", ")
			posts = append(posts, m)
		}

	}

	if err = rows.Err(); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return posts, nil
}

func (p *PostRepo) GetMyPosts(user models.User) ([]*models.Post, error) {
	Myposts := []*models.Post{}

	//stmt := `SELECT
	//	id,
	//	username,
	//	category,
	//	title,
	//	content,
	//	like,
	//	dislike
	//FROM
	//	post
	//	WHERE p.user_id = ?
	//`
	stmt := `SELECT
	p.id,
		u.username,
		p.category,
		p.title,
		p.content,
		p.like,
		p.dislike
	FROM
	post p
	JOIN
	users u
	ON
	p.user_id = u.id
	WHERE p.user_id = ?`

	rows, err := p.db.Query(stmt, user.Id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var categoryStr string

	for rows.Next() {
		m := &models.Post{}

		err = rows.Scan(&m.Id, &m.Author, &categoryStr, &m.Title, &m.Content, &m.Like, &m.Dislike)
		fmt.Println(m.Content)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		m.Category = strings.Split(categoryStr, ", ")

		Myposts = append(Myposts, m)

	}
	return Myposts, nil
}
func (p *PostRepo) GetMylikedPosts(user models.User) ([]*models.Post, error) {
	stmt := `SELECT
	p.id,
		u.username,
		p.category,
		p.title,
		p.content,
		p.like,
		p.dislike
	FROM
	post p
	JOIN
	users u
	ON
	p.user_id = u.id
	JOIN
	likes l
	ON
	p.id = l.post_id
	WHERE l.user_id = ?`

	rows, err := p.db.Query(stmt, user.Id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var categoryStr string
	posts := []*models.Post{}
	for rows.Next() {
		m := &models.Post{}

		err = rows.Scan(&m.Id, &m.Author, &categoryStr, &m.Title, &m.Content, &m.Like, &m.Dislike)
		fmt.Println(m.Content)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		m.Category = strings.Split(categoryStr, ", ")

		posts = append(posts, m)

	}
	return posts, nil
}
