package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/messages"
)

type messageTemplateUsecase struct {
	messageTemplateRepository domain.MessageTemplateRepository
	contextTimeout            time.Duration
}

func NewMessageTemplateUsecase(messageTemplateRepository domain.MessageTemplateRepository, timeout time.Duration) domain.MessageTemplateUsecase {
	return &messageTemplateUsecase{
		messageTemplateRepository: messageTemplateRepository,
		contextTimeout:            timeout,
	}
}

func (tu *messageTemplateUsecase) Create(c context.Context, messageTemplate *domain.MessageTemplate) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.messageTemplateRepository.Create(ctx, messageTemplate)
}

func (tu *messageTemplateUsecase) Update(c context.Context, messageTemplate *domain.MessageTemplate) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.messageTemplateRepository.Update(ctx, messageTemplate)
}

func (tu *messageTemplateUsecase) Delete(c context.Context, messageTemplate *domain.MessageTemplate) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.messageTemplateRepository.Delete(ctx, messageTemplate)
}

func (lu *messageTemplateUsecase) FetchByID(c context.Context, messageTemplateID string) (domain.MessageTemplate, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.messageTemplateRepository.FetchByID(ctx, messageTemplateID)
}

func (lu *messageTemplateUsecase) Fetch(c context.Context) ([]domain.MessageTemplate, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.messageTemplateRepository.Fetch(ctx)
}
