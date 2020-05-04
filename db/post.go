package db

import (
	"fmt"
	"github.com/Matemateg/blog/entities"
	"github.com/jmoiron/sqlx"
	"time"
)

type Post struct {
	db *sqlx.DB
}

func NewPost(db *sqlx.DB) *Post {
	return &Post{db: db}
}

func (p *Post) ListByUserID(userID int64) ([]entities.Post, error) {
	rows, err := p.db.Queryx("SELECT * FROM posts WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}

	var posts []entities.Post
	for rows.Next() {
		var post entities.Post
		if err = rows.StructScan(&post); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (p *Post) AddPost(userID int64, text string) error {
	_, err := p.db.Exec(
		"INSERT INTO posts (text, user_id, created_at)VALUES (?,?,?)",
		text,
		userID,
		time.Now(),
	)
	if err != nil {
		return fmt.Errorf("insert post in DB, %v", err)
	}
	return nil
}

type PostUser struct {
	Post entities.Post
	User entities.User
}

func (p *Post) ReturnLastNPosts(n int) ([]PostUser, error) {
	res := []PostUser{}
	query := "SELECT * FROM posts INNER JOIN users ON posts.user_id = users.id ORDER BY posts.id DESC LIMIT ?"
	rows, err := p.db.Queryx(query, n)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		postUser := PostUser{}
		err = rows.Scan(
			&postUser.Post.ID,  &postUser.Post.Text, &postUser.Post.UserID, &postUser.Post.CreatedAt,
			&postUser.User.ID, &postUser.User.Name, &postUser.User.Login, &postUser.User.Password, &postUser.User.SessionID,
		)
		if err != nil {
			return nil, err
		}
		res = append(res, postUser)
	}

	return res, nil
}
