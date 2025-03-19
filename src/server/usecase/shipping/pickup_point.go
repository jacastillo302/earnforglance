package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/shipping"
)

type pickuppointUsecase struct {
	pickuppointRepository domain.PickupPointRepository
	contextTimeout        time.Duration
}

func NewPickupPointUsecase(pickuppointRepository domain.PickupPointRepository, timeout time.Duration) domain.PickupPointUsecase {
	return &pickuppointUsecase{
		pickuppointRepository: pickuppointRepository,
		contextTimeout:        timeout,
	}
}

func (tu *pickuppointUsecase) Create(c context.Context, pickuppoint *domain.PickupPoint) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.pickuppointRepository.Create(ctx, pickuppoint)
}

func (tu *pickuppointUsecase) Update(c context.Context, pickuppoint *domain.PickupPoint) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.pickuppointRepository.Update(ctx, pickuppoint)
}

func (tu *pickuppointUsecase) Delete(c context.Context, pickuppoint *domain.PickupPoint) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.pickuppointRepository.Delete(ctx, pickuppoint)
}

func (lu *pickuppointUsecase) FetchByID(c context.Context, pickuppointID string) (domain.PickupPoint, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.pickuppointRepository.FetchByID(ctx, pickuppointID)
}

func (lu *pickuppointUsecase) Fetch(c context.Context) ([]domain.PickupPoint, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.pickuppointRepository.Fetch(ctx)
}
