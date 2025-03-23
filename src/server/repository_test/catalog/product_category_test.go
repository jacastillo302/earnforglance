package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/catalog"
	"earnforglance/server/mongo/mocks"
	repository "earnforglance/server/repository/catalog"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultProductCategory struct {
	mock.Mock
}

func (m *MockSingleResultProductCategory) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.ProductCategory); ok {
		*v.(*domain.ProductCategory) = *result
	}
	return args.Error(1)
}

var mockItemProductCategory = &domain.ProductCategory{
	ID:                primitive.NewObjectID(),
	ProductID:         primitive.NewObjectID(),
	CategoryID:        primitive.NewObjectID(),
	IsFeaturedProduct: false,
	DisplayOrder:      2,
}

func TestProductCategoryRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionProductCategory

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultProductCategory{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemProductCategory, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewProductCategoryRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemProductCategory.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultProductCategory{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewProductCategoryRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemProductCategory.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestProductCategoryRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionProductCategory

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemProductCategory).Return(nil, nil).Once()

	repo := repository.NewProductCategoryRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemProductCategory)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestProductCategoryRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionProductCategory

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemProductCategory.ID}
	update := bson.M{"$set": mockItemProductCategory}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewProductCategoryRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemProductCategory)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
