package login

import (
	"github.com/Matemateg/blog/middlewares"
	"html/template"
	"net/http"
)

type pageLogin struct {
	tpl *template.Template
}

func NewPageLogin() *pageLogin {
	tpl := template.Must(template.ParseFiles("templates/basePage.gohtml", "templates/login.html"))
	return &pageLogin{tpl: tpl}
}

func (u *pageLogin) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	currentUser := middlewares.GetCurrentUser(r.Context())
	if currentUser != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	err := u.tpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
