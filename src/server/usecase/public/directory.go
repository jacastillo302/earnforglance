package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/public"
)

type directoryUsecase struct {
	itemRepository domain.DirectoryRepository
	contextTimeout time.Duration
}

func NewdirectoryUsecase(itemRepository domain.DirectoryRepository, timeout time.Duration) domain.DirectoryUsecase {
	return &directoryUsecase{
		itemRepository: itemRepository,
		contextTimeout: timeout,
	}
}

func (r *directoryUsecase) GetCurrencies(c context.Context, filter domain.CurrencyRequest) ([]domain.CurrenciesResponse, error) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()
	return r.itemRepository.GetCurrencies(ctx, filter)
}

func (r *directoryUsecase) GetCountries(c context.Context, filter domain.CountryRequest) ([]domain.CountriesResponse, error) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()
	return r.itemRepository.GetCountries(ctx, filter)
}
