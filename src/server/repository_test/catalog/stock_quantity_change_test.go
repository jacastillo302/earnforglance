package repository_test

import (
	"context"
	"errors"
	"testing"
	"time"

	domain "earnforglance/server/domain/catalog"
	repository "earnforglance/server/repository/catalog"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultStockQuantityChange struct {
	mock.Mock
}

func (m *MockSingleResultStockQuantityChange) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.StockQuantityChange); ok {
		*v.(*domain.StockQuantityChange) = *result
	}
	return args.Error(1)
}

var mockItemStockQuantityChange = &domain.StockQuantityChange{
	ID:                 primitive.NewObjectID(), // Existing ID of the record to update
	QuantityAdjustment: -5,
	StockQuantity:      95,
	Message:            "Stock reduced due to sale.",
	CreatedOnUtc:       time.Now(),
	ProductID:          primitive.NewObjectID(),
	CombinationID:      nil,
	WarehouseID:        primitive.NewObjectID(),
}

func TestStockQuantityChangeRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionStockQuantityChange

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultStockQuantityChange{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemStockQuantityChange, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewStockQuantityChangeRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemStockQuantityChange.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultStockQuantityChange{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewStockQuantityChangeRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemStockQuantityChange.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestStockQuantityChangeRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionStockQuantityChange

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemStockQuantityChange).Return(nil, nil).Once()

	repo := repository.NewStockQuantityChangeRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemStockQuantityChange)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestStockQuantityChangeRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionStockQuantityChange

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemStockQuantityChange.ID}
	update := bson.M{"$set": mockItemStockQuantityChange}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewStockQuantityChangeRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemStockQuantityChange)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
