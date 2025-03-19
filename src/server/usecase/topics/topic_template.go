package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/topics"
)

type topicTemplateUsecase struct {
	topicTemplateRepository domain.TopicTemplateRepository
	contextTimeout          time.Duration
}

func NewTopicTemplateUsecase(topicTemplateRepository domain.TopicTemplateRepository, timeout time.Duration) domain.TopicTemplateUsecase {
	return &topicTemplateUsecase{
		topicTemplateRepository: topicTemplateRepository,
		contextTimeout:          timeout,
	}
}

func (tu *topicTemplateUsecase) Create(c context.Context, topicTemplate *domain.TopicTemplate) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.topicTemplateRepository.Create(ctx, topicTemplate)
}

func (tu *topicTemplateUsecase) Update(c context.Context, topicTemplate *domain.TopicTemplate) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.topicTemplateRepository.Update(ctx, topicTemplate)
}

func (tu *topicTemplateUsecase) Delete(c context.Context, topicTemplate *domain.TopicTemplate) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.topicTemplateRepository.Delete(ctx, topicTemplate)
}

func (lu *topicTemplateUsecase) FetchByID(c context.Context, topicTemplateID string) (domain.TopicTemplate, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.topicTemplateRepository.FetchByID(ctx, topicTemplateID)
}

func (lu *topicTemplateUsecase) Fetch(c context.Context) ([]domain.TopicTemplate, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.topicTemplateRepository.Fetch(ctx)
}
