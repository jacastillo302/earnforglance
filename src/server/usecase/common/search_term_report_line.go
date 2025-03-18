package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/common"
)

type SearchTermReportLineUsecase struct {
	SearchTermReportLineRepository domain.SearchTermReportLineRepository
	contextTimeout                 time.Duration
}

func NewSearchTermReportLineUsecase(SearchTermReportLineRepository domain.SearchTermReportLineRepository, timeout time.Duration) domain.SearchTermReportLineUsecase {
	return &SearchTermReportLineUsecase{
		SearchTermReportLineRepository: SearchTermReportLineRepository,
		contextTimeout:                 timeout,
	}
}

func (tu *SearchTermReportLineUsecase) Create(c context.Context, SearchTermReportLine *domain.SearchTermReportLine) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.SearchTermReportLineRepository.Create(ctx, SearchTermReportLine)
}

func (tu *SearchTermReportLineUsecase) Update(c context.Context, SearchTermReportLine *domain.SearchTermReportLine) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.SearchTermReportLineRepository.Update(ctx, SearchTermReportLine)
}

func (tu *SearchTermReportLineUsecase) Delete(c context.Context, SearchTermReportLine *domain.SearchTermReportLine) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.SearchTermReportLineRepository.Delete(ctx, SearchTermReportLine)
}

func (lu *SearchTermReportLineUsecase) FetchByID(c context.Context, SearchTermReportLineID string) (domain.SearchTermReportLine, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.SearchTermReportLineRepository.FetchByID(ctx, SearchTermReportLineID)
}

func (lu *SearchTermReportLineUsecase) Fetch(c context.Context) ([]domain.SearchTermReportLine, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.SearchTermReportLineRepository.Fetch(ctx)
}
