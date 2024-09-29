package user

import (
	"github.com/gocql/gocql"
)

// User Storage
type userStore struct {
	db_session *gocql.Session
}

func NewUserStore(db_session *gocql.Session) *userStore {
	return &userStore{db_session: db_session}
}
