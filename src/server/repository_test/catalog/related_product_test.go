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

	"go.mongodb.org/mongo-driver/v2/bson"
)

type MockSingleResultRelatedProduct struct {
	mock.Mock
}

func (m *MockSingleResultRelatedProduct) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.RelatedProduct); ok {
		*v.(*domain.RelatedProduct) = *result
	}
	return args.Error(1)
}

var mockItemRelatedProduct = &domain.RelatedProduct{
	ID:           bson.NewObjectID(), // Existing ID of the record to update
	ProductID1:   bson.NewObjectID(),
	ProductID2:   bson.NewObjectID(),
	DisplayOrder: 2,
}

func TestRelatedProductRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionRelatedProduct

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultRelatedProduct{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemRelatedProduct, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewRelatedProductRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemRelatedProduct.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultRelatedProduct{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewRelatedProductRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemRelatedProduct.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestRelatedProductRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionRelatedProduct

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemRelatedProduct).Return(nil, nil).Once()

	repo := repository.NewRelatedProductRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemRelatedProduct)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestRelatedProductRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionRelatedProduct

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemRelatedProduct.ID}
	update := bson.M{"$set": mockItemRelatedProduct}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewRelatedProductRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemRelatedProduct)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
