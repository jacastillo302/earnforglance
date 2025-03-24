package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/gdpr"
)

type GdprLogUsecase struct {
	GdprLogRepository domain.GdprLogRepository
	contextTimeout    time.Duration
}

func NewGdprLogUsecase(GdprLogRepository domain.GdprLogRepository, timeout time.Duration) domain.GdprLogUsecase {
	return &GdprLogUsecase{
		GdprLogRepository: GdprLogRepository,
		contextTimeout:    timeout,
	}
}

func (tu *GdprLogUsecase) CreateMany(c context.Context, items []domain.GdprLog) error {
	ctx,
		cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.GdprLogRepository.CreateMany(ctx, items)
}

func (tu *GdprLogUsecase) Create(c context.Context, GdprLog *domain.GdprLog) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.GdprLogRepository.Create(ctx, GdprLog)
}

func (tu *GdprLogUsecase) Update(c context.Context, GdprLog *domain.GdprLog) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.GdprLogRepository.Update(ctx, GdprLog)
}

func (tu *GdprLogUsecase) Delete(c context.Context, GdprLog string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.GdprLogRepository.Delete(ctx, GdprLog)
}

func (lu *GdprLogUsecase) FetchByID(c context.Context, GdprLogID string) (domain.GdprLog, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.GdprLogRepository.FetchByID(ctx, GdprLogID)
}

func (lu *GdprLogUsecase) Fetch(c context.Context) ([]domain.GdprLog, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.GdprLogRepository.Fetch(ctx)
}
