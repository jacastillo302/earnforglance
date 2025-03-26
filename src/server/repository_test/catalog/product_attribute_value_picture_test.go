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

type MockSingleResultProductAttributeValuePicture struct {
	mock.Mock
}

func (m *MockSingleResultProductAttributeValuePicture) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.ProductAttributeValuePicture); ok {
		*v.(*domain.ProductAttributeValuePicture) = *result
	}
	return args.Error(1)
}

var mockItemProductAttributeValuePicture = &domain.ProductAttributeValuePicture{
	ID:                      primitive.NewObjectID(), // Existing ID of the record to update
	ProductAttributeValueID: primitive.NewObjectID(),
	PictureID:               primitive.NewObjectID(),
}

func TestProductAttributeValuePictureRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionProductAttributeValuePicture

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultProductAttributeValuePicture{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemProductAttributeValuePicture, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewProductAttributeValuePictureRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemProductAttributeValuePicture.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultProductAttributeValuePicture{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewProductAttributeValuePictureRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemProductAttributeValuePicture.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestProductAttributeValuePictureRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionProductAttributeValuePicture

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemProductAttributeValuePicture).Return(nil, nil).Once()

	repo := repository.NewProductAttributeValuePictureRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemProductAttributeValuePicture)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestProductAttributeValuePictureRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionProductAttributeValuePicture

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemProductAttributeValuePicture.ID}
	update := bson.M{"$set": mockItemProductAttributeValuePicture}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewProductAttributeValuePictureRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemProductAttributeValuePicture)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
