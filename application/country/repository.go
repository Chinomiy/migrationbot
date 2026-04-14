package country

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

	errCountryNotFound = errors.New("country not found")
)

// country
const (
	countryTableName  = "country"
	countryColumnID   = "id"
	countryColumnCode = "code"
	countryColumnName = "name"
	countryColumnDesc = "description"
)

// trip_type
const (
	tripTypesTableName     = "trip_type"
	tripTypesColumnID      = "id"
	tripTypeColumnName     = "name"
	tripTypeColumnCallback = "callback"
)

// country_trip_type
const (
	countryTripTypeTableName       = "country_trip_type"
	countryTripTypeCountryColumnID = "country_id"
	cuntryTripTypeColumnID         = "trip_type_id"
)
const (
	countryTripContentTableName = "country_trip_content"
)

type CountryRepository struct {
	db *pgxpool.Pool
	//ctxGetter *trmpgx.CtxGetter
}

func NewCountryRepository(db *pgxpool.Pool) *CountryRepository {
	return &CountryRepository{db: db}
}
func (c *CountryRepository) GetByCode(ctx context.Context, code string) (*Country, error) {
	return c.getByCode(ctx, code, false)
}
func (c *CountryRepository) GetByCodeUpdate(ctx context.Context, code string) (*Country, error) {
	return c.getByCode(ctx, code, true)
}
func (c *CountryRepository) getByCode(ctx context.Context, code string, forUpdate bool) (*Country, error) {
	const op = "CountryRepository.getByCode"
	builder := psql.
		Select(
			"c.id",
			"c.code",
			"c.description",
			"c.name",
		).
		From(countryTableName + " AS c").
		Where(sq.Eq{"c.code": code})

	if forUpdate {
		builder = builder.Suffix("FOR UPDATE")
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	var country Country
	//r.ctxGetter.DefaultTrOrD
	err = c.db.QueryRow(ctx, query, args...).Scan(
		&country.ID,
		&country.Code,
		&country.Description,
		&country.Name,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errCountryNotFound
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &country, nil
}
func (c *CountryRepository) GetCountryTrip(ctx context.Context, code string) (TripType, error) {
	const op = "CountryRepository.GetCountryTrip"
	builder := psql.
		Select(
			"tt.name",
			"tt.callback",
		).
		From(countryTripTypeTableName + " AS ctp").
		InnerJoin(tripTypesTableName + " AS tt ON tt.id = ctp.trip_type_id").
		InnerJoin(countryTableName + " AS c ON c.id = ctp.country_id").
		Where(sq.Eq{"c.code": code}).
		OrderBy("tt.name ASC")
	query, args, err := builder.ToSql()
	if err != nil {
		return TripType{}, fmt.Errorf("%s: %w", op, err)
	}
	var tripType TripType
	trip := make(map[string]string)
	rows, err := c.db.Query(ctx, query, args...)
	if err != nil {
		return TripType{}, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		var callback string
		if err := rows.Scan(&name, &callback); err != nil {
			return TripType{}, fmt.Errorf("%s: %w", op, err)
		}
		trip[callback] = name
	}
	if err := rows.Err(); err != nil {
		return TripType{}, fmt.Errorf("%s: %w", op, err)
	}
	tripType.Data = trip
	return tripType, nil
}
func (c *CountryRepository) List(ctx context.Context) (*[]Country, error) {
	const op = "CountryRepository.List"
	builder := psql.
		Select(
			countryColumnID,
			countryColumnCode,
			countryColumnName,
			countryColumnDesc,
		).
		From(countryTableName).
		OrderBy(countryColumnID + " ASC")

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	rows, err := c.db.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()
	var countries []Country
	for rows.Next() {
		var country Country
		if err = rows.Scan(
			&country.ID,
			&country.Code,
			&country.Name,
			&country.Description,
		); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		countries = append(countries, country)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &countries, nil
}

func (c *CountryRepository) GetAllTrip(ctx context.Context) (TripType, error) {
	const op = "CountryRepository.GetAllTrip"
	builder := psql.
		Select(
			tripTypeColumnCallback,
			tripTypeColumnName).
		From(tripTypesTableName)
	query, args, err := builder.ToSql()
	if err != nil {
		return TripType{}, fmt.Errorf("%s: %w", op, err)
	}
	var tripType TripType
	tt := make(map[string]string)
	rows, err := c.db.Query(ctx, query, args...)
	if err != nil {
		return TripType{}, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		var callback string
		if err := rows.Scan(&name, &callback); err != nil {
			return TripType{}, fmt.Errorf("%s: %w", op, err)
		}
		tt[callback] = name
	}
	if err := rows.Err(); err != nil {
		return TripType{}, fmt.Errorf("%s: %w", op, err)
	}
	tripType.Data = tt

	return tripType, nil
}

func (c *CountryRepository) GetContentByCallback(ctx context.Context, code, callback string) (string, error) {
	const op = "CountryRepository.GetContentByCallback"
	builder := psql.
		Select(
			"ctc.content").
		From(countryTripContentTableName + " AS ctc").
		LeftJoin("country AS c ON c.id = ctc.country_id").
		LeftJoin("trip_type AS tt ON tt.id = ctc.trip_type_id").
		Where(sq.Eq{"tt.callback": callback, "c.code": code})
	query, args, err := builder.ToSql()
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	var content string
	err = c.db.QueryRow(ctx, query, args...).Scan(&content)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", fmt.Errorf("%s: %w", op, err) // добавить контент не найден
		}
	}
	return content, nil
}

func (c *CountryRepository) GetContentByCodeAndCallback(ctx context.Context, code, callback string) (string, error) {
	const op = "CountryRepository.GetContentByCodeAndCallback"
	builder := psql.
		Select("")
}
