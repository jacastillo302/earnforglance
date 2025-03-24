package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/scheduleTasks"
)

type scheduletaskUsecase struct {
	scheduletaskRepository domain.ScheduleTaskRepository
	contextTimeout         time.Duration
}

func NewScheduleTaskUsecase(scheduletaskRepository domain.ScheduleTaskRepository, timeout time.Duration) domain.ScheduleTaskUsecase {
	return &scheduletaskUsecase{
		scheduletaskRepository: scheduletaskRepository,
		contextTimeout:         timeout,
	}
}

func (tu *scheduletaskUsecase) CreateMany(c context.Context, MessageTemplatesSettingsList []domain.ScheduleTask) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.scheduletaskRepository.CreateMany(ctx, MessageTemplatesSettingsList)
}

func (tu *scheduletaskUsecase) Create(c context.Context, scheduletask *domain.ScheduleTask) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.scheduletaskRepository.Create(ctx, scheduletask)
}

func (tu *scheduletaskUsecase) Update(c context.Context, scheduletask *domain.ScheduleTask) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.scheduletaskRepository.Update(ctx, scheduletask)
}

func (tu *scheduletaskUsecase) Delete(c context.Context, scheduletask string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.scheduletaskRepository.Delete(ctx, scheduletask)
}

func (lu *scheduletaskUsecase) FetchByID(c context.Context, scheduletaskID string) (domain.ScheduleTask, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.scheduletaskRepository.FetchByID(ctx, scheduletaskID)
}

func (lu *scheduletaskUsecase) Fetch(c context.Context) ([]domain.ScheduleTask, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.scheduletaskRepository.Fetch(ctx)
}
