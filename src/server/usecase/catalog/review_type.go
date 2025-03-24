package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/catalog"
)

type reviewtypeUsecase struct {
	reviewtypeRepository domain.ReviewTypeRepository
	contextTimeout       time.Duration
}

func NewReviewTypeUsecase(reviewtypeRepository domain.ReviewTypeRepository, timeout time.Duration) domain.ReviewTypeUsecase {
	return &reviewtypeUsecase{
		reviewtypeRepository: reviewtypeRepository,
		contextTimeout:       timeout,
	}
}

func (tu *reviewtypeUsecase) CreateMany(c context.Context, items []domain.ReviewType) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.reviewtypeRepository.CreateMany(ctx, items)
}

func (tu *reviewtypeUsecase) Create(c context.Context, reviewtype *domain.ReviewType) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.reviewtypeRepository.Create(ctx, reviewtype)
}

func (tu *reviewtypeUsecase) Update(c context.Context, reviewtype *domain.ReviewType) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.reviewtypeRepository.Update(ctx, reviewtype)
}

func (tu *reviewtypeUsecase) Delete(c context.Context, reviewtype string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.reviewtypeRepository.Delete(ctx, reviewtype)
}

func (lu *reviewtypeUsecase) FetchByID(c context.Context, reviewtypeID string) (domain.ReviewType, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.reviewtypeRepository.FetchByID(ctx, reviewtypeID)
}

func (lu *reviewtypeUsecase) Fetch(c context.Context) ([]domain.ReviewType, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.reviewtypeRepository.Fetch(ctx)
}
