package bookmark

import (
	"context"
	"fmt"
	"migtationbot/application/country"
	"migtationbot/application/user"
)

type BookmarkServiceImpl struct {
	BookMarkRepo BookmarkRepository

	UserSVC    *user.UserServiceImpl
	CountrySVC *country.CountryService
}

func NewBookMarkService(
	userSvc *user.UserServiceImpl,
	countrySvc *country.CountryService,
	BookmarkRepo BookmarkRepository,
) BookmarkService {
	return &BookmarkServiceImpl{
		UserSVC:      userSvc,
		CountrySVC:   countrySvc,
		BookMarkRepo: BookmarkRepo,
	}
}

func (s *BookmarkServiceImpl) GetUserFavorites(ctx context.Context, userID int64) ([]UserFavorite, error) {
	user, err := s.BookMarkRepo.GetUserFavorite(ctx, userID)

	if err != nil {
		return nil, fmt.Errorf("failed to get user favorite: %w", err)
	}
	/*
		Беру юзера
		Беру его добавленные в favorite_table id стран + id trip_type
		Достаю trip_type.callback + country.name , country.code
		Отрисовываю в личном кабинете
	*/
	return user, nil
}

/*
	builder := psql.
		Insert(UserFavoriteTableName).
		Columns(UserIDColumnName, CountryTripContentColumnName).
		Select(
			psql. //  не понимаю как сюда вставить сквирелом ПОПРАВИТЬ
				Select().
				Column(sq.Expr("?", userID)).
				Column("ctc.id").
				From("country_trip_content AS ctc").
				Join("country AS c ON c.id = ctc.country_id").
				Join("trip_type AS t ON t.id = ctc.trip_type_id").
				Where(sq.Eq{
					"c.code":     code,
					"t.callback": trip,
				}),
		).Suffix("ON CONFLICT (user_id, country_trip_content_id) DO NOTHING")
*/
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
