package db

import (
	"github.com/Matemateg/blog/entities"
	"github.com/jmoiron/sqlx"
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
