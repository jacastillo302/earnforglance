package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/orders"
)

type orderUsecase struct {
	orderRepository domain.OrderRepository
	contextTimeout  time.Duration
}

func NewOrderUsecase(orderRepository domain.OrderRepository, timeout time.Duration) domain.OrderUsecase {
	return &orderUsecase{
		orderRepository: orderRepository,
		contextTimeout:  timeout,
	}
}

func (ou *orderUsecase) Create(c context.Context, order *domain.Order) error {
	ctx, cancel := context.WithTimeout(c, ou.contextTimeout)
	defer cancel()
	return ou.orderRepository.Create(ctx, order)
}

func (ou *orderUsecase) Update(c context.Context, order *domain.Order) error {
	ctx, cancel := context.WithTimeout(c, ou.contextTimeout)
	defer cancel()
	return ou.orderRepository.Update(ctx, order)
}

func (ou *orderUsecase) Delete(c context.Context, order string) error {
	ctx, cancel := context.WithTimeout(c, ou.contextTimeout)
	defer cancel()
	return ou.orderRepository.Delete(ctx, order)
}

func (ou *orderUsecase) FetchByID(c context.Context, orderID string) (domain.Order, error) {
	ctx, cancel := context.WithTimeout(c, ou.contextTimeout)
	defer cancel()
	return ou.orderRepository.FetchByID(ctx, orderID)
}

func (ou *orderUsecase) Fetch(c context.Context) ([]domain.Order, error) {
	ctx, cancel := context.WithTimeout(c, ou.contextTimeout)
	defer cancel()
	return ou.orderRepository.Fetch(ctx)
}
