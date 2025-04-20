package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/public"
)

type topicUsecase struct {
	itemRepository domain.TopicRepository
	contextTimeout time.Duration
}

func NewtopicUsecase(itemRepository domain.TopicRepository, timeout time.Duration) domain.TopictUsecase {
	return &topicUsecase{
		itemRepository: itemRepository,
		contextTimeout: timeout,
	}
}

func (r *topicUsecase) GetTopics(c context.Context, filter domain.TopicRequest) ([]domain.TopicsResponse, error) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()
	return r.itemRepository.GetTopics(ctx, filter)
}
