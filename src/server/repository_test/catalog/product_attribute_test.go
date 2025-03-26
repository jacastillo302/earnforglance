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

type MockSingleResultProductAttribute struct {
	mock.Mock
}

func (m *MockSingleResultProductAttribute) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.ProductAttribute); ok {
		*v.(*domain.ProductAttribute) = *result
	}
	return args.Error(1)
}

var mockItemProductAttribute = &domain.ProductAttribute{
	ID:          primitive.NewObjectID(),
	Name:        "Size",
	Description: "Defines the size of the product",
}

func TestProductAttributeRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionProductAttribute

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultProductAttribute{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemProductAttribute, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewProductAttributeRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemProductAttribute.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultProductAttribute{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewProductAttributeRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemProductAttribute.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestProductAttributeRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionProductAttribute

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemProductAttribute).Return(nil, nil).Once()

	repo := repository.NewProductAttributeRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemProductAttribute)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestProductAttributeRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionProductAttribute

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)

	filter := bson.M{"_id": mockItemProductAttribute.ID}
	update := bson.M{"$set": mockItemProductAttribute}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewProductAttributeRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemProductAttribute)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
