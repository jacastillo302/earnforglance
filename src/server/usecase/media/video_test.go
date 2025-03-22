package usecase

import (
	"context"
	domain "earnforglance/server/domain/media"
	mocks "earnforglance/server/domain/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestVideoUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.VideoRepository)
	timeout := time.Duration(10)
	usecase := NewVideoUsecase(mockRepo, timeout)

	videoID := primitive.NewObjectID().Hex()

	updatedVideo := domain.Video{
		ID:       primitive.NewObjectID(), // Existing ID of the record to update
		VideoUrl: "https://example.com/updated-video.mp4",
	}

	mockRepo.On("FetchByID", mock.Anything, videoID).Return(updatedVideo, nil)

	result, err := usecase.FetchByID(context.Background(), videoID)

	assert.NoError(t, err)
	assert.Equal(t, updatedVideo, result)
	mockRepo.AssertExpectations(t)
}

func TestVideoUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.VideoRepository)
	timeout := time.Duration(10)
	usecase := NewVideoUsecase(mockRepo, timeout)

	newVideo := &domain.Video{
		VideoUrl: "https://example.com/video.mp4",
	}

	mockRepo.On("Create", mock.Anything, newVideo).Return(nil)

	err := usecase.Create(context.Background(), newVideo)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestVideoUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.VideoRepository)
	timeout := time.Duration(10)
	usecase := NewVideoUsecase(mockRepo, timeout)

	updatedVideo := &domain.Video{
		ID:       primitive.NewObjectID(), // Existing ID of the record to update
		VideoUrl: "https://example.com/updated-video.mp4",
	}

	mockRepo.On("Update", mock.Anything, updatedVideo).Return(nil)

	err := usecase.Update(context.Background(), updatedVideo)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestVideoUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.VideoRepository)
	timeout := time.Duration(10)
	usecase := NewVideoUsecase(mockRepo, timeout)

	videoID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, videoID).Return(nil)

	err := usecase.Delete(context.Background(), videoID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestVideoUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.VideoRepository)
	timeout := time.Duration(10)
	usecase := NewVideoUsecase(mockRepo, timeout)

	fetchedVideos := []domain.Video{
		{
			ID:       primitive.NewObjectID(),
			VideoUrl: "https://example.com/video1.mp4",
		},
		{
			ID:       primitive.NewObjectID(),
			VideoUrl: "https://example.com/video2.mp4",
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedVideos, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedVideos, result)
	mockRepo.AssertExpectations(t)
}
