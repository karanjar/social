package store

import (
	"context"
	"database/sql"
	"social/internal/models/users"
)

type UsersStore struct {
	db *sql.DB
}

func (s *UsersStore) Create(ctx context.Context, user *users.User) error {
	query := `INSERT INTO users (username,password,email) VALUES($1,$2,$3) RETURNING id,created_at`

	err := s.db.QueryRowContext(
		ctx,
		query,
		user.Username,
		user.Email,
		user.Password,
	).Scan(
		&user.ID,
		&user.CreatedAt,
	)
	if err != nil {
		return err
	}
	return nil

}
