package handlers

import (
	"github.com/Matemateg/blog/entities"
	"github.com/Matemateg/blog/middlewares"
	"github.com/Matemateg/blog/service"
	"html/template"
	"net/http"
)

type mainPage struct {
	service *service.PostService
	tpl     *template.Template
}

func NewMainPage(service *service.PostService) *mainPage {
	tpl := template.Must(template.ParseFiles("templates/basePage.gohtml", "templates/mainPage.html"))
	return &mainPage{service: service, tpl: tpl}
}

func (h *mainPage) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	posts, err := h.service.ReturnLastNPosts(10)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	currentUser := middlewares.GetCurrentUser(r.Context())

	pageData := struct {
		CurrentUser  *entities.User
		UserPostList []service.PostUser
	}{
		CurrentUser:  currentUser,
		UserPostList: posts,
	}

	err = h.tpl.Execute(w, pageData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}