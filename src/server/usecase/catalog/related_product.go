package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/catalog"
)

type relatedproductUsecase struct {
	relatedproductRepository domain.RelatedProductRepository
	contextTimeout           time.Duration
}

func NewRelatedProductUsecase(relatedproductRepository domain.RelatedProductRepository, timeout time.Duration) domain.RelatedProductUsecase {
	return &relatedproductUsecase{
		relatedproductRepository: relatedproductRepository,
		contextTimeout:           timeout,
	}
}

func (tu *relatedproductUsecase) CreateMany(c context.Context, items []domain.RelatedProduct) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.relatedproductRepository.CreateMany(ctx, items)
}

func (tu *relatedproductUsecase) Create(c context.Context, relatedproduct *domain.RelatedProduct) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.relatedproductRepository.Create(ctx, relatedproduct)
}

func (tu *relatedproductUsecase) Update(c context.Context, relatedproduct *domain.RelatedProduct) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.relatedproductRepository.Update(ctx, relatedproduct)
}

func (tu *relatedproductUsecase) Delete(c context.Context, relatedproduct string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.relatedproductRepository.Delete(ctx, relatedproduct)
}

func (lu *relatedproductUsecase) FetchByID(c context.Context, relatedproductID string) (domain.RelatedProduct, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.relatedproductRepository.FetchByID(ctx, relatedproductID)
}

func (lu *relatedproductUsecase) Fetch(c context.Context) ([]domain.RelatedProduct, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.relatedproductRepository.Fetch(ctx)
}
