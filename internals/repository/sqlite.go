package repository

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const (
	userTable = `CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT UNIQUE,
		email TEXT UNIQUE,
		password TEXT,
		token TEXT,
		token_duration DATETIME);`

	postTable = `CREATE TABLE IF NOT EXISTS post(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		category TEXT,
		title TEXT,
		content TEXT,
		like INTEGER DEFAULT 0,
		dislike INTEGER DEFAULT 0,
		FOREIGN KEY (user_id) REFERENCES users (id)
		);`
	commentTable = `CREATE TABLE IF NOT EXISTS comment(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT,
		content TEXT,
		likes INTEGER DEFAULT 0,
		dislikes INTEGER DEFAULT 0,
		post_id int,
		FOREIGN KEY (post_id) REFERENCES post (id)
		);`
	likesTable = `CREATE TABLE IF NOT EXISTS likes(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		post_id INTEGER,
		comment_id INTEGER,
		status INTEGER,
		FOREIGN KEY (post_id) REFERENCES post (id),
		FOREIGN KEY (comment_id) REFERENCES comment (id),
		FOREIGN KEY (user_id) REFERENCES users (id)
		);`
	dislikesTable = `CREATE TABLE IF NOT EXISTS dislikes(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		post_id INTEGER,
		comment_id INTEGER,
		status INTEGER,
		FOREIGN KEY (post_id) REFERENCES post (id),
		FOREIGN KEY (comment_id) REFERENCES comment (id),
		FOREIGN KEY (user_id) REFERENCES users (id)
		);`
)

func InitRepo() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func CreateTables(db *sql.DB) error {
	tables := []string{userTable, postTable, commentTable, likesTable, dislikesTable}
	for _, table := range tables {
		_, err := db.Exec(table)
		if err != nil {
			return err
		}
	}
	return nil
}
