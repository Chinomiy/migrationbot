package bookmark

import (
	"context"
	"fmt"
	"migtationbot/logger"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
)

var (
	UserFavoriteTableName        = "user_favorite"
	UserIDColumnName             = "user_id"
	CountryTripContentColumnName = "country_trip_content_id"
)

type BookmarkRepositoryImpl struct {
	db *pgxpool.Pool
}

func NewBookmarkRepository(db *pgxpool.Pool) BookmarkRepository {
	return &BookmarkRepositoryImpl{db: db}
}

func (r *BookmarkRepositoryImpl) GetUserFavorite(ctx context.Context, userID int64) ([]UserFavorite, error) {
	const op = "BookmarkRepository.GetUserFavorite"
	builder := psql.
		Select(
			"c.code",
			"c.name AS country_name",
			"t.callback",
			"t.name AS trip_name").
		From(UserFavoriteTableName + " AS uf").
		LeftJoin("country_trip_content AS ctc  ON uf." + CountryTripContentColumnName + "= ctc.id").
		LeftJoin("country AS c ON c.id = ctc.country_id").
		LeftJoin("trip_type AS t ON t.id = ctc.trip_type_id").
		Where(sq.Eq{"uf." + UserIDColumnName: userID})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	var userFavorites []UserFavorite
	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()
	for rows.Next() {
		var userFavorite UserFavorite
		if err = rows.Scan(
			&userFavorite.CountryCode,
			&userFavorite.CountryName,
			&userFavorite.TripCallback,
			&userFavorite.TripType,
		); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		userFavorites = append(userFavorites, userFavorite)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return userFavorites, nil
}

func (r *BookmarkRepositoryImpl) AddFavorite(ctx context.Context, userID int64, code, trip string) error {
	const op = "BookmarkRepository.AddFavorite"
	builder := psql.
		Insert(UserFavoriteTableName).
		Columns(UserIDColumnName, CountryTripContentColumnName).
		Select(
			psql.
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

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	logger.Infof("added to favorite: userId:%d code:%s trip:%s", userID, code, trip)
	return nil
}

func (r *BookmarkRepositoryImpl) RemoveFavorite(ctx context.Context, userID int64, code, trip string) error {
	const op = "BookmarkRepository.RemoveFavorite"
	builder := psql.
		Delete(UserFavoriteTableName).
		Where(sq.Eq{UserIDColumnName: userID}).
		Where(
			sq.Expr(CountryTripContentColumnName+" = (?)",
				sq.Select("ctc.id").
					From("country_trip_content AS ctc").
					Join("country AS c ON c.id = ctc.country_id").
					Join("trip_type AS t ON t.id = ctc.trip_type_id").
					Where(sq.Eq{
						"c.code":     code,
						"t.callback": trip,
					},
					),
			),
		)
	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	if _, err := r.db.Exec(ctx, query, args...); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
