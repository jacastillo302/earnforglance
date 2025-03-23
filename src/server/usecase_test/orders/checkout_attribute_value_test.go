package usecase_test

import (
	"context"
	mocks "earnforglance/server/domain/mocks"
	domain "earnforglance/server/domain/orders"
	test "earnforglance/server/usecase/orders"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCheckoutAttributeValueUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.CheckoutAttributeValueRepository)
	timeout := time.Duration(10)
	usecase := test.NewCheckoutAttributeValueUsecase(mockRepo, timeout)

	checkoutAttributeValueID := primitive.NewObjectID().Hex()
	updatedCheckoutAttributeValue := domain.CheckoutAttributeValue{
		ID:               primitive.NewObjectID(), // Existing ID of the record to update
		ColorSquaresRgb:  "#33FF57",
		PriceAdjustment:  15.75,
		WeightAdjustment: 0.50,
	}

	mockRepo.On("FetchByID", mock.Anything, checkoutAttributeValueID).Return(updatedCheckoutAttributeValue, nil)

	result, err := usecase.FetchByID(context.Background(), checkoutAttributeValueID)

	assert.NoError(t, err)
	assert.Equal(t, updatedCheckoutAttributeValue, result)
	mockRepo.AssertExpectations(t)
}

func TestCheckoutAttributeValueUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.CheckoutAttributeValueRepository)
	timeout := time.Duration(10)
	usecase := test.NewCheckoutAttributeValueUsecase(mockRepo, timeout)

	newCheckoutAttributeValue := &domain.CheckoutAttributeValue{
		ColorSquaresRgb:  "#FF5733",
		PriceAdjustment:  10.50,
		WeightAdjustment: 0.25,
	}

	mockRepo.On("Create", mock.Anything, newCheckoutAttributeValue).Return(nil)

	err := usecase.Create(context.Background(), newCheckoutAttributeValue)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCheckoutAttributeValueUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.CheckoutAttributeValueRepository)
	timeout := time.Duration(10)
	usecase := test.NewCheckoutAttributeValueUsecase(mockRepo, timeout)

	updatedCheckoutAttributeValue := &domain.CheckoutAttributeValue{
		ID:               primitive.NewObjectID(), // Existing ID of the record to update
		ColorSquaresRgb:  "#33FF57",
		PriceAdjustment:  15.75,
		WeightAdjustment: 0.50,
	}

	mockRepo.On("Update", mock.Anything, updatedCheckoutAttributeValue).Return(nil)

	err := usecase.Update(context.Background(), updatedCheckoutAttributeValue)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCheckoutAttributeValueUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.CheckoutAttributeValueRepository)
	timeout := time.Duration(10)
	usecase := test.NewCheckoutAttributeValueUsecase(mockRepo, timeout)

	checkoutAttributeValueID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, checkoutAttributeValueID).Return(nil)

	err := usecase.Delete(context.Background(), checkoutAttributeValueID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCheckoutAttributeValueUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.CheckoutAttributeValueRepository)
	timeout := time.Duration(10)
	usecase := test.NewCheckoutAttributeValueUsecase(mockRepo, timeout)

	fetchedCheckoutAttributeValues := []domain.CheckoutAttributeValue{
		{
			ID:               primitive.NewObjectID(),
			ColorSquaresRgb:  "#FF5733",
			PriceAdjustment:  10.50,
			WeightAdjustment: 0.25,
		},
		{
			ID:               primitive.NewObjectID(),
			ColorSquaresRgb:  "#3357FF",
			PriceAdjustment:  20.00,
			WeightAdjustment: 0.75,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedCheckoutAttributeValues, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedCheckoutAttributeValues, result)
	mockRepo.AssertExpectations(t)
}
