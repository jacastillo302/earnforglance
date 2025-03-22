package usecase

import (
	"context"
	mocks "earnforglance/server/domain/mocks"
	domain "earnforglance/server/domain/shipping"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestShippingMethodUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.ShippingMethodRepository)
	timeout := time.Duration(10)
	usecase := NewShippingMethodUsecase(mockRepo, timeout)

	shippingMethodID := primitive.NewObjectID().Hex()

	updatedShippingMethod := domain.ShippingMethod{
		ID:           primitive.NewObjectID(), // Existing ID of the record to update
		Name:         "Express Shipping",
		Description:  "Delivery within 1-2 business days.",
		DisplayOrder: 2,
	}

	mockRepo.On("FetchByID", mock.Anything, shippingMethodID).Return(updatedShippingMethod, nil)

	result, err := usecase.FetchByID(context.Background(), shippingMethodID)

	assert.NoError(t, err)
	assert.Equal(t, updatedShippingMethod, result)
	mockRepo.AssertExpectations(t)
}

func TestShippingMethodUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.ShippingMethodRepository)
	timeout := time.Duration(10)
	usecase := NewShippingMethodUsecase(mockRepo, timeout)

	newShippingMethod := &domain.ShippingMethod{
		Name:         "Standard Shipping",
		Description:  "Delivery within 5-7 business days.",
		DisplayOrder: 1,
	}

	mockRepo.On("Create", mock.Anything, newShippingMethod).Return(nil)

	err := usecase.Create(context.Background(), newShippingMethod)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestShippingMethodUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.ShippingMethodRepository)
	timeout := time.Duration(10)
	usecase := NewShippingMethodUsecase(mockRepo, timeout)

	updatedShippingMethod := &domain.ShippingMethod{
		ID:           primitive.NewObjectID(), // Existing ID of the record to update
		Name:         "Express Shipping",
		Description:  "Delivery within 1-2 business days.",
		DisplayOrder: 2,
	}

	mockRepo.On("Update", mock.Anything, updatedShippingMethod).Return(nil)

	err := usecase.Update(context.Background(), updatedShippingMethod)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestShippingMethodUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.ShippingMethodRepository)
	timeout := time.Duration(10)
	usecase := NewShippingMethodUsecase(mockRepo, timeout)

	shippingMethodID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, shippingMethodID).Return(nil)

	err := usecase.Delete(context.Background(), shippingMethodID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestShippingMethodUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.ShippingMethodRepository)
	timeout := time.Duration(10)
	usecase := NewShippingMethodUsecase(mockRepo, timeout)

	fetchedShippingMethods := []domain.ShippingMethod{
		{
			ID:           primitive.NewObjectID(),
			Name:         "Standard Shipping",
			Description:  "Delivery within 5-7 business days.",
			DisplayOrder: 1,
		},
		{
			ID:           primitive.NewObjectID(),
			Name:         "Express Shipping",
			Description:  "Delivery within 1-2 business days.",
			DisplayOrder: 2,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedShippingMethods, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedShippingMethods, result)
	mockRepo.AssertExpectations(t)
}
