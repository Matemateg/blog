package handlers

import (
	"github.com/Matemateg/blog/service"
	"html/template"
	"net/http"
	"strconv"
)

type UserProfile struct {
	service *service.UserService
	tpl     *template.Template
}

func NewUserProfile(service *service.UserService) *UserProfile {
	tpl := template.Must(template.ParseFiles("templates/userPage.gohtml"))
	return &UserProfile{service: service, tpl: tpl}
}

func (u *UserProfile) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(r.FormValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "invalid user id: "+err.Error(), http.StatusBadRequest)
		return
	}

	profile, err := u.service.GetUserProfile(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = u.tpl.Execute(w, profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
