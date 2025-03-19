package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/media"
)

type PictureHashItemUsecase struct {
	PictureHashItemRepository domain.PictureHashItemRepository
	contextTimeout            time.Duration
}

func NewPictureHashItemUsecase(PictureHashItemRepository domain.PictureHashItemRepository, timeout time.Duration) domain.PictureHashItemUsecase {
	return &PictureHashItemUsecase{
		PictureHashItemRepository: PictureHashItemRepository,
		contextTimeout:            timeout,
	}
}

func (tu *PictureHashItemUsecase) Create(c context.Context, PictureHashItem *domain.PictureHashItem) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.PictureHashItemRepository.Create(ctx, PictureHashItem)
}

func (tu *PictureHashItemUsecase) Update(c context.Context, PictureHashItem *domain.PictureHashItem) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.PictureHashItemRepository.Update(ctx, PictureHashItem)
}

func (tu *PictureHashItemUsecase) Delete(c context.Context, PictureHashItem *domain.PictureHashItem) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.PictureHashItemRepository.Delete(ctx, PictureHashItem)
}

func (lu *PictureHashItemUsecase) FetchByID(c context.Context, PictureHashItemID string) (domain.PictureHashItem, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.PictureHashItemRepository.FetchByID(ctx, PictureHashItemID)
}

func (lu *PictureHashItemUsecase) Fetch(c context.Context) ([]domain.PictureHashItem, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.PictureHashItemRepository.Fetch(ctx)
}
