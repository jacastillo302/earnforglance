package repository_test

import (
	"context"
	"errors"
	"testing"
	"time"

	domain "earnforglance/server/domain/directory"
	repository "earnforglance/server/repository/directory"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultExchangeRate struct {
	mock.Mock
}

func (m *MockSingleResultExchangeRate) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.ExchangeRate); ok {
		*v.(*domain.ExchangeRate) = *result
	}
	return args.Error(1)
}

var mockItemExchangeRate = &domain.ExchangeRate{
	ID:           primitive.NewObjectID(), // Existing ID of the record to update
	CurrencyCode: "EUR",
	Rate:         0.85,
	UpdatedOn:    time.Now(),
}

func TestExchangeRateRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionExchangeRate

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultExchangeRate{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemExchangeRate, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewExchangeRateRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemExchangeRate.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultExchangeRate{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewExchangeRateRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemExchangeRate.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestExchangeRateRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionExchangeRate

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemExchangeRate).Return(nil, nil).Once()

	repo := repository.NewExchangeRateRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemExchangeRate)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestExchangeRateRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionExchangeRate

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemExchangeRate.ID}
	update := bson.M{"$set": mockItemExchangeRate}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewExchangeRateRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemExchangeRate)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
