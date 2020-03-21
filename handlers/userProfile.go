package handlers

import (
	"fmt"
	"github.com/Matemateg/blog/service"
	"net/http"
)

type UserProfile struct {
	service *service.UserService
}

func NewUserProfile(service *service.UserService) *UserProfile {
	return &UserProfile{service: service}
}

func (u *UserProfile) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	profile := u.service.GetUserProfile(222)
	fmt.Fprint(w, profile.User)
}

