package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/directory"
)

type countryUsecase struct {
	countryRepository domain.CountryRepository
	contextTimeout    time.Duration
}

func NewCountryUsecase(countryRepository domain.CountryRepository, timeout time.Duration) domain.CountryUsecase {
	return &countryUsecase{
		countryRepository: countryRepository,
		contextTimeout:    timeout,
	}
}

func (tu *countryUsecase) Create(c context.Context, country *domain.Country) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.countryRepository.Create(ctx, country)
}

func (tu *countryUsecase) Update(c context.Context, country *domain.Country) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.countryRepository.Update(ctx, country)
}

func (tu *countryUsecase) Delete(c context.Context, country *domain.Country) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.countryRepository.Delete(ctx, country)
}

func (lu *countryUsecase) FetchByID(c context.Context, countryID string) (domain.Country, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.countryRepository.FetchByID(ctx, countryID)
}

func (lu *countryUsecase) Fetch(c context.Context) ([]domain.Country, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.countryRepository.Fetch(ctx)
}
