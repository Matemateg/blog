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
	user, err := us.userDB.GetByLoginPassword(login, password)
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

func (us *UserService) NewPost(userID int64, text string) error {
	err := us.postDB.AddPost(userID, text)
	if err != nil {
		return fmt.Errorf("adding post, %v", err)
	}
	return nil
}

func (us *UserService) Registration(name, login, password string) (*entities.User, error) {
	if name == "" && login == "" && password == "" {
		return nil, fmt.Errorf("name, login or password are too short")
	}
	err := us.userDB.RegWithLoginPass(name, login, password)
	if err != nil {
		return nil, fmt.Errorf("register user, %v", err)
	}
	user, err := us.userDB.GetByLoginPassword(login, password)
	if err != nil {
		return nil, fmt.Errorf("getting user, %v", err)
	}
	return user, nil
}

func (us *UserService) SearchUser(searchString string) ([]entities.User, error) {
	users, err := us.userDB.SearchUsers(searchString)
	if err != nil {
		return nil, fmt.Errorf("search user, %v", err)
	}
	return users, nil
}