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

func TestDiscountUsageHistoryUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.DiscountUsageHistoryRepository)
	timeout := time.Duration(10)
	usecase := test.NewDiscountUsageHistoryUsecase(mockRepo, timeout)

	discountUsageHistoryID := bson.NewObjectID().Hex()

	updatedDiscountUsageHistory := domain.DiscountUsageHistory{
		ID:           bson.NewObjectID(), // Existing ID of the record to update
		DiscountID:   bson.NewObjectID(),
		OrderID:      bson.NewObjectID(),
		CreatedOnUtc: time.Now().AddDate(0, 0, -7), // Created 7 days ago
	}

	mockRepo.On("FetchByID", mock.Anything, discountUsageHistoryID).Return(updatedDiscountUsageHistory, nil)

	result, err := usecase.FetchByID(context.Background(), discountUsageHistoryID)

	assert.NoError(t, err)
	assert.Equal(t, updatedDiscountUsageHistory, result)
	mockRepo.AssertExpectations(t)
}

func TestDiscountUsageHistoryUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.DiscountUsageHistoryRepository)
	timeout := time.Duration(10)
	usecase := test.NewDiscountUsageHistoryUsecase(mockRepo, timeout)

	newDiscountUsageHistory := &domain.DiscountUsageHistory{
		DiscountID:   bson.NewObjectID(),
		OrderID:      bson.NewObjectID(),
		CreatedOnUtc: time.Now(),
	}

	mockRepo.On("Create", mock.Anything, newDiscountUsageHistory).Return(nil)

	err := usecase.Create(context.Background(), newDiscountUsageHistory)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDiscountUsageHistoryUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.DiscountUsageHistoryRepository)
	timeout := time.Duration(10)
	usecase := test.NewDiscountUsageHistoryUsecase(mockRepo, timeout)

	updatedDiscountUsageHistory := &domain.DiscountUsageHistory{
		ID:           bson.NewObjectID(), // Existing ID of the record to update
		DiscountID:   bson.NewObjectID(),
		OrderID:      bson.NewObjectID(),
		CreatedOnUtc: time.Now().AddDate(0, 0, -7), // Created 7 days ago
	}

	mockRepo.On("Update", mock.Anything, updatedDiscountUsageHistory).Return(nil)

	err := usecase.Update(context.Background(), updatedDiscountUsageHistory)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDiscountUsageHistoryUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.DiscountUsageHistoryRepository)
	timeout := time.Duration(10)
	usecase := test.NewDiscountUsageHistoryUsecase(mockRepo, timeout)

	discountUsageHistoryID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, discountUsageHistoryID).Return(nil)

	err := usecase.Delete(context.Background(), discountUsageHistoryID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDiscountUsageHistoryUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.DiscountUsageHistoryRepository)
	timeout := time.Duration(10)
	usecase := test.NewDiscountUsageHistoryUsecase(mockRepo, timeout)

	fetchedDiscountUsageHistories := []domain.DiscountUsageHistory{
		{
			ID:           bson.NewObjectID(),
			DiscountID:   bson.NewObjectID(),
			OrderID:      bson.NewObjectID(),
			CreatedOnUtc: time.Now().AddDate(0, 0, -10), // Created 10 days ago
		},
		{
			ID:           bson.NewObjectID(),
			DiscountID:   bson.NewObjectID(),
			OrderID:      bson.NewObjectID(),
			CreatedOnUtc: time.Now().AddDate(0, 0, -5), // Created 5 days ago
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedDiscountUsageHistories, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedDiscountUsageHistories, result)
	mockRepo.AssertExpectations(t)
}
