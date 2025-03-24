package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/messages"
)

// Compare this snippet from src/server/usecase/messages/message_templates_settings.go:
type MessageTemplatesSettingsUsecase struct {
	MessageTemplatesSettingsRepository domain.MessageTemplatesSettingsRepository
	contextTimeout                     time.Duration
}

func NewMessageTemplatesSettingsUsecase(MessageTemplatesSettingsRepository domain.MessageTemplatesSettingsRepository, timeout time.Duration) domain.MessageTemplatesSettingsRepository {
	return &MessageTemplatesSettingsUsecase{
		MessageTemplatesSettingsRepository: MessageTemplatesSettingsRepository,
		contextTimeout:                     timeout,
	}
}

func (tu *MessageTemplatesSettingsUsecase) CreateMany(c context.Context, MessageTemplatesSettingsList []domain.MessageTemplatesSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.MessageTemplatesSettingsRepository.CreateMany(ctx, MessageTemplatesSettingsList)
}

func (tu *MessageTemplatesSettingsUsecase) Create(c context.Context, MessageTemplatesSettings *domain.MessageTemplatesSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.MessageTemplatesSettingsRepository.Create(ctx, MessageTemplatesSettings)
}

func (tu *MessageTemplatesSettingsUsecase) Update(c context.Context, MessageTemplatesSettings *domain.MessageTemplatesSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.MessageTemplatesSettingsRepository.Update(ctx, MessageTemplatesSettings)
}

func (tu *MessageTemplatesSettingsUsecase) Delete(c context.Context, MessageTemplatesSettings string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.MessageTemplatesSettingsRepository.Delete(ctx, MessageTemplatesSettings)
}

func (lu *MessageTemplatesSettingsUsecase) FetchByID(c context.Context, MessageTemplatesSettingsID string) (domain.MessageTemplatesSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.MessageTemplatesSettingsRepository.FetchByID(ctx, MessageTemplatesSettingsID)
}

func (lu *MessageTemplatesSettingsUsecase) Fetch(c context.Context) ([]domain.MessageTemplatesSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.MessageTemplatesSettingsRepository.Fetch(ctx)
}
