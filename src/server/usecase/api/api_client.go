package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/api"
)

type apiclientUsecase struct {
	apiclientRepository domain.ApiClientRepository
	contextTimeout      time.Duration
}

func NewApiClientUsecase(apiclientRepository domain.ApiClientRepository, timeout time.Duration) domain.ApiClientUsecase {
	return &apiclientUsecase{
		apiclientRepository: apiclientRepository,
		contextTimeout:      timeout,
	}
}

func (tu *apiclientUsecase) CreateMany(c context.Context, items []domain.ApiClient) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.apiclientRepository.CreateMany(ctx, items)
}

func (tu *apiclientUsecase) Create(c context.Context, apiclient *domain.ApiClient) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.apiclientRepository.Create(ctx, apiclient)
}

func (tu *apiclientUsecase) Update(c context.Context, apiclient *domain.ApiClient) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.apiclientRepository.Update(ctx, apiclient)
}

func (tu *apiclientUsecase) Delete(c context.Context, apiclient string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.apiclientRepository.Delete(ctx, apiclient)
}

func (lu *apiclientUsecase) FetchByID(c context.Context, apiclientID string) (domain.ApiClient, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.apiclientRepository.FetchByID(ctx, apiclientID)
}

func (lu *apiclientUsecase) Fetch(c context.Context) ([]domain.ApiClient, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.apiclientRepository.Fetch(ctx)
}
