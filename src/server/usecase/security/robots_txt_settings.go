package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/security"
)

type robotsTxtSettingsUsecase struct {
	robotsTxtSettingsRepository domain.RobotsTxtSettingsRepository
	contextTimeout              time.Duration
}

func NewRobotsTxtSettingsUsecase(robotsTxtSettingsRepository domain.RobotsTxtSettingsRepository, timeout time.Duration) domain.RobotsTxtSettingsUsecase {
	return &robotsTxtSettingsUsecase{
		robotsTxtSettingsRepository: robotsTxtSettingsRepository,
		contextTimeout:              timeout,
	}
}

func (tu *robotsTxtSettingsUsecase) Create(c context.Context, robotsTxtSettings *domain.RobotsTxtSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.robotsTxtSettingsRepository.Create(ctx, robotsTxtSettings)
}

func (tu *robotsTxtSettingsUsecase) Update(c context.Context, robotsTxtSettings *domain.RobotsTxtSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.robotsTxtSettingsRepository.Update(ctx, robotsTxtSettings)
}

func (tu *robotsTxtSettingsUsecase) Delete(c context.Context, robotsTxtSettings string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.robotsTxtSettingsRepository.Delete(ctx, robotsTxtSettings)
}

func (lu *robotsTxtSettingsUsecase) FetchByID(c context.Context, robotsTxtSettingsID string) (domain.RobotsTxtSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.robotsTxtSettingsRepository.FetchByID(ctx, robotsTxtSettingsID)
}

func (lu *robotsTxtSettingsUsecase) Fetch(c context.Context) ([]domain.RobotsTxtSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.robotsTxtSettingsRepository.Fetch(ctx)
}
