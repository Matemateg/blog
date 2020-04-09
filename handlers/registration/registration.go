package registration

import (
	"html/template"
	"net/http"
)

type pageRegistration struct {
	tpl     *template.Template
}

func NewPageRegistration() *pageRegistration {
	tpl := template.Must(template.ParseFiles("templates/registration.html"))
	return &pageRegistration{tpl: tpl}
}

func (u *pageRegistration) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := u.tpl.Execute(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}