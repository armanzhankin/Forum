package models

type Comment struct {
	Id       int
	User_Id  int
	Author   string
	Content  string
	Likes    int
	Dislikes int
	PostId   int
}
