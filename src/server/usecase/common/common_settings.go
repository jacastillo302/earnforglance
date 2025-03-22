package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/common"
)

type commonsettingsUsecase struct {
	commonsettingsRepository domain.CommonSettingsRepository
	contextTimeout           time.Duration
}

func NewCommonSettingsUsecase(commonsettingsRepository domain.CommonSettingsRepository, timeout time.Duration) domain.CommonSettingsUsecase {
	return &commonsettingsUsecase{
		commonsettingsRepository: commonsettingsRepository,
		contextTimeout:           timeout,
	}
}

func (tu *commonsettingsUsecase) Create(c context.Context, commonsettings *domain.CommonSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.commonsettingsRepository.Create(ctx, commonsettings)
}

func (tu *commonsettingsUsecase) Update(c context.Context, commonsettings *domain.CommonSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.commonsettingsRepository.Update(ctx, commonsettings)
}

func (tu *commonsettingsUsecase) Delete(c context.Context, commonsettings string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.commonsettingsRepository.Delete(ctx, commonsettings)
}

func (lu *commonsettingsUsecase) FetchByID(c context.Context, commonsettingsID string) (domain.CommonSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.commonsettingsRepository.FetchByID(ctx, commonsettingsID)
}

func (lu *commonsettingsUsecase) Fetch(c context.Context) ([]domain.CommonSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.commonsettingsRepository.Fetch(ctx)
}
