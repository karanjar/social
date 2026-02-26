package store

import (
	"context"
	"database/sql"
	"social/internal/models/posts"
	"social/internal/models/users"
)

type Storage struct {
	Post interface {
		Create(ctx context.Context, post *posts.Post) error
	}
	Users interface {
		Create(ctx context.Context, user *users.User) error
	}
}

func NewPostgresStorage(db *sql.DB) Storage {
	return Storage{
		Post:  &PostsStore{db},
		Users: &UsersStore{db},
	}

}
