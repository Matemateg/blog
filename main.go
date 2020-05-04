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
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mysqlDSN := os.Getenv("MYSQL_DSN")
	if mysqlDSN == "" {
		mysqlDSN = "root:123@/blog?parseTime=true"
	}

	sqlDB, err := sqlx.Connect("mysql", mysqlDSN)
	if err != nil {
		log.Fatalln(err)
	}
	err = sqlDB.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	userProfileSrv := service.NewUserProfile(db.NewUser(sqlDB), db.NewPost(sqlDB))
	postSrv := service.NewPostService(db.NewUser(sqlDB), db.NewPost(sqlDB))

	http.Handle("/", mw.Auth(handlers.NewMainPage(postSrv), userProfileSrv))

	http.Handle("/user/", mw.Auth(handlers.NewUserProfile(userProfileSrv), userProfileSrv))

	http.Handle("/registration/", registration.NewPageRegistration())
	http.Handle("/signup/", registration.NewUserSignup(userProfileSrv))

	http.Handle("/login/", login.NewPageLogin())
	http.Handle("/auth/", login.NewUserAuth(userProfileSrv))

	http.Handle("/logout/", handlers.NewPageLogout())

	http.Handle("/addPost/", mw.Auth(handlers.NewNewPost(userProfileSrv), userProfileSrv))

	http.Handle("/search", mw.Auth(handlers.NewPageSearch(userProfileSrv), userProfileSrv))

	wd, _ := os.Getwd()
	log.Printf("working in directory: %s", wd)
	log.Printf("starting linsen on port: %s", port)
	fmt.Println(http.ListenAndServe(":"+port, nil))
}
