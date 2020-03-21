package service

import (
	"github.com/Matemateg/blog/db"
	"github.com/Matemateg/blog/entities"
)

type UserService struct {
	userDB *db.User
	postDB *db.Post
}

func NewUserProfile(userDB *db.User, postDB *db.Post) *UserService {
	return &UserService{userDB: userDB, postDB: postDB}
}

type UserProfileData struct {
	User *entities.User
	Posts []entities.Post
}

func (us *UserService) GetUserProfile(id int64) *UserProfileData {
	return &UserProfileData{
		User:  us.userDB.GetByID(id),
		Posts: us.postDB.ListByUserID(id),
	}
}