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
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestVendorAttributeUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.VendorAttributeRepository)
	timeout := time.Duration(10)
	usecase := test.NewVendorAttributeUsecase(mockRepo, timeout)

	vendorAttributeID := bson.NewObjectID().Hex()

	updatedVendorAttribute := domain.VendorAttribute{
		ID: bson.NewObjectID(), // Existing ID of the record to update
	}

	mockRepo.On("FetchByID", mock.Anything, vendorAttributeID).Return(updatedVendorAttribute, nil)

	result, err := usecase.FetchByID(context.Background(), vendorAttributeID)

	assert.NoError(t, err)
	assert.Equal(t, updatedVendorAttribute, result)
	mockRepo.AssertExpectations(t)
}

func TestVendorAttributeUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.VendorAttributeRepository)
	timeout := time.Duration(10)
	usecase := test.NewVendorAttributeUsecase(mockRepo, timeout)

	newVendorAttribute := &domain.VendorAttribute{
		ID: bson.NewObjectID(),
	}

	mockRepo.On("Create", mock.Anything, newVendorAttribute).Return(nil)

	err := usecase.Create(context.Background(), newVendorAttribute)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestVendorAttributeUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.VendorAttributeRepository)
	timeout := time.Duration(10)
	usecase := test.NewVendorAttributeUsecase(mockRepo, timeout)

	updatedVendorAttribute := &domain.VendorAttribute{
		ID: bson.NewObjectID(), // Existing ID of the record to update
	}

	mockRepo.On("Update", mock.Anything, updatedVendorAttribute).Return(nil)

	err := usecase.Update(context.Background(), updatedVendorAttribute)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestVendorAttributeUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.VendorAttributeRepository)
	timeout := time.Duration(10)
	usecase := test.NewVendorAttributeUsecase(mockRepo, timeout)

	vendorAttributeID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, vendorAttributeID).Return(nil)

	err := usecase.Delete(context.Background(), vendorAttributeID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestVendorAttributeUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.VendorAttributeRepository)
	timeout := time.Duration(10)
	usecase := test.NewVendorAttributeUsecase(mockRepo, timeout)

	fetchedVendorAttributes := []domain.VendorAttribute{
		{
			ID: bson.NewObjectID(),
		},
		{
			ID: bson.NewObjectID(),
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedVendorAttributes, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedVendorAttributes, result)
	mockRepo.AssertExpectations(t)
}
