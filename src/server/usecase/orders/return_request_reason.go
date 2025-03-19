package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/orders"
)

type returnrequestreasonUsecase struct {
	returnrequestreasonRepository domain.ReturnRequestReasonRepository
	contextTimeout                time.Duration
}

func NewReturnRequestReasonUsecase(returnrequestreasonRepository domain.ReturnRequestReasonRepository, timeout time.Duration) domain.ReturnRequestReasonUsecase {
	return &returnrequestreasonUsecase{
		returnrequestreasonRepository: returnrequestreasonRepository,
		contextTimeout:                timeout,
	}
}

func (tu *returnrequestreasonUsecase) Create(c context.Context, returnrequestreason *domain.ReturnRequestReason) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.returnrequestreasonRepository.Create(ctx, returnrequestreason)
}

func (tu *returnrequestreasonUsecase) Update(c context.Context, returnrequestreason *domain.ReturnRequestReason) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.returnrequestreasonRepository.Update(ctx, returnrequestreason)
}

func (tu *returnrequestreasonUsecase) Delete(c context.Context, returnrequestreason *domain.ReturnRequestReason) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.returnrequestreasonRepository.Delete(ctx, returnrequestreason)
}

func (lu *returnrequestreasonUsecase) FetchByID(c context.Context, returnrequestreasonID string) (domain.ReturnRequestReason, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.returnrequestreasonRepository.FetchByID(ctx, returnrequestreasonID)
}

func (lu *returnrequestreasonUsecase) Fetch(c context.Context) ([]domain.ReturnRequestReason, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.returnrequestreasonRepository.Fetch(ctx)
}
