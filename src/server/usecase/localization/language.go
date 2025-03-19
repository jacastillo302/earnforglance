package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/localization"
)

type languageUsecase struct {
	languageRepository domain.LanguageRepository
	contextTimeout     time.Duration
}

func NewLanguageUsecase(languageRepository domain.LanguageRepository, timeout time.Duration) domain.LanguageUsecase {
	return &languageUsecase{
		languageRepository: languageRepository,
		contextTimeout:     timeout,
	}
}

func (lu *languageUsecase) Create(c context.Context, language *domain.Language) error {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.languageRepository.Create(ctx, language)
}

func (lu *languageUsecase) Update(c context.Context, language *domain.Language) error {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.languageRepository.Update(ctx, language)
}

func (lu *languageUsecase) Delete(c context.Context, language *domain.Language) error {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.languageRepository.Delete(ctx, language)
}

func (lu *languageUsecase) FetchByID(c context.Context, languageID string) (domain.Language, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.languageRepository.FetchByID(ctx, languageID)
}

func (lu *languageUsecase) Fetch(c context.Context) ([]domain.Language, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.languageRepository.Fetch(ctx)
}
