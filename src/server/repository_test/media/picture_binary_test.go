package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/media"
	"earnforglance/server/mongo/mocks"
	repository "earnforglance/server/repository/media"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultPictureBinary struct {
	mock.Mock
}

func (m *MockSingleResultPictureBinary) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.PictureBinary); ok {
		*v.(*domain.PictureBinary) = *result
	}
	return args.Error(1)
}

var mockItemPictureBinary = &domain.PictureBinary{
	ID:         primitive.NewObjectID(), // Existing ID of the record to update
	BinaryData: []byte("updated binary image data"),
	PictureID:  primitive.NewObjectID(),
}

func TestPictureBinaryRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionPictureBinary

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultPictureBinary{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemPictureBinary, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewPictureBinaryRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemPictureBinary.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultPictureBinary{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewPictureBinaryRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemPictureBinary.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestPictureBinaryRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionPictureBinary

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemPictureBinary).Return(nil, nil).Once()

	repo := repository.NewPictureBinaryRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemPictureBinary)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestPictureBinaryRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionPictureBinary

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemPictureBinary.ID}
	update := bson.M{"$set": mockItemPictureBinary}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewPictureBinaryRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemPictureBinary)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
