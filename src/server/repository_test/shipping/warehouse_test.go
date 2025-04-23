package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/shipping"
	repository "earnforglance/server/repository/shipping"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type MockSingleResultWarehouse struct {
	mock.Mock
}

func (m *MockSingleResultWarehouse) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.Warehouse); ok {
		*v.(*domain.Warehouse) = *result
	}
	return args.Error(1)
}

var mockItemWarehouse = &domain.Warehouse{
	ID:           bson.NewObjectID(), // Existing ID of the record to update
	Name:         "Secondary Warehouse",
	AdminComment: "Backup storage facility.",
	AddressID:    bson.NewObjectID(),
}

func TestWarehouseRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionWarehouse

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultWarehouse{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemWarehouse, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewWarehouseRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemWarehouse.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultWarehouse{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewWarehouseRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemWarehouse.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestWarehouseRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionWarehouse

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemWarehouse).Return(nil, nil).Once()

	repo := repository.NewWarehouseRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemWarehouse)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestWarehouseRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionWarehouse

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemWarehouse.ID}
	update := bson.M{"$set": mockItemWarehouse}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewWarehouseRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemWarehouse)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
