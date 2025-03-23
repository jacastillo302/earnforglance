package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/directory"
	"earnforglance/server/mongo/mocks"
	repository "earnforglance/server/repository/directory"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultCurrencySettings struct {
	mock.Mock
}

func (m *MockSingleResultCurrencySettings) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.CurrencySettings); ok {
		*v.(*domain.CurrencySettings) = *result
	}
	return args.Error(1)
}

var mockItemCurrencySettings = &domain.CurrencySettings{
	ID:                                   primitive.NewObjectID(), // Existing ID of the record to update
	DisplayCurrencyLabel:                 false,
	PrimaryStoreCurrencyID:               primitive.NewObjectID(),
	PrimaryExchangeRateCurrencyID:        primitive.NewObjectID(),
	ActiveExchangeRateProviderSystemName: "UpdatedExchangeRateProvider",
	AutoUpdateEnabled:                    false,
}

func TestCurrencySettingsRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionCurrencySettings

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultCurrencySettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemCurrencySettings, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewCurrencySettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemCurrencySettings.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultCurrencySettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewCurrencySettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemCurrencySettings.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestCurrencySettingsRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionCurrencySettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemCurrencySettings).Return(nil, nil).Once()

	repo := repository.NewCurrencySettingsRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemCurrencySettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestCurrencySettingsRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionCurrencySettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemCurrencySettings.ID}
	update := bson.M{"$set": mockItemCurrencySettings}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewCurrencySettingsRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemCurrencySettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
