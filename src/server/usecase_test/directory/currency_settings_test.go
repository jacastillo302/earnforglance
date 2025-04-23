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

func TestCurrencySettingsUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.CurrencySettingsRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewCurrencySettingsUsecase(mockRepo, timeout)

	currencySettingsID := bson.NewObjectID().Hex()

	updatedCurrencySettings := domain.CurrencySettings{
		ID:                                   bson.NewObjectID(), // Existing ID of the record to update
		DisplayCurrencyLabel:                 false,
		PrimaryStoreCurrencyID:               bson.NewObjectID(),
		PrimaryExchangeRateCurrencyID:        bson.NewObjectID(),
		ActiveExchangeRateProviderSystemName: "UpdatedExchangeRateProvider",
		AutoUpdateEnabled:                    false,
	}

	mockRepo.On("FetchByID", mock.Anything, currencySettingsID).Return(updatedCurrencySettings, nil)

	result, err := usecase.FetchByID(context.Background(), currencySettingsID)

	assert.NoError(t, err)
	assert.Equal(t, updatedCurrencySettings, result)
	mockRepo.AssertExpectations(t)
}

func TestCurrencySettingsUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.CurrencySettingsRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewCurrencySettingsUsecase(mockRepo, timeout)

	newCurrencySettings := &domain.CurrencySettings{
		DisplayCurrencyLabel:                 true,
		PrimaryStoreCurrencyID:               bson.NewObjectID(),
		PrimaryExchangeRateCurrencyID:        bson.NewObjectID(),
		ActiveExchangeRateProviderSystemName: "DefaultExchangeRateProvider",
		AutoUpdateEnabled:                    true,
	}
	mockRepo.On("Create", mock.Anything, newCurrencySettings).Return(nil)

	err := usecase.Create(context.Background(), newCurrencySettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCurrencySettingsUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.CurrencySettingsRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewCurrencySettingsUsecase(mockRepo, timeout)

	updatedCurrencySettings := &domain.CurrencySettings{
		ID:                                   bson.NewObjectID(), // Existing ID of the record to update
		DisplayCurrencyLabel:                 false,
		PrimaryStoreCurrencyID:               bson.NewObjectID(),
		PrimaryExchangeRateCurrencyID:        bson.NewObjectID(),
		ActiveExchangeRateProviderSystemName: "UpdatedExchangeRateProvider",
		AutoUpdateEnabled:                    false,
	}

	mockRepo.On("Update", mock.Anything, updatedCurrencySettings).Return(nil)

	err := usecase.Update(context.Background(), updatedCurrencySettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCurrencySettingsUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.CurrencySettingsRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewCurrencySettingsUsecase(mockRepo, timeout)

	currencySettingsID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, currencySettingsID).Return(nil)

	err := usecase.Delete(context.Background(), currencySettingsID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCurrencySettingsUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.CurrencySettingsRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewCurrencySettingsUsecase(mockRepo, timeout)

	fetchedCurrencySettings := []domain.CurrencySettings{
		{
			ID:                                   bson.NewObjectID(),
			DisplayCurrencyLabel:                 true,
			PrimaryStoreCurrencyID:               bson.NewObjectID(),
			PrimaryExchangeRateCurrencyID:        bson.NewObjectID(),
			ActiveExchangeRateProviderSystemName: "DefaultExchangeRateProvider",
			AutoUpdateEnabled:                    true,
		},
		{
			ID:                                   bson.NewObjectID(),
			DisplayCurrencyLabel:                 false,
			PrimaryStoreCurrencyID:               bson.NewObjectID(),
			PrimaryExchangeRateCurrencyID:        bson.NewObjectID(),
			ActiveExchangeRateProviderSystemName: "UpdatedExchangeRateProvider",
			AutoUpdateEnabled:                    false,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedCurrencySettings, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedCurrencySettings, result)
	mockRepo.AssertExpectations(t)
}
