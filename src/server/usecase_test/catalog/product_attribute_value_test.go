package usecase_test

import (
	"context"
	domain "earnforglance/server/domain/catalog"
	mocks "earnforglance/server/domain/mocks"
	test "earnforglance/server/usecase/catalog"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestProductAttributeValueUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.ProductAttributeValueRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewProductAttributeValueUsecase(mockRepo, timeout)

	productAttributeValueID := bson.NewObjectID().Hex()

	updatedProductAttributeValue := domain.ProductAttributeValue{
		ID:                           bson.NewObjectID(), // Existing ID of the record to update
		ProductAttributeMappingID:    bson.NewObjectID(),
		AttributeValueTypeID:         2,
		AssociatedProductID:          bson.NewObjectID(),
		Name:                         "Size - Large",
		ColorSquaresRgb:              "",
		ImageSquaresPictureID:        bson.NewObjectID(),
		PriceAdjustment:              15.0,
		PriceAdjustmentUsePercentage: true,
		WeightAdjustment:             1.0,
		Cost:                         7.0,
		CustomerEntersQty:            true,
		Quantity:                     50,
		IsPreSelected:                false,
		DisplayOrder:                 2,
		PictureID:                    nil, // Deprecated field
	}

	mockRepo.On("FetchByID", mock.Anything, productAttributeValueID).Return(updatedProductAttributeValue, nil)

	result, err := usecase.FetchByID(context.Background(), productAttributeValueID)

	assert.NoError(t, err)
	assert.Equal(t, updatedProductAttributeValue, result)
	mockRepo.AssertExpectations(t)
}

func TestProductAttributeValueUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.ProductAttributeValueRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewProductAttributeValueUsecase(mockRepo, timeout)

	newProductAttributeValue := &domain.ProductAttributeValue{
		ProductAttributeMappingID:    bson.NewObjectID(),
		AttributeValueTypeID:         1,
		AssociatedProductID:          bson.NewObjectID(),
		Name:                         "Color - Red",
		ColorSquaresRgb:              "#FF0000",
		ImageSquaresPictureID:        bson.NewObjectID(),
		PriceAdjustment:              10.0,
		PriceAdjustmentUsePercentage: false,
		WeightAdjustment:             0.5,
		Cost:                         5.0,
		CustomerEntersQty:            false,
		Quantity:                     100,
		IsPreSelected:                true,
		DisplayOrder:                 1,
		PictureID:                    nil, // Deprecated field
	}

	mockRepo.On("Create", mock.Anything, newProductAttributeValue).Return(nil)

	err := usecase.Create(context.Background(), newProductAttributeValue)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductAttributeValueUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.ProductAttributeValueRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewProductAttributeValueUsecase(mockRepo, timeout)

	updatedProductAttributeValue := &domain.ProductAttributeValue{
		ID:                           bson.NewObjectID(), // Existing ID of the record to update
		ProductAttributeMappingID:    bson.NewObjectID(),
		AttributeValueTypeID:         2,
		AssociatedProductID:          bson.NewObjectID(),
		Name:                         "Size - Large",
		ColorSquaresRgb:              "",
		ImageSquaresPictureID:        bson.NewObjectID(),
		PriceAdjustment:              15.0,
		PriceAdjustmentUsePercentage: true,
		WeightAdjustment:             1.0,
		Cost:                         7.0,
		CustomerEntersQty:            true,
		Quantity:                     50,
		IsPreSelected:                false,
		DisplayOrder:                 2,
		PictureID:                    nil, // Deprecated field
	}

	mockRepo.On("Update", mock.Anything, updatedProductAttributeValue).Return(nil)

	err := usecase.Update(context.Background(), updatedProductAttributeValue)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductAttributeValueUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.ProductAttributeValueRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewProductAttributeValueUsecase(mockRepo, timeout)

	productAttributeValueID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, productAttributeValueID).Return(nil)

	err := usecase.Delete(context.Background(), productAttributeValueID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductAttributeValueUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.ProductAttributeValueRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewProductAttributeValueUsecase(mockRepo, timeout)

	fetchedProductAttributeValues := []domain.ProductAttributeValue{
		{
			ID:                           bson.NewObjectID(),
			ProductAttributeMappingID:    bson.NewObjectID(),
			AttributeValueTypeID:         1,
			AssociatedProductID:          bson.NewObjectID(),
			Name:                         "Color - Blue",
			ColorSquaresRgb:              "#0000FF",
			ImageSquaresPictureID:        bson.NewObjectID(),
			PriceAdjustment:              5.0,
			PriceAdjustmentUsePercentage: false,
			WeightAdjustment:             0.2,
			Cost:                         3.0,
			CustomerEntersQty:            false,
			Quantity:                     200,
			IsPreSelected:                true,
			DisplayOrder:                 1,
			PictureID:                    nil, // Deprecated field
		},
		{
			ID:                           bson.NewObjectID(),
			ProductAttributeMappingID:    bson.NewObjectID(),
			AttributeValueTypeID:         2,
			AssociatedProductID:          bson.NewObjectID(),
			Name:                         "Size - Medium",
			ColorSquaresRgb:              "",
			ImageSquaresPictureID:        bson.NewObjectID(),
			PriceAdjustment:              20.0,
			PriceAdjustmentUsePercentage: true,
			WeightAdjustment:             0.8,
			Cost:                         10.0,
			CustomerEntersQty:            true,
			Quantity:                     30,
			IsPreSelected:                false,
			DisplayOrder:                 2,
			PictureID:                    nil, // Deprecated field
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedProductAttributeValues, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedProductAttributeValues, result)
	mockRepo.AssertExpectations(t)
}
