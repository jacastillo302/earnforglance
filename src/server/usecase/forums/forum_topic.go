package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/forums"
)

type forumtopicUsecase struct {
	forumtopicRepository domain.ForumTopicRepository
	contextTimeout       time.Duration
}

func NewForumTopicUsecase(forumtopicRepository domain.ForumTopicRepository, timeout time.Duration) domain.ForumTopicUsecase {
	return &forumtopicUsecase{
		forumtopicRepository: forumtopicRepository,
		contextTimeout:       timeout,
	}
}

func (tu *forumtopicUsecase) Create(c context.Context, forumtopic *domain.ForumTopic) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.forumtopicRepository.Create(ctx, forumtopic)
}

func (tu *forumtopicUsecase) Update(c context.Context, forumtopic *domain.ForumTopic) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.forumtopicRepository.Update(ctx, forumtopic)
}

func (tu *forumtopicUsecase) Delete(c context.Context, forumtopic *domain.ForumTopic) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.forumtopicRepository.Delete(ctx, forumtopic)
}

func (lu *forumtopicUsecase) FetchByID(c context.Context, forumtopicID string) (domain.ForumTopic, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.forumtopicRepository.FetchByID(ctx, forumtopicID)
}

func (lu *forumtopicUsecase) Fetch(c context.Context) ([]domain.ForumTopic, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.forumtopicRepository.Fetch(ctx)
}
