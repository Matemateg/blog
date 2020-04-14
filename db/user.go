package db

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Matemateg/blog/entities"
	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
	"log"
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

func (u *User) getByLogin(login string) (*entities.User, error) {
	row := u.db.QueryRowx("SELECT * FROM users WHERE login = ?", login)
	var p entities.User
	if err := row.StructScan(&p); err != nil {
		return nil, fmt.Errorf("getting user from db with login, %v", err)
	}
	return &p, nil
}

func (u *User) GetByLoginPassword(login, password string) (*entities.User, error) {
	user, err := u.getByLogin(login)
	if err != nil {
		return nil, fmt.Errorf("getting user from db with login, %v", err)
	}
	checkPass := comparePasswords(user.Password, password)
	if !checkPass {
		return nil, fmt.Errorf("uncorrect password, %v", err)
	}
	return user, nil
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

func (u *User) RegWithLoginPass(name, login, password string) error {
	ssid := uuid.New().String()
	passwordHashString, err := hashWithSalt(password)
	if err != nil {
		return fmt.Errorf("hashing password, %v", err)
	}
	_, err = u.db.Exec("INSERT INTO users (name, login, password, session_id) VALUES (?, ?, ?, ?)", name, login, passwordHashString, ssid)
	if err == nil {
		return nil
	}

	if me, ok := err.(*mysql.MySQLError); ok && me.Number == 1062 {
		return errors.New("User already exists in a database.")
	}

	return fmt.Errorf("insert user into db with name, login, password, %v", err)
}
func hashWithSalt(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func comparePasswords(hashedPwd string, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	bytePlainPwd := []byte(plainPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePlainPwd)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (u *User) SearchUsers(searchString string) ([]entities.User, error) {
	rows, err := u.db.Queryx("SELECT * FROM users WHERE MATCH (login,name) AGAINST (?)", searchString)
	if err != nil {
		return nil, err
	}

	var users []entities.User
	for rows.Next() {
		var user entities.User
		if err = rows.StructScan(&user); err != nil {
			return nil, fmt.Errorf("search user from db, %v", err)
		}
		users = append(users, user)
	}
	return users, nil
}