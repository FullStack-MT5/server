package postgresql

import (
	"database/sql"

	"github.com/benchttp/server/benchttp"
)

// Ensure service implements interface.
var _ benchttp.UserService = (*UserService)(nil)

type UserService struct {
	db *sql.DB
}

func NewUserService(conn Connection) UserService {
	return UserService{conn.db}
}

func (s UserService) Create(name, email string) (string, error) {
	return "1", nil
}

func (s UserService) Retrieve(id string) (benchttp.User, error) {
	return benchttp.User{}, nil
}

func (s UserService) Exists(name, email string) bool {
	return false
}
