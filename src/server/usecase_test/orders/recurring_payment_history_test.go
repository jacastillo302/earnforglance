package usecase_test

import (
	"context"
	"testing"
	"time"

	mocks "earnforglance/server/domain/mocks"
	domain "earnforglance/server/domain/orders"
	test "earnforglance/server/usecase/orders"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestRecurringPaymentHistoryUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.RecurringPaymentHistoryRepository)
	timeout := time.Duration(10)
	usecase := test.NewRecurringPaymentHistoryUsecase(mockRepo, timeout)

	recurringPaymentHistoryID := bson.NewObjectID().Hex()

	updatedRecurringPaymentHistory := domain.RecurringPaymentHistory{
		ID:                 bson.NewObjectID(), // Existing ID of the record to update
		RecurringPaymentID: bson.NewObjectID(),
		OrderID:            bson.NewObjectID(),
		CreatedOnUtc:       time.Now().AddDate(0, 0, -7), // Created 7 days ago
	}

	mockRepo.On("FetchByID", mock.Anything, recurringPaymentHistoryID).Return(updatedRecurringPaymentHistory, nil)

	result, err := usecase.FetchByID(context.Background(), recurringPaymentHistoryID)

	assert.NoError(t, err)
	assert.Equal(t, updatedRecurringPaymentHistory, result)
	mockRepo.AssertExpectations(t)
}

func TestRecurringPaymentHistoryUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.RecurringPaymentHistoryRepository)
	timeout := time.Duration(10)
	usecase := test.NewRecurringPaymentHistoryUsecase(mockRepo, timeout)

	newRecurringPaymentHistory := &domain.RecurringPaymentHistory{
		RecurringPaymentID: bson.NewObjectID(),
		OrderID:            bson.NewObjectID(),
		CreatedOnUtc:       time.Now(),
	}

	mockRepo.On("Create", mock.Anything, newRecurringPaymentHistory).Return(nil)

	err := usecase.Create(context.Background(), newRecurringPaymentHistory)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestRecurringPaymentHistoryUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.RecurringPaymentHistoryRepository)
	timeout := time.Duration(10)
	usecase := test.NewRecurringPaymentHistoryUsecase(mockRepo, timeout)

	updatedRecurringPaymentHistory := &domain.RecurringPaymentHistory{
		ID:                 bson.NewObjectID(), // Existing ID of the record to update
		RecurringPaymentID: bson.NewObjectID(),
		OrderID:            bson.NewObjectID(),
		CreatedOnUtc:       time.Now().AddDate(0, 0, -7), // Created 7 days ago
	}

	mockRepo.On("Update", mock.Anything, updatedRecurringPaymentHistory).Return(nil)

	err := usecase.Update(context.Background(), updatedRecurringPaymentHistory)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestRecurringPaymentHistoryUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.RecurringPaymentHistoryRepository)
	timeout := time.Duration(10)
	usecase := test.NewRecurringPaymentHistoryUsecase(mockRepo, timeout)

	recurringPaymentHistoryID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, recurringPaymentHistoryID).Return(nil)

	err := usecase.Delete(context.Background(), recurringPaymentHistoryID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestRecurringPaymentHistoryUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.RecurringPaymentHistoryRepository)
	timeout := time.Duration(10)
	usecase := test.NewRecurringPaymentHistoryUsecase(mockRepo, timeout)

	fetchedRecurringPaymentHistories := []domain.RecurringPaymentHistory{
		{
			ID:                 bson.NewObjectID(),
			RecurringPaymentID: bson.NewObjectID(),
			OrderID:            bson.NewObjectID(),
			CreatedOnUtc:       time.Now().AddDate(0, 0, -10), // Created 10 days ago
		},
		{
			ID:                 bson.NewObjectID(),
			RecurringPaymentID: bson.NewObjectID(),
			OrderID:            bson.NewObjectID(),
			CreatedOnUtc:       time.Now().AddDate(0, 0, -5), // Created 5 days ago
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedRecurringPaymentHistories, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedRecurringPaymentHistories, result)
	mockRepo.AssertExpectations(t)
}
