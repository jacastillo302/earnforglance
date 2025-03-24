package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/security"
)

type captchaSettingsUsecase struct {
	captchaSettingsRepository domain.CaptchaSettingsRepository
	contextTimeout            time.Duration
}

func NewCaptchaSettingsUsecase(captchaSettingsRepository domain.CaptchaSettingsRepository, timeout time.Duration) domain.CaptchaSettingsUsecase {
	return &captchaSettingsUsecase{
		captchaSettingsRepository: captchaSettingsRepository,
		contextTimeout:            timeout,
	}
}

func (tu *captchaSettingsUsecase) CreateMany(c context.Context, MessageTemplatesSettingsList []domain.CaptchaSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.captchaSettingsRepository.CreateMany(ctx, MessageTemplatesSettingsList)
}

func (tu *captchaSettingsUsecase) Create(c context.Context, captchaSettings *domain.CaptchaSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.captchaSettingsRepository.Create(ctx, captchaSettings)
}

func (tu *captchaSettingsUsecase) Update(c context.Context, captchaSettings *domain.CaptchaSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.captchaSettingsRepository.Update(ctx, captchaSettings)
}

func (tu *captchaSettingsUsecase) Delete(c context.Context, captchaSettings string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.captchaSettingsRepository.Delete(ctx, captchaSettings)
}

func (lu *captchaSettingsUsecase) FetchByID(c context.Context, captchaSettingsID string) (domain.CaptchaSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.captchaSettingsRepository.FetchByID(ctx, captchaSettingsID)
}

func (lu *captchaSettingsUsecase) Fetch(c context.Context) ([]domain.CaptchaSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.captchaSettingsRepository.Fetch(ctx)
}
