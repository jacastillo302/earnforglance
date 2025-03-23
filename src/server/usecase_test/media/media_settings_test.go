package usecase_test

import (
	"context"
	domain "earnforglance/server/domain/media"
	mocks "earnforglance/server/domain/mocks"
	test "earnforglance/server/usecase/media"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestMediaSettingsUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.MediaSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewMediaSettingsUsecase(mockRepo, timeout)

	mediaID := primitive.NewObjectID().Hex()
	updatedMediaSettings := domain.MediaSettings{
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

	mockRepo.On("FetchByID", mock.Anything, mediaID).Return(updatedMediaSettings, nil)

	result, err := usecase.FetchByID(context.Background(), mediaID)

	assert.NoError(t, err)
	assert.Equal(t, updatedMediaSettings, result)
	mockRepo.AssertExpectations(t)
}

func TestMediaSettingsUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.MediaSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewMediaSettingsUsecase(mockRepo, timeout)

	newMediaSettings := &domain.MediaSettings{
		AvatarPictureSize:                           100,
		ProductThumbPictureSize:                     200,
		ProductDetailsPictureSize:                   400,
		ProductThumbPictureSizeOnProductDetailsPage: 150,
		AssociatedProductPictureSize:                120,
		CategoryThumbPictureSize:                    180,
		ManufacturerThumbPictureSize:                160,
		VendorThumbPictureSize:                      140,
		CartThumbPictureSize:                        80,
		OrderThumbPictureSize:                       90,
		MiniCartThumbPictureSize:                    70,
		AutoCompleteSearchThumbPictureSize:          60,
		ImageSquarePictureSize:                      50,
		DefaultPictureZoomEnabled:                   true,
		AllowSVGUploads:                             true,
		MaximumImageSize:                            5000,
		DefaultImageQuality:                         90,
		MultipleThumbDirectories:                    false,
		ImportProductImagesUsingHash:                true,
		AzureCacheControlHeader:                     "public, max-age=31536000",
		UseAbsoluteImagePath:                        false,
		VideoIframeAllow:                            "fullscreen",
		VideoIframeWidth:                            640,
		VideoIframeHeight:                           360,
		ProductDefaultImageID:                       primitive.NewObjectID(),
		AutoOrientImage:                             true,
	}

	mockRepo.On("Create", mock.Anything, newMediaSettings).Return(nil)

	err := usecase.Create(context.Background(), newMediaSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestMediaSettingsUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.MediaSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewMediaSettingsUsecase(mockRepo, timeout)

	updatedMediaSettings := &domain.MediaSettings{
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

	mockRepo.On("Update", mock.Anything, updatedMediaSettings).Return(nil)

	err := usecase.Update(context.Background(), updatedMediaSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestMediaSettingsUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.MediaSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewMediaSettingsUsecase(mockRepo, timeout)

	mediaID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, mediaID).Return(nil)

	err := usecase.Delete(context.Background(), mediaID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestMediaSettingsUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.MediaSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewMediaSettingsUsecase(mockRepo, timeout)

	fetchedMediaSettings := []domain.MediaSettings{
		{
			ID:                        primitive.NewObjectID(),
			AvatarPictureSize:         100,
			ProductThumbPictureSize:   200,
			ProductDetailsPictureSize: 400,
			ProductThumbPictureSizeOnProductDetailsPage: 150,
			AssociatedProductPictureSize:                120,
			CategoryThumbPictureSize:                    180,
			ManufacturerThumbPictureSize:                160,
			VendorThumbPictureSize:                      140,
			CartThumbPictureSize:                        80,
			OrderThumbPictureSize:                       90,
			MiniCartThumbPictureSize:                    70,
			AutoCompleteSearchThumbPictureSize:          60,
			ImageSquarePictureSize:                      50,
			DefaultPictureZoomEnabled:                   true,
			AllowSVGUploads:                             true,
			MaximumImageSize:                            5000,
			DefaultImageQuality:                         90,
			MultipleThumbDirectories:                    false,
			ImportProductImagesUsingHash:                true,
			AzureCacheControlHeader:                     "public, max-age=31536000",
			UseAbsoluteImagePath:                        false,
			VideoIframeAllow:                            "fullscreen",
			VideoIframeWidth:                            640,
			VideoIframeHeight:                           360,
			ProductDefaultImageID:                       primitive.NewObjectID(),
			AutoOrientImage:                             true,
		},
		{
			ID:                        primitive.NewObjectID(),
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
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedMediaSettings, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedMediaSettings, result)
	mockRepo.AssertExpectations(t)
}
