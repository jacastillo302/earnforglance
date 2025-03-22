package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/localization"
)

type localestringresourceUsecase struct {
	localestringresourceRepository domain.LocaleStringResourceRepository
	contextTimeout                 time.Duration
}

func NewLocaleStringResourceUsecase(localestringresourceRepository domain.LocaleStringResourceRepository, timeout time.Duration) domain.LocaleStringResourceUsecase {
	return &localestringresourceUsecase{
		localestringresourceRepository: localestringresourceRepository,
		contextTimeout:                 timeout,
	}
}

func (tu *localestringresourceUsecase) Create(c context.Context, localestringresource *domain.LocaleStringResource) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.localestringresourceRepository.Create(ctx, localestringresource)
}

func (tu *localestringresourceUsecase) Update(c context.Context, localestringresource *domain.LocaleStringResource) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.localestringresourceRepository.Update(ctx, localestringresource)
}

func (tu *localestringresourceUsecase) Delete(c context.Context, localestringresource string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.localestringresourceRepository.Delete(ctx, localestringresource)
}

func (lu *localestringresourceUsecase) FetchByID(c context.Context, localestringresourceID string) (domain.LocaleStringResource, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.localestringresourceRepository.FetchByID(ctx, localestringresourceID)
}

func (lu *localestringresourceUsecase) Fetch(c context.Context) ([]domain.LocaleStringResource, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.localestringresourceRepository.Fetch(ctx)
}
