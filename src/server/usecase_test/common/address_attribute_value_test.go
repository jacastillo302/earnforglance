package usecase_test

import (
	"context"
	domain "earnforglance/server/domain/common"
	mocks "earnforglance/server/domain/mocks"
	test "earnforglance/server/usecase/common"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestAddressAttributeValueUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.AddressAttributeValueRepository)
	timeout := time.Duration(10)
	usecase := test.NewAddressAttributeValueUsecase(mockRepo, timeout)

	addressAttributeValueID := bson.NewObjectID().Hex()

	updatedAddressAttributeValue := domain.AddressAttributeValue{
		ID:                 bson.NewObjectID(), // Existing ID of the record to update
		AddressAttributeID: bson.NewObjectID(),
		Name:               "State",
		IsPreSelected:      false,
		DisplayOrder:       2,
	}

	mockRepo.On("FetchByID", mock.Anything, addressAttributeValueID).Return(updatedAddressAttributeValue, nil)

	result, err := usecase.FetchByID(context.Background(), addressAttributeValueID)

	assert.NoError(t, err)
	assert.Equal(t, updatedAddressAttributeValue, result)
	mockRepo.AssertExpectations(t)
}

func TestAddressAttributeValueUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.AddressAttributeValueRepository)
	timeout := time.Duration(10)
	usecase := test.NewAddressAttributeValueUsecase(mockRepo, timeout)

	newAddressAttributeValue := &domain.AddressAttributeValue{
		AddressAttributeID: bson.NewObjectID(),
		Name:               "Country",
		IsPreSelected:      true,
		DisplayOrder:       1,
	}

	mockRepo.On("Create", mock.Anything, newAddressAttributeValue).Return(nil)

	err := usecase.Create(context.Background(), newAddressAttributeValue)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestAddressAttributeValueUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.AddressAttributeValueRepository)
	timeout := time.Duration(10)
	usecase := test.NewAddressAttributeValueUsecase(mockRepo, timeout)

	updatedAddressAttributeValue := &domain.AddressAttributeValue{
		ID:                 bson.NewObjectID(), // Existing ID of the record to update
		AddressAttributeID: bson.NewObjectID(),
		Name:               "State",
		IsPreSelected:      false,
		DisplayOrder:       2,
	}

	mockRepo.On("Update", mock.Anything, updatedAddressAttributeValue).Return(nil)

	err := usecase.Update(context.Background(), updatedAddressAttributeValue)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestAddressAttributeValueUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.AddressAttributeValueRepository)
	timeout := time.Duration(10)
	usecase := test.NewAddressAttributeValueUsecase(mockRepo, timeout)

	addressAttributeValueID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, addressAttributeValueID).Return(nil)

	err := usecase.Delete(context.Background(), addressAttributeValueID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestAddressAttributeValueUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.AddressAttributeValueRepository)
	timeout := time.Duration(10)
	usecase := test.NewAddressAttributeValueUsecase(mockRepo, timeout)

	fetchedAddressAttributeValues := []domain.AddressAttributeValue{
		{
			ID:                 bson.NewObjectID(),
			AddressAttributeID: bson.NewObjectID(),
			Name:               "Country",
			IsPreSelected:      true,
			DisplayOrder:       1,
		},
		{
			ID:                 bson.NewObjectID(),
			AddressAttributeID: bson.NewObjectID(),
			Name:               "State",
			IsPreSelected:      false,
			DisplayOrder:       2,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedAddressAttributeValues, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedAddressAttributeValues, result)
	mockRepo.AssertExpectations(t)
}
