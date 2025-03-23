package repository_test

import (
	"context"
	"errors"
	"testing"
	"time"

	domain "earnforglance/server/domain/customers"
	"earnforglance/server/mongo/mocks"
	repository "earnforglance/server/repository/customers"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultRewardPointsHistory struct {
	mock.Mock
}

func (m *MockSingleResultRewardPointsHistory) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.RewardPointsHistory); ok {
		*v.(*domain.RewardPointsHistory) = *result
	}
	return args.Error(1)
}

var mockItemRewardPointsHistory = &domain.RewardPointsHistory{
	ID:            primitive.NewObjectID(), // Existing ID of the record to update
	CustomerID:    primitive.NewObjectID(),
	StoreID:       primitive.NewObjectID(),
	Points:        -50,
	PointsBalance: new(int),
	UsedAmount:    25.0,
	Message:       "Reward points used for order",
	CreatedOnUtc:  time.Now().AddDate(0, 0, -7), // Created 7 days ago
	EndDateUtc:    new(time.Time),
	ValidPoints:   new(int),
	UsedWithOrder: nil,
}

func TestRewardPointsHistoryRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionRewardPointsHistory

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultRewardPointsHistory{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemRewardPointsHistory, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewRewardPointsHistoryRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemRewardPointsHistory.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultRewardPointsHistory{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewRewardPointsHistoryRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemRewardPointsHistory.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestRewardPointsHistoryRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionRewardPointsHistory

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemRewardPointsHistory).Return(nil, nil).Once()

	repo := repository.NewRewardPointsHistoryRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemRewardPointsHistory)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestRewardPointsHistoryRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionRewardPointsHistory

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemRewardPointsHistory.ID}
	update := bson.M{"$set": mockItemRewardPointsHistory}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewRewardPointsHistoryRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemRewardPointsHistory)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
