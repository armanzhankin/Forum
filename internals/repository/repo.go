package repository

import "database/sql"

type Repository struct {
	Post     *PostRepo
	Comment  *CommentRepo
	Likes    *LikesRepo
	Dislikes *DislikesRepo
	User     *UserRepo
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Post:     CreatePostRepo(db),
		Comment:  CreateCommentRepo(db),
		Likes:    CreateLikesRepo(db),
		Dislikes: CreateDislikesRepo(db),
		User:     CreateUserRepo(db),
	}
}
