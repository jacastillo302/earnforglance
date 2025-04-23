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

type MockSingleResultProductAttributeCombinationPicture struct {
	mock.Mock
}

func (m *MockSingleResultProductAttributeCombinationPicture) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.ProductAttributeCombinationPicture); ok {
		*v.(*domain.ProductAttributeCombinationPicture) = *result
	}
	return args.Error(1)
}

func TestProductAttributeCombinationPictureRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionProductAttributeCombinationPicture

	mockItem := domain.ProductAttributeCombinationPicture{
		ID:                            bson.NewObjectID(),
		ProductAttributeCombinationID: bson.NewObjectID(),
		PictureID:                     bson.NewObjectID(),
	}

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultProductAttributeCombinationPicture{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItem, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewProductAttributeCombinationPictureRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItem.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultProductAttributeCombinationPicture{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewProductAttributeCombinationPictureRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItem.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestProductAttributeCombinationPictureRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionProductAttributeCombinationPicture

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)

	mockProductAttributeCombinationPicture := &domain.ProductAttributeCombinationPicture{
		ID:                            bson.NewObjectID(),
		ProductAttributeCombinationID: bson.NewObjectID(),
		PictureID:                     bson.NewObjectID(),
	}

	collectionHelper.On("InsertOne", mock.Anything, mockProductAttributeCombinationPicture).Return(nil, nil).Once()

	repo := repository.NewProductAttributeCombinationPictureRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockProductAttributeCombinationPicture)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestProductAttributeCombinationPictureRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionProductAttributeCombinationPicture

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)

	mockProductAttributeCombinationPicture := &domain.ProductAttributeCombinationPicture{
		ID:                            bson.NewObjectID(),
		ProductAttributeCombinationID: bson.NewObjectID(),
		PictureID:                     bson.NewObjectID(),
	}

	filter := bson.M{"_id": mockProductAttributeCombinationPicture.ID}
	update := bson.M{"$set": mockProductAttributeCombinationPicture}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewProductAttributeCombinationPictureRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockProductAttributeCombinationPicture)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
