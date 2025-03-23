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

type MockSingleResultMediaSettings struct {
	mock.Mock
}

func (m *MockSingleResultMediaSettings) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.MediaSettings); ok {
		*v.(*domain.MediaSettings) = *result
	}
	return args.Error(1)
}

var mockItemMediaSettings = &domain.MediaSettings{
	ID:                        primitive.NewObjectID(), // Existing ID of the record to update
	AvatarPictureSize:         120,
	ProductThumbPictureSize:   250,
	ProductDetailsPictureSize: 450,
	ProductThumbPictureSizeOnProductDetailsPage: 180,
	AssociatedProductPictureSize:                140,
	CategoryThumbPictureSize:                    200,
	ManufacturerThumbPictureSize:                180,
	VendorThumbPictureSize:                      160,
	CartThumbPictureSize:                        100,
	OrderThumbPictureSize:                       110,
	MiniCartThumbPictureSize:                    90,
	AutoCompleteSearchThumbPictureSize:          80,
	ImageSquarePictureSize:                      60,
	DefaultPictureZoomEnabled:                   false,
	AllowSVGUploads:                             false,
	MaximumImageSize:                            6000,
	DefaultImageQuality:                         85,
	MultipleThumbDirectories:                    true,
	ImportProductImagesUsingHash:                false,
	AzureCacheControlHeader:                     "private, max-age=86400",
	UseAbsoluteImagePath:                        true,
	VideoIframeAllow:                            "autoplay",
	VideoIframeWidth:                            800,
	VideoIframeHeight:                           450,
	ProductDefaultImageID:                       primitive.NewObjectID(),
	AutoOrientImage:                             false,
}

func TestMediaSettingsRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionMediaSettings

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultMediaSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemMediaSettings, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewMediaSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemMediaSettings.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultMediaSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewMediaSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemMediaSettings.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestMediaSettingsRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionMediaSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemMediaSettings).Return(nil, nil).Once()

	repo := repository.NewMediaSettingsRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemMediaSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestMediaSettingsRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionMediaSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemMediaSettings.ID}
	update := bson.M{"$set": mockItemMediaSettings}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewMediaSettingsRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemMediaSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
