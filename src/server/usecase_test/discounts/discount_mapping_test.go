package usecase_test

import (
	"context"
	domain "earnforglance/server/domain/discounts"
	mocks "earnforglance/server/domain/mocks"
	test "earnforglance/server/usecase/discounts"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestDiscountMappingUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.DiscountMappingRepository)
	timeout := time.Duration(10)
	usecase := test.NewDiscountMappingUsecase(mockRepo, timeout)

	discountID := bson.NewObjectID().Hex()

	updatedDiscountMapping := domain.DiscountMapping{
		ID:         bson.NewObjectID(), // Existing ID of the record to update
		DiscountID: bson.NewObjectID(),
		EntityID:   bson.NewObjectID(),
	}

	mockRepo.On("FetchByID", mock.Anything, discountID).Return(updatedDiscountMapping, nil)

	result, err := usecase.FetchByID(context.Background(), discountID)

	assert.NoError(t, err)
	assert.Equal(t, updatedDiscountMapping, result)
	mockRepo.AssertExpectations(t)
}

func TestDiscountMappingUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.DiscountMappingRepository)
	timeout := time.Duration(10)
	usecase := test.NewDiscountMappingUsecase(mockRepo, timeout)

	newDiscountMapping := &domain.DiscountMapping{
		DiscountID: bson.NewObjectID(),
		EntityID:   bson.NewObjectID(),
	}

	mockRepo.On("Create", mock.Anything, newDiscountMapping).Return(nil)

	err := usecase.Create(context.Background(), newDiscountMapping)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDiscountMappingUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.DiscountMappingRepository)
	timeout := time.Duration(10)
	usecase := test.NewDiscountMappingUsecase(mockRepo, timeout)

	updatedDiscountMapping := &domain.DiscountMapping{
		ID:         bson.NewObjectID(), // Existing ID of the record to update
		DiscountID: bson.NewObjectID(),
		EntityID:   bson.NewObjectID(),
	}

	mockRepo.On("Update", mock.Anything, updatedDiscountMapping).Return(nil)

	err := usecase.Update(context.Background(), updatedDiscountMapping)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDiscountMappingUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.DiscountMappingRepository)
	timeout := time.Duration(10)
	usecase := test.NewDiscountMappingUsecase(mockRepo, timeout)

	discountID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, discountID).Return(nil)

	err := usecase.Delete(context.Background(), discountID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDiscountMappingUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.DiscountMappingRepository)
	timeout := time.Duration(10)
	usecase := test.NewDiscountMappingUsecase(mockRepo, timeout)

	fetchedDiscountMappings := []domain.DiscountMapping{
		{
			ID:         bson.NewObjectID(),
			DiscountID: bson.NewObjectID(),
			EntityID:   bson.NewObjectID(),
		},
		{
			ID:         bson.NewObjectID(),
			DiscountID: bson.NewObjectID(),
			EntityID:   bson.NewObjectID(),
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedDiscountMappings, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedDiscountMappings, result)
	mockRepo.AssertExpectations(t)
}
