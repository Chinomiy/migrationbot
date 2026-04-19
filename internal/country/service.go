package country

import (
	"context"
	"errors"
	"migtationbot/internal/app"

	"github.com/avito-tech/go-transaction-manager/trm/v2/manager"
)

type ServiceImpl struct {
	repo      Repository
	trManager *manager.Manager
}

func NewCountryService(repo Repository, trManager *manager.Manager) Service {
	return &ServiceImpl{repo: repo, trManager: trManager}
}

func (s *ServiceImpl) List(ctx context.Context) (*[]Country, error) {
	countries, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}
	return countries, nil
}

func (s *ServiceImpl) GetCountryWithTrip(ctx context.Context, code string) (*Country, error) {
	var country *Country
	existing, err := s.repo.GetCountryByCode(ctx, code)
	if err != nil {
		return nil, err
	}
	trip, err := s.repo.GetCountryTrip(ctx, code)
	if err != nil {
		return nil, err
	}
	country = &Country{
		ID:          existing.ID,
		Code:        existing.Code,
		Name:        existing.Name,
		Description: existing.Description,
		TripTypes:   trip,
	}
	return country, nil
}

func (s *ServiceImpl) GetAllTrips(ctx context.Context) (TripType, error) {
	tt, err := s.repo.GetAllTrip(ctx)

	if err != nil {
		return TripType{}, err
	}
	return tt, nil
}

func (s *ServiceImpl) GetCountryContentByTrip(ctx context.Context, code, callback string) (string, error) {
	content, err := s.repo.GetContentByCallback(ctx, code, callback)
	if err != nil {
		return "", err
	}
	return content, nil
}

func (s *ServiceImpl) CreateCountry(ctx context.Context, code, name, desc string) error {
	if code == "" || name == "" || desc == "" {
		return app.ErrOneOfRequiredFieldsEmpty
	}
	err := s.trManager.Do(ctx, func(ctx context.Context) error {
		_, err := s.repo.GetCountryByCode(ctx, code)
		if err == nil {
			return app.ErrCountryWithGivenCodeAlreadyExists
		}
		newCountry := Country{
			Code:        code,
			Name:        name,
			Description: desc,
		}
		err = s.repo.CreateCountry(ctx, &newCountry)
		if err != nil {
			return err
		}
		return nil
	})

	return err
}
func (s *ServiceImpl) CreateTrip(ctx context.Context, name, callback string) error {
	if name == "" || callback == "" {
		return app.ErrOneOfRequiredFieldsEmpty
	}
	err := s.trManager.Do(ctx, func(ctx context.Context) error {
		_, err := s.repo.GetTripByCallback(ctx, callback)
		if err == nil {
			return app.ErrTripWithGivenCodeAlreadyExists
		}
		err = s.repo.CreateTrip(ctx, name, callback)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (s *ServiceImpl) SetCountryTrip(ctx context.Context, code, cb string) error {
	if code == "" || cb == "" {
		return app.ErrOneOfRequiredFieldsEmpty
	}
	err := s.trManager.Do(ctx, func(ctx context.Context) error {
		country, err := s.repo.GetCountryByCode(ctx, code)
		if err != nil {
			if errors.Is(err, app.ErrCountryNotFound) {
				return app.ErrCountryNotFound
			}
			return err
		}
		trip, err := s.repo.GetTripByCallback(ctx, cb)
		if err != nil {
			if errors.Is(err, app.ErrTripWithGivenCallbackNotFound) {
				return app.ErrTripWithGivenCallbackNotFound
			}
			return err
		}
		err = s.repo.SetCountryTripType(ctx, country.ID, trip.Id)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}
func (s *ServiceImpl) SetCountryContent(ctx context.Context, countryCode, callback, content string) error {
	if countryCode == "" || callback == "" || content == "" {
		return app.ErrOneOfRequiredFieldsEmpty
	}
	err := s.trManager.Do(ctx, func(ctx context.Context) error {
		country, err := s.repo.GetCountryByCode(ctx, countryCode)
		if err != nil {
			if errors.Is(err, app.ErrCountryNotFound) {
				return app.ErrCountryNotFound
			}
			return err
		}
		trip, err := s.repo.GetTripByCallback(ctx, callback)
		if err != nil {
			if errors.Is(err, app.ErrTripWithGivenCallbackNotFound) {
				return app.ErrTripWithGivenCallbackNotFound
			}
			return err
		}
		err = s.repo.SetCountryTripContent(ctx, country.ID, trip.Id, content)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (s *ServiceImpl) PublishCountry(ctx context.Context, countryCode string) error {
	if countryCode == "" {
		return app.ErrOneOfRequiredFieldsEmpty
	}
	err := s.trManager.Do(ctx, func(ctx context.Context) error {
		country, err := s.repo.GetCountryByCode(ctx, countryCode)
		if err != nil {
			if errors.Is(err, app.ErrCountryNotFound) {
				return app.ErrCountryNotFound
			}
			return err
		}
		err = s.repo.PublishCountry(ctx, country.ID)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

/*
На подумать: стране добавить флаг published:
смысл в том что когда создали страну, то она еще не заполнена контентом, пользователь видит пустую страну - такого быть не должно

1. Создали страну
2. Опционально. создали тип или взяли из имеющихсся
3. Добавили стране типы
4. Добавили текст типам
5. Опубликовали
func setcountrytriptype (

	trmanager.do() {
{
*/
