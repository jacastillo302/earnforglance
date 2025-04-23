package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/media"
	repository "earnforglance/server/repository/media"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type MockSingleResultDownload struct {
	mock.Mock
}

func (m *MockSingleResultDownload) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.Download); ok {
		*v.(*domain.Download) = *result
	}
	return args.Error(1)
}

var mockItemDownload = &domain.Download{
	ID:             bson.NewObjectID(), // Existing ID of the record to update
	DownloadGuid:   uuid.New(),
	UseDownloadUrl: false,
	DownloadUrl:    "",
	DownloadBinary: []byte("binary data"),
	ContentType:    "image/png",
	Filename:       "image.png",
	Extension:      ".png",
	IsNew:          false,
}

func TestDownloadRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionDownload

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultDownload{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemDownload, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewDownloadRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemDownload.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultDownload{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewDownloadRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemDownload.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestDownloadRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionDownload

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemDownload).Return(nil, nil).Once()

	repo := repository.NewDownloadRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemDownload)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestDownloadRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionDownload

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemDownload.ID}
	update := bson.M{"$set": mockItemDownload}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewDownloadRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemDownload)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
