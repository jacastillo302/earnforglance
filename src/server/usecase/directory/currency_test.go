package usecase

import (
	"context"
	domain "earnforglance/server/domain/directory"
	mocks "earnforglance/server/domain/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCurrencyUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.CurrencyRepository)
	timeout := time.Duration(10)
	usecase := NewCurrencyUsecase(mockRepo, timeout)

	currencyID := primitive.NewObjectID().Hex()

	updatedCurrency := domain.Currency{
		ID:               primitive.NewObjectID(), // Existing ID of the record to update
		Name:             "Euro",
		CurrencyCode:     "EUR",
		Rate:             0.85,
		DisplayLocale:    "de-DE",
		CustomFormatting: "€#,##0.00",
		LimitedToStores:  true,
		Published:        false,
		DisplayOrder:     2,
		CreatedOnUtc:     time.Now().AddDate(0, 0, -30), // Created 30 days ago
		UpdatedOnUtc:     time.Now(),
		RoundingTypeID:   2,
		RoundingType:     5,
	}

	mockRepo.On("FetchByID", mock.Anything, currencyID).Return(updatedCurrency, nil)

	result, err := usecase.FetchByID(context.Background(), currencyID)

	assert.NoError(t, err)
	assert.Equal(t, updatedCurrency, result)
	mockRepo.AssertExpectations(t)
}

func TestCurrencyUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.CurrencyRepository)
	timeout := time.Duration(10)
	usecase := NewCurrencyUsecase(mockRepo, timeout)

	newCurrency := &domain.Currency{
		Name:             "US Dollar",
		CurrencyCode:     "USD",
		Rate:             1.0,
		DisplayLocale:    "en-US",
		CustomFormatting: "$#,##0.00",
		LimitedToStores:  false,
		Published:        true,
		DisplayOrder:     1,
		CreatedOnUtc:     time.Now(),
		UpdatedOnUtc:     time.Now(),
		RoundingTypeID:   1,
		RoundingType:     4,
	}

	mockRepo.On("Create", mock.Anything, newCurrency).Return(nil)

	err := usecase.Create(context.Background(), newCurrency)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCurrencyUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.CurrencyRepository)
	timeout := time.Duration(10)
	usecase := NewCurrencyUsecase(mockRepo, timeout)

	updatedCurrency := &domain.Currency{
		ID:               primitive.NewObjectID(), // Existing ID of the record to update
		Name:             "Euro",
		CurrencyCode:     "EUR",
		Rate:             0.85,
		DisplayLocale:    "de-DE",
		CustomFormatting: "€#,##0.00",
		LimitedToStores:  true,
		Published:        false,
		DisplayOrder:     2,
		CreatedOnUtc:     time.Now().AddDate(0, 0, -30), // Created 30 days ago
		UpdatedOnUtc:     time.Now(),
		RoundingTypeID:   2,
		RoundingType:     4,
	}

	mockRepo.On("Update", mock.Anything, updatedCurrency).Return(nil)

	err := usecase.Update(context.Background(), updatedCurrency)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCurrencyUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.CurrencyRepository)
	timeout := time.Duration(10)
	usecase := NewCurrencyUsecase(mockRepo, timeout)

	currencyID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, currencyID).Return(nil)

	err := usecase.Delete(context.Background(), currencyID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCurrencyUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.CurrencyRepository)
	timeout := time.Duration(10)
	usecase := NewCurrencyUsecase(mockRepo, timeout)

	fetchedCurrencies := []domain.Currency{
		{
			ID:               primitive.NewObjectID(),
			Name:             "US Dollar",
			CurrencyCode:     "USD",
			Rate:             1.0,
			DisplayLocale:    "en-US",
			CustomFormatting: "$#,##0.00",
			LimitedToStores:  false,
			Published:        true,
			DisplayOrder:     1,
			CreatedOnUtc:     time.Now().AddDate(0, 0, -10), // Created 10 days ago
			UpdatedOnUtc:     time.Now(),
			RoundingTypeID:   1,
			RoundingType:     6,
		},
		{
			ID:               primitive.NewObjectID(),
			Name:             "Euro",
			CurrencyCode:     "EUR",
			Rate:             0.85,
			DisplayLocale:    "de-DE",
			CustomFormatting: "€#,##0.00",
			LimitedToStores:  true,
			Published:        false,
			DisplayOrder:     2,
			CreatedOnUtc:     time.Now().AddDate(0, 0, -30), // Created 30 days ago
			UpdatedOnUtc:     time.Now(),
			RoundingTypeID:   2,
			RoundingType:     7,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedCurrencies, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedCurrencies, result)
	mockRepo.AssertExpectations(t)
}
