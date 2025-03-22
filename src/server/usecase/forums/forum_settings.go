package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/forums"
)

type forumSettingsUsecase struct {
	forumSettingsRepository domain.ForumSettingsRepository
	contextTimeout          time.Duration
}

func NewForumSettingsUsecase(forumSettingsRepository domain.ForumSettingsRepository, timeout time.Duration) domain.ForumSettingsUsecase {
	return &forumSettingsUsecase{
		forumSettingsRepository: forumSettingsRepository,
		contextTimeout:          timeout,
	}
}

func (tu *forumSettingsUsecase) Create(c context.Context, forumSettings *domain.ForumSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.forumSettingsRepository.Create(ctx, forumSettings)
}

func (tu *forumSettingsUsecase) Update(c context.Context, forumSettings *domain.ForumSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.forumSettingsRepository.Update(ctx, forumSettings)
}

func (tu *forumSettingsUsecase) Delete(c context.Context, forumSettings string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.forumSettingsRepository.Delete(ctx, forumSettings)
}

func (lu *forumSettingsUsecase) FetchByID(c context.Context, forumSettingsID string) (domain.ForumSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.forumSettingsRepository.FetchByID(ctx, forumSettingsID)
}

func (lu *forumSettingsUsecase) Fetch(c context.Context) ([]domain.ForumSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.forumSettingsRepository.Fetch(ctx)
}
