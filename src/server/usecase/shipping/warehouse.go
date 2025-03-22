package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/shipping"
)

type warehouseUsecase struct {
	warehouseRepository domain.WarehouseRepository
	contextTimeout      time.Duration
}

func NewWarehouseUsecase(warehouseRepository domain.WarehouseRepository, timeout time.Duration) domain.WarehouseUsecase {
	return &warehouseUsecase{
		warehouseRepository: warehouseRepository,
		contextTimeout:      timeout,
	}
}

func (tu *warehouseUsecase) Create(c context.Context, warehouse *domain.Warehouse) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.warehouseRepository.Create(ctx, warehouse)
}

func (tu *warehouseUsecase) Update(c context.Context, warehouse *domain.Warehouse) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.warehouseRepository.Update(ctx, warehouse)
}

func (tu *warehouseUsecase) Delete(c context.Context, warehouse string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.warehouseRepository.Delete(ctx, warehouse)
}

func (lu *warehouseUsecase) FetchByID(c context.Context, warehouseID string) (domain.Warehouse, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.warehouseRepository.FetchByID(ctx, warehouseID)
}

func (lu *warehouseUsecase) Fetch(c context.Context) ([]domain.Warehouse, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.warehouseRepository.Fetch(ctx)
}
