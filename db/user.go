package db

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Matemateg/blog/entities"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
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
		return nil, fmt.Errorf("getting user from db with id, %v", err)
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

func (u *User) GetBySSID(ssid string) (*entities.User, error) {
	row := u.db.QueryRowx("SELECT * FROM users WHERE session_id = ?", ssid)
	var p entities.User
	if err := row.StructScan(&p); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("getting user from db with ssid, %v", err)
	}
	return &p, nil
}

func (u *User) RegWithLoginPass(name, login, password string) (error) {
	ssid := uuid.New().String()
	_, err := u.db.Exec("INSERT INTO users (name, login, password, session_id) VALUES (?, ?, ?, ?)", name, login, password, ssid)

	if err == nil {
		return  nil
	}

	if me, ok := err.(*mysql.MySQLError); ok && me.Number == 1062 {
			return errors.New("User already exists in a database.")
	}

	return fmt.Errorf("insert user into db with name, login, password, %v", err)
}