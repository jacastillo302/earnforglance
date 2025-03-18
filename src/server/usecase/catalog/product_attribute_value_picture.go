package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/catalog"
)

type productattributevaluepictureUsecase struct {
	productattributevaluepictureRepository domain.ProductAttributeValuePictureRepository
	contextTimeout                         time.Duration
}

func NewProductAttributeValuePictureUsecase(productattributevaluepictureRepository domain.ProductAttributeValuePictureRepository, timeout time.Duration) domain.ProductAttributeValuePictureUsecase {
	return &productattributevaluepictureUsecase{
		productattributevaluepictureRepository: productattributevaluepictureRepository,
		contextTimeout:                         timeout,
	}
}

func (tu *productattributevaluepictureUsecase) Create(c context.Context, productattributevaluepicture *domain.ProductAttributeValuePicture) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.productattributevaluepictureRepository.Create(ctx, productattributevaluepicture)
}

func (tu *productattributevaluepictureUsecase) Update(c context.Context, productattributevaluepicture *domain.ProductAttributeValuePicture) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.productattributevaluepictureRepository.Update(ctx, productattributevaluepicture)
}

func (tu *productattributevaluepictureUsecase) Delete(c context.Context, productattributevaluepicture *domain.ProductAttributeValuePicture) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.productattributevaluepictureRepository.Delete(ctx, productattributevaluepicture)
}

func (lu *productattributevaluepictureUsecase) FetchByID(c context.Context, productattributevaluepictureID string) (domain.ProductAttributeValuePicture, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.productattributevaluepictureRepository.FetchByID(ctx, productattributevaluepictureID)
}

func (lu *productattributevaluepictureUsecase) Fetch(c context.Context) ([]domain.ProductAttributeValuePicture, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.productattributevaluepictureRepository.Fetch(ctx)
}
