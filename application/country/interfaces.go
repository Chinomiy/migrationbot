package country

import "context"

type CountryRepository interface {
	GetByCode(ctx context.Context, code string) (*Country, error)
	GetCountryTrip(ctx context.Context, code string) (TripType, error)
	List(ctx context.Context) (*[]Country, error)
	GetAllTrip(ctx context.Context) (TripType, error)
	GetContentByCallback(ctx context.Context, code, callback string) (string, error)
}

type CountryService interface {
	FindByCode(ctx context.Context, code string) (*Country, error)
	List(ctx context.Context) (*[]Country, error)
	GetCountryWithTrip(ctx context.Context, code string) (*Country, error)
	GetAllTrips(ctx context.Context) (TripType, error)
	GetCountryContentByTrip(ctx context.Context, code, callback string) (string, error)
}
