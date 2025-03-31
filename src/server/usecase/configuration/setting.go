package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/configuration"
)

type settingUsecase struct {
	settingRepository domain.SettingRepository
	contextTimeout    time.Duration
}

func NewSettingUsecase(settingRepository domain.SettingRepository, timeout time.Duration) domain.SettingUsecase {
	return &settingUsecase{
		settingRepository: settingRepository,
		contextTimeout:    timeout,
	}
}

func (tu *settingUsecase) CreateMany(c context.Context, items []domain.Setting) error {
	ctx,
		cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.settingRepository.CreateMany(ctx, items)
}

func (tu *settingUsecase) Create(c context.Context, setting *domain.Setting) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.settingRepository.Create(ctx, setting)
}

func (tu *settingUsecase) Update(c context.Context, setting *domain.Setting) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.settingRepository.Update(ctx, setting)
}

func (tu *settingUsecase) Delete(c context.Context, setting string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.settingRepository.Delete(ctx, setting)
}

func (lu *settingUsecase) FetchByID(c context.Context, settingID string) (domain.Setting, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.settingRepository.FetchByID(ctx, settingID)
}

func (lu *settingUsecase) Fetch(c context.Context) ([]domain.Setting, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.settingRepository.Fetch(ctx)
}

func (lu *settingUsecase) FetchByName(c context.Context, name string) (domain.Setting, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.settingRepository.FetchByName(ctx, name)
}

func (lu *settingUsecase) FetchByNames(c context.Context, names []string) ([]domain.Setting, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.settingRepository.FetchByNames(ctx, names)
}
