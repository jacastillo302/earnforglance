package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/localization"
)

type localizedpropertyUsecase struct {
	localizedpropertyRepository domain.LocalizedPropertyRepository
	contextTimeout              time.Duration
}

func NewLocalizedPropertyUsecase(localizedpropertyRepository domain.LocalizedPropertyRepository, timeout time.Duration) domain.LocalizedPropertyUsecase {
	return &localizedpropertyUsecase{
		localizedpropertyRepository: localizedpropertyRepository,
		contextTimeout:              timeout,
	}
}

func (tu *localizedpropertyUsecase) Create(c context.Context, localizedproperty *domain.LocalizedProperty) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.localizedpropertyRepository.Create(ctx, localizedproperty)
}

func (tu *localizedpropertyUsecase) Update(c context.Context, localizedproperty *domain.LocalizedProperty) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.localizedpropertyRepository.Update(ctx, localizedproperty)
}

func (tu *localizedpropertyUsecase) Delete(c context.Context, localizedproperty *domain.LocalizedProperty) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.localizedpropertyRepository.Delete(ctx, localizedproperty)
}

func (lu *localizedpropertyUsecase) FetchByID(c context.Context, localizedpropertyID string) (domain.LocalizedProperty, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.localizedpropertyRepository.FetchByID(ctx, localizedpropertyID)
}

func (lu *localizedpropertyUsecase) Fetch(c context.Context) ([]domain.LocalizedProperty, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.localizedpropertyRepository.Fetch(ctx)
}
