package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/common"
)

type searchtermUsecase struct {
	searchtermRepository domain.SearchTermRepository
	contextTimeout       time.Duration
}

func NewSearchTermUsecase(searchtermRepository domain.SearchTermRepository, timeout time.Duration) domain.SearchTermUsecase {
	return &searchtermUsecase{
		searchtermRepository: searchtermRepository,
		contextTimeout:       timeout,
	}
}

func (tu *searchtermUsecase) CreateMany(c context.Context, items []domain.SearchTerm) error {
	ctx,
		cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.searchtermRepository.CreateMany(ctx, items)
}

func (tu *searchtermUsecase) Create(c context.Context, searchterm *domain.SearchTerm) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.searchtermRepository.Create(ctx, searchterm)
}

func (tu *searchtermUsecase) Update(c context.Context, searchterm *domain.SearchTerm) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.searchtermRepository.Update(ctx, searchterm)
}

func (tu *searchtermUsecase) Delete(c context.Context, searchterm string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.searchtermRepository.Delete(ctx, searchterm)
}

func (lu *searchtermUsecase) FetchByID(c context.Context, searchtermID string) (domain.SearchTerm, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.searchtermRepository.FetchByID(ctx, searchtermID)
}

func (lu *searchtermUsecase) Fetch(c context.Context) ([]domain.SearchTerm, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.searchtermRepository.Fetch(ctx)
}
