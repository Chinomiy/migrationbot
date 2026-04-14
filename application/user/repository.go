package user

import (
	"context"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
)

var (
	userTableName  = "user"
	userTgIDColumn = "tg_id"
	userRoleColumn = "role"
)

type UserRepositoryImpl struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) CreateUser(ctx context.Context, user *User) error {
	const op = "UserRepositoryImpl.CreateUser"
	builder := psql.
		Insert(userTableName).
		Columns(userTgIDColumn, userRoleColumn).
		Values(userTgIDColumn, user.Role)

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
func (r *UserRepositoryImpl) GetUser(ctx context.Context, id int64) (*User, error) {
	const op = "UserRepositoryImpl.GetUser"
	builder := psql.
		Select(userTgIDColumn, userRoleColumn).
		From(userTableName).
		Where(sq.Eq{userTgIDColumn: id})
	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	var user User
	err = r.db.QueryRow(ctx, query, args...).Scan(
		&user.TelegramID,
		&user.Role,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, pgx.ErrNoRows
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &user, nil
}

func (r *UserRepositoryImpl) UpdateUserRole(ctx context.Context, id int64, newRole string) error {
	const op = "UserRepositoryImpl.UpdateUserRole"
	builder := psql.
		Update(userTableName).
		Set(userRoleColumn, newRole).
		Where(sq.Eq{userTgIDColumn: id})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
