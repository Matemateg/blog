package handlers

import (
	"github.com/Matemateg/blog/middlewares"
	"github.com/Matemateg/blog/service"
	"net/http"
)

type newPost struct {
	service *service.UserService
}

func NewNewPost(service *service.UserService) *newPost {
	return &newPost{service: service}
}

func (h *newPost) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	newPostText := r.PostFormValue("text")

	currentUser := middlewares.GetCurrentUser(r.Context())
	//var currentUser *entities.User
	//cookie, err := r.Cookie("SSID")
	//if err == nil && cookie != nil {
	//	currentUser, err = h.service.GetBySSID(cookie.Value)
	//	if err != nil {
	//		http.Error(w, err.Error(), http.StatusInternalServerError)
	//		return
	//	}
	//}
	if currentUser == nil {
		http.Error(w, "You are unauthorized", http.StatusUnauthorized)
		return
	}

	err := h.service.NewPost(currentUser.ID, newPostText)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, r.Referer(), http.StatusFound)
}

