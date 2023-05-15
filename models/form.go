package models

type PostForm struct {
	User     User
	Post     *Post
	Comments []*Comment
	Err      string
}

type HomeForm struct {
	User User
	Post []*Post
	Err  string
}

type SignUpForm struct {
	User User
	Errs string
}

type ErrForm struct {
	Status  int
	Text    string
	Message string
}
