package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/common"
)

type addressAttributeValueUsecase struct {
	addressAttributeValueRepository domain.AddressAttributeValueRepository
	contextTimeout                  time.Duration
}

func NewAddressAttributeValueUsecase(addressAttributeValueRepository domain.AddressAttributeValueRepository, timeout time.Duration) domain.AddressAttributeValueUsecase {
	return &addressAttributeValueUsecase{
		addressAttributeValueRepository: addressAttributeValueRepository,
		contextTimeout:                  timeout,
	}
}

func (tu *addressAttributeValueUsecase) Create(c context.Context, addressAttributeValue *domain.AddressAttributeValue) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.addressAttributeValueRepository.Create(ctx, addressAttributeValue)
}

func (tu *addressAttributeValueUsecase) Update(c context.Context, addressAttributeValue *domain.AddressAttributeValue) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.addressAttributeValueRepository.Update(ctx, addressAttributeValue)
}

func (tu *addressAttributeValueUsecase) Delete(c context.Context, addressAttributeValue *domain.AddressAttributeValue) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.addressAttributeValueRepository.Delete(ctx, addressAttributeValue)
}

func (lu *addressAttributeValueUsecase) FetchByID(c context.Context, addressAttributeValueID string) (domain.AddressAttributeValue, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.addressAttributeValueRepository.FetchByID(ctx, addressAttributeValueID)
}

func (lu *addressAttributeValueUsecase) Fetch(c context.Context) ([]domain.AddressAttributeValue, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.addressAttributeValueRepository.Fetch(ctx)
}
