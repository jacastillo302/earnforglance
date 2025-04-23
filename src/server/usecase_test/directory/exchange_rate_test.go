package usecase_test

import (
	"context"
	domain "earnforglance/server/domain/directory"
	mocks "earnforglance/server/domain/mocks"
	test "earnforglance/server/usecase/directory"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestExchangeRateUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.ExchangeRateRepository)
	timeout := time.Duration(10)
	usecase := test.NewExchangeRateUsecase(mockRepo, timeout)

	exchangeRateID := bson.NewObjectID().Hex()

	updatedExchangeRate := domain.ExchangeRate{
		ID:           bson.NewObjectID(), // Existing ID of the record to update
		CurrencyCode: "EUR",
		Rate:         0.85,
		UpdatedOn:    time.Now(),
	}

	mockRepo.On("FetchByID", mock.Anything, exchangeRateID).Return(updatedExchangeRate, nil)

	result, err := usecase.FetchByID(context.Background(), exchangeRateID)

	assert.NoError(t, err)
	assert.Equal(t, updatedExchangeRate, result)
	mockRepo.AssertExpectations(t)
}

func TestExchangeRateUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.ExchangeRateRepository)
	timeout := time.Duration(10)
	usecase := test.NewExchangeRateUsecase(mockRepo, timeout)

	newExchangeRate := &domain.ExchangeRate{
		CurrencyCode: "USD",
		Rate:         1.0,
		UpdatedOn:    time.Now(),
	}

	mockRepo.On("Create", mock.Anything, newExchangeRate).Return(nil)

	err := usecase.Create(context.Background(), newExchangeRate)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestExchangeRateUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.ExchangeRateRepository)
	timeout := time.Duration(10)
	usecase := test.NewExchangeRateUsecase(mockRepo, timeout)

	updatedExchangeRate := &domain.ExchangeRate{
		ID:           bson.NewObjectID(), // Existing ID of the record to update
		CurrencyCode: "EUR",
		Rate:         0.85,
		UpdatedOn:    time.Now(),
	}

	mockRepo.On("Update", mock.Anything, updatedExchangeRate).Return(nil)

	err := usecase.Update(context.Background(), updatedExchangeRate)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestExchangeRateUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.ExchangeRateRepository)
	timeout := time.Duration(10)
	usecase := test.NewExchangeRateUsecase(mockRepo, timeout)

	exchangeRateID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, exchangeRateID).Return(nil)

	err := usecase.Delete(context.Background(), exchangeRateID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestExchangeRateUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.ExchangeRateRepository)
	timeout := time.Duration(10)
	usecase := test.NewExchangeRateUsecase(mockRepo, timeout)

	fetchedExchangeRates := []domain.ExchangeRate{
		{
			ID:           bson.NewObjectID(),
			CurrencyCode: "USD",
			Rate:         1.0,
			UpdatedOn:    time.Now().AddDate(0, 0, -10), // Updated 10 days ago
		},
		{
			ID:           bson.NewObjectID(),
			CurrencyCode: "EUR",
			Rate:         0.85,
			UpdatedOn:    time.Now().AddDate(0, 0, -5), // Updated 5 days ago
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedExchangeRates, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedExchangeRates, result)
	mockRepo.AssertExpectations(t)
}
