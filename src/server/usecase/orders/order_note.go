package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/orders"
)

type orderNoteUsecase struct {
	orderNoteRepository domain.OrderNoteRepository
	contextTimeout      time.Duration
}

func NewOrderNoteUsecase(orderNoteRepository domain.OrderNoteRepository, timeout time.Duration) domain.OrderNoteUsecase {
	return &orderNoteUsecase{
		orderNoteRepository: orderNoteRepository,
		contextTimeout:      timeout,
	}
}

func (tu *orderNoteUsecase) Create(c context.Context, orderNote *domain.OrderNote) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.orderNoteRepository.Create(ctx, orderNote)
}

func (tu *orderNoteUsecase) Update(c context.Context, orderNote *domain.OrderNote) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.orderNoteRepository.Update(ctx, orderNote)
}

func (tu *orderNoteUsecase) Delete(c context.Context, orderNote *domain.OrderNote) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.orderNoteRepository.Delete(ctx, orderNote)
}

func (lu *orderNoteUsecase) FetchByID(c context.Context, orderNoteID string) (domain.OrderNote, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.orderNoteRepository.FetchByID(ctx, orderNoteID)
}

func (lu *orderNoteUsecase) Fetch(c context.Context) ([]domain.OrderNote, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.orderNoteRepository.Fetch(ctx)
}
