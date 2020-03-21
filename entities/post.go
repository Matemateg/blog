package entities

import "time"

type Post struct {
	ID int64
	UserID int64
	Text string
	CreatedAt time.Time
}
