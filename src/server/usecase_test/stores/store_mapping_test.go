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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestStoreMappingUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.StoreMappingRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewStoreMappingUsecase(mockRepo, timeout)

	storeMappingID := primitive.NewObjectID().Hex()

	updatedStoreMapping := domain.StoreMapping{
		ID:         primitive.NewObjectID(), // Existing ID of the record to update
		EntityID:   primitive.NewObjectID(),
		EntityName: "Category",
		StoreID:    primitive.NewObjectID(),
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
		EntityID:   primitive.NewObjectID(),
		EntityName: "Product",
		StoreID:    primitive.NewObjectID(),
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
		ID:         primitive.NewObjectID(), // Existing ID of the record to update
		EntityID:   primitive.NewObjectID(),
		EntityName: "Category",
		StoreID:    primitive.NewObjectID(),
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

	storeMappingID := primitive.NewObjectID().Hex()

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
			ID:         primitive.NewObjectID(),
			EntityID:   primitive.NewObjectID(),
			EntityName: "Product",
			StoreID:    primitive.NewObjectID(),
		},
		{
			ID:         primitive.NewObjectID(),
			EntityID:   primitive.NewObjectID(),
			EntityName: "Category",
			StoreID:    primitive.NewObjectID(),
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedStoreMappings, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedStoreMappings, result)
	mockRepo.AssertExpectations(t)
}
