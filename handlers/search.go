package handlers

import (
	"github.com/Matemateg/blog/entities"
	"github.com/Matemateg/blog/middlewares"
	"github.com/Matemateg/blog/service"
	"html/template"
	"net/http"
)

type pageSearch struct {
	service *service.UserService
	tpl     *template.Template
}

func NewPageSearch(service *service.UserService) *pageSearch {
	tpl := template.Must(template.ParseFiles("templates/basePage.gohtml", "templates/search.html"))
	return &pageSearch{service: service, tpl: tpl}
}

func (h *pageSearch) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//	get the parameter q from r (search string)
	searchKey := r.FormValue("q")
	//	from service get users
	users, err := h.service.SearchUser(searchKey)
		if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	currentUser := middlewares.GetCurrentUser(r.Context())

	pageData := struct {
		CurrentUser *entities.User
		FoundUsers []entities.User
	}{
		CurrentUser: currentUser,
		FoundUsers: users,
	}
	//	put in template
	err = h.tpl.Execute(w, pageData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}