package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/catalog"
)

type StockQuantityChangeUsecase struct {
	StockQuantityChangeRepository domain.StockQuantityChangeRepository
	contextTimeout                time.Duration
}

func NewStockQuantityChangeUsecase(StockQuantityChangeRepository domain.StockQuantityChangeRepository, timeout time.Duration) domain.StockQuantityChangeUsecase {
	return &StockQuantityChangeUsecase{
		StockQuantityChangeRepository: StockQuantityChangeRepository,
		contextTimeout:                timeout,
	}
}

func (tu *StockQuantityChangeUsecase) CreateMany(c context.Context, items []domain.StockQuantityChange) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.StockQuantityChangeRepository.CreateMany(ctx, items)
}

func (tu *StockQuantityChangeUsecase) Create(c context.Context, StockQuantityChange *domain.StockQuantityChange) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.StockQuantityChangeRepository.Create(ctx, StockQuantityChange)
}

func (tu *StockQuantityChangeUsecase) Update(c context.Context, StockQuantityChange *domain.StockQuantityChange) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.StockQuantityChangeRepository.Update(ctx, StockQuantityChange)
}

func (tu *StockQuantityChangeUsecase) Delete(c context.Context, StockQuantityChange string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.StockQuantityChangeRepository.Delete(ctx, StockQuantityChange)
}

func (lu *StockQuantityChangeUsecase) FetchByID(c context.Context, StockQuantityChangeID string) (domain.StockQuantityChange, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.StockQuantityChangeRepository.FetchByID(ctx, StockQuantityChangeID)
}

func (lu *StockQuantityChangeUsecase) Fetch(c context.Context) ([]domain.StockQuantityChange, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.StockQuantityChangeRepository.Fetch(ctx)
}
