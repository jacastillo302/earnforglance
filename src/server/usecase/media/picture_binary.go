package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/media"
)

type picturebinaryUsecase struct {
	picturebinaryRepository domain.PictureBinaryRepository
	contextTimeout          time.Duration
}

func NewPictureBinaryUsecase(picturebinaryRepository domain.PictureBinaryRepository, timeout time.Duration) domain.PictureBinaryUsecase {
	return &picturebinaryUsecase{
		picturebinaryRepository: picturebinaryRepository,
		contextTimeout:          timeout,
	}
}

func (tu *picturebinaryUsecase) Create(c context.Context, picturebinary *domain.PictureBinary) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.picturebinaryRepository.Create(ctx, picturebinary)
}

func (tu *picturebinaryUsecase) Update(c context.Context, picturebinary *domain.PictureBinary) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.picturebinaryRepository.Update(ctx, picturebinary)
}

func (tu *picturebinaryUsecase) Delete(c context.Context, picturebinary string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.picturebinaryRepository.Delete(ctx, picturebinary)
}

func (lu *picturebinaryUsecase) FetchByID(c context.Context, picturebinaryID string) (domain.PictureBinary, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.picturebinaryRepository.FetchByID(ctx, picturebinaryID)
}

func (lu *picturebinaryUsecase) Fetch(c context.Context) ([]domain.PictureBinary, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.picturebinaryRepository.Fetch(ctx)
}
