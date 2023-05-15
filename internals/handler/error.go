package handler

import (
	"fmt"
	"net/http"
	"text/template"

	"forum/models"
)

func (h *Handler) ErrorHandler(w http.ResponseWriter, message string, status int) {
	errPage := models.ErrForm{Status: status, Text: http.StatusText(status), Message: message}
	tmpl, err := template.ParseFiles("./ui/templates/error.html")
	if err != nil {
		fmt.Println("dsds")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	if err := tmpl.Execute(w, errPage); err != nil {
		fmt.Println("ssss")

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
