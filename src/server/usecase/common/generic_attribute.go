package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/common"
)

type genericattributeUsecase struct {
	genericattributeRepository domain.GenericAttributeRepository
	contextTimeout             time.Duration
}

func NewGenericAttributeUsecase(genericattributeRepository domain.GenericAttributeRepository, timeout time.Duration) domain.GenericAttributeUsecase {
	return &genericattributeUsecase{
		genericattributeRepository: genericattributeRepository,
		contextTimeout:             timeout,
	}
}

func (tu *genericattributeUsecase) Create(c context.Context, genericattribute *domain.GenericAttribute) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.genericattributeRepository.Create(ctx, genericattribute)
}

func (tu *genericattributeUsecase) Update(c context.Context, genericattribute *domain.GenericAttribute) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.genericattributeRepository.Update(ctx, genericattribute)
}

func (tu *genericattributeUsecase) Delete(c context.Context, genericattribute *domain.GenericAttribute) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.genericattributeRepository.Delete(ctx, genericattribute)
}

func (lu *genericattributeUsecase) FetchByID(c context.Context, genericattributeID string) (domain.GenericAttribute, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.genericattributeRepository.FetchByID(ctx, genericattributeID)
}

func (lu *genericattributeUsecase) Fetch(c context.Context) ([]domain.GenericAttribute, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.genericattributeRepository.Fetch(ctx)
}
