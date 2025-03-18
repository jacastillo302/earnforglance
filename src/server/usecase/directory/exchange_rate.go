package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/directory"
)

type exchangeRateUsecase struct {
	exchangeRateRepository domain.ExchangeRateRepository
	contextTimeout         time.Duration
}

func NewExchangeRateUsecase(exchangeRateRepository domain.ExchangeRateRepository, timeout time.Duration) domain.ExchangeRateUsecase {
	return &exchangeRateUsecase{
		exchangeRateRepository: exchangeRateRepository,
		contextTimeout:         timeout,
	}
}

func (tu *exchangeRateUsecase) Create(c context.Context, exchangeRate *domain.ExchangeRate) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.exchangeRateRepository.Create(ctx, exchangeRate)
}

func (tu *exchangeRateUsecase) Update(c context.Context, exchangeRate *domain.ExchangeRate) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.exchangeRateRepository.Update(ctx, exchangeRate)
}

func (tu *exchangeRateUsecase) Delete(c context.Context, exchangeRate *domain.ExchangeRate) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.exchangeRateRepository.Delete(ctx, exchangeRate)
}

func (lu *exchangeRateUsecase) FetchByID(c context.Context, exchangeRateID string) (domain.ExchangeRate, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.exchangeRateRepository.FetchByID(ctx, exchangeRateID)
}

func (lu *exchangeRateUsecase) Fetch(c context.Context) ([]domain.ExchangeRate, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.exchangeRateRepository.Fetch(ctx)
}
