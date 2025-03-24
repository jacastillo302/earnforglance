package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/logging"
)

type activityLogUsecase struct {
	activityLogRepository domain.ActivityLogRepository
	contextTimeout        time.Duration
}

func NewActivityLogUsecase(activityLogRepository domain.ActivityLogRepository, timeout time.Duration) domain.ActivityLogUsecase {
	return &activityLogUsecase{
		activityLogRepository: activityLogRepository,
		contextTimeout:        timeout,
	}
}

func (tu *activityLogUsecase) CreateMany(c context.Context, items []domain.ActivityLog) error {
	ctx,
		cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.activityLogRepository.CreateMany(ctx, items)
}

func (tu *activityLogUsecase) Create(c context.Context, activityLog *domain.ActivityLog) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.activityLogRepository.Create(ctx, activityLog)
}

func (tu *activityLogUsecase) Update(c context.Context, activityLog *domain.ActivityLog) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.activityLogRepository.Update(ctx, activityLog)
}

func (tu *activityLogUsecase) Delete(c context.Context, activityLog string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.activityLogRepository.Delete(ctx, activityLog)
}

func (lu *activityLogUsecase) FetchByID(c context.Context, activityLogID string) (domain.ActivityLog, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.activityLogRepository.FetchByID(ctx, activityLogID)
}

func (lu *activityLogUsecase) Fetch(c context.Context) ([]domain.ActivityLog, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.activityLogRepository.Fetch(ctx)
}
