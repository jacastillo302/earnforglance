package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/forums"
)

type forumgroupUsecase struct {
	forumgroupRepository domain.ForumGroupRepository
	contextTimeout       time.Duration
}

func NewForumGroupUsecase(forumgroupRepository domain.ForumGroupRepository, timeout time.Duration) domain.ForumGroupUsecase {
	return &forumgroupUsecase{
		forumgroupRepository: forumgroupRepository,
		contextTimeout:       timeout,
	}
}

func (tu *forumgroupUsecase) CreateMany(c context.Context, items []domain.ForumGroup) error {
	ctx,
		cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.forumgroupRepository.CreateMany(ctx, items)
}

func (tu *forumgroupUsecase) Create(c context.Context, forumgroup *domain.ForumGroup) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.forumgroupRepository.Create(ctx, forumgroup)
}

func (tu *forumgroupUsecase) Update(c context.Context, forumgroup *domain.ForumGroup) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.forumgroupRepository.Update(ctx, forumgroup)
}

func (tu *forumgroupUsecase) Delete(c context.Context, forumgroup string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.forumgroupRepository.Delete(ctx, forumgroup)
}

func (lu *forumgroupUsecase) FetchByID(c context.Context, forumgroupID string) (domain.ForumGroup, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.forumgroupRepository.FetchByID(ctx, forumgroupID)
}

func (lu *forumgroupUsecase) Fetch(c context.Context) ([]domain.ForumGroup, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.forumgroupRepository.Fetch(ctx)
}
