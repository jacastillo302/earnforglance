package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/shipping"
	"earnforglance/server/mongo/mocks"
	repository "earnforglance/server/repository/shipping"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultShipmentItem struct {
	mock.Mock
}

func (m *MockSingleResultShipmentItem) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.ShipmentItem); ok {
		*v.(*domain.ShipmentItem) = *result
	}
	return args.Error(1)
}

var mockItemShipmentItem = &domain.ShipmentItem{
	ID:          primitive.NewObjectID(), // Existing ID of the record to update
	ShipmentID:  primitive.NewObjectID(),
	OrderItemID: primitive.NewObjectID(),
	Quantity:    15,
	WarehouseID: primitive.NewObjectID(),
}

func TestShipmentItemRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionShipmentItem

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultShipmentItem{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemShipmentItem, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewShipmentItemRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemShipmentItem.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultShipmentItem{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewShipmentItemRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemShipmentItem.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestShipmentItemRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionShipmentItem

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemShipmentItem).Return(nil, nil).Once()

	repo := repository.NewShipmentItemRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemShipmentItem)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestShipmentItemRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionShipmentItem

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemShipmentItem.ID}
	update := bson.M{"$set": mockItemShipmentItem}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewShipmentItemRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemShipmentItem)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
