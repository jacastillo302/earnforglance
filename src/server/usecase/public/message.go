package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/public"
)

type newsLetterUsecase struct {
	itemRepository domain.NewsLetterRepository
	contextTimeout time.Duration
}

func NewNewsLetterUsecase(itemRepository domain.NewsLetterRepository, timeout time.Duration) domain.NewsLetterRepository {
	return &newsLetterUsecase{
		itemRepository: itemRepository,
		contextTimeout: timeout,
	}
}

func (cu *newsLetterUsecase) NewsLetterSubscription(c context.Context, filter domain.NewsLetterRequest, IpAdress string) (domain.NewsLetterResponse, error) {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.itemRepository.NewsLetterSubscription(ctx, filter, IpAdress)
}
