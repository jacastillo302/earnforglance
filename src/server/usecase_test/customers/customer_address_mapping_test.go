package usecase_test

import (
	"context"
	domain "earnforglance/server/domain/customers"
	mocks "earnforglance/server/domain/mocks"
	test "earnforglance/server/usecase/customers"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestCustomerAddressMappingUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.CustomerAddressMappingRepository)
	timeout := time.Duration(10)
	usecase := test.NewCustomerAddressMappingUsecase(mockRepo, timeout)

	customerAddressMappingID := bson.NewObjectID().Hex()

	updatedCustomerAddressMapping := domain.CustomerAddressMapping{
		ID:         bson.NewObjectID(), // Existing ID of the record to update
		CustomerID: bson.NewObjectID(),
		AddressID:  bson.NewObjectID(),
	}
	mockRepo.On("FetchByID", mock.Anything, customerAddressMappingID).Return(updatedCustomerAddressMapping, nil)

	result, err := usecase.FetchByID(context.Background(), customerAddressMappingID)

	assert.NoError(t, err)
	assert.Equal(t, updatedCustomerAddressMapping, result)
	mockRepo.AssertExpectations(t)
}

func TestCustomerAddressMappingUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.CustomerAddressMappingRepository)
	timeout := time.Duration(10)
	usecase := test.NewCustomerAddressMappingUsecase(mockRepo, timeout)

	newCustomerAddressMapping := &domain.CustomerAddressMapping{
		CustomerID: bson.NewObjectID(),
		AddressID:  bson.NewObjectID(),
	}

	mockRepo.On("Create", mock.Anything, newCustomerAddressMapping).Return(nil)

	err := usecase.Create(context.Background(), newCustomerAddressMapping)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCustomerAddressMappingUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.CustomerAddressMappingRepository)
	timeout := time.Duration(10)
	usecase := test.NewCustomerAddressMappingUsecase(mockRepo, timeout)

	updatedCustomerAddressMapping := &domain.CustomerAddressMapping{
		ID:         bson.NewObjectID(), // Existing ID of the record to update
		CustomerID: bson.NewObjectID(),
		AddressID:  bson.NewObjectID(),
	}

	mockRepo.On("Update", mock.Anything, updatedCustomerAddressMapping).Return(nil)

	err := usecase.Update(context.Background(), updatedCustomerAddressMapping)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCustomerAddressMappingUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.CustomerAddressMappingRepository)
	timeout := time.Duration(10)
	usecase := test.NewCustomerAddressMappingUsecase(mockRepo, timeout)

	customerAddressMappingID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, customerAddressMappingID).Return(nil)

	err := usecase.Delete(context.Background(), customerAddressMappingID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCustomerAddressMappingUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.CustomerAddressMappingRepository)
	timeout := time.Duration(10)
	usecase := test.NewCustomerAddressMappingUsecase(mockRepo, timeout)

	fetchedCustomerAddressMappings := []domain.CustomerAddressMapping{
		{
			ID:         bson.NewObjectID(),
			CustomerID: bson.NewObjectID(),
			AddressID:  bson.NewObjectID(),
		},
		{
			ID:         bson.NewObjectID(),
			CustomerID: bson.NewObjectID(),
			AddressID:  bson.NewObjectID(),
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedCustomerAddressMappings, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedCustomerAddressMappings, result)
	mockRepo.AssertExpectations(t)
}
