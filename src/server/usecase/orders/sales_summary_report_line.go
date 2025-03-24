package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/orders"
)

type salesSummaryReportLineUsecase struct {
	salesSummaryReportLineRepository domain.SalesSummaryReportLineRepository
	contextTimeout                   time.Duration
}

func NewSalesSummaryReportLineUsecase(salesSummaryReportLineRepository domain.SalesSummaryReportLineRepository, timeout time.Duration) domain.SalesSummaryReportLineUsecase {
	return &salesSummaryReportLineUsecase{
		salesSummaryReportLineRepository: salesSummaryReportLineRepository,
		contextTimeout:                   timeout,
	}
}

func (tu *salesSummaryReportLineUsecase) CreateMany(c context.Context, MessageTemplatesSettingsList []domain.SalesSummaryReportLine) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.salesSummaryReportLineRepository.CreateMany(ctx, MessageTemplatesSettingsList)
}

func (tu *salesSummaryReportLineUsecase) Create(c context.Context, salesSummaryReportLine *domain.SalesSummaryReportLine) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.salesSummaryReportLineRepository.Create(ctx, salesSummaryReportLine)
}

func (tu *salesSummaryReportLineUsecase) Update(c context.Context, salesSummaryReportLine *domain.SalesSummaryReportLine) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.salesSummaryReportLineRepository.Update(ctx, salesSummaryReportLine)
}

func (tu *salesSummaryReportLineUsecase) Delete(c context.Context, salesSummaryReportLine string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.salesSummaryReportLineRepository.Delete(ctx, salesSummaryReportLine)
}

func (lu *salesSummaryReportLineUsecase) FetchByID(c context.Context, salesSummaryReportLineID string) (domain.SalesSummaryReportLine, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.salesSummaryReportLineRepository.FetchByID(ctx, salesSummaryReportLineID)
}

func (lu *salesSummaryReportLineUsecase) Fetch(c context.Context) ([]domain.SalesSummaryReportLine, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.salesSummaryReportLineRepository.Fetch(ctx)
}
