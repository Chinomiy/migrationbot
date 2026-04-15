package country

import (
	"context"
	"migtationbot/application/app"
	"strings"

	"github.com/avito-tech/go-transaction-manager/trm/v2/manager"
)

type CountryServiceImpl struct {
	repo      CountryRepository
	trManager *manager.Manager
}

func NewCountryService(repo CountryRepository, trManager *manager.Manager) CountryService {
	return &CountryServiceImpl{repo: repo, trManager: trManager}
}

func (s *CountryServiceImpl) FindByCode(ctx context.Context, code string) (*Country, error) {
	if code == "" {
		return nil, app.ErrEmptyCountryCode
	}
	strings.ToUpper(code)
	country, err := s.repo.GetByCode(ctx, code)
	if err != nil {
		return nil, err
	}
	return country, nil
}

func (s *CountryServiceImpl) List(ctx context.Context) (*[]Country, error) {
	countries, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}
	return countries, nil
}

func (s *CountryServiceImpl) GetCountryWithTrip(ctx context.Context, code string) (*Country, error) {
	var country *Country

	err := s.trManager.Do(ctx, func(ctx context.Context) error {
		existing, err := s.repo.GetByCode(ctx, code)
		if err != nil {
			return err
		}
		trip, err := s.repo.GetCountryTrip(ctx, code)
		if err != nil {
			return err
		}
		country = &Country{
			ID:          existing.ID,
			Code:        existing.Code,
			Name:        existing.Name,
			Description: existing.Description,
			TripTypes:   trip,
		}
		return nil
	})
	return country, err
}

func (s *CountryServiceImpl) GetAllTrips(ctx context.Context) (TripType, error) {
	tt, err := s.repo.GetAllTrip(ctx)

	if err != nil {
		return TripType{}, err
	}
	return tt, nil
}

func (s *CountryServiceImpl) GetCountryContentByTrip(ctx context.Context, code, callback string) (string, error) {
	content, err := s.repo.GetContentByCallback(ctx, code, callback)
	if err != nil {
		return "", err
	}
	return content, nil
}
