package db

import (
	"github.com/Matemateg/blog/entities"
	"time"
)

type Post struct {

}

func(p *Post) ListByUserID(userID int64) []entities.Post {
	return []entities.Post{
		{
			ID:        1,
			UserID:    22,
			Text:      "The sun is down",
			CreatedAt: time.Now(),
		},
		{
			ID:        2,
			UserID:    22,
			Text:      "The sun rose",
			CreatedAt: time.Now(),
		},
	}
}