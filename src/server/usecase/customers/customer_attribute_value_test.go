package usecase

import (
	"context"
	domain "earnforglance/server/domain/customers"
	mocks "earnforglance/server/domain/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCustomerAttributeValueUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.CustomerAttributeValueRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewCustomerAttributeValueUsecase(mockRepo, timeout)

	customerID := primitive.NewObjectID().Hex()

	updatedCustomerAttributeValue := domain.CustomerAttributeValue{
		ID:                  primitive.NewObjectID(), // Existing ID of the record to update
		CustomerAttributeID: primitive.NewObjectID(),
		Name:                "Preferred Currency",
		IsPreSelected:       false,
		DisplayOrder:        2,
	}

	mockRepo.On("FetchByID", mock.Anything, customerID).Return(updatedCustomerAttributeValue, nil)

	result, err := usecase.FetchByID(context.Background(), customerID)

	assert.NoError(t, err)
	assert.Equal(t, updatedCustomerAttributeValue, result)
	mockRepo.AssertExpectations(t)
}

func TestCustomerAttributeValueUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.CustomerAttributeValueRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewCustomerAttributeValueUsecase(mockRepo, timeout)

	newCustomerAttributeValue := &domain.CustomerAttributeValue{
		CustomerAttributeID: primitive.NewObjectID(),
		Name:                "Preferred Language",
		IsPreSelected:       true,
		DisplayOrder:        1,
	}

	mockRepo.On("Create", mock.Anything, newCustomerAttributeValue).Return(nil)

	err := usecase.Create(context.Background(), newCustomerAttributeValue)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCustomerAttributeValueUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.CustomerAttributeValueRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewCustomerAttributeValueUsecase(mockRepo, timeout)

	updatedCustomerAttributeValue := &domain.CustomerAttributeValue{
		ID:                  primitive.NewObjectID(), // Existing ID of the record to update
		CustomerAttributeID: primitive.NewObjectID(),
		Name:                "Preferred Currency",
		IsPreSelected:       false,
		DisplayOrder:        2,
	}

	mockRepo.On("Update", mock.Anything, updatedCustomerAttributeValue).Return(nil)

	err := usecase.Update(context.Background(), updatedCustomerAttributeValue)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCustomerAttributeValueUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.CustomerAttributeValueRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewCustomerAttributeValueUsecase(mockRepo, timeout)

	customerID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, customerID).Return(nil)

	err := usecase.Delete(context.Background(), customerID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCustomerAttributeValueUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.CustomerAttributeValueRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewCustomerAttributeValueUsecase(mockRepo, timeout)

	fetchedCustomerAttributeValues := []domain.CustomerAttributeValue{
		{
			ID:                  primitive.NewObjectID(),
			CustomerAttributeID: primitive.NewObjectID(),
			Name:                "Preferred Language",
			IsPreSelected:       true,
			DisplayOrder:        1,
		},
		{
			ID:                  primitive.NewObjectID(),
			CustomerAttributeID: primitive.NewObjectID(),
			Name:                "Preferred Currency",
			IsPreSelected:       false,
			DisplayOrder:        2,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedCustomerAttributeValues, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedCustomerAttributeValues, result)
	mockRepo.AssertExpectations(t)
}
