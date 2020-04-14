package handlers

import (
	"github.com/Matemateg/blog/entities"
	"github.com/Matemateg/blog/middlewares"
	"github.com/Matemateg/blog/service"
	"html/template"
	"net/http"
	"strconv"
)

type userProfile struct {
	service *service.UserService
	tpl     *template.Template
}

func NewUserProfile(service *service.UserService) *userProfile {
	tpl := template.Must(template.ParseFiles("templates/basePage.gohtml", "templates/userProfile.html"))
	return &userProfile{service: service, tpl: tpl}
}

func (u *userProfile) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	currentUser := middlewares.GetCurrentUser(r.Context())
	//var currentUser *entities.User
	//cookie, err := r.Cookie("SSID")
	//if err == nil && cookie != nil {
	//	currentUser, err = u.service.GetBySSID(cookie.Value)
	//	if err != nil {
	//		http.Error(w, err.Error(), http.StatusInternalServerError)
	//		return
	//	}
	//}

	pageData := struct {
		CurrentUser *entities.User
		UserProfile *service.UserProfileData
		CanWrite bool
	}{
		CurrentUser: currentUser,
		UserProfile: profile,
		CanWrite: currentUser != nil && currentUser.ID == profile.User.ID,
	}

	err = u.tpl.Execute(w, pageData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
