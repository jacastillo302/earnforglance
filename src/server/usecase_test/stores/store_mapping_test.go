package usecase_test

import (
	"context"
	mocks "earnforglance/server/domain/mocks"
	domain "earnforglance/server/domain/stores"
	test "earnforglance/server/usecase/stores"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestStoreMappingUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.StoreMappingRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewStoreMappingUsecase(mockRepo, timeout)

	storeMappingID := bson.NewObjectID().Hex()

	updatedStoreMapping := domain.StoreMapping{
		ID:         bson.NewObjectID(), // Existing ID of the record to update
		EntityID:   bson.NewObjectID(),
		EntityName: "Category",
		StoreID:    bson.NewObjectID(),
	}

	mockRepo.On("FetchByID", mock.Anything, storeMappingID).Return(updatedStoreMapping, nil)

	result, err := usecase.FetchByID(context.Background(), storeMappingID)

	assert.NoError(t, err)
	assert.Equal(t, updatedStoreMapping, result)
	mockRepo.AssertExpectations(t)
}

func TestStoreMappingUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.StoreMappingRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewStoreMappingUsecase(mockRepo, timeout)

	newStoreMapping := &domain.StoreMapping{
		EntityID:   bson.NewObjectID(),
		EntityName: "Product",
		StoreID:    bson.NewObjectID(),
	}

	mockRepo.On("Create", mock.Anything, newStoreMapping).Return(nil)

	err := usecase.Create(context.Background(), newStoreMapping)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestStoreMappingUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.StoreMappingRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewStoreMappingUsecase(mockRepo, timeout)

	updatedStoreMapping := &domain.StoreMapping{
		ID:         bson.NewObjectID(), // Existing ID of the record to update
		EntityID:   bson.NewObjectID(),
		EntityName: "Category",
		StoreID:    bson.NewObjectID(),
	}

	mockRepo.On("Update", mock.Anything, updatedStoreMapping).Return(nil)

	err := usecase.Update(context.Background(), updatedStoreMapping)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestStoreMappingUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.StoreMappingRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewStoreMappingUsecase(mockRepo, timeout)

	storeMappingID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, storeMappingID).Return(nil)

	err := usecase.Delete(context.Background(), storeMappingID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestStoreMappingUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.StoreMappingRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewStoreMappingUsecase(mockRepo, timeout)

	fetchedStoreMappings := []domain.StoreMapping{
		{
			ID:         bson.NewObjectID(),
			EntityID:   bson.NewObjectID(),
			EntityName: "Product",
			StoreID:    bson.NewObjectID(),
		},
		{
			ID:         bson.NewObjectID(),
			EntityID:   bson.NewObjectID(),
			EntityName: "Category",
			StoreID:    bson.NewObjectID(),
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedStoreMappings, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedStoreMappings, result)
	mockRepo.AssertExpectations(t)
}
