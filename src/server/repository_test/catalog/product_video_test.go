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

type MockSingleResultProductVideo struct {
	mock.Mock
}

func (m *MockSingleResultProductVideo) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.ProductVideo); ok {
		*v.(*domain.ProductVideo) = *result
	}
	return args.Error(1)
}

var mockItemProductVideo = &domain.ProductVideo{
	ID:           bson.NewObjectID(), // Existing ID of the record to update
	ProductID:    bson.NewObjectID(),
	VideoID:      bson.NewObjectID(),
	DisplayOrder: 2,
}

func TestProductVideoRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionProductVideo

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultProductVideo{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemProductVideo, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewProductVideoRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemProductVideo.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultProductVideo{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewProductVideoRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemProductVideo.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestProductVideoRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionProductVideo

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemProductVideo).Return(nil, nil).Once()

	repo := repository.NewProductVideoRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemProductVideo)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestProductVideoRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionProductVideo

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemProductVideo.ID}
	update := bson.M{"$set": mockItemProductVideo}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewProductVideoRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemProductVideo)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
