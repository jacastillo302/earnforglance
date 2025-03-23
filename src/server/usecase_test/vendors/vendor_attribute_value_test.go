package usecase_test

import (
	"context"
	mocks "earnforglance/server/domain/mocks"
	domain "earnforglance/server/domain/vendors"
	test "earnforglance/server/usecase/vendors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestVendorAttributeValueUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.VendorAttributeValueRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewVendorAttributeValueUsecase(mockRepo, timeout)

	vendorAttributeValueID := primitive.NewObjectID().Hex()

	updatedVendorAttributeValue := domain.VendorAttributeValue{
		ID: primitive.NewObjectID(), // Existing ID of the record to update
	}

	mockRepo.On("FetchByID", mock.Anything, vendorAttributeValueID).Return(updatedVendorAttributeValue, nil)

	result, err := usecase.FetchByID(context.Background(), vendorAttributeValueID)

	assert.NoError(t, err)
	assert.Equal(t, updatedVendorAttributeValue, result)
	mockRepo.AssertExpectations(t)
}

func TestVendorAttributeValueUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.VendorAttributeValueRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewVendorAttributeValueUsecase(mockRepo, timeout)

	newVendorAttributeValue := &domain.VendorAttributeValue{
		ID: primitive.NewObjectID(),
	}

	mockRepo.On("Create", mock.Anything, newVendorAttributeValue).Return(nil)

	err := usecase.Create(context.Background(), newVendorAttributeValue)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestVendorAttributeValueUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.VendorAttributeValueRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewVendorAttributeValueUsecase(mockRepo, timeout)

	updatedVendorAttributeValue := &domain.VendorAttributeValue{
		ID: primitive.NewObjectID(), // Existing ID of the record to update
	}

	mockRepo.On("Update", mock.Anything, updatedVendorAttributeValue).Return(nil)

	err := usecase.Update(context.Background(), updatedVendorAttributeValue)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestVendorAttributeValueUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.VendorAttributeValueRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewVendorAttributeValueUsecase(mockRepo, timeout)

	vendorAttributeValueID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, vendorAttributeValueID).Return(nil)

	err := usecase.Delete(context.Background(), vendorAttributeValueID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestVendorAttributeValueUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.VendorAttributeValueRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewVendorAttributeValueUsecase(mockRepo, timeout)

	fetchedVendorAttributeValues := []domain.VendorAttributeValue{
		{
			ID: primitive.NewObjectID(),
		},
		{
			ID: primitive.NewObjectID(),
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedVendorAttributeValues, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedVendorAttributeValues, result)
	mockRepo.AssertExpectations(t)
}
