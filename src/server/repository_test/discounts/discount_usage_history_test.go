package repository_test

import (
	"context"
	"errors"
	"testing"
	"time"

	domain "earnforglance/server/domain/discounts"
	"earnforglance/server/mongo/mocks"
	repository "earnforglance/server/repository/discounts"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultDiscountUsageHistory struct {
	mock.Mock
}

func (m *MockSingleResultDiscountUsageHistory) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.DiscountUsageHistory); ok {
		*v.(*domain.DiscountUsageHistory) = *result
	}
	return args.Error(1)
}

var mockItemDiscountUsageHistory = &domain.DiscountUsageHistory{
	ID:           primitive.NewObjectID(), // Existing ID of the record to update
	DiscountID:   primitive.NewObjectID(),
	OrderID:      primitive.NewObjectID(),
	CreatedOnUtc: time.Now().AddDate(0, 0, -7), // Created 7 days ago
}

func TestDiscountUsageHistoryRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionDiscountUsageHistory

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultDiscountUsageHistory{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemDiscountUsageHistory, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewDiscountUsageHistoryRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemDiscountUsageHistory.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultDiscountUsageHistory{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewDiscountUsageHistoryRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemDiscountUsageHistory.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestDiscountUsageHistoryRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionDiscountUsageHistory

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemDiscountUsageHistory).Return(nil, nil).Once()

	repo := repository.NewDiscountUsageHistoryRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemDiscountUsageHistory)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestDiscountUsageHistoryRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionDiscountUsageHistory

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemDiscountUsageHistory.ID}
	update := bson.M{"$set": mockItemDiscountUsageHistory}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewDiscountUsageHistoryRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemDiscountUsageHistory)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
