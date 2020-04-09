package main

import (
	"fmt"
	"github.com/Matemateg/blog/db"
	"github.com/Matemateg/blog/handlers"
	"github.com/Matemateg/blog/handlers/login"
	"github.com/Matemateg/blog/handlers/registration"
	mw "github.com/Matemateg/blog/middlewares"
	"github.com/Matemateg/blog/service"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
)

func main() {
	sqlDB, err := sqlx.Connect("mysql", "root:123@/blog?parseTime=true")
	if err != nil {
		log.Fatalln(err)
	}
	err = sqlDB.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	userProfileSrv := service.NewUserProfile(db.NewUser(sqlDB), db.NewPost(sqlDB))
	http.Handle("/user/", mw.Auth(handlers.NewUserProfile(userProfileSrv), userProfileSrv))

	http.Handle("/registration/", registration.NewPageRegistration())
	http.Handle("/signup/", registration.NewUserSignup(userProfileSrv))

	http.Handle("/login/", login.NewPageLogin())
	http.Handle("/auth/", login.NewUserAuth(userProfileSrv))

	http.Handle("/logout/", handlers.NewPageLogout())

	http.Handle("/addPost/", mw.Auth(handlers.NewNewPost(userProfileSrv), userProfileSrv))

	fmt.Println(http.ListenAndServe(":8080", nil))
}
