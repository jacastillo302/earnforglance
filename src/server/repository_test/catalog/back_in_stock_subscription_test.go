package repository_test

import (
	"context"
	domain "earnforglance/server/domain/catalog"
	"earnforglance/server/mongo/mocks"
	repository "earnforglance/server/repository/catalog"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultBackInStockSubscription struct {
	mock.Mock
}

func (m *MockSingleResultBackInStockSubscription) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.BackInStockSubscription); ok {
		*v.(*domain.BackInStockSubscription) = *result
	}
	return args.Error(1)
}

func TestBackInStockSubscriptionRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionBackInStockSubscription

	mockItem := domain.BackInStockSubscription{
		ID:           primitive.ObjectID{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, // Existing ID of the record to update
		StoreID:      primitive.ObjectID{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
		ProductID:    primitive.ObjectID{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
		CustomerID:   primitive.ObjectID{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
		CreatedOnUtc: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
	}

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultBackInStockSubscription{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItem, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewBackInStockSubscriptionRepository(databaseHelper, collectionName)

		result, err := ur.FetchByID(context.Background(), mockItem.ID.Hex())

		assert.NoError(t, err)
		assert.Equal(t, mockItem, result)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultBackInStockSubscription{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewBackInStockSubscriptionRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItem.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestBackInStockSubscriptionRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionBackInStockSubscription

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)

	mockBackInStockSubscription := &domain.BackInStockSubscription{
		ID:           primitive.NewObjectID(), // Existing ID of the record to update
		StoreID:      primitive.NewObjectID(),
		ProductID:    primitive.NewObjectID(),
		CustomerID:   primitive.NewObjectID(),
		CreatedOnUtc: time.Now(),
	}

	collectionHelper.On("InsertOne", mock.Anything, mockBackInStockSubscription).Return(nil, nil).Once()

	repo := repository.NewBackInStockSubscriptionRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockBackInStockSubscription)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestBackInStockSubscriptionRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionBackInStockSubscription

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)

	mockBackInStockSubscription := &domain.BackInStockSubscription{
		ID:           primitive.NewObjectID(), // Existing ID of the record to update
		StoreID:      primitive.NewObjectID(),
		ProductID:    primitive.NewObjectID(),
		CustomerID:   primitive.NewObjectID(),
		CreatedOnUtc: time.Now(),
	}

	filter := bson.M{"_id": mockBackInStockSubscription.ID}
	update := bson.M{"$set": mockBackInStockSubscription}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewBackInStockSubscriptionRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockBackInStockSubscription)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
