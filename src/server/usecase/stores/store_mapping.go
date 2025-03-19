package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/stores"
)

type storemappingUsecase struct {
	storemappingRepository domain.StoreMappingRepository
	contextTimeout         time.Duration
}

func NewStoreMappingUsecase(storemappingRepository domain.StoreMappingRepository, timeout time.Duration) domain.StoreMappingUsecase {
	return &storemappingUsecase{
		storemappingRepository: storemappingRepository,
		contextTimeout:         timeout,
	}
}

func (tu *storemappingUsecase) Create(c context.Context, storemapping *domain.StoreMapping) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.storemappingRepository.Create(ctx, storemapping)
}

func (tu *storemappingUsecase) Update(c context.Context, storemapping *domain.StoreMapping) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.storemappingRepository.Update(ctx, storemapping)
}

func (tu *storemappingUsecase) Delete(c context.Context, storemapping *domain.StoreMapping) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.storemappingRepository.Delete(ctx, storemapping)
}

func (lu *storemappingUsecase) FetchByID(c context.Context, storemappingID string) (domain.StoreMapping, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.storemappingRepository.FetchByID(ctx, storemappingID)
}

func (lu *storemappingUsecase) Fetch(c context.Context) ([]domain.StoreMapping, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.storemappingRepository.Fetch(ctx)
}
