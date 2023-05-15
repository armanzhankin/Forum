package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"forum/models"
)

func (h *Handler) createPost(w http.ResponseWriter, r *http.Request) {
	var user models.User

	user = r.Context().Value("user").(models.User)

	if user.Id == 0 {
		h.ErrorHandler(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	switch r.Method {
	case http.MethodGet:
		err := h.Render(w, r, nil, "./ui/templates/createPost.html")
		if err != nil {
			h.ErrorHandler(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	case http.MethodPost:

		err := r.ParseForm()
		if err != nil {
			h.ErrorHandler(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		category, ok := r.Form["category"]

		if !ok {
			h.ErrorHandler(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		newPost := &models.Post{
			User_Id:  user.Id,
			Category: category,
			Title:    r.FormValue("title"),
			Content:  r.FormValue("content"),
		}
		newPost, err = h.service.ValidPost(newPost)
		if err != nil {
			newForm := models.ErrInvalidPost.Error()
			h.Render(w, r, newForm, "./ui/templates/createPost.html")
			return
		}
		id, err := h.service.CreatePost(newPost)
		newPost.Id = id
		if err != nil {
			fmt.Println(err)
			h.ErrorHandler(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, fmt.Sprintf("/post/%v", id), http.StatusSeeOther)
	default:
		h.ErrorHandler(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}

func (h *Handler) showPost(w http.ResponseWriter, r *http.Request) {
	var user models.User

	user = r.Context().Value("user").(models.User)
	url := strings.Split(r.URL.Path, "/")
	if r.URL.Path != "/post/"+string(url[2]) {
		h.ErrorHandler(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	id, err := strconv.Atoi(url[2])
	if err != nil || id < 1 {
		h.ErrorHandler(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	m, err := h.service.GetPostByID(id)
	if err == models.ErrNoRecord {
		h.ErrorHandler(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	} else if err != nil {
		h.ErrorHandler(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	author, err := h.service.User.GetUsernameById(m.User_Id)
	if err != nil {
		h.ErrorHandler(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	m.Author = author
	data := &models.PostForm{
		User:     user,
		Post:     m,
		Comments: []*models.Comment{},
	}

	switch r.Method {
	case http.MethodGet:

		comments, err := h.service.GetCommentsByPostId(m.Id)
		// fmt.Printf("likes of comment: %d\n", comments[0].Likes)
		// fmt.Printf("likes of comment: %d\n", comments[0].Dislikes)
		if err != nil {
			h.ErrorHandler(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		data.Comments = comments

		if err != nil {
			h.ErrorHandler(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		if user.Id == 0 {
			h.ErrorHandler(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		err := r.ParseForm()
		if err != nil {
			h.ErrorHandler(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		newComment := &models.Comment{
			Content: r.FormValue("comment"),
			PostId:  id,
			Author:  user.Username,
		}

		newComment, err = h.service.ValidComment(newComment)

		if err != nil {

			h.ErrorHandler(w, models.ErrInvalidComment.Error(), http.StatusBadRequest)
			return
		}

		err = h.service.CreateComment(newComment)
		if err != nil {
			h.ErrorHandler(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, r.URL.Path, http.StatusSeeOther)
	default:
		h.ErrorHandler(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	if err := h.Render(w, r, data, "./ui/templates/showPost.html"); err != nil {
		h.ErrorHandler(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
