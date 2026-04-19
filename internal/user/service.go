package user

import (
	"context"
	"errors"
	"migtationbot/internal/app"
	"migtationbot/logger"

	"github.com/avito-tech/go-transaction-manager/trm/v2/manager"
)

type ServiceImpl struct {
	UserRepo  Repository
	trManager *manager.Manager
}

func NewUserService(userRepository Repository, trManager *manager.Manager) Service {
	return &ServiceImpl{
		UserRepo:  userRepository,
		trManager: trManager,
	}
}

func (u *ServiceImpl) GetOrCreateUser(ctx context.Context, id int64, tgUsername string) (*User, error) {
	user := &User{
		TelegramID:       id,
		Role:             RoleUser,
		TelegramUsername: tgUsername,
	}
	err := u.trManager.Do(ctx, func(context.Context) error {
		err := u.UserRepo.Create(ctx, user)
		if err == nil {
			return nil
		}
		if errors.Is(err, app.ErrUserAlreadyExists) {
			return nil
		}
		user, err = u.UserRepo.Get(ctx, id)
		logger.Info(user.TelegramUsername)
		if err != nil {
			return err
		}
		return nil
	})
	return user, err
}

func (u *ServiceImpl) UpdateUserRole(ctx context.Context, id int64, role string) error {
	if role == "" {
		return app.ErrEmptyGivenRole
	}

	err := u.trManager.Do(ctx, func(context.Context) error {
		user, err := u.UserRepo.Get(ctx, id)
		if err != nil {
			return err
		}
		err = u.UserRepo.UpdateRole(ctx, user.TelegramUsername, role)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}
