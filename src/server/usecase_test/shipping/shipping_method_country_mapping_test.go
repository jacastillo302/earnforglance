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
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestShippingMethodCountryMappingUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.ShippingMethodCountryMappingRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewShippingMethodCountryMappingUsecase(mockRepo, timeout)

	shippingMethodCountryMappingID := bson.NewObjectID().Hex()

	updatedShippingMethodCountryMapping := domain.ShippingMethodCountryMapping{
		ID:               bson.NewObjectID(), // Existing ID of the record to update
		ShippingMethodID: bson.NewObjectID(),
		CountryID:        bson.NewObjectID(),
	}

	mockRepo.On("FetchByID", mock.Anything, shippingMethodCountryMappingID).Return(updatedShippingMethodCountryMapping, nil)

	result, err := usecase.FetchByID(context.Background(), shippingMethodCountryMappingID)

	assert.NoError(t, err)
	assert.Equal(t, updatedShippingMethodCountryMapping, result)
	mockRepo.AssertExpectations(t)
}

func TestShippingMethodCountryMappingUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.ShippingMethodCountryMappingRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewShippingMethodCountryMappingUsecase(mockRepo, timeout)

	newShippingMethodCountryMapping := &domain.ShippingMethodCountryMapping{
		ShippingMethodID: bson.NewObjectID(),
		CountryID:        bson.NewObjectID(),
	}

	mockRepo.On("Create", mock.Anything, newShippingMethodCountryMapping).Return(nil)

	err := usecase.Create(context.Background(), newShippingMethodCountryMapping)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestShippingMethodCountryMappingUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.ShippingMethodCountryMappingRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewShippingMethodCountryMappingUsecase(mockRepo, timeout)

	updatedShippingMethodCountryMapping := &domain.ShippingMethodCountryMapping{
		ID:               bson.NewObjectID(), // Existing ID of the record to update
		ShippingMethodID: bson.NewObjectID(),
		CountryID:        bson.NewObjectID(),
	}

	mockRepo.On("Update", mock.Anything, updatedShippingMethodCountryMapping).Return(nil)

	err := usecase.Update(context.Background(), updatedShippingMethodCountryMapping)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestShippingMethodCountryMappingUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.ShippingMethodCountryMappingRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewShippingMethodCountryMappingUsecase(mockRepo, timeout)

	shippingMethodCountryMappingID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, shippingMethodCountryMappingID).Return(nil)

	err := usecase.Delete(context.Background(), shippingMethodCountryMappingID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestShippingMethodCountryMappingUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.ShippingMethodCountryMappingRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewShippingMethodCountryMappingUsecase(mockRepo, timeout)

	fetchedShippingMethodCountryMappings := []domain.ShippingMethodCountryMapping{
		{
			ID:               bson.NewObjectID(),
			ShippingMethodID: bson.NewObjectID(),
			CountryID:        bson.NewObjectID(),
		},
		{
			ID:               bson.NewObjectID(),
			ShippingMethodID: bson.NewObjectID(),
			CountryID:        bson.NewObjectID(),
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedShippingMethodCountryMappings, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedShippingMethodCountryMappings, result)
	mockRepo.AssertExpectations(t)
}
