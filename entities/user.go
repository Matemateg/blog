package entities

type User struct {
	ID   int64
	Name string
	Login string
	Password string
	SessionID string `db:"session_id"`
}
