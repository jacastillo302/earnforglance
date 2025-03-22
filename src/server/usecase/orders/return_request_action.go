package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/orders"
)

type returnrequestactionUsecase struct {
	returnrequestactionRepository domain.ReturnRequestActionRepository
	contextTimeout                time.Duration
}

func NewReturnRequestActionUsecase(returnrequestactionRepository domain.ReturnRequestActionRepository, timeout time.Duration) domain.ReturnRequestActionUsecase {
	return &returnrequestactionUsecase{
		returnrequestactionRepository: returnrequestactionRepository,
		contextTimeout:                timeout,
	}
}

func (tu *returnrequestactionUsecase) Create(c context.Context, returnrequestaction *domain.ReturnRequestAction) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.returnrequestactionRepository.Create(ctx, returnrequestaction)
}

func (tu *returnrequestactionUsecase) Update(c context.Context, returnrequestaction *domain.ReturnRequestAction) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.returnrequestactionRepository.Update(ctx, returnrequestaction)
}

func (tu *returnrequestactionUsecase) Delete(c context.Context, returnrequestaction string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.returnrequestactionRepository.Delete(ctx, returnrequestaction)
}

func (lu *returnrequestactionUsecase) FetchByID(c context.Context, returnrequestactionID string) (domain.ReturnRequestAction, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.returnrequestactionRepository.FetchByID(ctx, returnrequestactionID)
}

func (lu *returnrequestactionUsecase) Fetch(c context.Context) ([]domain.ReturnRequestAction, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.returnrequestactionRepository.Fetch(ctx)
}
