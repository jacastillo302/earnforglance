package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/shipping"
)

type shipmentitemUsecase struct {
	shipmentitemRepository domain.ShipmentItemRepository
	contextTimeout         time.Duration
}

func NewShipmentItemUsecase(shipmentitemRepository domain.ShipmentItemRepository, timeout time.Duration) domain.ShipmentItemUsecase {
	return &shipmentitemUsecase{
		shipmentitemRepository: shipmentitemRepository,
		contextTimeout:         timeout,
	}
}

func (tu *shipmentitemUsecase) Create(c context.Context, shipmentitem *domain.ShipmentItem) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.shipmentitemRepository.Create(ctx, shipmentitem)
}

func (tu *shipmentitemUsecase) Update(c context.Context, shipmentitem *domain.ShipmentItem) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.shipmentitemRepository.Update(ctx, shipmentitem)
}

func (tu *shipmentitemUsecase) Delete(c context.Context, shipmentitem *domain.ShipmentItem) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.shipmentitemRepository.Delete(ctx, shipmentitem)
}

func (lu *shipmentitemUsecase) FetchByID(c context.Context, shipmentitemID string) (domain.ShipmentItem, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.shipmentitemRepository.FetchByID(ctx, shipmentitemID)
}

func (lu *shipmentitemUsecase) Fetch(c context.Context) ([]domain.ShipmentItem, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.shipmentitemRepository.Fetch(ctx)
}
