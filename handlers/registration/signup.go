package registration

import (
	"fmt"
	"github.com/Matemateg/blog/service"
	"net/http"
)

type UserSignup struct {
	service *service.UserService
}

func NewUserSignup(service *service.UserService) *UserSignup {
	return &UserSignup{service: service}
}

func (h *UserSignup) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userName := r.PostFormValue("name")
	userLogin := r.PostFormValue("login")
	userPassword := r.PostFormValue("password")

	user, err := h.service.Registration(userName, userLogin, userPassword)
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