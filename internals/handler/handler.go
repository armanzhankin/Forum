package handler

import (
	"net/http"

	"forum/internals/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", h.requireAuthenticatedUser(h.home))
	mux.HandleFunc("/create", h.requireAuthenticatedUser(h.createPost))
	mux.HandleFunc("/post/", h.requireAuthenticatedUser(h.showPost))
	mux.HandleFunc("/like-post", h.requireAuthenticatedUser(h.postLike))
	mux.HandleFunc("/dislike-post", h.requireAuthenticatedUser(h.postDislike))
	mux.HandleFunc("/like-comment", h.requireAuthenticatedUser(h.commentLike))
	mux.HandleFunc("/dislike-comment", h.requireAuthenticatedUser(h.commentDislike))
	mux.HandleFunc("/categories", h.requireAuthenticatedUser(h.categoriesFilter))
	mux.HandleFunc("/signup", h.signUp)
	mux.HandleFunc("/login", h.logIn)
	mux.HandleFunc("/myposts", h.requireAuthenticatedUser(h.MyPosts))
	mux.HandleFunc("/logout", h.logOut)
	mux.HandleFunc("/likedpost", h.requireAuthenticatedUser(h.LikedPosts))
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	return mux
}
