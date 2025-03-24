package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/directory"
)

type stateprovinceUsecase struct {
	stateprovinceRepository domain.StateProvinceRepository
	contextTimeout          time.Duration
}

func NewStateProvinceUsecase(stateprovinceRepository domain.StateProvinceRepository, timeout time.Duration) domain.StateProvinceUsecase {
	return &stateprovinceUsecase{
		stateprovinceRepository: stateprovinceRepository,
		contextTimeout:          timeout,
	}
}

func (tu *stateprovinceUsecase) CreateMany(c context.Context, items []domain.StateProvince) error {
	ctx,
		cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.stateprovinceRepository.CreateMany(ctx, items)
}

func (tu *stateprovinceUsecase) Create(c context.Context, stateprovince *domain.StateProvince) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.stateprovinceRepository.Create(ctx, stateprovince)
}

func (tu *stateprovinceUsecase) Update(c context.Context, stateprovince *domain.StateProvince) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.stateprovinceRepository.Update(ctx, stateprovince)
}

func (tu *stateprovinceUsecase) Delete(c context.Context, stateprovince string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.stateprovinceRepository.Delete(ctx, stateprovince)
}

func (lu *stateprovinceUsecase) FetchByID(c context.Context, stateprovinceID string) (domain.StateProvince, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.stateprovinceRepository.FetchByID(ctx, stateprovinceID)
}

func (lu *stateprovinceUsecase) Fetch(c context.Context) ([]domain.StateProvince, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.stateprovinceRepository.Fetch(ctx)
}
