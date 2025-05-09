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

func TestRecurringPaymentUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.RecurringPaymentRepository)
	timeout := time.Duration(10)
	usecase := test.NewRecurringPaymentUsecase(mockRepo, timeout)

	recurringPaymentID := bson.NewObjectID().Hex()

	updatedRecurringPayment := domain.RecurringPayment{
		ID:                            bson.NewObjectID(), // Existing ID of the record to update
		CycleLength:                   15,
		RecurringProductCyclePeriodID: 2,
		TotalCycles:                   6,
		StartDateUtc:                  time.Now().AddDate(0, 0, -30), // Started 30 days ago
		IsActive:                      false,
		LastPaymentFailed:             true,
		Deleted:                       true,
		OrderID:                       1002,
		CreatedOnUtc:                  time.Now().AddDate(0, 0, -60), // Created 60 days ago
		CyclePeriod:                   1,
	}

	mockRepo.On("FetchByID", mock.Anything, recurringPaymentID).Return(updatedRecurringPayment, nil)

	result, err := usecase.FetchByID(context.Background(), recurringPaymentID)

	assert.NoError(t, err)
	assert.Equal(t, updatedRecurringPayment, result)
	mockRepo.AssertExpectations(t)
}

func TestRecurringPaymentUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.RecurringPaymentRepository)
	timeout := time.Duration(10)
	usecase := test.NewRecurringPaymentUsecase(mockRepo, timeout)

	newRecurringPayment := &domain.RecurringPayment{
		CycleLength:                   30,
		RecurringProductCyclePeriodID: 1,
		TotalCycles:                   12,
		StartDateUtc:                  time.Now(),
		IsActive:                      true,
		LastPaymentFailed:             false,
		Deleted:                       false,
		OrderID:                       1001,
		CreatedOnUtc:                  time.Now(),
		CyclePeriod:                   2,
	}

	mockRepo.On("Create", mock.Anything, newRecurringPayment).Return(nil)

	err := usecase.Create(context.Background(), newRecurringPayment)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestRecurringPaymentUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.RecurringPaymentRepository)
	timeout := time.Duration(10)
	usecase := test.NewRecurringPaymentUsecase(mockRepo, timeout)

	updatedRecurringPayment := &domain.RecurringPayment{
		ID:                            bson.NewObjectID(), // Existing ID of the record to update
		CycleLength:                   15,
		RecurringProductCyclePeriodID: 2,
		TotalCycles:                   6,
		StartDateUtc:                  time.Now().AddDate(0, 0, -30), // Started 30 days ago
		IsActive:                      false,
		LastPaymentFailed:             true,
		Deleted:                       true,
		OrderID:                       1002,
		CreatedOnUtc:                  time.Now().AddDate(0, 0, -60), // Created 60 days ago
		CyclePeriod:                   1,
	}

	mockRepo.On("Update", mock.Anything, updatedRecurringPayment).Return(nil)

	err := usecase.Update(context.Background(), updatedRecurringPayment)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestRecurringPaymentUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.RecurringPaymentRepository)
	timeout := time.Duration(10)
	usecase := test.NewRecurringPaymentUsecase(mockRepo, timeout)

	recurringPaymentID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, recurringPaymentID).Return(nil)

	err := usecase.Delete(context.Background(), recurringPaymentID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestRecurringPaymentUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.RecurringPaymentRepository)
	timeout := time.Duration(10)
	usecase := test.NewRecurringPaymentUsecase(mockRepo, timeout)

	fetchedRecurringPayments := []domain.RecurringPayment{
		{
			ID:                            bson.NewObjectID(),
			CycleLength:                   30,
			RecurringProductCyclePeriodID: 1,
			TotalCycles:                   12,
			StartDateUtc:                  time.Now().AddDate(0, 0, -90), // Started 90 days ago
			IsActive:                      true,
			LastPaymentFailed:             false,
			Deleted:                       false,
			OrderID:                       1001,
			CreatedOnUtc:                  time.Now().AddDate(0, 0, -120), // Created 120 days ago
			CyclePeriod:                   2,
		},
		{
			ID:                            bson.NewObjectID(),
			CycleLength:                   15,
			RecurringProductCyclePeriodID: 2,
			TotalCycles:                   6,
			StartDateUtc:                  time.Now().AddDate(0, 0, -30), // Started 30 days ago
			IsActive:                      false,
			LastPaymentFailed:             true,
			Deleted:                       true,
			OrderID:                       1002,
			CreatedOnUtc:                  time.Now().AddDate(0, 0, -60), // Created 60 days ago
			CyclePeriod:                   1,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedRecurringPayments, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedRecurringPayments, result)
	mockRepo.AssertExpectations(t)
}
