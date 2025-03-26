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

type MockSingleResultProductTag struct {
	mock.Mock
}

func (m *MockSingleResultProductTag) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.ProductTag); ok {
		*v.(*domain.ProductTag) = *result
	}
	return args.Error(1)
}

var mockItemProductTag = &domain.ProductTag{
	ID:              primitive.NewObjectID(), // Existing ID of the record to update
	Name:            "Home Appliances",
	MetaDescription: "Tags related to home appliances",
	MetaKeywords:    "appliances, home, kitchen",
	MetaTitle:       "Home Appliances Tag",
}

func TestProductTagRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionProductTag

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultProductTag{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemProductTag, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewProductTagRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemProductTag.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultProductTag{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewProductTagRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemProductTag.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestProductTagRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionProductTag

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemProductTag).Return(nil, nil).Once()

	repo := repository.NewProductTagRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemProductTag)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestProductTagRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionProductTag

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemProductTag.ID}
	update := bson.M{"$set": mockItemProductTag}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewProductTagRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemProductTag)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
