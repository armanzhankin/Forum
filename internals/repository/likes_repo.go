package repository

import (
	"database/sql"
	"fmt"

	"forum/models"
)

type LikesRepo struct {
	db *sql.DB
}

type Likes interface {
	LikePost(*models.Reaction) error
	DeletePostLike(*models.Reaction) error
	CheckPostLike(*models.Reaction) (bool, error)
	CheckPostDislike(int, int) (bool, error)
	DeletePostDislike(int, int) error
	UpdatePostReaction(int) error
	LikeComment(*models.Reaction) error
	DeleteCommentLike(*models.Reaction) error
	CheckCommentLike(*models.Reaction) (bool, error)
	CheckCommentDislike(int, int) (bool, error)
	DeleteCommentDislike(int, int) error
	UpdateCommentReaction(int) error
}

func CreateLikesRepo(db *sql.DB) *LikesRepo {
	return &LikesRepo{
		db: db,
	}
}

func (r *LikesRepo) LikePost(reaction *models.Reaction) error {
	stmt := `INSERT OR REPLACE INTO likes(
		user_id,
		post_id,
		status)
		VALUES(?, ?, ?)`
	_, err := r.db.Exec(stmt, reaction.User_Id, reaction.PostId, reaction.Status)
	if err != nil {
		fmt.Println("like post repo: %w", err)
		return err
	}

	return nil
}

func (r *LikesRepo) DeletePostLike(reaction *models.Reaction) error {
	stmt, err := r.db.Prepare(`DELETE FROM likes WHERE post_id = ? AND user_id = ?`)
	if err != nil {
		fmt.Println("delete post like repo: %w", err)
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(reaction.PostId, reaction.User_Id)
	if err != nil {
		fmt.Println("delete post like repo stmt exec: %w", err)
		return err
	}
	return nil
}

func (r *LikesRepo) CheckPostLike(reaction *models.Reaction) (bool, error) {
	var exists int

	stmt := `SELECT COUNT(*) FROM likes WHERE post_id = ? AND user_id = ?`
	err := r.db.QueryRow(stmt, reaction.PostId, reaction.User_Id).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println("check post like: %w", err)
		return false, err
	}

	return exists > 0, nil
}

func (r *LikesRepo) CheckPostDislike(postId, userId int) (bool, error) {
	var exists int

	stmt := `SELECT COUNT(*) FROM dislikes WHERE post_id = ? AND user_id = ?`
	err := r.db.QueryRow(stmt, postId, userId).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println("check post like: %w", err)
		return false, err
	}

	return exists > 0, nil
}

func (r *LikesRepo) DeletePostDislike(postId, userId int) error {
	stmt, err := r.db.Prepare(`DELETE FROM dislikes WHERE post_id = ? AND user_id = ?`)
	if err != nil {
		fmt.Println("delete post dislike: %w", err)
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(postId, userId)
	if err != nil {
		return err
	}

	return nil
}

func (r *LikesRepo) UpdatePostReaction(postId int) error {
	stmt := `SELECT COUNT(post_id) FROM likes WHERE post_id = ? AND status = 1`
	row := r.db.QueryRow(stmt, postId, 1)
	var likesCounter int
	err := row.Scan(&likesCounter)
	if err != nil {
		fmt.Println("update post reaction first stmt: %w", err)
		return err
	}

	stmt = `SELECT COUNT(post_id) FROM dislikes WHERE post_id = ? AND status = 1`
	row = r.db.QueryRow(stmt, postId, 1)
	var dislikesCounter int
	err = row.Scan(&dislikesCounter)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println("update post reaction second stmt: %w", err)
		return err
	}

	stmt = `UPDATE post SET like = ? WHERE id = ?`
	_, err = r.db.Exec(stmt, likesCounter, postId)
	if err != nil {
		fmt.Println("update post reaction third stmt: %w", err)
		return err
	}

	stmt = `UPDATE post SET dislike = ? WHERE id = ?`
	_, err = r.db.Exec(stmt, dislikesCounter, postId)
	if err != nil {
		fmt.Println("update post reaction third stmt: %w", err)
		return err
	}

	return nil
}

////////////LIKE COMMENTS//////////////////////

func (r *LikesRepo) LikeComment(reaction *models.Reaction) error {
	stmt := `INSERT OR REPLACE INTO likes(
		user_id,
		comment_id,
		status)
		VALUES(?, ?, ?)`
	_, err := r.db.Exec(stmt, reaction.User_Id, reaction.CommentId, reaction.Status)
	if err != nil {
		fmt.Println("like post repo: %w", err)
		return err
	}

	return nil
}

func (r *LikesRepo) DeleteCommentLike(reaction *models.Reaction) error {
	stmt, err := r.db.Prepare(`DELETE FROM likes WHERE comment_id = ? AND user_id = ?`)
	if err != nil {
		fmt.Println("delete post like repo: %w", err)
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(reaction.CommentId, reaction.User_Id)
	if err != nil {
		fmt.Println("delete post like repo stmt exec: %w", err)
		return err
	}
	return nil
}

func (r *LikesRepo) CheckCommentLike(reaction *models.Reaction) (bool, error) {
	var exists int

	stmt := `SELECT COUNT(*) FROM likes WHERE comment_id = ? AND user_id = ?`
	err := r.db.QueryRow(stmt, reaction.CommentId, reaction.User_Id).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println("check post like: %w", err)
		return false, err
	}

	return exists > 0, nil
}

func (r *LikesRepo) CheckCommentDislike(commentId, userId int) (bool, error) {
	var exists int

	stmt := `SELECT COUNT(*) FROM dislikes WHERE comment_id = ? AND user_id = ?`
	err := r.db.QueryRow(stmt, commentId, userId).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println("check post like: %w", err)
		return false, err
	}

	return exists > 0, nil
}

func (r *LikesRepo) DeleteCommentDislike(commentId, userId int) error {
	stmt, err := r.db.Prepare(`DELETE FROM dislikes WHERE comment_id = ? AND user_id = ?`)
	if err != nil {
		fmt.Println("delete post dislike: %w", err)
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(commentId, userId)
	if err != nil {
		return err
	}

	return nil
}

func (r *LikesRepo) UpdateCommentReaction(commentId int) error {
	stmt := `SELECT COUNT(comment_id) FROM likes WHERE comment_id = ? AND status = 1`
	row := r.db.QueryRow(stmt, commentId, 1)
	var likesCounter int
	err := row.Scan(&likesCounter)
	if err != nil {
		fmt.Println("update post reaction first stmt: %w", err)
		return err
	}

	fmt.Printf("like counter %d\n", likesCounter)

	stmt = `SELECT COUNT(comment_id) FROM dislikes WHERE comment_id = ? AND status = 1`
	row = r.db.QueryRow(stmt, commentId, 1)
	var dislikesCounter int
	err = row.Scan(&dislikesCounter)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println("update post reaction second stmt: %w", err)
		return err
	}

	fmt.Printf("dislikes counter %d\n", dislikesCounter)

	stmt = `UPDATE comment SET likes = ? WHERE id = ?`
	_, err = r.db.Exec(stmt, likesCounter, commentId)
	if err != nil {
		fmt.Println("update post reaction third stmt: %w", err)
		return err
	}

	stmt = `UPDATE comment SET dislikes = ? WHERE id = ?`
	_, err = r.db.Exec(stmt, dislikesCounter, commentId)
	if err != nil {
		fmt.Println("update post reaction third stmt: %w", err)
		return err
	}

	return nil
}
