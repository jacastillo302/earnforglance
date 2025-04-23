package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/public"
)

type catalogtUsecase struct {
	itemRepository domain.CatalogRepository
	contextTimeout time.Duration
}

func NewCatalogtUsecase(itemRepository domain.CatalogRepository, timeout time.Duration) domain.CatalogtUsecase {
	return &catalogtUsecase{
		itemRepository: itemRepository,
		contextTimeout: timeout,
	}
}

func (cu *catalogtUsecase) GetProducts(c context.Context, filter domain.ProductRequest) ([]domain.ProductsResponse, error) {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.itemRepository.GetProducts(ctx, filter)
}

func (cu *catalogtUsecase) GetCategories(c context.Context, filter domain.CategoryRequest) ([]domain.CategoriesResponse, error) {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.itemRepository.GetCategories(ctx, filter)
}

func (cu *catalogtUsecase) GetManufacturers(c context.Context, filter domain.ManufacturerRequest) ([]domain.ManufacturersResponse, error) {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.itemRepository.GetManufacturers(ctx, filter)
}
