package user

import (
	"context"
	"fmt"
)

type UserServiceImpl struct {
	UserRepo UserRepository
}

func NewUserService(userRepository UserRepository) UserService {
	return &UserServiceImpl{
		UserRepo: userRepository,
	}
}

func (u *UserServiceImpl) CreateUser(ctx context.Context, id int64) (*User, error) {
	if id == 0 {
		return nil, fmt.Errorf("invalid user id")
	}
	user := &User{
		TelegramID: id,
		Role:       "user",
	}
	err := u.UserRepo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserServiceImpl) GetUser(ctx context.Context, id int64) (*User, error) {
	if id == 0 {
		return nil, fmt.Errorf("invalid user id")
	}
	user, err := u.UserRepo.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserServiceImpl) UpdateUserRole(ctx context.Context, id int64, role string) error {
	if role == "" {
		return fmt.Errorf("invalid user role")
	}
	err := u.UserRepo.UpdateUserRole(ctx, id, role)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserServiceImpl) ExistsUser(ctx context.Context, id int64) (bool, error) {
	if id == 0 {
		return false, fmt.Errorf("invalid user id")
	}
	user, err := u.GetUser(ctx, id)
	if err != nil {
		return false, err
	}
	return user != nil, nil
}
