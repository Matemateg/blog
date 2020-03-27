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
		return nil, fmt.Errorf("getting user's profile, %v", err)
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

func (us *UserService) Login(login, password string) (*entities.User, error) {
	user, err := us.userDB.GetByLogin(login, password)
	if err != nil {
		return nil, fmt.Errorf("getting user, %v", err)
	}
	return user, nil
}

func (us *UserService) GetBySSID(ssid string) (*entities.User, error) {
	user, err := us.userDB.GetBySSID(ssid)
	if err != nil {
		return nil, fmt.Errorf("getting user, %v", err)
	}
	return user, nil
}