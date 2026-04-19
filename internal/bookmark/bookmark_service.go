package bookmark

import (
	"context"
	"fmt"
	"migtationbot/internal/country"
	"migtationbot/internal/user"
)

type ServiceImpl struct {
	BookMarkRepo Repository

	UserSVC    user.Service
	CountrySVC country.Service
}

func NewBookMarkService(
	userSvc user.Service,
	countrySvc country.Service,
	BookmarkRepo Repository,
) Service {
	return &ServiceImpl{
		UserSVC:      userSvc,
		CountrySVC:   countrySvc,
		BookMarkRepo: BookmarkRepo,
	}
}

// TODO ПЕРЕДВАТЬ В РЕПОЗИТОРИЙ УЖЕ ГОТЫЙ ОБЬЕКТ А НЕ СОБИРАТЬ ЕГО В РЕПОЗИТОРИИ + ДОБАВИТЬ БИЗНЕСОВЫЕ ПРОВЕРКИ ЕСЛИ ЕСТЬ
func (s *ServiceImpl) GetUserFavorites(ctx context.Context, userID int64) ([]UserFavorite, error) {
	userFavorite, err := s.BookMarkRepo.GetUserFavorite(ctx, userID)

	if err != nil {
		return nil, fmt.Errorf("failed to get user favorite: %w", err)
	}

	return userFavorite, nil
}

func (s *ServiceImpl) AddFavorite(ctx context.Context, userID int64, code, trip string) error {

	err := s.BookMarkRepo.AddFavorite(ctx, userID, code, trip)
	if err != nil {
		return err
	}

	return nil
}

func (s *ServiceImpl) RemoveFavorite(ctx context.Context, userID int64, code, trip string) error {
	err := s.BookMarkRepo.RemoveFavorite(ctx, userID, code, trip)
	if err != nil {
		return err
	}
	return nil
}
