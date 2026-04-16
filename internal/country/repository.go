package country

import (
	"context"
	"errors"
	"fmt"
	"migtationbot/internal/app"

	sq "github.com/Masterminds/squirrel"
	trmpgx "github.com/avito-tech/go-transaction-manager/drivers/pgxv5/v2"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
)

// country
const (
	countryTableName       = "country"
	countryColumnID        = "id"
	countryColumnCode      = "code"
	countryColumnName      = "name"
	countryColumnDesc      = "description"
	countryColumnPublished = "published"
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
	countryTripTypeColumnID        = "trip_type_id"
)
const (
	countryTripContentTableName        = "country_trip_content"
	countryTripContentColumnContent    = "content"
	countryTripContentColumnCountryID  = "country_id"
	countryTripContentColumnTripTypeID = "trip_type_id"
)

type CountryRepositoryImpl struct {
	db        *pgxpool.Pool
	ctxGetter *trmpgx.CtxGetter
}

func NewCountryRepository(db *pgxpool.Pool) CountryRepository {
	return &CountryRepositoryImpl{db: db, ctxGetter: trmpgx.DefaultCtxGetter}
}

func (c *CountryRepositoryImpl) GetCountryByCode(ctx context.Context, code string) (*Country, error) {
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

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	var country Country
	err = c.ctxGetter.DefaultTrOrDB(ctx, c.db).QueryRow(ctx, query, args...).Scan(
		&country.ID,
		&country.Code,
		&country.Description,
		&country.Name,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, app.ErrCountryNotFound
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &country, nil
}

func (c *CountryRepositoryImpl) GetCountryTrip(ctx context.Context, code string) (TripType, error) {
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
	rows, err := c.ctxGetter.DefaultTrOrDB(ctx, c.db).Query(ctx, query, args...)

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
func (c *CountryRepositoryImpl) List(ctx context.Context) (*[]Country, error) {
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

func (c *CountryRepositoryImpl) GetAllTrip(ctx context.Context) (TripType, error) {
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

func (c *CountryRepositoryImpl) GetTripByCallback(ctx context.Context, callback string) (TripType, error) {
	const op = "CountryRepository.GetTripByCallback"
	builder := psql.
		Select(tripTypesColumnID, tripTypeColumnName, tripTypeColumnCallback).
		From(tripTypesTableName).
		Where(sq.Eq{tripTypeColumnCallback: callback})
	query, args, err := builder.ToSql()
	if err != nil {
		return TripType{}, fmt.Errorf("%s: %w", op, err)
	}
	tripType := TripType{
		Data: make(map[string]string),
	}
	var name string
	var cb string
	var id int
	err = c.ctxGetter.DefaultTrOrDB(ctx, c.db).QueryRow(ctx, query, args...).Scan(&id, &name, &cb)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return TripType{}, app.ErrTripWithGivenCallbackNotFound
		}
		return TripType{}, fmt.Errorf("%s: %w", op, err)
	}
	tripType.Id = id
	tripType.Data[cb] = name
	return tripType, nil
}

func (c *CountryRepositoryImpl) GetContentByCallback(ctx context.Context, code, callback string) (string, error) {
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
			return "", app.ErrContentNotFound
		}
		return "", fmt.Errorf("%s: %w", op, err)
	}
	return content, nil
}

func (c *CountryRepositoryImpl) CreateCountry(ctx context.Context, country *Country) error {
	const op = "CountryRepository.CreateCountry"

	builder := psql.
		Insert(countryTableName).
		Columns(countryColumnName, countryColumnCode, countryColumnDesc).
		Values(country.Name, country.Code, country.Description)
	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	_, err = c.ctxGetter.DefaultTrOrDB(ctx, c.db).Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (c *CountryRepositoryImpl) CreateTrip(ctx context.Context, name, callback string) error {
	const op = "CountryRepository.CreateTrip"

	builder := psql.
		Insert(tripTypesTableName).
		Columns(tripTypeColumnName, tripTypeColumnCallback).
		Values(name, callback)
	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	_, err = c.db.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (c *CountryRepositoryImpl) SetCountryTripType(ctx context.Context, countryId, tripTypeId int) error {
	const op = "CountryRepository.SetTripToCountry"
	builder := psql.
		Insert(countryTripTypeTableName).
		Columns(countryTripTypeCountryColumnID, countryTripTypeColumnID).
		Values(tripTypeId, countryId)
	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	_, err = c.ctxGetter.DefaultTrOrDB(ctx, c.db).Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (c *CountryRepositoryImpl) SetCountryTripContent(ctx context.Context, countryId, tripId int, content string) error {
	const op = "CountryRepository.SetCountryContent"

	builder := psql.
		Insert(countryTripContentTableName).
		Columns(countryTripContentColumnContent, countryTripContentColumnCountryID, countryTripContentColumnTripTypeID).
		Values(content, countryId, tripId)
	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	_, err = c.ctxGetter.DefaultTrOrDB(ctx, c.db).Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
func (c *CountryRepositoryImpl) PublishCountry(ctx context.Context, countryID int) error {
	const op = "CountryRepository.PublishCountry"
	builder := psql.
		Update(countryTableName).
		Set(countryColumnPublished, true).
		Where(sq.Eq{"id": countryID})
	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	_, err = c.ctxGetter.DefaultTrOrDB(ctx, c.db).Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
