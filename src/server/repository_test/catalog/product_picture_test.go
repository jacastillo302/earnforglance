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

type MockSingleResultProductPicture struct {
	mock.Mock
}

func (m *MockSingleResultProductPicture) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.ProductPicture); ok {
		*v.(*domain.ProductPicture) = *result
	}
	return args.Error(1)
}

var mockItemProductPicture = &domain.ProductPicture{
	ID:           primitive.NewObjectID(), // Existing ID of the record to update
	ProductID:    primitive.NewObjectID(),
	PictureID:    102,
	DisplayOrder: 2,
}

func TestProductPictureRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionProductPicture

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultProductPicture{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemProductPicture, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewProductPictureRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemProductPicture.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultProductPicture{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewProductPictureRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemProductPicture.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestProductPictureRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionProductPicture

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemProductPicture).Return(nil, nil).Once()

	repo := repository.NewProductPictureRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemProductPicture)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestProductPictureRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionProductPicture

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemProductPicture.ID}
	update := bson.M{"$set": mockItemProductPicture}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewProductPictureRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemProductPicture)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
