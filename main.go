package main

import (
	"fmt"
	"github.com/Matemateg/blog/db"
	"github.com/Matemateg/blog/handlers"
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
	http.Handle("/user/", handlers.NewUserProfile(userProfileSrv))

	http.Handle("/login/", handlers.NewPageLogin())

	http.Handle("/auth/", handlers.NewUserAuth(userProfileSrv))

	fmt.Println(http.ListenAndServe(":8080", nil))
}
