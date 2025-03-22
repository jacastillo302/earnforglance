package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/catalog"
)

type productreviewhelpfulnessUsecase struct {
	productreviewhelpfulnessRepository domain.ProductReviewHelpfulnessRepository
	contextTimeout                     time.Duration
}

func NewProductReviewHelpfulnessUsecase(productreviewhelpfulnessRepository domain.ProductReviewHelpfulnessRepository, timeout time.Duration) domain.ProductReviewHelpfulnessUsecase {
	return &productreviewhelpfulnessUsecase{
		productreviewhelpfulnessRepository: productreviewhelpfulnessRepository,
		contextTimeout:                     timeout,
	}
}

func (tu *productreviewhelpfulnessUsecase) Create(c context.Context, productreviewhelpfulness *domain.ProductReviewHelpfulness) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.productreviewhelpfulnessRepository.Create(ctx, productreviewhelpfulness)
}

func (tu *productreviewhelpfulnessUsecase) Update(c context.Context, productreviewhelpfulness *domain.ProductReviewHelpfulness) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.productreviewhelpfulnessRepository.Update(ctx, productreviewhelpfulness)
}

func (tu *productreviewhelpfulnessUsecase) Delete(c context.Context, productreviewhelpfulness string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.productreviewhelpfulnessRepository.Delete(ctx, productreviewhelpfulness)
}

func (lu *productreviewhelpfulnessUsecase) FetchByID(c context.Context, productreviewhelpfulnessID string) (domain.ProductReviewHelpfulness, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.productreviewhelpfulnessRepository.FetchByID(ctx, productreviewhelpfulnessID)
}

func (lu *productreviewhelpfulnessUsecase) Fetch(c context.Context) ([]domain.ProductReviewHelpfulness, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.productreviewhelpfulnessRepository.Fetch(ctx)
}
