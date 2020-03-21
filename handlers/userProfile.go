package handlers

import (
	"github.com/Matemateg/blog/service"
	"html/template"
	"net/http"
)

type UserProfile struct {
	service *service.UserService
	tpl *template.Template
}

func NewUserProfile(service *service.UserService) *UserProfile {
	tpl := template.Must(template.ParseFiles("templates/userPage.gohtml"))
	return &UserProfile{service: service, tpl: tpl}
}

func (u *UserProfile) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	profile := u.service.GetUserProfile(222)

	err := u.tpl.Execute(w, profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

