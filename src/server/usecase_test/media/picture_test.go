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

func TestPictureUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.PictureRepository)
	timeout := time.Duration(10)
	usecase := test.NewPictureUsecase(mockRepo, timeout)

	pictureID := primitive.NewObjectID().Hex()

	updatedPicture := domain.Picture{
		ID:             primitive.NewObjectID(), // Existing ID of the record to update
		MimeType:       "image/png",
		SeoFilename:    "updated-image",
		AltAttribute:   "Updated Image",
		TitleAttribute: "Updated Title",
		IsNew:          false,
		VirtualPath:    "/images/updated-image.png",
	}

	mockRepo.On("FetchByID", mock.Anything, pictureID).Return(updatedPicture, nil)

	result, err := usecase.FetchByID(context.Background(), pictureID)

	assert.NoError(t, err)
	assert.Equal(t, updatedPicture, result)
	mockRepo.AssertExpectations(t)
}

func TestPictureUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.PictureRepository)
	timeout := time.Duration(10)
	usecase := test.NewPictureUsecase(mockRepo, timeout)

	newPicture := &domain.Picture{
		MimeType:       "image/jpeg",
		SeoFilename:    "example-image",
		AltAttribute:   "Example Image",
		TitleAttribute: "Example Title",
		IsNew:          true,
		VirtualPath:    "/images/example-image.jpg",
	}

	mockRepo.On("Create", mock.Anything, newPicture).Return(nil)

	err := usecase.Create(context.Background(), newPicture)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestPictureUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.PictureRepository)
	timeout := time.Duration(10)
	usecase := test.NewPictureUsecase(mockRepo, timeout)

	updatedPicture := &domain.Picture{
		ID:             primitive.NewObjectID(), // Existing ID of the record to update
		MimeType:       "image/png",
		SeoFilename:    "updated-image",
		AltAttribute:   "Updated Image",
		TitleAttribute: "Updated Title",
		IsNew:          false,
		VirtualPath:    "/images/updated-image.png",
	}

	mockRepo.On("Update", mock.Anything, updatedPicture).Return(nil)

	err := usecase.Update(context.Background(), updatedPicture)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestPictureUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.PictureRepository)
	timeout := time.Duration(10)
	usecase := test.NewPictureUsecase(mockRepo, timeout)

	pictureID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, pictureID).Return(nil)

	err := usecase.Delete(context.Background(), pictureID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestPictureUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.PictureRepository)
	timeout := time.Duration(10)
	usecase := test.NewPictureUsecase(mockRepo, timeout)

	fetchedPictures := []domain.Picture{
		{
			ID:             primitive.NewObjectID(),
			MimeType:       "image/jpeg",
			SeoFilename:    "example-image",
			AltAttribute:   "Example Image",
			TitleAttribute: "Example Title",
			IsNew:          true,
			VirtualPath:    "/images/example-image.jpg",
		},
		{
			ID:             primitive.NewObjectID(),
			MimeType:       "image/png",
			SeoFilename:    "updated-image",
			AltAttribute:   "Updated Image",
			TitleAttribute: "Updated Title",
			IsNew:          false,
			VirtualPath:    "/images/updated-image.png",
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedPictures, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedPictures, result)
	mockRepo.AssertExpectations(t)
}
