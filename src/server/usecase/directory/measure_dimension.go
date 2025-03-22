package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/directory"
)

type measuredimensionUsecase struct {
	MeasureDimensionRepository domain.MeasureDimensionRepository
	contextTimeout             time.Duration
}

func NewMeasureDimensionUsecase(MeasureDimensionRepository domain.MeasureDimensionRepository, timeout time.Duration) domain.MeasureDimensionUsecase {
	return &measuredimensionUsecase{
		MeasureDimensionRepository: MeasureDimensionRepository,
		contextTimeout:             timeout,
	}
}

func (tu *measuredimensionUsecase) Create(c context.Context, measuredimension *domain.MeasureDimension) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.MeasureDimensionRepository.Create(ctx, measuredimension)
}

func (tu *measuredimensionUsecase) Update(c context.Context, measuredimension *domain.MeasureDimension) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.MeasureDimensionRepository.Update(ctx, measuredimension)
}

func (tu *measuredimensionUsecase) Delete(c context.Context, measuredimension string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.MeasureDimensionRepository.Delete(ctx, measuredimension)
}

func (lu *measuredimensionUsecase) FetchByID(c context.Context, measuredimensionID string) (domain.MeasureDimension, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.MeasureDimensionRepository.FetchByID(ctx, measuredimensionID)
}

func (lu *measuredimensionUsecase) Fetch(c context.Context) ([]domain.MeasureDimension, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.MeasureDimensionRepository.Fetch(ctx)
}
