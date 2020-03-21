package entities

import "time"

type Post struct {
	ID        int64
	UserID    int64 `db:"user_id"`
	Text      string
	CreatedAt time.Time `db:"created_at"`
}
