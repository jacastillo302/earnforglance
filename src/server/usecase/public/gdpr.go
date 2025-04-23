package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/public"
)

type gdprConsentUsecase struct {
	itemRepository domain.GdprConsentRepository
	contextTimeout time.Duration
}

func NewGdprConsentUsecase(itemRepository domain.GdprConsentRepository, timeout time.Duration) domain.GdprConsentRepository {
	return &gdprConsentUsecase{
		itemRepository: itemRepository,
		contextTimeout: timeout,
	}
}

func (cu *gdprConsentUsecase) GetGdprConsents(c context.Context, filter domain.GdprConsentRequest) ([]domain.GdprConsentsResponse, error) {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.itemRepository.GetGdprConsents(ctx, filter)
}
