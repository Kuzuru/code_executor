package models

import "context"

type Storage interface {
	CreateUser(ctx context.Context, snippet *User) error
	FindUserByName(ctx context.Context, name string) (*User, error)
	FindUserByID(ctx context.Context, id int) (*User, error)
	GetUsers(ctx context.Context, limit int64) error
}
