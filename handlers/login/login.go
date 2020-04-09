package login

import (
	"html/template"
	"net/http"
)

type pageLogin struct {
	tpl     *template.Template
}

func NewPageLogin() *pageLogin {
	tpl := template.Must(template.ParseFiles("templates/login.html"))
	return &pageLogin{tpl: tpl}
}

func (u *pageLogin) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := u.tpl.Execute(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}