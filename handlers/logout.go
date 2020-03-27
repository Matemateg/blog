package handlers

import (
	"net/http"
)

type pageLogout struct {
}

func NewPageLogout() *pageLogout {
	return &pageLogout{}
}

func (u *pageLogout) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name: "SSID",
		Path: "/",
		MaxAge: -1,
		HttpOnly: true,
	})


	http.Redirect(w, r, r.Referer(), http.StatusFound)
}
