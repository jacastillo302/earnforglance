package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/catalog"
)

type productcategoryUsecase struct {
	productcategoryRepository domain.ProductCategoryRepository
	contextTimeout            time.Duration
}

func NewProductCategoryUsecase(productcategoryRepository domain.ProductCategoryRepository, timeout time.Duration) domain.ProductCategoryUsecase {
	return &productcategoryUsecase{
		productcategoryRepository: productcategoryRepository,
		contextTimeout:            timeout,
	}
}

func (tu *productcategoryUsecase) Create(c context.Context, productcategory *domain.ProductCategory) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.productcategoryRepository.Create(ctx, productcategory)
}

func (tu *productcategoryUsecase) Update(c context.Context, productcategory *domain.ProductCategory) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.productcategoryRepository.Update(ctx, productcategory)
}

func (tu *productcategoryUsecase) Delete(c context.Context, productcategory *domain.ProductCategory) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.productcategoryRepository.Delete(ctx, productcategory)
}

func (lu *productcategoryUsecase) FetchByID(c context.Context, productcategoryID string) (domain.ProductCategory, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.productcategoryRepository.FetchByID(ctx, productcategoryID)
}

func (lu *productcategoryUsecase) Fetch(c context.Context) ([]domain.ProductCategory, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.productcategoryRepository.Fetch(ctx)
}
