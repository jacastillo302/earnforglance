package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/public"
)

type localizationUsecase struct {
	itemRepository domain.LocalizationRepository
	contextTimeout time.Duration
}

func NewlocalizationUsecase(itemRepository domain.LocalizationRepository, timeout time.Duration) domain.LocalizationUsecase {
	return &localizationUsecase{
		itemRepository: itemRepository,
		contextTimeout: timeout,
	}
}

func (r *localizationUsecase) GetLocalizations(c context.Context, filter domain.LocalizationRequest) ([]domain.LocalizationsResponse, error) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()
	return r.itemRepository.GetLocalizations(ctx, filter)
}
