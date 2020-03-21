package main

import (
	"fmt"
	"github.com/Matemateg/blog/db"
	"github.com/Matemateg/blog/handlers"
	"github.com/Matemateg/blog/service"
	"net/http"
)

func main() {
	userProfileSrv := service.NewUserProfile(&db.User{}, &db.Post{})
	http.Handle("/user/", handlers.NewUserProfile(userProfileSrv))
	fmt.Println(http.ListenAndServe(":8080", nil))
}
