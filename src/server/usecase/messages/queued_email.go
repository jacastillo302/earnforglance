package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/messages"
)

type queuedEmailUsecase struct {
	queuedEmailRepository domain.QueuedEmailRepository
	contextTimeout        time.Duration
}

func NewQueuedEmailUsecase(queuedEmailRepository domain.QueuedEmailRepository, timeout time.Duration) domain.QueuedEmailUsecase {
	return &queuedEmailUsecase{
		queuedEmailRepository: queuedEmailRepository,
		contextTimeout:        timeout,
	}
}

func (tu *queuedEmailUsecase) CreateMany(c context.Context, MessageTemplatesSettingsList []domain.QueuedEmail) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.queuedEmailRepository.CreateMany(ctx, MessageTemplatesSettingsList)
}

func (tu *queuedEmailUsecase) Create(c context.Context, queuedEmail *domain.QueuedEmail) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.queuedEmailRepository.Create(ctx, queuedEmail)
}

func (tu *queuedEmailUsecase) Update(c context.Context, queuedEmail *domain.QueuedEmail) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.queuedEmailRepository.Update(ctx, queuedEmail)
}

func (tu *queuedEmailUsecase) Delete(c context.Context, queuedEmail string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.queuedEmailRepository.Delete(ctx, queuedEmail)
}

func (lu *queuedEmailUsecase) FetchByID(c context.Context, queuedEmailID string) (domain.QueuedEmail, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.queuedEmailRepository.FetchByID(ctx, queuedEmailID)
}

func (lu *queuedEmailUsecase) Fetch(c context.Context) ([]domain.QueuedEmail, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.queuedEmailRepository.Fetch(ctx)
}
