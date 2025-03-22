package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/orders"
)

type returnrequestUsecase struct {
	returnrequestRepository domain.ReturnRequestRepository
	contextTimeout          time.Duration
}

func NewReturnRequestUsecase(returnrequestRepository domain.ReturnRequestRepository, timeout time.Duration) domain.ReturnRequestUsecase {
	return &returnrequestUsecase{
		returnrequestRepository: returnrequestRepository,
		contextTimeout:          timeout,
	}
}

func (tu *returnrequestUsecase) Create(c context.Context, returnrequest *domain.ReturnRequest) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.returnrequestRepository.Create(ctx, returnrequest)
}

func (tu *returnrequestUsecase) Update(c context.Context, returnrequest *domain.ReturnRequest) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.returnrequestRepository.Update(ctx, returnrequest)
}

func (tu *returnrequestUsecase) Delete(c context.Context, returnrequest string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.returnrequestRepository.Delete(ctx, returnrequest)
}

func (lu *returnrequestUsecase) FetchByID(c context.Context, returnrequestID string) (domain.ReturnRequest, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.returnrequestRepository.FetchByID(ctx, returnrequestID)
}

func (lu *returnrequestUsecase) Fetch(c context.Context) ([]domain.ReturnRequest, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.returnrequestRepository.Fetch(ctx)
}
