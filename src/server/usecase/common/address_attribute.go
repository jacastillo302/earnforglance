package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/common"
)

type addressattributeUsecase struct {
	addressattributeRepository domain.AddressAttributeRepository
	contextTimeout             time.Duration
}

func NewAddressAttributeUsecase(addressattributeRepository domain.AddressAttributeRepository, timeout time.Duration) domain.AddressAttributeUsecase {
	return &addressattributeUsecase{
		addressattributeRepository: addressattributeRepository,
		contextTimeout:             timeout,
	}
}

func (tu *addressattributeUsecase) Create(c context.Context, addressattribute *domain.AddressAttribute) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.addressattributeRepository.Create(ctx, addressattribute)
}

func (tu *addressattributeUsecase) Update(c context.Context, addressattribute *domain.AddressAttribute) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.addressattributeRepository.Update(ctx, addressattribute)
}

func (tu *addressattributeUsecase) Delete(c context.Context, addressattribute *domain.AddressAttribute) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.addressattributeRepository.Delete(ctx, addressattribute)
}

func (lu *addressattributeUsecase) FetchByID(c context.Context, addressattributeID string) (domain.AddressAttribute, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.addressattributeRepository.FetchByID(ctx, addressattributeID)
}

func (lu *addressattributeUsecase) Fetch(c context.Context) ([]domain.AddressAttribute, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.addressattributeRepository.Fetch(ctx)
}
