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

	"go.mongodb.org/mongo-driver/v2/bson"
)

type MockSingleResultGiftCardUsageHistory struct {
	mock.Mock
}

func (m *MockSingleResultGiftCardUsageHistory) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.GiftCardUsageHistory); ok {
		*v.(*domain.GiftCardUsageHistory) = *result
	}
	return args.Error(1)
}

var mockItemGiftCardUsageHistory = &domain.GiftCardUsageHistory{
	ID:           bson.NewObjectID(), // Existing ID of the record to update
	GiftCardID:   bson.NewObjectID(),
	OrderID:      bson.NewObjectID(),
	UsedValue:    75.00,
	CreatedOnUtc: time.Now().AddDate(0, 0, -7), // Created 7 days ago
}

func TestGiftCardUsageHistoryRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionGiftCardUsageHistory

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultGiftCardUsageHistory{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemGiftCardUsageHistory, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewGiftCardUsageHistoryRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemGiftCardUsageHistory.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultGiftCardUsageHistory{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewGiftCardUsageHistoryRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemGiftCardUsageHistory.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestGiftCardUsageHistoryRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionGiftCardUsageHistory

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemGiftCardUsageHistory).Return(nil, nil).Once()

	repo := repository.NewGiftCardUsageHistoryRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemGiftCardUsageHistory)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestGiftCardUsageHistoryRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionGiftCardUsageHistory

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemGiftCardUsageHistory.ID}
	update := bson.M{"$set": mockItemGiftCardUsageHistory}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewGiftCardUsageHistoryRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemGiftCardUsageHistory)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
