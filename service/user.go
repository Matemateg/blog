package service

import (
	"fmt"
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
	User  *entities.User
	Posts []entities.Post
}

func (us *UserService) GetUserProfile(id int64) (*UserProfileData, error) {
	user, err := us.userDB.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("getting users, %v", err)
	}
	posts, err := us.postDB.ListByUserID(id)
	if err != nil {
		return nil, fmt.Errorf("getting posts, %v", err)
	}
	return &UserProfileData{
		User:  user,
		Posts: posts,
	}, nil
}
