package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/directory"
)

type measureweightUsecase struct {
	measureweightRepository domain.MeasureWeightRepository
	contextTimeout          time.Duration
}

func NewMeasureWeightUsecase(measureweightRepository domain.MeasureWeightRepository, timeout time.Duration) domain.MeasureWeightUsecase {
	return &measureweightUsecase{
		measureweightRepository: measureweightRepository,
		contextTimeout:          timeout,
	}
}

func (tu *measureweightUsecase) CreateMany(c context.Context, items []domain.MeasureWeight) error {
	ctx,
		cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.measureweightRepository.CreateMany(ctx, items)
}

func (tu *measureweightUsecase) Create(c context.Context, measureweight *domain.MeasureWeight) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.measureweightRepository.Create(ctx, measureweight)
}

func (tu *measureweightUsecase) Update(c context.Context, measureweight *domain.MeasureWeight) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.measureweightRepository.Update(ctx, measureweight)
}

func (tu *measureweightUsecase) Delete(c context.Context, measureweight string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.measureweightRepository.Delete(ctx, measureweight)
}

func (lu *measureweightUsecase) FetchByID(c context.Context, measureweightID string) (domain.MeasureWeight, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.measureweightRepository.FetchByID(ctx, measureweightID)
}

func (lu *measureweightUsecase) Fetch(c context.Context) ([]domain.MeasureWeight, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.measureweightRepository.Fetch(ctx)
}
