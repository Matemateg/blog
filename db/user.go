package db

import (
	"fmt"
	"github.com/Matemateg/blog/entities"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type User struct {
	db *sqlx.DB
}

func NewUser(db *sqlx.DB) *User {
	return &User{db: db}
}

func (u *User) GetByID(id int64) (*entities.User, error) {
	row := u.db.QueryRowx("SELECT * FROM users WHERE id = ?", id)
	var p entities.User
	if err := row.StructScan(&p); err != nil {
		return nil, fmt.Errorf("getting user from db, %v", err)
	}
	return &p, nil
}

func (u *User) GetByLogin(login, password string) (*entities.User, error) {
	row := u.db.QueryRowx("SELECT * FROM users WHERE login = ? AND password = ?", login, password)
	var p entities.User
	if err := row.StructScan(&p); err != nil {
		return nil, fmt.Errorf("getting user from db with login, password, %v", err)
	}
	return &p, nil
}