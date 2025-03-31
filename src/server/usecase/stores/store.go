package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/stores"
)

type storeUsecase struct {
	storeRepository domain.StoreRepository
	contextTimeout  time.Duration
}

func NewStoreUsecase(storeRepository domain.StoreRepository, timeout time.Duration) domain.StoreUsecase {
	return &storeUsecase{
		storeRepository: storeRepository,
		contextTimeout:  timeout,
	}
}

func (tu *storeUsecase) CreateMany(c context.Context, items []domain.Store) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.storeRepository.CreateMany(ctx, items)
}

func (tu *storeUsecase) Create(c context.Context, store *domain.Store) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.storeRepository.Create(ctx, store)
}

func (tu *storeUsecase) Update(c context.Context, store *domain.Store) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.storeRepository.Update(ctx, store)
}

func (tu *storeUsecase) Delete(c context.Context, store string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.storeRepository.Delete(ctx, store)
}

func (lu *storeUsecase) FetchByID(c context.Context, storeID string) (domain.Store, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.storeRepository.FetchByID(ctx, storeID)
}

func (lu *storeUsecase) Fetch(c context.Context) ([]domain.Store, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.storeRepository.Fetch(ctx)
}
