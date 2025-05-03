package usecase_test

import (
	"context"
	domain "earnforglance/server/domain/attributes"
	mocks "earnforglance/server/domain/mocks"
	test "earnforglance/server/usecase/attributes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestCustomerAttributeValueUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.CustomerAttributeValueRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewCustomerAttributeValueUsecase(mockRepo, timeout)

	customerID := bson.NewObjectID().Hex()

	updatedCustomerAttributeValue := domain.CustomerAttributeValue{
		ID:                  bson.NewObjectID(), // Existing ID of the record to update
		CustomerAttributeID: bson.NewObjectID(),
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
	usecase := test.NewCustomerAttributeValueUsecase(mockRepo, timeout)

	newCustomerAttributeValue := &domain.CustomerAttributeValue{
		CustomerAttributeID: bson.NewObjectID(),
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
	usecase := test.NewCustomerAttributeValueUsecase(mockRepo, timeout)

	updatedCustomerAttributeValue := &domain.CustomerAttributeValue{
		ID:                  bson.NewObjectID(), // Existing ID of the record to update
		CustomerAttributeID: bson.NewObjectID(),
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
	usecase := test.NewCustomerAttributeValueUsecase(mockRepo, timeout)

	customerID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, customerID).Return(nil)

	err := usecase.Delete(context.Background(), customerID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCustomerAttributeValueUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.CustomerAttributeValueRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewCustomerAttributeValueUsecase(mockRepo, timeout)

	fetchedCustomerAttributeValues := []domain.CustomerAttributeValue{
		{
			ID:                  bson.NewObjectID(),
			CustomerAttributeID: bson.NewObjectID(),
			Name:                "Preferred Language",
			IsPreSelected:       true,
			DisplayOrder:        1,
		},
		{
			ID:                  bson.NewObjectID(),
			CustomerAttributeID: bson.NewObjectID(),
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
