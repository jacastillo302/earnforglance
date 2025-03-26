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

type MockSingleResultProductProductTagMapping struct {
	mock.Mock
}

func (m *MockSingleResultProductProductTagMapping) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.ProductProductTagMapping); ok {
		*v.(*domain.ProductProductTagMapping) = *result
	}
	return args.Error(1)
}

var mockItemProductProductTagMapping = &domain.ProductProductTagMapping{
	ID:           primitive.NewObjectID(), // Existing ID of the record to update
	ProductID:    primitive.NewObjectID(),
	ProductTagID: primitive.NewObjectID(),
}

func TestProductProductTagMappingRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionProductProductTagMapping

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultProductProductTagMapping{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemProductProductTagMapping, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewProductProductTagMappingRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemProductProductTagMapping.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultProductProductTagMapping{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewProductProductTagMappingRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemProductProductTagMapping.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestProductProductTagMappingRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionProductProductTagMapping

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemProductProductTagMapping).Return(nil, nil).Once()

	repo := repository.NewProductProductTagMappingRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemProductProductTagMapping)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestProductProductTagMappingRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionProductProductTagMapping

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemProductProductTagMapping.ID}
	update := bson.M{"$set": mockItemProductProductTagMapping}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewProductProductTagMappingRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemProductProductTagMapping)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
