package models

type Post struct {
	Id       int
	User_Id  int
	Author   string
	Category []string
	Title    string
	Content  string
	Like     int
	Dislike  int
}
