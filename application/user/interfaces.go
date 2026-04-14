package user

import "context"

type UserService interface {
	CreateUser(ctx context.Context, id int64) (*User, error)
	GetUser(ctx context.Context, id int64) (*User, error)
	UpdateUserRole(ctx context.Context, id int64, role string) error
	ExistsUser(ctx context.Context, id int64) (bool, error)
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *User) error
	UpdateUserRole(ctx context.Context, id int64, role string) error
	GetUser(ctx context.Context, id int64) (*User, error)
}
