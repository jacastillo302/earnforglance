package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/media"
)

type pictureUsecase struct {
	pictureRepository domain.PictureRepository
	contextTimeout    time.Duration
}

func NewPictureUsecase(pictureRepository domain.PictureRepository, timeout time.Duration) domain.PictureUsecase {
	return &pictureUsecase{
		pictureRepository: pictureRepository,
		contextTimeout:    timeout,
	}
}

func (tu *pictureUsecase) CreateMany(c context.Context, items []domain.Picture) error {
	ctx,
		cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.pictureRepository.CreateMany(ctx, items)
}

func (pu *pictureUsecase) Create(c context.Context, picture *domain.Picture) error {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	return pu.pictureRepository.Create(ctx, picture)
}

func (pu *pictureUsecase) Update(c context.Context, picture *domain.Picture) error {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	return pu.pictureRepository.Update(ctx, picture)
}

func (pu *pictureUsecase) Delete(c context.Context, picture string) error {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	return pu.pictureRepository.Delete(ctx, picture)
}

func (pu *pictureUsecase) FetchByID(c context.Context, pictureID string) (domain.Picture, error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	return pu.pictureRepository.FetchByID(ctx, pictureID)
}

func (pu *pictureUsecase) Fetch(c context.Context) ([]domain.Picture, error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	return pu.pictureRepository.Fetch(ctx)
}
