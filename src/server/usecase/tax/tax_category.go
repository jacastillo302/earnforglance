package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/tax"
)

type taxcategoryUsecase struct {
	taxcategoryRepository domain.TaxCategoryRepository
	contextTimeout        time.Duration
}

func NewTaxCategoryUsecase(taxcategoryRepository domain.TaxCategoryRepository, timeout time.Duration) domain.TaxCategoryUsecase {
	return &taxcategoryUsecase{
		taxcategoryRepository: taxcategoryRepository,
		contextTimeout:        timeout,
	}
}

func (tu *taxcategoryUsecase) CreateMany(c context.Context, MessageTemplatesSettingsList []domain.TaxCategory) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taxcategoryRepository.CreateMany(ctx, MessageTemplatesSettingsList)
}

func (tu *taxcategoryUsecase) Create(c context.Context, taxcategory *domain.TaxCategory) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taxcategoryRepository.Create(ctx, taxcategory)
}

func (tu *taxcategoryUsecase) Update(c context.Context, taxcategory *domain.TaxCategory) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taxcategoryRepository.Update(ctx, taxcategory)
}

func (tu *taxcategoryUsecase) Delete(c context.Context, taxcategory string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taxcategoryRepository.Delete(ctx, taxcategory)
}

func (lu *taxcategoryUsecase) FetchByID(c context.Context, taxcategoryID string) (domain.TaxCategory, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.taxcategoryRepository.FetchByID(ctx, taxcategoryID)
}

func (lu *taxcategoryUsecase) Fetch(c context.Context) ([]domain.TaxCategory, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.taxcategoryRepository.Fetch(ctx)
}
