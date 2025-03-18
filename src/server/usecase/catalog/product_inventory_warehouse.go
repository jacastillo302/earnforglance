package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/catalog"
)

type ProductWarehouseInventoryUsecase struct {
	ProductWarehouseInventoryRepository domain.ProductWarehouseInventoryRepository
	contextTimeout                      time.Duration
}

func NewProductWarehouseInventoryUsecase(ProductWarehouseInventoryRepository domain.ProductWarehouseInventoryRepository, timeout time.Duration) domain.ProductWarehouseInventoryUsecase {
	return &ProductWarehouseInventoryUsecase{
		ProductWarehouseInventoryRepository: ProductWarehouseInventoryRepository,
		contextTimeout:                      timeout,
	}
}

func (tu *ProductWarehouseInventoryUsecase) Create(c context.Context, ProductWarehouseInventory *domain.ProductWarehouseInventory) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.ProductWarehouseInventoryRepository.Create(ctx, ProductWarehouseInventory)
}

func (tu *ProductWarehouseInventoryUsecase) Update(c context.Context, ProductWarehouseInventory *domain.ProductWarehouseInventory) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.ProductWarehouseInventoryRepository.Update(ctx, ProductWarehouseInventory)
}

func (tu *ProductWarehouseInventoryUsecase) Delete(c context.Context, ProductWarehouseInventory *domain.ProductWarehouseInventory) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.ProductWarehouseInventoryRepository.Delete(ctx, ProductWarehouseInventory)
}

func (lu *ProductWarehouseInventoryUsecase) FetchByID(c context.Context, ProductWarehouseInventoryID string) (domain.ProductWarehouseInventory, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.ProductWarehouseInventoryRepository.FetchByID(ctx, ProductWarehouseInventoryID)
}

func (lu *ProductWarehouseInventoryUsecase) Fetch(c context.Context) ([]domain.ProductWarehouseInventory, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.ProductWarehouseInventoryRepository.Fetch(ctx)
}
