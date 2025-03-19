package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/orders"
)

type giftcardUsecase struct {
	giftcardRepository domain.GiftCardRepository
	contextTimeout     time.Duration
}

func NewGiftCardUsecase(giftcardRepository domain.GiftCardRepository, timeout time.Duration) domain.GiftCardUsecase {
	return &giftcardUsecase{
		giftcardRepository: giftcardRepository,
		contextTimeout:     timeout,
	}
}

func (tu *giftcardUsecase) Create(c context.Context, giftcard *domain.GiftCard) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.giftcardRepository.Create(ctx, giftcard)
}

func (tu *giftcardUsecase) Update(c context.Context, giftcard *domain.GiftCard) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.giftcardRepository.Update(ctx, giftcard)
}

func (tu *giftcardUsecase) Delete(c context.Context, giftcard *domain.GiftCard) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.giftcardRepository.Delete(ctx, giftcard)
}

func (lu *giftcardUsecase) FetchByID(c context.Context, giftcardID string) (domain.GiftCard, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.giftcardRepository.FetchByID(ctx, giftcardID)
}

func (lu *giftcardUsecase) Fetch(c context.Context) ([]domain.GiftCard, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.giftcardRepository.Fetch(ctx)
}
