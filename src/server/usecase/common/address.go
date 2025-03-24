package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/common"
)

type addressUsecase struct {
	addressRepository domain.AddressRepository
	contextTimeout    time.Duration
}

func NewAddressUsecase(addressRepository domain.AddressRepository, timeout time.Duration) domain.AddressUsecase {
	return &addressUsecase{
		addressRepository: addressRepository,
		contextTimeout:    timeout,
	}
}

func (tu *addressUsecase) CreateMany(c context.Context, items []domain.Address) error {
	ctx,
		cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.addressRepository.CreateMany(ctx, items)
}

func (tu *addressUsecase) Create(c context.Context, address *domain.Address) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.addressRepository.Create(ctx, address)
}

func (tu *addressUsecase) Update(c context.Context, address *domain.Address) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.addressRepository.Update(ctx, address)
}

func (tu *addressUsecase) Delete(c context.Context, address string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.addressRepository.Delete(ctx, address)
}

func (lu *addressUsecase) FetchByID(c context.Context, addressID string) (domain.Address, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.addressRepository.FetchByID(ctx, addressID)
}

func (lu *addressUsecase) Fetch(c context.Context) ([]domain.Address, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.addressRepository.Fetch(ctx)
}
