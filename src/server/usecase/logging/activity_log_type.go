package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/logging"
)

type activitylogtypeUsecase struct {
	activitylogtypeRepository domain.ActivityLogTypeRepository
	contextTimeout            time.Duration
}

func NewActivityLogTypeUsecase(activitylogtypeRepository domain.ActivityLogTypeRepository, timeout time.Duration) domain.ActivityLogTypeUsecase {
	return &activitylogtypeUsecase{
		activitylogtypeRepository: activitylogtypeRepository,
		contextTimeout:            timeout,
	}
}

func (tu *activitylogtypeUsecase) CreateMany(c context.Context, items []domain.ActivityLogType) error {
	ctx,
		cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.activitylogtypeRepository.CreateMany(ctx, items)
}

func (tu *activitylogtypeUsecase) Create(c context.Context, activitylogtype *domain.ActivityLogType) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.activitylogtypeRepository.Create(ctx, activitylogtype)
}

func (tu *activitylogtypeUsecase) Update(c context.Context, activitylogtype *domain.ActivityLogType) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.activitylogtypeRepository.Update(ctx, activitylogtype)
}

func (tu *activitylogtypeUsecase) Delete(c context.Context, activitylogtype string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.activitylogtypeRepository.Delete(ctx, activitylogtype)
}

func (lu *activitylogtypeUsecase) FetchByID(c context.Context, activitylogtypeID string) (domain.ActivityLogType, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.activitylogtypeRepository.FetchByID(ctx, activitylogtypeID)
}

func (lu *activitylogtypeUsecase) Fetch(c context.Context) ([]domain.ActivityLogType, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.activitylogtypeRepository.Fetch(ctx)
}
