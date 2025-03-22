package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/shipping"
)

type deliverydateUsecase struct {
	deliverydateRepository domain.DeliveryDateRepository
	contextTimeout         time.Duration
}

func NewDeliveryDateUsecase(deliverydateRepository domain.DeliveryDateRepository, timeout time.Duration) domain.DeliveryDateUsecase {
	return &deliverydateUsecase{
		deliverydateRepository: deliverydateRepository,
		contextTimeout:         timeout,
	}
}

func (tu *deliverydateUsecase) Create(c context.Context, deliverydate *domain.DeliveryDate) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.deliverydateRepository.Create(ctx, deliverydate)
}

func (tu *deliverydateUsecase) Update(c context.Context, deliverydate *domain.DeliveryDate) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.deliverydateRepository.Update(ctx, deliverydate)
}

func (tu *deliverydateUsecase) Delete(c context.Context, deliverydate string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.deliverydateRepository.Delete(ctx, deliverydate)
}

func (lu *deliverydateUsecase) FetchByID(c context.Context, deliverydateID string) (domain.DeliveryDate, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.deliverydateRepository.FetchByID(ctx, deliverydateID)
}

func (lu *deliverydateUsecase) Fetch(c context.Context) ([]domain.DeliveryDate, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.deliverydateRepository.Fetch(ctx)
}
