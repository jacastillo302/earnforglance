package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/directory"
)

type currencyUsecase struct {
	currencyRepository domain.CurrencyRepository
	contextTimeout     time.Duration
}

func NewCurrencyUsecase(currencyRepository domain.CurrencyRepository, timeout time.Duration) domain.CurrencyUsecase {
	return &currencyUsecase{
		currencyRepository: currencyRepository,
		contextTimeout:     timeout,
	}
}

func (cu *currencyUsecase) Create(c context.Context, currency *domain.Currency) error {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.currencyRepository.Create(ctx, currency)
}

func (cu *currencyUsecase) Update(c context.Context, currency *domain.Currency) error {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.currencyRepository.Update(ctx, currency)
}

func (cu *currencyUsecase) Delete(c context.Context, currency *domain.Currency) error {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.currencyRepository.Delete(ctx, currency)
}

func (cu *currencyUsecase) FetchByID(c context.Context, currencyID string) (domain.Currency, error) {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.currencyRepository.FetchByID(ctx, currencyID)
}

func (cu *currencyUsecase) Fetch(c context.Context) ([]domain.Currency, error) {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.currencyRepository.Fetch(ctx)
}
