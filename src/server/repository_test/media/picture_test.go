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

type MockSingleResultPicture struct {
	mock.Mock
}

func (m *MockSingleResultPicture) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.Picture); ok {
		*v.(*domain.Picture) = *result
	}
	return args.Error(1)
}

var mockItemPicture = &domain.Picture{
	ID:             bson.NewObjectID(), // Existing ID of the record to update
	MimeType:       "image/png",
	SeoFilename:    "updated-image",
	AltAttribute:   "Updated Image",
	TitleAttribute: "Updated Title",
	IsNew:          false,
	VirtualPath:    "/images/updated-image.png",
}

func TestPictureRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionPicture

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultPicture{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemPicture, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewPictureRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemPicture.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultPicture{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewPictureRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemPicture.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestPictureRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionPicture

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemPicture).Return(nil, nil).Once()

	repo := repository.NewPictureRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemPicture)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestPictureRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionPicture

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemPicture.ID}
	update := bson.M{"$set": mockItemPicture}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewPictureRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemPicture)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
