package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/catalog"
)

type productreviewUsecase struct {
	productreviewRepository domain.ProductReviewRepository
	contextTimeout          time.Duration
}

func NewProductReviewUsecase(productreviewRepository domain.ProductReviewRepository, timeout time.Duration) domain.ProductReviewUsecase {
	return &productreviewUsecase{
		productreviewRepository: productreviewRepository,
		contextTimeout:          timeout,
	}
}

func (tu *productreviewUsecase) Create(c context.Context, productreview *domain.ProductReview) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.productreviewRepository.Create(ctx, productreview)
}

func (tu *productreviewUsecase) Update(c context.Context, productreview *domain.ProductReview) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.productreviewRepository.Update(ctx, productreview)
}

func (tu *productreviewUsecase) Delete(c context.Context, productreview string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.productreviewRepository.Delete(ctx, productreview)
}

func (lu *productreviewUsecase) FetchByID(c context.Context, productreviewID string) (domain.ProductReview, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.productreviewRepository.FetchByID(ctx, productreviewID)
}

func (lu *productreviewUsecase) Fetch(c context.Context) ([]domain.ProductReview, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.productreviewRepository.Fetch(ctx)
}
