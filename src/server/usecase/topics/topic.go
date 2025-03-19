package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/topics"
)

type topicUsecase struct {
	topicRepository domain.TopicRepository
	contextTimeout  time.Duration
}

func NewTopicUsecase(topicRepository domain.TopicRepository, timeout time.Duration) domain.TopicUsecase {
	return &topicUsecase{
		topicRepository: topicRepository,
		contextTimeout:  timeout,
	}
}

func (tu *topicUsecase) Create(c context.Context, topic *domain.Topic) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.topicRepository.Create(ctx, topic)
}

func (tu *topicUsecase) Update(c context.Context, topic *domain.Topic) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.topicRepository.Update(ctx, topic)
}

func (tu *topicUsecase) Delete(c context.Context, topic *domain.Topic) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.topicRepository.Delete(ctx, topic)
}

func (lu *topicUsecase) FetchByID(c context.Context, topicID string) (domain.Topic, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.topicRepository.FetchByID(ctx, topicID)
}

func (lu *topicUsecase) Fetch(c context.Context) ([]domain.Topic, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.topicRepository.Fetch(ctx)
}
