package country

import "context"

type CountryRepository interface {
	GetCountryByCode(ctx context.Context, code string) (*Country, error)
	GetCountryTrip(ctx context.Context, code string) (TripType, error)
	List(ctx context.Context) (*[]Country, error)
	GetAllTrip(ctx context.Context) (TripType, error)
	GetContentByCallback(ctx context.Context, code, callback string) (string, error)
	CreateCountry(ctx context.Context, country *Country) error
	CreateTrip(ctx context.Context, name, callback string) error
	GetTripByCallback(ctx context.Context, callback string) (TripType, error)
	SetCountryTripType(ctx context.Context, countryId, tripTypeId int) error
	SetCountryTripContent(ctx context.Context, countryId, tripId int, content string) error
	PublishCountry(ctx context.Context, countryID int) error
}

type CountryService interface {
	FindByCode(ctx context.Context, code string) (*Country, error)
	List(ctx context.Context) (*[]Country, error)
	GetCountryWithTrip(ctx context.Context, code string) (*Country, error)
	GetAllTrips(ctx context.Context) (TripType, error)
	GetCountryContentByTrip(ctx context.Context, code, callback string) (string, error)
	CreateCountry(ctx context.Context, code, name, desc string) error
	CreateTrip(ctx context.Context, name, callback string) error
	SetCountryTrip(ctx context.Context, code, callback string) error
	SetCountryContent(ctx context.Context, countryCode, callback, content string) error
	PublishCountry(ctx context.Context, countryCode string) error
}
