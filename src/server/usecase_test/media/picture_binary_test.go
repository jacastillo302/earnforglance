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
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestPictureBinaryUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.PictureBinaryRepository)
	timeout := time.Duration(10)
	usecase := test.NewPictureBinaryUsecase(mockRepo, timeout)

	pictureBinaryID := bson.NewObjectID().Hex()

	updatedPictureBinary := domain.PictureBinary{
		ID:         bson.NewObjectID(), // Existing ID of the record to update
		BinaryData: []byte("updated binary image data"),
		PictureID:  bson.NewObjectID(),
	}

	mockRepo.On("FetchByID", mock.Anything, pictureBinaryID).Return(updatedPictureBinary, nil)

	result, err := usecase.FetchByID(context.Background(), pictureBinaryID)

	assert.NoError(t, err)
	assert.Equal(t, updatedPictureBinary, result)
	mockRepo.AssertExpectations(t)
}

func TestPictureBinaryUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.PictureBinaryRepository)
	timeout := time.Duration(10)
	usecase := test.NewPictureBinaryUsecase(mockRepo, timeout)

	newPictureBinary := &domain.PictureBinary{
		BinaryData: []byte("binary image data"),
		PictureID:  bson.NewObjectID(),
	}

	mockRepo.On("Create", mock.Anything, newPictureBinary).Return(nil)

	err := usecase.Create(context.Background(), newPictureBinary)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestPictureBinaryUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.PictureBinaryRepository)
	timeout := time.Duration(10)
	usecase := test.NewPictureBinaryUsecase(mockRepo, timeout)

	updatedPictureBinary := &domain.PictureBinary{
		ID:         bson.NewObjectID(), // Existing ID of the record to update
		BinaryData: []byte("updated binary image data"),
		PictureID:  bson.NewObjectID(),
	}

	mockRepo.On("Update", mock.Anything, updatedPictureBinary).Return(nil)

	err := usecase.Update(context.Background(), updatedPictureBinary)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestPictureBinaryUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.PictureBinaryRepository)
	timeout := time.Duration(10)
	usecase := test.NewPictureBinaryUsecase(mockRepo, timeout)

	pictureBinaryID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, pictureBinaryID).Return(nil)

	err := usecase.Delete(context.Background(), pictureBinaryID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestPictureBinaryUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.PictureBinaryRepository)
	timeout := time.Duration(10)
	usecase := test.NewPictureBinaryUsecase(mockRepo, timeout)

	fetchedPictureBinaries := []domain.PictureBinary{
		{
			ID:         bson.NewObjectID(),
			BinaryData: []byte("binary image data 1"),
			PictureID:  bson.NewObjectID(),
		},
		{
			ID:         bson.NewObjectID(),
			BinaryData: []byte("binary image data 2"),
			PictureID:  bson.NewObjectID(),
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedPictureBinaries, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedPictureBinaries, result)
	mockRepo.AssertExpectations(t)
}
