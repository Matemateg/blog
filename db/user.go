package db

import "github.com/Matemateg/blog/entities"

type User struct {
}

func (u *User) GetByID(id int64) *entities.User {
	return &entities.User{
		ID:   id,
		Name: "Pes",
	}
}
