package registration

import (
	"github.com/Matemateg/blog/middlewares"
	"html/template"
	"net/http"
)

type pageRegistration struct {
	tpl     *template.Template
}

func NewPageRegistration() *pageRegistration {
	tpl := template.Must(template.ParseFiles("templates/basePage.gohtml", "templates/registration.html"))
	return &pageRegistration{tpl: tpl}
}

func (u *pageRegistration) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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