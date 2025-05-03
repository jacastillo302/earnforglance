package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/public"
)

type customerUsecase struct {
	newsRepository domain.NewsLetterRepository
	itemRepository domain.CustomerRepository
	contextTimeout time.Duration
}

func NewCustomerUsecase(itemRepository domain.CustomerRepository, newsRepository domain.NewsLetterRepository, timeout time.Duration) domain.CustomerRepository {
	return &customerUsecase{
		itemRepository: itemRepository,
		contextTimeout: timeout,
	}
}

func (cu *customerUsecase) SingIn(c context.Context, sigin domain.SingInRequest) (domain.SingInResponse, error) {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.itemRepository.SingIn(ctx, sigin)
}

func (cu *customerUsecase) GetSlugs(c context.Context, record string) ([]string, error) {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.itemRepository.GetSlugs(ctx, record)
}
