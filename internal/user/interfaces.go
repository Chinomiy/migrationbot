package user

import "context"

type Service interface {
	//GetUser(ctx context.Context, id int64) (*User, error)
	UpdateUserRole(ctx context.Context, id int64, role string) error
	// GetOrCreateUser ExistsUser(ctx context.Context, id int64) (bool, error)
	GetOrCreateUser(ctx context.Context, id int64, tgUsername string) (*User, error)
}

type Repository interface {
	Create(ctx context.Context, user *User) error
	UpdateRole(ctx context.Context, tgUsername string, role string) error
	Get(ctx context.Context, id int64) (*User, error)
}
