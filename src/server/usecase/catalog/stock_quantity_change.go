package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/catalog"
)

type StockQuantityHistoryUsecase struct {
	StockQuantityHistoryRepository domain.StockQuantityHistoryRepository
	contextTimeout                 time.Duration
}

func NewStockQuantityHistoryUsecase(StockQuantityHistoryRepository domain.StockQuantityHistoryRepository, timeout time.Duration) domain.StockQuantityHistoryUsecase {
	return &StockQuantityHistoryUsecase{
		StockQuantityHistoryRepository: StockQuantityHistoryRepository,
		contextTimeout:                 timeout,
	}
}

func (tu *StockQuantityHistoryUsecase) Create(c context.Context, StockQuantityHistory *domain.StockQuantityHistory) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.StockQuantityHistoryRepository.Create(ctx, StockQuantityHistory)
}

func (tu *StockQuantityHistoryUsecase) Update(c context.Context, StockQuantityHistory *domain.StockQuantityHistory) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.StockQuantityHistoryRepository.Update(ctx, StockQuantityHistory)
}

func (tu *StockQuantityHistoryUsecase) Delete(c context.Context, StockQuantityHistory *domain.StockQuantityHistory) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.StockQuantityHistoryRepository.Delete(ctx, StockQuantityHistory)
}

func (lu *StockQuantityHistoryUsecase) FetchByID(c context.Context, StockQuantityHistoryID string) (domain.StockQuantityHistory, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.StockQuantityHistoryRepository.FetchByID(ctx, StockQuantityHistoryID)
}

func (lu *StockQuantityHistoryUsecase) Fetch(c context.Context) ([]domain.StockQuantityHistory, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.StockQuantityHistoryRepository.Fetch(ctx)
}
