package bookmark

import "context"

type (
	BookmarkService interface {
		GetUserFavorites(ctx context.Context, userID int64) ([]UserFavorite, error)
		AddFavorite(ctx context.Context, userID int64, code, trip string) error
		RemoveFavorite(ctx context.Context, userID int64, code, trip string) error
	}
	BookmarkRepository interface {
		GetUserFavorite(ctx context.Context, userID int64) ([]UserFavorite, error)
		AddFavorite(ctx context.Context, userID int64, code, trip string) error
		RemoveFavorite(ctx context.Context, userID int64, code, trip string) error
	}
)
