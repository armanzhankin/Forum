package handler

import (
	"net/http"

	"forum/models"
)

func (h *Handler) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	var user models.User

	user = r.Context().Value("user").(models.User)

	switch r.Method {
	case http.MethodGet:
		m, err := h.service.Post.AllPosts()
		if err != nil {

			h.ErrorHandler(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		homeForm := models.HomeForm{
			User: user,
			Post: m,
		}

		h.Render(w, r, homeForm, "./ui/templates/index.html")

	default:
		h.ErrorHandler(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}
