package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/orders"
)

type shoppingcartitemUsecase struct {
	shoppingcartitemRepository domain.ShoppingCartItemRepository
	contextTimeout             time.Duration
}

func NewShoppingCartItemUsecase(shoppingcartitemRepository domain.ShoppingCartItemRepository, timeout time.Duration) domain.ShoppingCartItemUsecase {
	return &shoppingcartitemUsecase{
		shoppingcartitemRepository: shoppingcartitemRepository,
		contextTimeout:             timeout,
	}
}

func (tu *shoppingcartitemUsecase) Create(c context.Context, shoppingcartitem *domain.ShoppingCartItem) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.shoppingcartitemRepository.Create(ctx, shoppingcartitem)
}

func (tu *shoppingcartitemUsecase) Update(c context.Context, shoppingcartitem *domain.ShoppingCartItem) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.shoppingcartitemRepository.Update(ctx, shoppingcartitem)
}

func (tu *shoppingcartitemUsecase) Delete(c context.Context, shoppingcartitem *domain.ShoppingCartItem) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.shoppingcartitemRepository.Delete(ctx, shoppingcartitem)
}

func (lu *shoppingcartitemUsecase) FetchByID(c context.Context, shoppingcartitemID string) (domain.ShoppingCartItem, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.shoppingcartitemRepository.FetchByID(ctx, shoppingcartitemID)
}

func (lu *shoppingcartitemUsecase) Fetch(c context.Context) ([]domain.ShoppingCartItem, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.shoppingcartitemRepository.Fetch(ctx)
}
