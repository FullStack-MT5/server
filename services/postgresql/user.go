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
	user := benchttp.User{Name: name, Email: email}
	stmt := `
INSERT INTO public.users(name, email)
VALUES ($1, $2)`[1:]
	result, err := s.db.Exec(stmt, name, email)
	if err != nil {
		return user, ErrExecutingPreparedStmt
	}
	id, err := result.LastInsertId()
	if err != nil {
		return user, ErrGettingIDInsertion
	}
	user.ID = id
	return user, nil
}

func (s UserService) GetByEmail(email string) (benchttp.User, error) {
	user := benchttp.User{Email: email}
	stmt := `
SELECT id, name
FROM public.users
WHERE email = $1`[1:]
	row := s.db.QueryRow(stmt, email)
	err := row.Scan(
		&user.ID,
		&user.Name)
	if err != nil {
		return user, ErrScanningRows
	}
	return user, nil
}

func (s UserService) Exists(email string) bool {
	var exists bool
	stmt := `
SELECT EXISTS(
SELECT email
FROM public.users WHERE email=$1)`[1:]
	err := s.db.QueryRow(stmt, email).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return false
	}
	return exists
}
