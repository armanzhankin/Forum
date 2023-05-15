package handler

import (
	"errors"
	"fmt"
	"net/http"

	"forum/models"
)

func (h *Handler) categoriesFilter(w http.ResponseWriter, r *http.Request) {
	var user models.User

	user = r.Context().Value("user").(models.User)

	if r.Method != http.MethodGet {
		h.ErrorHandler(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err := r.ParseForm()
	if err != nil {
		h.ErrorHandler(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	category := r.FormValue("category")
	fmt.Println(category)
	posts, err := h.service.GetPostsByCategories(category)
	if errors.Is(err, models.ErrNoPost) {
		homeForm := models.HomeForm{
			User: user,
			Post: nil,
			Err:  models.ErrNoPost.Error(),
		}
		h.Render(w, r, homeForm, "")
		return
	}

	fmt.Println(posts)

	homeForm := models.HomeForm{
		User: user,
		Post: posts,
	}

	h.Render(w, r, homeForm, "./ui/templates/index.html")
}
