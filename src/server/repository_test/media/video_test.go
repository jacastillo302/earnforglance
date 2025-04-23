package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/media"
	repository "earnforglance/server/repository/media"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type MockSingleResultVideo struct {
	mock.Mock
}

func (m *MockSingleResultVideo) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.Video); ok {
		*v.(*domain.Video) = *result
	}
	return args.Error(1)
}

var mockItemVideo = &domain.Video{
	ID:       bson.NewObjectID(), // Existing ID of the record to update
	VideoUrl: "https://example.com/updated-video.mp4",
}

func TestVideoRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionVideo

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultVideo{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemVideo, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewVideoRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemVideo.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultVideo{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewVideoRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemVideo.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestVideoRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionVideo

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemVideo).Return(nil, nil).Once()

	repo := repository.NewVideoRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemVideo)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestVideoRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionVideo

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemVideo.ID}
	update := bson.M{"$set": mockItemVideo}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewVideoRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemVideo)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
