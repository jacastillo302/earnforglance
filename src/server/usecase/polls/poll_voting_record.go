package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/polls"
)

type pollVotingRecordUsecase struct {
	pollVotingRecordRepository domain.PollVotingRecordRepository
	contextTimeout             time.Duration
}

func NewPollVotingRecordUsecase(pollVotingRecordRepository domain.PollVotingRecordRepository, timeout time.Duration) domain.PollVotingRecordUsecase {
	return &pollVotingRecordUsecase{
		pollVotingRecordRepository: pollVotingRecordRepository,
		contextTimeout:             timeout,
	}
}

func (tu *pollVotingRecordUsecase) Create(c context.Context, pollVotingRecord *domain.PollVotingRecord) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.pollVotingRecordRepository.Create(ctx, pollVotingRecord)
}

func (tu *pollVotingRecordUsecase) Update(c context.Context, pollVotingRecord *domain.PollVotingRecord) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.pollVotingRecordRepository.Update(ctx, pollVotingRecord)
}

func (tu *pollVotingRecordUsecase) Delete(c context.Context, pollVotingRecord *domain.PollVotingRecord) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.pollVotingRecordRepository.Delete(ctx, pollVotingRecord)
}

func (lu *pollVotingRecordUsecase) FetchByID(c context.Context, pollVotingRecordID string) (domain.PollVotingRecord, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.pollVotingRecordRepository.FetchByID(ctx, pollVotingRecordID)
}

func (lu *pollVotingRecordUsecase) Fetch(c context.Context) ([]domain.PollVotingRecord, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.pollVotingRecordRepository.Fetch(ctx)
}
