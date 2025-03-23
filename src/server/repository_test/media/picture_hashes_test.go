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

type MockSingleResultPictureHashes struct {
	mock.Mock
}

func (m *MockSingleResultPictureHashes) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.PictureHashes); ok {
		*v.(*domain.PictureHashes) = *result
	}
	return args.Error(1)
}

var mockItemPictureHashes = &domain.PictureHashes{
	PictureID: primitive.NewObjectID(), // Existing PictureID of the record to update
	Hash:      []byte("updated_hash_data"),
}

func TestPictureHashesRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionPictureHashes

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultPictureHashes{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemPictureHashes, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewPictureHashesRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemPictureHashes.PictureID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultPictureHashes{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewPictureHashesRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemPictureHashes.PictureID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestPictureHashesRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionPictureHashes

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemPictureHashes).Return(nil, nil).Once()

	repo := repository.NewPictureHashesRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemPictureHashes)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestPictureHashesRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionPictureHashes

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemPictureHashes.PictureID}
	update := bson.M{"$set": mockItemPictureHashes}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewPictureHashesRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemPictureHashes)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
