package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/catalog"
)

type productpictureUsecase struct {
	productpictureRepository domain.ProductPictureRepository
	contextTimeout           time.Duration
}

func NewProductPictureUsecase(productpictureRepository domain.ProductPictureRepository, timeout time.Duration) domain.ProductPictureUsecase {
	return &productpictureUsecase{
		productpictureRepository: productpictureRepository,
		contextTimeout:           timeout,
	}
}

func (tu *productpictureUsecase) Create(c context.Context, productpicture *domain.ProductPicture) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.productpictureRepository.Create(ctx, productpicture)
}

func (tu *productpictureUsecase) Update(c context.Context, productpicture *domain.ProductPicture) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.productpictureRepository.Update(ctx, productpicture)
}

func (tu *productpictureUsecase) Delete(c context.Context, productpicture *domain.ProductPicture) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.productpictureRepository.Delete(ctx, productpicture)
}

func (lu *productpictureUsecase) FetchByID(c context.Context, productpictureID string) (domain.ProductPicture, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.productpictureRepository.FetchByID(ctx, productpictureID)
}

func (lu *productpictureUsecase) Fetch(c context.Context) ([]domain.ProductPicture, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.productpictureRepository.Fetch(ctx)
}
