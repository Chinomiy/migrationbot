package bookmark

import "context"

type (
	Service interface {
		GetUserFavorites(ctx context.Context, userID int64) ([]UserFavorite, error)
		AddFavorite(ctx context.Context, userID int64, code, trip string) error
		RemoveFavorite(ctx context.Context, userID int64, code, trip string) error
	}
	Repository interface {
		GetUserFavorite(ctx context.Context, userID int64) ([]UserFavorite, error)
		AddFavorite(ctx context.Context, userID int64, code, trip string) error
		RemoveFavorite(ctx context.Context, userID int64, code, trip string) error
	}
)
