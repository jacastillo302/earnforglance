package usecase_test

import (
	"context"
	mocks "earnforglance/server/domain/mocks"
	domain "earnforglance/server/domain/orders"
	test "earnforglance/server/usecase/orders"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestGiftCardUsageHistoryUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.GiftCardUsageHistoryRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewGiftCardUsageHistoryUsecase(mockRepo, timeout)

	giftCardUsageHistoryID := bson.NewObjectID().Hex()

	updatedGiftCardUsageHistory := domain.GiftCardUsageHistory{
		ID:           bson.NewObjectID(), // Existing ID of the record to update
		GiftCardID:   bson.NewObjectID(),
		OrderID:      bson.NewObjectID(),
		UsedValue:    75.00,
		CreatedOnUtc: time.Now().AddDate(0, 0, -7), // Created 7 days ago
	}

	mockRepo.On("FetchByID", mock.Anything, giftCardUsageHistoryID).Return(updatedGiftCardUsageHistory, nil)

	result, err := usecase.FetchByID(context.Background(), giftCardUsageHistoryID)

	assert.NoError(t, err)
	assert.Equal(t, updatedGiftCardUsageHistory, result)
	mockRepo.AssertExpectations(t)
}

func TestGiftCardUsageHistoryUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.GiftCardUsageHistoryRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewGiftCardUsageHistoryUsecase(mockRepo, timeout)

	newGiftCardUsageHistory := &domain.GiftCardUsageHistory{
		GiftCardID:   bson.NewObjectID(),
		OrderID:      bson.NewObjectID(),
		UsedValue:    50.00,
		CreatedOnUtc: time.Now(),
	}

	mockRepo.On("Create", mock.Anything, newGiftCardUsageHistory).Return(nil)

	err := usecase.Create(context.Background(), newGiftCardUsageHistory)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGiftCardUsageHistoryUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.GiftCardUsageHistoryRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewGiftCardUsageHistoryUsecase(mockRepo, timeout)

	updatedGiftCardUsageHistory := &domain.GiftCardUsageHistory{
		ID:           bson.NewObjectID(), // Existing ID of the record to update
		GiftCardID:   bson.NewObjectID(),
		OrderID:      bson.NewObjectID(),
		UsedValue:    75.00,
		CreatedOnUtc: time.Now().AddDate(0, 0, -7), // Created 7 days ago
	}

	mockRepo.On("Update", mock.Anything, updatedGiftCardUsageHistory).Return(nil)

	err := usecase.Update(context.Background(), updatedGiftCardUsageHistory)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGiftCardUsageHistoryUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.GiftCardUsageHistoryRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewGiftCardUsageHistoryUsecase(mockRepo, timeout)

	giftCardUsageHistoryID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, giftCardUsageHistoryID).Return(nil)

	err := usecase.Delete(context.Background(), giftCardUsageHistoryID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGiftCardUsageHistoryUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.GiftCardUsageHistoryRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewGiftCardUsageHistoryUsecase(mockRepo, timeout)

	fetchedGiftCardUsageHistories := []domain.GiftCardUsageHistory{
		{
			ID:           bson.NewObjectID(),
			GiftCardID:   bson.NewObjectID(),
			OrderID:      bson.NewObjectID(),
			UsedValue:    50.00,
			CreatedOnUtc: time.Now().AddDate(0, 0, -10), // Created 10 days ago
		},
		{
			ID:           bson.NewObjectID(),
			GiftCardID:   bson.NewObjectID(),
			OrderID:      bson.NewObjectID(),
			UsedValue:    100.00,
			CreatedOnUtc: time.Now().AddDate(0, 0, -5), // Created 5 days ago
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedGiftCardUsageHistories, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedGiftCardUsageHistories, result)
	mockRepo.AssertExpectations(t)
}
