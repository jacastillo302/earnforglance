package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/catalog"
)

type productUsecase struct {
	productRepository domain.ProductRepository
	contextTimeout    time.Duration
}

func NewProductUsecase(productRepository domain.ProductRepository, timeout time.Duration) domain.ProductUsecase {
	return &productUsecase{
		productRepository: productRepository,
		contextTimeout:    timeout,
	}
}

func (tu *productUsecase) Create(c context.Context, product *domain.Product) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.productRepository.Create(ctx, product)
}

func (tu *productUsecase) Update(c context.Context, product *domain.Product) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.productRepository.Update(ctx, product)
}

func (tu *productUsecase) Delete(c context.Context, product *domain.Product) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.productRepository.Delete(ctx, product)
}

func (lu *productUsecase) FetchByID(c context.Context, productID string) (domain.Product, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.productRepository.FetchByID(ctx, productID)
}

func (lu *productUsecase) Fetch(c context.Context) ([]domain.Product, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.productRepository.Fetch(ctx)
}
