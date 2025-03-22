package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/orders"
)

type bestsellersreportlineUsecase struct {
	bestsellersreportlineRepository domain.BestSellersReportLineRepository
	contextTimeout                  time.Duration
}

func NewBestSellersReportLineUsecase(bestsellersreportlineRepository domain.BestSellersReportLineRepository, timeout time.Duration) domain.BestSellersReportLineUsecase {
	return &bestsellersreportlineUsecase{
		bestsellersreportlineRepository: bestsellersreportlineRepository,
		contextTimeout:                  timeout,
	}
}

func (tu *bestsellersreportlineUsecase) Create(c context.Context, bestsellersreportline *domain.BestSellersReportLine) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.bestsellersreportlineRepository.Create(ctx, bestsellersreportline)
}

func (tu *bestsellersreportlineUsecase) Update(c context.Context, bestsellersreportline *domain.BestSellersReportLine) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.bestsellersreportlineRepository.Update(ctx, bestsellersreportline)
}

func (tu *bestsellersreportlineUsecase) Delete(c context.Context, bestsellersreportline string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.bestsellersreportlineRepository.Delete(ctx, bestsellersreportline)
}

func (lu *bestsellersreportlineUsecase) FetchByID(c context.Context, bestsellersreportlineID string) (domain.BestSellersReportLine, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.bestsellersreportlineRepository.FetchByID(ctx, bestsellersreportlineID)
}

func (lu *bestsellersreportlineUsecase) Fetch(c context.Context) ([]domain.BestSellersReportLine, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.bestsellersreportlineRepository.Fetch(ctx)
}
