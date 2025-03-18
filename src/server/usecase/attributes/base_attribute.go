package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/attributes"
)

type baseattributeUsecase struct {
	baseattributeRepository domain.BaseAttributeRepository
	contextTimeout          time.Duration
}

func NewBaseAttributeUsecase(baseattributeRepository domain.BaseAttributeRepository, timeout time.Duration) domain.BaseAttributeUsecase {
	return &baseattributeUsecase{
		baseattributeRepository: baseattributeRepository,
		contextTimeout:          timeout,
	}
}

func (tu *baseattributeUsecase) Create(c context.Context, baseattribute *domain.BaseAttribute) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.baseattributeRepository.Create(ctx, baseattribute)
}

func (tu *baseattributeUsecase) Update(c context.Context, baseattribute *domain.BaseAttribute) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.baseattributeRepository.Update(ctx, baseattribute)
}

func (tu *baseattributeUsecase) Delete(c context.Context, baseattribute *domain.BaseAttribute) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.baseattributeRepository.Delete(ctx, baseattribute)
}

func (lu *baseattributeUsecase) FetchByID(c context.Context, baseattributeID string) (domain.BaseAttribute, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.baseattributeRepository.FetchByID(ctx, baseattributeID)
}

func (lu *baseattributeUsecase) Fetch(c context.Context) ([]domain.BaseAttribute, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.baseattributeRepository.Fetch(ctx)
}
