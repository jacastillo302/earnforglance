package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/catalog"
	repository "earnforglance/server/repository/catalog"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultProductWarehouseInventory struct {
	mock.Mock
}

func (m *MockSingleResultProductWarehouseInventory) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.ProductWarehouseInventory); ok {
		*v.(*domain.ProductWarehouseInventory) = *result
	}
	return args.Error(1)
}

var mockItemProductWarehouseInventory = &domain.ProductWarehouseInventory{
	ProductID:        primitive.NewObjectID(),
	WarehouseID:      primitive.NewObjectID(),
	StockQuantity:    150,
	ReservedQuantity: 20,
}

func TestProductWarehouseInventoryRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionProductWarehouseInventory

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultProductWarehouseInventory{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemProductWarehouseInventory, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewProductWarehouseInventoryRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemProductWarehouseInventory.ProductID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultProductWarehouseInventory{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewProductWarehouseInventoryRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), string(mockItemProductWarehouseInventory.ProductID.Hex()))

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestProductWarehouseInventoryRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionProductWarehouseInventory

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemProductWarehouseInventory).Return(nil, nil).Once()

	repo := repository.NewProductWarehouseInventoryRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemProductWarehouseInventory)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestProductWarehouseInventoryRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionProductWarehouseInventory

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemProductWarehouseInventory.ProductID}
	update := bson.M{"$set": mockItemProductWarehouseInventory}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewProductWarehouseInventoryRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemProductWarehouseInventory)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
