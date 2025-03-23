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

func TestPictureHashesUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.PictureHashesRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewPictureHashesUsecase(mockRepo, timeout)

	pictureHashID := primitive.NewObjectID().Hex()

	updatedPictureHashes := domain.PictureHashes{
		PictureID: primitive.NewObjectID(), // Existing PictureID of the record to update
		Hash:      []byte("updated_hash_data"),
	}

	mockRepo.On("FetchByID", mock.Anything, pictureHashID).Return(updatedPictureHashes, nil)

	result, err := usecase.FetchByID(context.Background(), pictureHashID)

	assert.NoError(t, err)
	assert.Equal(t, updatedPictureHashes, result)
	mockRepo.AssertExpectations(t)
}

func TestPictureHashesUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.PictureHashesRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewPictureHashesUsecase(mockRepo, timeout)

	newPictureHashes := &domain.PictureHashes{
		PictureID: primitive.NewObjectID(),
		Hash:      []byte("hash_data_1"),
	}

	mockRepo.On("Create", mock.Anything, newPictureHashes).Return(nil)

	err := usecase.Create(context.Background(), newPictureHashes)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestPictureHashesUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.PictureHashesRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewPictureHashesUsecase(mockRepo, timeout)

	updatedPictureHashes := &domain.PictureHashes{
		PictureID: primitive.NewObjectID(), // Existing PictureID of the record to update
		Hash:      []byte("updated_hash_data"),
	}

	mockRepo.On("Update", mock.Anything, updatedPictureHashes).Return(nil)

	err := usecase.Update(context.Background(), updatedPictureHashes)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestPictureHashesUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.PictureHashesRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewPictureHashesUsecase(mockRepo, timeout)

	pictureHashID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, pictureHashID).Return(nil)

	err := usecase.Delete(context.Background(), pictureHashID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestPictureHashesUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.PictureHashesRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewPictureHashesUsecase(mockRepo, timeout)

	fetchedPictureHashes := []domain.PictureHashes{
		{
			PictureID: primitive.NewObjectID(),
			Hash:      []byte("hash_data_1"),
		},
		{
			PictureID: primitive.NewObjectID(),
			Hash:      []byte("hash_data_2"),
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedPictureHashes, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedPictureHashes, result)
	mockRepo.AssertExpectations(t)
}
