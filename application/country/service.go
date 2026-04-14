package country

import (
	"context"
)

type CountryService struct {
	repo *CountryRepository
	//trManager *manager.Manager
}

func NewCountryService(repo *CountryRepository) *CountryService {
	return &CountryService{repo: repo}
}

func (s *CountryService) FindByCode(ctx context.Context, code string) (*Country, error) {

	country, err := s.repo.GetByCode(ctx, code)
	if err != nil {
		return nil, err
	}
	return country, nil
}

func (s *CountryService) List(ctx context.Context) (*[]Country, error) {
	countries, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}
	return countries, nil
}

func (s *CountryService) GetCountryWithTrip(ctx context.Context, code string) (*Country, error) {
	country, err := s.repo.GetByCode(ctx, code)
	if err != nil {
		return nil, err
	}
	trip, err := s.repo.GetCountryTrip(ctx, code)
	if err != nil {
		return nil, err
	}
	country.TripTypes = trip
	return country, nil
}

func (s *CountryService) GetAllTrips(ctx context.Context) (TripType, error) {
	tt, err := s.repo.GetAllTrip(ctx)

	if err != nil {
		return TripType{}, err
	}
	return tt, nil
}

func (s *CountryService) GetCountryContentByTrip(ctx context.Context, code, callback string) (string, error) {
	content, err := s.repo.GetContentByCallback(ctx, code, callback)
	if err != nil {
		return "", err
	}
	return content, nil
}
