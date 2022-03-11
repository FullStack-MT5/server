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

func (s UserService) Create(name, email string) (benchttp.User, error) {
	return benchttp.User{}, nil
}

func (s UserService) GetByID(id string) (benchttp.User, error) {
	return benchttp.User{}, nil
}

func (s UserService) GetByCred(name, email string) (benchttp.User, error) {
	return benchttp.User{}, nil
}

func (s UserService) Exists(name, email string) bool {
	return false
}
