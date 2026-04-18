package bookmark

import (
	"context"
	"fmt"
	"migtationbot/internal/country"
	"migtationbot/internal/user"
)

type BookmarkServiceImpl struct {
	BookMarkRepo BookmarkRepository

	UserSVC    user.UserService
	CountrySVC country.CountryService
}

func NewBookMarkService(
	userSvc user.UserService,
	countrySvc country.CountryService,
	BookmarkRepo BookmarkRepository,
) BookmarkService {
	return &BookmarkServiceImpl{
		UserSVC:      userSvc,
		CountrySVC:   countrySvc,
		BookMarkRepo: BookmarkRepo,
	}
}

func (s *BookmarkServiceImpl) GetUserFavorites(ctx context.Context, userID int64) ([]UserFavorite, error) {
	userFavorite, err := s.BookMarkRepo.GetUserFavorite(ctx, userID)

	if err != nil {
		return nil, fmt.Errorf("failed to get user favorite: %w", err)
	}

	return userFavorite, nil
}

func (s *BookmarkServiceImpl) AddFavorite(ctx context.Context, userID int64, code, trip string) error {

	err := s.BookMarkRepo.AddFavorite(ctx, userID, code, trip)
	if err != nil {
		return err
	}

	return nil
}

func (s *BookmarkServiceImpl) RemoveFavorite(ctx context.Context, userID int64, code, trip string) error {
	err := s.BookMarkRepo.RemoveFavorite(ctx, userID, code, trip)
	if err != nil {
		return err
	}
	return nil
}
