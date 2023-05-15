package handler

import (
	"forum/models"
	"net/http"
)

func (h *Handler) LikedPosts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.ErrorHandler(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	var user models.User

	user = r.Context().Value("user").(models.User)
	if user.Id == 0 {
		h.ErrorHandler(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	posts, err := h.service.GetLikedPosts(user)
	if err != nil {
		h.ErrorHandler(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	form := models.HomeForm{
		Post: posts,
		User: user,
	}

	if err := h.Render(w, r, form, "./ui/templates/likedPosts.html"); err != nil {
		h.ErrorHandler(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
