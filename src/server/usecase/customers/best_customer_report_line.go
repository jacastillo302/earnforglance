package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/customers"
)

type bestcustomerreportlineUsecase struct {
	bestcustomerreportlineRepository domain.BestCustomerReportLineRepository
	contextTimeout                   time.Duration
}

func NewBestCustomerReportLineUsecase(bestcustomerreportlineRepository domain.BestCustomerReportLineRepository, timeout time.Duration) domain.BestCustomerReportLineUsecase {
	return &bestcustomerreportlineUsecase{
		bestcustomerreportlineRepository: bestcustomerreportlineRepository,
		contextTimeout:                   timeout,
	}
}

func (tu *bestcustomerreportlineUsecase) Create(c context.Context, bestcustomerreportline *domain.BestCustomerReportLine) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.bestcustomerreportlineRepository.Create(ctx, bestcustomerreportline)
}

func (tu *bestcustomerreportlineUsecase) Update(c context.Context, bestcustomerreportline *domain.BestCustomerReportLine) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.bestcustomerreportlineRepository.Update(ctx, bestcustomerreportline)
}

func (tu *bestcustomerreportlineUsecase) Delete(c context.Context, bestcustomerreportline string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.bestcustomerreportlineRepository.Delete(ctx, bestcustomerreportline)
}

func (lu *bestcustomerreportlineUsecase) FetchByID(c context.Context, bestcustomerreportlineID string) (domain.BestCustomerReportLine, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.bestcustomerreportlineRepository.FetchByID(ctx, bestcustomerreportlineID)
}

func (lu *bestcustomerreportlineUsecase) Fetch(c context.Context) ([]domain.BestCustomerReportLine, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.bestcustomerreportlineRepository.Fetch(ctx)
}
