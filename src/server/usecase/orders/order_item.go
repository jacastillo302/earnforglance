package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/orders"
)

type orderitemUsecase struct {
	orderitemRepository domain.OrderItemRepository
	contextTimeout      time.Duration
}

func NewOrderItemUsecase(orderitemRepository domain.OrderItemRepository, timeout time.Duration) domain.OrderItemUsecase {
	return &orderitemUsecase{
		orderitemRepository: orderitemRepository,
		contextTimeout:      timeout,
	}
}

func (tu *orderitemUsecase) Create(c context.Context, orderitem *domain.OrderItem) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.orderitemRepository.Create(ctx, orderitem)
}

func (tu *orderitemUsecase) Update(c context.Context, orderitem *domain.OrderItem) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.orderitemRepository.Update(ctx, orderitem)
}

func (tu *orderitemUsecase) Delete(c context.Context, orderitem string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.orderitemRepository.Delete(ctx, orderitem)
}

func (lu *orderitemUsecase) FetchByID(c context.Context, orderitemID string) (domain.OrderItem, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.orderitemRepository.FetchByID(ctx, orderitemID)
}

func (lu *orderitemUsecase) Fetch(c context.Context) ([]domain.OrderItem, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.orderitemRepository.Fetch(ctx)
}
