package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPost_ReturnLastNPosts(t *testing.T) {
	db, err := sqlx.Open("mysql", "root:123@/blog?parseTime=true")
	require.NoError(t, err)
	postDB := NewPost(db)

	userpostList, err := postDB.ReturnLastNPosts(30)
	require.NoError(t, err)
	fmt.Println(userpostList)
}
