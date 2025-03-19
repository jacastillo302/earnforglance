package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/messages"
)

type emailaccountsettingsUsecase struct {
	emailaccountsettingsRepository domain.EmailAccountSettingsRepository
	contextTimeout                 time.Duration
}

func NewEmailAccountSettingsUsecase(emailaccountsettingsRepository domain.EmailAccountSettingsRepository, timeout time.Duration) domain.EmailAccountSettingsUsecase {
	return &emailaccountsettingsUsecase{
		emailaccountsettingsRepository: emailaccountsettingsRepository,
		contextTimeout:                 timeout,
	}
}

func (tu *emailaccountsettingsUsecase) Create(c context.Context, emailaccountsettings *domain.EmailAccountSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.emailaccountsettingsRepository.Create(ctx, emailaccountsettings)
}

func (tu *emailaccountsettingsUsecase) Update(c context.Context, emailaccountsettings *domain.EmailAccountSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.emailaccountsettingsRepository.Update(ctx, emailaccountsettings)
}

func (tu *emailaccountsettingsUsecase) Delete(c context.Context, emailaccountsettings *domain.EmailAccountSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.emailaccountsettingsRepository.Delete(ctx, emailaccountsettings)
}

func (lu *emailaccountsettingsUsecase) FetchByID(c context.Context, emailaccountsettingsID string) (domain.EmailAccountSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.emailaccountsettingsRepository.FetchByID(ctx, emailaccountsettingsID)
}

func (lu *emailaccountsettingsUsecase) Fetch(c context.Context) ([]domain.EmailAccountSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.emailaccountsettingsRepository.Fetch(ctx)
}
