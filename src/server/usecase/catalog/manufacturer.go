package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/catalog"
)

type manufacturerUsecase struct {
	manufacturerRepository domain.ManufacturerRepository
	contextTimeout         time.Duration
}

func NewManufacturerUsecase(manufacturerRepository domain.ManufacturerRepository, timeout time.Duration) domain.ManufacturerUsecase {
	return &manufacturerUsecase{
		manufacturerRepository: manufacturerRepository,
		contextTimeout:         timeout,
	}
}

func (tu *manufacturerUsecase) Create(c context.Context, manufacturer *domain.Manufacturer) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.manufacturerRepository.Create(ctx, manufacturer)
}

func (tu *manufacturerUsecase) Update(c context.Context, manufacturer *domain.Manufacturer) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.manufacturerRepository.Update(ctx, manufacturer)
}

func (tu *manufacturerUsecase) Delete(c context.Context, manufacturer string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.manufacturerRepository.Delete(ctx, manufacturer)
}

func (lu *manufacturerUsecase) FetchByID(c context.Context, manufacturerID string) (domain.Manufacturer, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.manufacturerRepository.FetchByID(ctx, manufacturerID)
}

func (lu *manufacturerUsecase) Fetch(c context.Context) ([]domain.Manufacturer, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.manufacturerRepository.Fetch(ctx)
}
