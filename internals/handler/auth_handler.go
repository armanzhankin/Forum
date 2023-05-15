package handler

import (
	"errors"
	"fmt"
	"net/http"

	"forum/models"

	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		err := h.Render(w, r, nil, "./ui/templates/signup.html")
		if err != nil {
			h.ErrorHandler(w, http.StatusText(http.StatusInternalServerError), http.StatusNotFound)
			return
		}
	case http.MethodPost:
		err := r.ParseForm()
		if err != nil {
			h.ErrorHandler(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		newUser := &models.User{
			Username: r.FormValue("username"),
			Email:    r.FormValue("email"),
			Password: r.FormValue("password"),
		}

		newForm := models.SignUpForm{
			Errs: "",
		}

		conf := r.FormValue("confirmation")
		if newUser.Password != conf {
			newForm.Errs = models.ErrPasswordConf.Error()
			// w.WriteHeader(http.StatusBadRequest)
			err = h.Render(w, r, newForm, "./ui/templates/signup.html")
			if err != nil {
				h.ErrorHandler(w, http.StatusText(http.StatusInternalServerError), http.StatusNotFound)
				return
			}
			return
		}

		err = h.service.User.ValidUser(newUser)
		if err != nil {
			fmt.Println(err)
			newForm.Errs = models.ErrInvalidData.Error()
			// w.WriteHeader(http.StatusBadRequest)
			h.Render(w, r, newForm, "./ui/templates/signup.html")

			return
		}

		err = h.service.User.CreateUser(newUser)
		if errors.Is(err, models.ErrDuplicate) {
			newForm.Errs = models.ErrDuplicate.Error()
			err = h.Render(w, r, newForm, "./ui/templates/signup.html")
			if err != nil {
				h.ErrorHandler(w, http.StatusText(http.StatusInternalServerError), http.StatusNotFound)
				return
			}
			return
		} else if err != nil {
			h.ErrorHandler(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/login", http.StatusSeeOther)
	default:
		h.ErrorHandler(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}

func (h *Handler) logIn(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		err := h.Render(w, r, nil, "./ui/templates/login.html")
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

		newForm := models.SignUpForm{
			Errs: "",
		}

		username := r.FormValue("login")
		password := r.FormValue("password")
		newUser, err := h.service.User.CheckUser(username, password)
		if errors.Is(err, models.ErrInvalidData) || errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			newForm.Errs = models.ErrInvalidData.Error()
			h.Render(w, r, newForm, "./ui/templates/login.html")
			return
		} else if err != nil {
			fmt.Println(err)
			h.ErrorHandler(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   newUser.Token,
			Path:    "/",
			Expires: newUser.TokenDuration,
		})

		newUser.Auth = 1

		http.Redirect(w, r, "/", http.StatusSeeOther)
	default:
		h.ErrorHandler(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}
