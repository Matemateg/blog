package service

import (
	"fmt"
	"github.com/Matemateg/blog/db"
	"github.com/Matemateg/blog/entities"
)

type PostService struct {
	userDB *db.User
	postDB *db.Post
}

func NewPostService(userDB *db.User, postDB *db.Post) *PostService {
	return &PostService{userDB: userDB, postDB: postDB}
}

type PostUser struct {
	Post entities.Post
	User entities.User
}

func (ps *PostService) ReturnLastNPosts(n int) ([]PostUser, error) {
	posts, err := ps.postDB.ReturnLastNPosts(n)
	if err != nil {
		return nil, fmt.Errorf("returnlast n posts, %v", err)
	}

	res := []PostUser{}
	for _, entry := range posts {
		res = append(res, PostUser(entry))
	}

	return res, nil
}
