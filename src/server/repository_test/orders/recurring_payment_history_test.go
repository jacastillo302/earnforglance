package repository_test

import (
	"context"
	"errors"
	"testing"
	"time"

	domain "earnforglance/server/domain/orders"
	repository "earnforglance/server/repository/orders"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultRecurringPaymentHistory struct {
	mock.Mock
}

func (m *MockSingleResultRecurringPaymentHistory) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.RecurringPaymentHistory); ok {
		*v.(*domain.RecurringPaymentHistory) = *result
	}
	return args.Error(1)
}

var mockItemRecurringPaymentHistory = &domain.RecurringPaymentHistory{
	ID:                 primitive.NewObjectID(), // Existing ID of the record to update
	RecurringPaymentID: primitive.NewObjectID(),
	OrderID:            primitive.NewObjectID(),
	CreatedOnUtc:       time.Now().AddDate(0, 0, -7), // Created 7 days ago
}

func TestRecurringPaymentHistoryRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionRecurringPaymentHistory

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultRecurringPaymentHistory{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemRecurringPaymentHistory, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewRecurringPaymentHistoryRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemRecurringPaymentHistory.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultRecurringPaymentHistory{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewRecurringPaymentHistoryRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemRecurringPaymentHistory.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestRecurringPaymentHistoryRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionRecurringPaymentHistory

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemRecurringPaymentHistory).Return(nil, nil).Once()

	repo := repository.NewRecurringPaymentHistoryRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemRecurringPaymentHistory)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestRecurringPaymentHistoryRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionRecurringPaymentHistory

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemRecurringPaymentHistory.ID}
	update := bson.M{"$set": mockItemRecurringPaymentHistory}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewRecurringPaymentHistoryRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemRecurringPaymentHistory)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
