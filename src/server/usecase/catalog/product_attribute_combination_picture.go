package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/catalog"
)

type productattributecombinationpictureUsecase struct {
	productattributecombinationpictureRepository domain.ProductAttributeCombinationPictureRepository
	contextTimeout                               time.Duration
}

func NewProductAttributeCombinationPictureUsecase(productattributecombinationpictureRepository domain.ProductAttributeCombinationPictureRepository, timeout time.Duration) domain.ProductAttributeCombinationPictureUsecase {
	return &productattributecombinationpictureUsecase{
		productattributecombinationpictureRepository: productattributecombinationpictureRepository,
		contextTimeout: timeout,
	}
}

func (tu *productattributecombinationpictureUsecase) Create(c context.Context, productattributecombinationpicture *domain.ProductAttributeCombinationPicture) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.productattributecombinationpictureRepository.Create(ctx, productattributecombinationpicture)
}

func (tu *productattributecombinationpictureUsecase) Update(c context.Context, productattributecombinationpicture *domain.ProductAttributeCombinationPicture) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.productattributecombinationpictureRepository.Update(ctx, productattributecombinationpicture)
}

func (tu *productattributecombinationpictureUsecase) Delete(c context.Context, productattributecombinationpicture string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.productattributecombinationpictureRepository.Delete(ctx, productattributecombinationpicture)
}

func (lu *productattributecombinationpictureUsecase) FetchByID(c context.Context, productattributecombinationpictureID string) (domain.ProductAttributeCombinationPicture, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.productattributecombinationpictureRepository.FetchByID(ctx, productattributecombinationpictureID)
}

func (lu *productattributecombinationpictureUsecase) Fetch(c context.Context) ([]domain.ProductAttributeCombinationPicture, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.productattributecombinationpictureRepository.Fetch(ctx)
}
