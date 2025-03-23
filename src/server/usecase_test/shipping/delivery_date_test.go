package usecase_test

import (
	"context"
	mocks "earnforglance/server/domain/mocks"
	domain "earnforglance/server/domain/shipping"
	test "earnforglance/server/usecase/shipping"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestDeliveryDateUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.DeliveryDateRepository)
	timeout := time.Duration(10)
	usecase := test.NewDeliveryDateUsecase(mockRepo, timeout)

	deliveryDateID := primitive.NewObjectID().Hex()

	updatedDeliveryDate := domain.DeliveryDate{
		ID:           primitive.NewObjectID(), // Existing ID of the record to update
		Name:         "Standard Delivery",
		DisplayOrder: 2,
	}

	mockRepo.On("FetchByID", mock.Anything, deliveryDateID).Return(updatedDeliveryDate, nil)

	result, err := usecase.FetchByID(context.Background(), deliveryDateID)

	assert.NoError(t, err)
	assert.Equal(t, updatedDeliveryDate, result)
	mockRepo.AssertExpectations(t)
}

func TestDeliveryDateUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.DeliveryDateRepository)
	timeout := time.Duration(10)
	usecase := test.NewDeliveryDateUsecase(mockRepo, timeout)

	newDeliveryDate := &domain.DeliveryDate{
		Name:         "Next Day Delivery",
		DisplayOrder: 1,
	}

	mockRepo.On("Create", mock.Anything, newDeliveryDate).Return(nil)

	err := usecase.Create(context.Background(), newDeliveryDate)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeliveryDateUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.DeliveryDateRepository)
	timeout := time.Duration(10)
	usecase := test.NewDeliveryDateUsecase(mockRepo, timeout)

	updatedDeliveryDate := &domain.DeliveryDate{
		ID:           primitive.NewObjectID(), // Existing ID of the record to update
		Name:         "Standard Delivery",
		DisplayOrder: 2,
	}

	mockRepo.On("Update", mock.Anything, updatedDeliveryDate).Return(nil)

	err := usecase.Update(context.Background(), updatedDeliveryDate)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeliveryDateUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.DeliveryDateRepository)
	timeout := time.Duration(10)
	usecase := test.NewDeliveryDateUsecase(mockRepo, timeout)

	deliveryDateID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, deliveryDateID).Return(nil)

	err := usecase.Delete(context.Background(), deliveryDateID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeliveryDateUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.DeliveryDateRepository)
	timeout := time.Duration(10)
	usecase := test.NewDeliveryDateUsecase(mockRepo, timeout)

	fetchedDeliveryDates := []domain.DeliveryDate{
		{
			ID:           primitive.NewObjectID(),
			Name:         "Next Day Delivery",
			DisplayOrder: 1,
		},
		{
			ID:           primitive.NewObjectID(),
			Name:         "Standard Delivery",
			DisplayOrder: 2,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedDeliveryDates, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedDeliveryDates, result)
	mockRepo.AssertExpectations(t)
}
