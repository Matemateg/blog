package login

import (
	"fmt"
	"github.com/Matemateg/blog/service"
	"net/http"
)

type UserAuth struct {
	service *service.UserService
}

func NewUserAuth(service *service.UserService) *UserAuth {
	return &UserAuth{service: service}
}

func (u *UserAuth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userLogin := r.PostFormValue("login")
	userPassword := r.PostFormValue("password")

	user, err := u.service.Login(userLogin, userPassword)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	url := fmt.Sprintf("/user/?id=%d", user.ID)

	http.SetCookie(w, &http.Cookie{
		Name: "SSID",
		Value: user.SessionID,
		Path: "/",
		//Expires: time.Now().Add(time.Minute * 3),
		MaxAge: 60*5,
		HttpOnly: true,
	})

	http.Redirect(w, r, url, http.StatusFound)
}