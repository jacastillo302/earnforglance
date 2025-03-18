package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/forums"
)

type privateMessageUsecase struct {
	privateMessageRepository domain.PrivateMessageRepository
	contextTimeout           time.Duration
}

func NewPrivateMessageUsecase(privateMessageRepository domain.PrivateMessageRepository, timeout time.Duration) domain.PrivateMessageUsecase {
	return &privateMessageUsecase{
		privateMessageRepository: privateMessageRepository,
		contextTimeout:           timeout,
	}
}

func (tu *privateMessageUsecase) Create(c context.Context, privateMessage *domain.PrivateMessage) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.privateMessageRepository.Create(ctx, privateMessage)
}

func (tu *privateMessageUsecase) Update(c context.Context, privateMessage *domain.PrivateMessage) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.privateMessageRepository.Update(ctx, privateMessage)
}

func (tu *privateMessageUsecase) Delete(c context.Context, privateMessage *domain.PrivateMessage) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.privateMessageRepository.Delete(ctx, privateMessage)
}

func (lu *privateMessageUsecase) FetchByID(c context.Context, privateMessageID string) (domain.PrivateMessage, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.privateMessageRepository.FetchByID(ctx, privateMessageID)
}

func (lu *privateMessageUsecase) Fetch(c context.Context) ([]domain.PrivateMessage, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.privateMessageRepository.Fetch(ctx)
}
