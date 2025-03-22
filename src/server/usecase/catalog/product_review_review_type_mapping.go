package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/catalog"
)

type productreviewreviewtypemappingUsecase struct {
	productreviewreviewtypemappingRepository domain.ProductReviewReviewTypeMappingRepository
	contextTimeout                           time.Duration
}

func NewProductReviewReviewTypeMappingUsecase(productreviewreviewtypemappingRepository domain.ProductReviewReviewTypeMappingRepository, timeout time.Duration) domain.ProductReviewReviewTypeMappingUsecase {
	return &productreviewreviewtypemappingUsecase{
		productreviewreviewtypemappingRepository: productreviewreviewtypemappingRepository,
		contextTimeout:                           timeout,
	}
}

func (tu *productreviewreviewtypemappingUsecase) Create(c context.Context, productreviewreviewtypemapping *domain.ProductReviewReviewTypeMapping) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.productreviewreviewtypemappingRepository.Create(ctx, productreviewreviewtypemapping)
}

func (tu *productreviewreviewtypemappingUsecase) Update(c context.Context, productreviewreviewtypemapping *domain.ProductReviewReviewTypeMapping) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.productreviewreviewtypemappingRepository.Update(ctx, productreviewreviewtypemapping)
}

func (tu *productreviewreviewtypemappingUsecase) Delete(c context.Context, productreviewreviewtypemapping string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.productreviewreviewtypemappingRepository.Delete(ctx, productreviewreviewtypemapping)
}

func (lu *productreviewreviewtypemappingUsecase) FetchByID(c context.Context, productreviewreviewtypemappingID string) (domain.ProductReviewReviewTypeMapping, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.productreviewreviewtypemappingRepository.FetchByID(ctx, productreviewreviewtypemappingID)
}

func (lu *productreviewreviewtypemappingUsecase) Fetch(c context.Context) ([]domain.ProductReviewReviewTypeMapping, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.productreviewreviewtypemappingRepository.Fetch(ctx)
}
