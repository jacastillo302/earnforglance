package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/public"
)

type configurationUsecase struct {
	itemRepository domain.ConfigurationRepository
	contextTimeout time.Duration
}

func NewConfigurationUsecase(itemRepository domain.ConfigurationRepository, timeout time.Duration) domain.ConfigurationRepository {
	return &configurationUsecase{
		itemRepository: itemRepository,
		contextTimeout: timeout,
	}
}

func (cu *configurationUsecase) GetConfigurations(c context.Context, filter domain.ConfigurationRequest) ([]domain.ConfigurationsResponse, error) {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.itemRepository.GetConfigurations(ctx, filter)
}
