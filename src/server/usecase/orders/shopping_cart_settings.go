package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/orders"
)

type shoppingCartSettingsUsecase struct {
	shoppingCartSettingsRepository domain.ShoppingCartSettingsRepository
	contextTimeout                 time.Duration
}

func NewShoppingCartSettingsUsecase(shoppingCartSettingsRepository domain.ShoppingCartSettingsRepository, timeout time.Duration) domain.ShoppingCartSettingsUsecase {
	return &shoppingCartSettingsUsecase{
		shoppingCartSettingsRepository: shoppingCartSettingsRepository,
		contextTimeout:                 timeout,
	}
}

func (tu *shoppingCartSettingsUsecase) Create(c context.Context, shoppingCartSettings *domain.ShoppingCartSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.shoppingCartSettingsRepository.Create(ctx, shoppingCartSettings)
}

func (tu *shoppingCartSettingsUsecase) Update(c context.Context, shoppingCartSettings *domain.ShoppingCartSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.shoppingCartSettingsRepository.Update(ctx, shoppingCartSettings)
}

func (tu *shoppingCartSettingsUsecase) Delete(c context.Context, shoppingCartSettings *domain.ShoppingCartSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.shoppingCartSettingsRepository.Delete(ctx, shoppingCartSettings)
}

func (lu *shoppingCartSettingsUsecase) FetchByID(c context.Context, shoppingCartSettingsID string) (domain.ShoppingCartSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.shoppingCartSettingsRepository.FetchByID(ctx, shoppingCartSettingsID)
}

func (lu *shoppingCartSettingsUsecase) Fetch(c context.Context) ([]domain.ShoppingCartSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.shoppingCartSettingsRepository.Fetch(ctx)
}
