package repository_test

import (
	"context"
	"errors"
	"testing"
	"time"

	domain "earnforglance/server/domain/directory"
	"earnforglance/server/mongo/mocks"
	repository "earnforglance/server/repository/directory"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultCurrency struct {
	mock.Mock
}

func (m *MockSingleResultCurrency) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.Currency); ok {
		*v.(*domain.Currency) = *result
	}
	return args.Error(1)
}

var mockItemCurrency = &domain.Currency{
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
}

func TestCurrencyRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionCurrency

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultCurrency{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemCurrency, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewCurrencyRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemCurrency.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultCurrency{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewCurrencyRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemCurrency.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestCurrencyRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionCurrency

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemCurrency).Return(nil, nil).Once()

	repo := repository.NewCurrencyRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemCurrency)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestCurrencyRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionCurrency

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemCurrency.ID}
	update := bson.M{"$set": mockItemCurrency}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewCurrencyRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemCurrency)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
