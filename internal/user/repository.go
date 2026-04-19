package user

import (
	"context"
	"errors"
	"fmt"
	"migtationbot/internal/app"

	sq "github.com/Masterminds/squirrel"
	trmpgx "github.com/avito-tech/go-transaction-manager/drivers/pgxv5/v2"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
)

var (
	userTableName  = "users"
	userTgIDColumn = "tg_id"
	userRoleColumn = "role"
	userTgUsername = "tg_username"
)

type RepositoryImpl struct {
	db        *pgxpool.Pool
	ctxGetter *trmpgx.CtxGetter
}

func NewUserRepository(db *pgxpool.Pool) Repository {
	return &RepositoryImpl{db: db, ctxGetter: trmpgx.DefaultCtxGetter}
}

func (r *RepositoryImpl) Create(ctx context.Context, user *User) error {
	const op = "UserRepositoryImpl.Create"

	builder := psql.
		Insert(userTableName).
		Columns(userTgIDColumn, userRoleColumn, userTgUsername).
		Values(user.TelegramID, user.Role, user.TelegramUsername)

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	_, err = r.ctxGetter.DefaultTrOrDB(ctx, r.db).Exec(ctx, query, args...)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return app.ErrUserAlreadyExists
		}
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
func (r *RepositoryImpl) Get(ctx context.Context, id int64) (*User, error) {
	const op = "UserRepositoryImpl.Get"

	builder := psql.
		Select(userTgIDColumn, userRoleColumn, userTgUsername).
		From(userTableName).
		Where(sq.Eq{userTgIDColumn: id})
	query, args, err := builder.ToSql()

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var user *User
	err = r.ctxGetter.DefaultTrOrDB(ctx, r.db).QueryRow(ctx, query, args...).
		Scan(
			&user.TelegramID,
			&user.Role,
			&user.TelegramUsername,
		)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, pgx.ErrNoRows
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}

func (r *RepositoryImpl) UpdateRole(ctx context.Context, tgUsername string, newRole string) error {
	const op = "UserRepositoryImpl.UpdateRole"

	builder := psql.
		Update(userTableName).
		Set(userRoleColumn, newRole).
		Where(sq.Eq{userTgUsername: tgUsername})
	query, args, err := builder.ToSql()

	if err != nil {
		return err
	}
	_, err = r.ctxGetter.DefaultTrOrDB(ctx, r.db).Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
