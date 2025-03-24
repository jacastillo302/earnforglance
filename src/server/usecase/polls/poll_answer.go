package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/polls"
)

type pollAnswerUsecase struct {
	pollAnswerRepository domain.PollAnswerRepository
	contextTimeout       time.Duration
}

func NewPollAnswerUsecase(pollAnswerRepository domain.PollAnswerRepository, timeout time.Duration) domain.PollAnswerUsecase {
	return &pollAnswerUsecase{
		pollAnswerRepository: pollAnswerRepository,
		contextTimeout:       timeout,
	}
}

func (tu *pollAnswerUsecase) CreateMany(c context.Context, MessageTemplatesSettingsList []domain.PollAnswer) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.pollAnswerRepository.CreateMany(ctx, MessageTemplatesSettingsList)
}

func (tu *pollAnswerUsecase) Create(c context.Context, pollAnswer *domain.PollAnswer) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.pollAnswerRepository.Create(ctx, pollAnswer)
}

func (tu *pollAnswerUsecase) Update(c context.Context, pollAnswer *domain.PollAnswer) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.pollAnswerRepository.Update(ctx, pollAnswer)
}

func (tu *pollAnswerUsecase) Delete(c context.Context, pollAnswer string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.pollAnswerRepository.Delete(ctx, pollAnswer)
}

func (lu *pollAnswerUsecase) FetchByID(c context.Context, pollAnswerID string) (domain.PollAnswer, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.pollAnswerRepository.FetchByID(ctx, pollAnswerID)
}

func (lu *pollAnswerUsecase) Fetch(c context.Context) ([]domain.PollAnswer, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.pollAnswerRepository.Fetch(ctx)
}
