package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/polls"
)

type pollUsecase struct {
	pollRepository domain.PollRepository
	contextTimeout time.Duration
}

func NewPollUsecase(pollRepository domain.PollRepository, timeout time.Duration) domain.PollUsecase {
	return &pollUsecase{
		pollRepository: pollRepository,
		contextTimeout: timeout,
	}
}

func (tu *pollUsecase) Create(c context.Context, poll *domain.Poll) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.pollRepository.Create(ctx, poll)
}

func (tu *pollUsecase) Update(c context.Context, poll *domain.Poll) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.pollRepository.Update(ctx, poll)
}

func (tu *pollUsecase) Delete(c context.Context, poll *domain.Poll) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.pollRepository.Delete(ctx, poll)
}

func (lu *pollUsecase) FetchByID(c context.Context, pollID string) (domain.Poll, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.pollRepository.FetchByID(ctx, pollID)
}

func (lu *pollUsecase) Fetch(c context.Context) ([]domain.Poll, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.pollRepository.Fetch(ctx)
}
