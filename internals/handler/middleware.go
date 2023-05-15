package handler

import (
	"context"
	"net/http"
	"time"

	"forum/models"
)

func (h *Handler) requireAuthenticatedUser(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("token")
		if err != nil {
			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "user", models.User{
				Id: 0,
			})))
			return
		}
		user, err := h.service.User.GetUserByToken(c.Value)
		if err != nil {
			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "user", models.User{Id: 0})))
			return
		}
		if user.TokenDuration.Before(time.Now()) {
			if err := h.service.User.DeleteToken(c.Value); err != nil {
				h.ErrorHandler(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "user", user)))
	}
}
