package handler

import (
	"fmt"
	"net/http"
	"text/template"
)

func (h *Handler) Render(w http.ResponseWriter, r *http.Request, data interface{}, path string) error {
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		fmt.Println(err)
		h.ErrorHandler(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return err
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		fmt.Println(err)
		h.ErrorHandler(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return err
	}
	return nil
}
