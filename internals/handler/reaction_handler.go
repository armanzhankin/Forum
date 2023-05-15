package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"forum/models"
)

func (h *Handler) postLike(w http.ResponseWriter, r *http.Request) {
	var user models.User

	user = r.Context().Value("user").(models.User)
	if user.Id == 0 {
		h.ErrorHandler(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	if r.Method != http.MethodPost {
		h.ErrorHandler(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		h.ErrorHandler(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(r.FormValue("value"))
	if err != nil {
		h.ErrorHandler(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	newLike := &models.Reaction{
		PostId:  id,
		User_Id: user.Id,
		Status:  1,
	}

	err = h.service.Likes.LikePost(newLike)
	if err != nil {
		h.ErrorHandler(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
}

func (h *Handler) postDislike(w http.ResponseWriter, r *http.Request) {
	var user models.User

	user = r.Context().Value("user").(models.User)

	if user.Id == 0 {
		h.ErrorHandler(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	if r.Method != http.MethodPost {
		h.ErrorHandler(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		h.ErrorHandler(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(r.FormValue("value"))
	if err != nil {
		h.ErrorHandler(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	newDislike := &models.Reaction{
		PostId:  id,
		User_Id: user.Id,
		Status:  1,
	}

	err = h.service.Dislikes.DislikePost(newDislike)
	if err != nil {
		h.ErrorHandler(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
}

func (h *Handler) commentLike(w http.ResponseWriter, r *http.Request) {
	var user models.User

	user = r.Context().Value("user").(models.User)
	if user.Id == 0 {
		h.ErrorHandler(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	if r.Method != http.MethodPost {
		h.ErrorHandler(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		h.ErrorHandler(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(r.FormValue("value"))
	fmt.Println(id)
	if err != nil {
		h.ErrorHandler(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	newLike := &models.Reaction{
		CommentId: id,
		User_Id:   user.Id,
		Status:    1,
	}

	err = h.service.Likes.LikeComment(newLike)
	if err != nil {
		h.ErrorHandler(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
}

func (h *Handler) commentDislike(w http.ResponseWriter, r *http.Request) {
	var user models.User

	user = r.Context().Value("user").(models.User)

	if user.Id == 0 {
		h.ErrorHandler(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	if r.Method != http.MethodPost {
		h.ErrorHandler(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		h.ErrorHandler(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(r.FormValue("value"))
	if err != nil {
		h.ErrorHandler(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	newDislike := &models.Reaction{
		CommentId: id,
		User_Id:   user.Id,
		Status:    1,
	}

	err = h.service.Dislikes.DislikeComment(newDislike)
	if err != nil {
		h.ErrorHandler(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
}
