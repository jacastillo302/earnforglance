package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/media"
)

type PictureHashesUsecase struct {
	PictureHashesRepository domain.PictureHashesRepository
	contextTimeout          time.Duration
}

func NewPictureHashesUsecase(PictureHashesRepository domain.PictureHashesRepository, timeout time.Duration) domain.PictureHashesUsecase {
	return &PictureHashesUsecase{
		PictureHashesRepository: PictureHashesRepository,
		contextTimeout:          timeout,
	}
}

func (tu *PictureHashesUsecase) Create(c context.Context, PictureHashes *domain.PictureHashes) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.PictureHashesRepository.Create(ctx, PictureHashes)
}

func (tu *PictureHashesUsecase) Update(c context.Context, PictureHashes *domain.PictureHashes) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.PictureHashesRepository.Update(ctx, PictureHashes)
}

func (tu *PictureHashesUsecase) Delete(c context.Context, PictureHashes string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.PictureHashesRepository.Delete(ctx, PictureHashes)
}

func (lu *PictureHashesUsecase) FetchByID(c context.Context, PictureHashesID string) (domain.PictureHashes, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.PictureHashesRepository.FetchByID(ctx, PictureHashesID)
}

func (lu *PictureHashesUsecase) Fetch(c context.Context) ([]domain.PictureHashes, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.PictureHashesRepository.Fetch(ctx)
}
