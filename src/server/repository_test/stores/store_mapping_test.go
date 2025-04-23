package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/stores"
	repository "earnforglance/server/repository/stores"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type MockSingleResultStoreMapping struct {
	mock.Mock
}

func (m *MockSingleResultStoreMapping) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.StoreMapping); ok {
		*v.(*domain.StoreMapping) = *result
	}
	return args.Error(1)
}

var mockItemStoreMapping = &domain.StoreMapping{
	ID:         bson.NewObjectID(), // Existing ID of the record to update
	EntityID:   bson.NewObjectID(),
	EntityName: "Category",
	StoreID:    bson.NewObjectID(),
}

func TestStoreMappingRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionStoreMapping

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultStoreMapping{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemStoreMapping, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewStoreMappingRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemStoreMapping.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultStoreMapping{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewStoreMappingRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemStoreMapping.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestStoreMappingRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionStoreMapping

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemStoreMapping).Return(nil, nil).Once()

	repo := repository.NewStoreMappingRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemStoreMapping)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestStoreMappingRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionStoreMapping

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemStoreMapping.ID}
	update := bson.M{"$set": mockItemStoreMapping}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewStoreMappingRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemStoreMapping)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
