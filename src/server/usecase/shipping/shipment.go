package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/shipping"
)

type shipmentUsecase struct {
	shipmentRepository domain.ShipmentRepository
	contextTimeout     time.Duration
}

func NewShipmentUsecase(shipmentRepository domain.ShipmentRepository, timeout time.Duration) domain.ShipmentUsecase {
	return &shipmentUsecase{
		shipmentRepository: shipmentRepository,
		contextTimeout:     timeout,
	}
}

func (tu *shipmentUsecase) Create(c context.Context, shipment *domain.Shipment) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.shipmentRepository.Create(ctx, shipment)
}

func (tu *shipmentUsecase) Update(c context.Context, shipment *domain.Shipment) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.shipmentRepository.Update(ctx, shipment)
}

func (tu *shipmentUsecase) Delete(c context.Context, shipment *domain.Shipment) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.shipmentRepository.Delete(ctx, shipment)
}

func (lu *shipmentUsecase) FetchByID(c context.Context, shipmentID string) (domain.Shipment, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.shipmentRepository.FetchByID(ctx, shipmentID)
}

func (lu *shipmentUsecase) Fetch(c context.Context) ([]domain.Shipment, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.shipmentRepository.Fetch(ctx)
}
