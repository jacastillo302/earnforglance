package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/common"
)

type addresssettingsUsecase struct {
	addresssettingsRepository domain.AddressSettingsRepository
	contextTimeout            time.Duration
}

func NewAddressSettingsUsecase(addresssettingsRepository domain.AddressSettingsRepository, timeout time.Duration) domain.AddressSettingsUsecase {
	return &addresssettingsUsecase{
		addresssettingsRepository: addresssettingsRepository,
		contextTimeout:            timeout,
	}
}

func (tu *addresssettingsUsecase) CreateMany(c context.Context, items []domain.AddressSettings) error {
	ctx,
		cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.addresssettingsRepository.CreateMany(ctx, items)
}

func (tu *addresssettingsUsecase) Create(c context.Context, addresssettings *domain.AddressSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.addresssettingsRepository.Create(ctx, addresssettings)
}

func (tu *addresssettingsUsecase) Update(c context.Context, addresssettings *domain.AddressSettings) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.addresssettingsRepository.Update(ctx, addresssettings)
}

func (tu *addresssettingsUsecase) Delete(c context.Context, addresssettings string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.addresssettingsRepository.Delete(ctx, addresssettings)
}

func (lu *addresssettingsUsecase) FetchByID(c context.Context, addresssettingsID string) (domain.AddressSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.addresssettingsRepository.FetchByID(ctx, addresssettingsID)
}

func (lu *addresssettingsUsecase) Fetch(c context.Context) ([]domain.AddressSettings, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.addresssettingsRepository.Fetch(ctx)
}
