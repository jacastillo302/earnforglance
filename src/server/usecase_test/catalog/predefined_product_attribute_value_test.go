package usecase_test

import (
	"context"
	test "earnforglance/server/usecase/catalog"
	"testing"
	"time"

	domain "earnforglance/server/domain/catalog"
	mocks "earnforglance/server/domain/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestPredefinedProductAttributeValueUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.PredefinedProductAttributeValueRepository)
	timeout := time.Duration(10)
	usecase := test.NewPredefinedProductAttributeValueUsecase(mockRepo, timeout)

	predefinedProductAttributeValueID := bson.NewObjectID().Hex()

	expectedPredefinedProductAttributeValue := domain.PredefinedProductAttributeValue{
		ID:                           bson.NewObjectID(),
		ProductAttributeID:           bson.NewObjectID(),
		Name:                         "Color - Blue",
		PriceAdjustment:              15.0,
		PriceAdjustmentUsePercentage: true,
		WeightAdjustment:             0.8,
		Cost:                         7.0,
		IsPreSelected:                false,
		DisplayOrder:                 2,
	}

	mockRepo.On("FetchByID", mock.Anything, predefinedProductAttributeValueID).Return(expectedPredefinedProductAttributeValue, nil)

	result, err := usecase.FetchByID(context.Background(), predefinedProductAttributeValueID)

	assert.NoError(t, err)
	assert.Equal(t, expectedPredefinedProductAttributeValue, result)
	mockRepo.AssertExpectations(t)
}

func TestPredefinedProductAttributeValueUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.PredefinedProductAttributeValueRepository)
	timeout := time.Duration(10)
	usecase := test.NewPredefinedProductAttributeValueUsecase(mockRepo, timeout)

	newPredefinedProductAttributeValue := &domain.PredefinedProductAttributeValue{
		ProductAttributeID:           bson.NewObjectID(),
		Name:                         "Color - Red",
		PriceAdjustment:              10.0,
		PriceAdjustmentUsePercentage: false,
		WeightAdjustment:             0.5,
		Cost:                         5.0,
		IsPreSelected:                true,
		DisplayOrder:                 1,
	}

	mockRepo.On("Create", mock.Anything, newPredefinedProductAttributeValue).Return(nil)

	err := usecase.Create(context.Background(), newPredefinedProductAttributeValue)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestPredefinedProductAttributeValueUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.PredefinedProductAttributeValueRepository)
	timeout := time.Duration(10)
	usecase := test.NewPredefinedProductAttributeValueUsecase(mockRepo, timeout)

	updatedPredefinedProductAttributeValue := &domain.PredefinedProductAttributeValue{
		ID:                           bson.NewObjectID(),
		ProductAttributeID:           bson.NewObjectID(),
		Name:                         "Color - Blue",
		PriceAdjustment:              15.0,
		PriceAdjustmentUsePercentage: true,
		WeightAdjustment:             0.8,
		Cost:                         7.0,
		IsPreSelected:                false,
		DisplayOrder:                 2,
	}

	mockRepo.On("Update", mock.Anything, updatedPredefinedProductAttributeValue).Return(nil)

	err := usecase.Update(context.Background(), updatedPredefinedProductAttributeValue)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestPredefinedProductAttributeValueUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.PredefinedProductAttributeValueRepository)
	timeout := time.Duration(10)
	usecase := test.NewPredefinedProductAttributeValueUsecase(mockRepo, timeout)

	predefinedProductAttributeValueID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, predefinedProductAttributeValueID).Return(nil)

	err := usecase.Delete(context.Background(), predefinedProductAttributeValueID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestPredefinedProductAttributeValueUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.PredefinedProductAttributeValueRepository)
	timeout := time.Duration(10)
	usecase := test.NewPredefinedProductAttributeValueUsecase(mockRepo, timeout)

	expectedPredefinedProductAttributeValues := []domain.PredefinedProductAttributeValue{
		{
			ID:                           bson.NewObjectID(),
			ProductAttributeID:           bson.NewObjectID(),
			Name:                         "Size - Small",
			PriceAdjustment:              5.0,
			PriceAdjustmentUsePercentage: false,
			WeightAdjustment:             0.2,
			Cost:                         2.0,
			IsPreSelected:                true,
			DisplayOrder:                 1,
		},
		{
			ID:                           bson.NewObjectID(),
			ProductAttributeID:           bson.NewObjectID(),
			Name:                         "Size - Large",
			PriceAdjustment:              20.0,
			PriceAdjustmentUsePercentage: true,
			WeightAdjustment:             1.0,
			Cost:                         10.0,
			IsPreSelected:                false,
			DisplayOrder:                 2,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(expectedPredefinedProductAttributeValues, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, expectedPredefinedProductAttributeValues, result)
	mockRepo.AssertExpectations(t)
}
