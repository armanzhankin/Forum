package handler

import (
	"fmt"
	"net/http"

	"forum/models"
)

func (h *Handler) MyPosts(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(models.User)
	if user.Id == 0 {
		h.ErrorHandler(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	switch r.Method {
	case http.MethodGet:

		Posts, err := h.service.GetMyPosts(user)
		if err != nil {
			h.ErrorHandler(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		MyPostForm := models.HomeForm{
			User: user,
			Post: Posts,
		}
		fmt.Println(Posts)
		err = h.Render(w, r, MyPostForm, "./ui/templates/myposts.html")
		if err != nil {
			h.ErrorHandler(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	default:
		h.ErrorHandler(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)

		return

	}
}
