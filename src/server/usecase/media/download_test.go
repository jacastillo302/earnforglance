package usecase

import (
	"context"
	domain "earnforglance/server/domain/media"
	mocks "earnforglance/server/domain/mocks"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestDownloadUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.DownloadRepository)
	timeout := time.Duration(10)
	usecase := NewDownloadUsecase(mockRepo, timeout)

	mediaID := primitive.NewObjectID().Hex()

	updatedDownload := domain.Download{
		ID:             primitive.NewObjectID(), // Existing ID of the record to update
		DownloadGuid:   uuid.New(),
		UseDownloadUrl: false,
		DownloadUrl:    "",
		DownloadBinary: []byte("binary data"),
		ContentType:    "image/png",
		Filename:       "image.png",
		Extension:      ".png",
		IsNew:          false,
	}

	mockRepo.On("FetchByID", mock.Anything, mediaID).Return(updatedDownload, nil)

	result, err := usecase.FetchByID(context.Background(), mediaID)

	assert.NoError(t, err)
	assert.Equal(t, updatedDownload, result)
	mockRepo.AssertExpectations(t)
}

func TestDownloadUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.DownloadRepository)
	timeout := time.Duration(10)
	usecase := NewDownloadUsecase(mockRepo, timeout)

	newDownload := &domain.Download{
		DownloadGuid:   uuid.New(),
		UseDownloadUrl: true,
		DownloadUrl:    "https://example.com/file.pdf",
		DownloadBinary: nil,
		ContentType:    "application/pdf",
		Filename:       "file.pdf",
		Extension:      ".pdf",
		IsNew:          true,
	}

	mockRepo.On("Create", mock.Anything, newDownload).Return(nil)

	err := usecase.Create(context.Background(), newDownload)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDownloadUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.DownloadRepository)
	timeout := time.Duration(10)
	usecase := NewDownloadUsecase(mockRepo, timeout)

	updatedDownload := &domain.Download{
		ID:             primitive.NewObjectID(), // Existing ID of the record to update
		DownloadGuid:   uuid.New(),
		UseDownloadUrl: false,
		DownloadUrl:    "",
		DownloadBinary: []byte("binary data"),
		ContentType:    "image/png",
		Filename:       "image.png",
		Extension:      ".png",
		IsNew:          false,
	}

	mockRepo.On("Update", mock.Anything, updatedDownload).Return(nil)

	err := usecase.Update(context.Background(), updatedDownload)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDownloadUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.DownloadRepository)
	timeout := time.Duration(10)
	usecase := NewDownloadUsecase(mockRepo, timeout)

	mediaID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, mediaID).Return(nil)

	err := usecase.Delete(context.Background(), mediaID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDownloadUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.DownloadRepository)
	timeout := time.Duration(10)
	usecase := NewDownloadUsecase(mockRepo, timeout)

	fetchedDownloads := []domain.Download{
		{
			ID:             primitive.NewObjectID(),
			DownloadGuid:   uuid.New(),
			UseDownloadUrl: true,
			DownloadUrl:    "https://example.com/file1.pdf",
			DownloadBinary: nil,
			ContentType:    "application/pdf",
			Filename:       "file1.pdf",
			Extension:      ".pdf",
			IsNew:          true,
		},
		{
			ID:             primitive.NewObjectID(),
			DownloadGuid:   uuid.New(),
			UseDownloadUrl: false,
			DownloadUrl:    "",
			DownloadBinary: []byte("binary data"),
			ContentType:    "image/jpeg",
			Filename:       "image.jpg",
			Extension:      ".jpg",
			IsNew:          false,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedDownloads, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedDownloads, result)
	mockRepo.AssertExpectations(t)
}
