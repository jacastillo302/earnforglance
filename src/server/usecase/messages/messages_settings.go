package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/messages"
)

type MessagesSettingsUsecase struct {
	MessagesSettingsRepository domain.MessagesSettingsRepository
	contextTimeout             time.Duration
}

func NewMessagesSettingsUsecase(MessagesSettingsRepository domain.MessagesSettingsRepository, timeout time.Duration) domain.MessagesSettingsUsecase {
	return &MessagesSettingsUsecase{
		MessagesSettingsRepository: MessagesSettingsRepository,
		contextTimeout:             timeout,
	}
}

func (tu *MessagesSettingsUsecase) Create(c context.Context, MessagesSettings *domain.MessagesSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.MessagesSettingsRepository.Create(ctx, MessagesSettings)
}

func (tu *MessagesSettingsUsecase) Update(c context.Context, MessagesSettings *domain.MessagesSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.MessagesSettingsRepository.Update(ctx, MessagesSettings)
}

func (tu *MessagesSettingsUsecase) Delete(c context.Context, MessagesSettings *domain.MessagesSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.MessagesSettingsRepository.Delete(ctx, MessagesSettings)
}

func (lu *MessagesSettingsUsecase) FetchByID(c context.Context, MessagesSettingsID string) (domain.MessagesSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.MessagesSettingsRepository.FetchByID(ctx, MessagesSettingsID)
}

func (lu *MessagesSettingsUsecase) Fetch(c context.Context) ([]domain.MessagesSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.MessagesSettingsRepository.Fetch(ctx)
}
