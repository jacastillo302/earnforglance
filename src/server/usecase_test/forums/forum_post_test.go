package usecase_test

import (
	"context"
	domain "earnforglance/server/domain/forums"
	mocks "earnforglance/server/domain/mocks"
	test "earnforglance/server/usecase/forums"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestForumPostUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.ForumPostRepository)
	timeout := time.Duration(10)
	usecase := test.NewForumPostUsecase(mockRepo, timeout)

	forumPostID := bson.NewObjectID().Hex()

	updatedForumPost := domain.ForumPost{
		ID:           bson.NewObjectID(), // Existing ID of the record to update
		TopicID:      bson.NewObjectID(),
		CustomerID:   bson.NewObjectID(),
		Text:         "This is an updated forum post.",
		IPAddress:    "192.168.1.2",
		CreatedOnUtc: time.Now().AddDate(0, 0, -7), // Created 7 days ago
		UpdatedOnUtc: time.Now(),
		VoteCount:    5,
	}

	mockRepo.On("FetchByID", mock.Anything, forumPostID).Return(updatedForumPost, nil)

	result, err := usecase.FetchByID(context.Background(), forumPostID)

	assert.NoError(t, err)
	assert.Equal(t, updatedForumPost, result)
	mockRepo.AssertExpectations(t)
}

func TestForumPostUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.ForumPostRepository)
	timeout := time.Duration(10)
	usecase := test.NewForumPostUsecase(mockRepo, timeout)

	newForumPost := &domain.ForumPost{
		TopicID:      bson.NewObjectID(),
		CustomerID:   bson.NewObjectID(),
		Text:         "This is a new forum post.",
		IPAddress:    "192.168.1.1",
		CreatedOnUtc: time.Now(),
		UpdatedOnUtc: time.Now(),
		VoteCount:    0,
	}

	mockRepo.On("Create", mock.Anything, newForumPost).Return(nil)

	err := usecase.Create(context.Background(), newForumPost)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestForumPostUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.ForumPostRepository)
	timeout := time.Duration(10)
	usecase := test.NewForumPostUsecase(mockRepo, timeout)

	updatedForumPost := &domain.ForumPost{
		ID:           bson.NewObjectID(), // Existing ID of the record to update
		TopicID:      bson.NewObjectID(),
		CustomerID:   bson.NewObjectID(),
		Text:         "This is an updated forum post.",
		IPAddress:    "192.168.1.2",
		CreatedOnUtc: time.Now().AddDate(0, 0, -7), // Created 7 days ago
		UpdatedOnUtc: time.Now(),
		VoteCount:    5,
	}

	mockRepo.On("Update", mock.Anything, updatedForumPost).Return(nil)

	err := usecase.Update(context.Background(), updatedForumPost)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestForumPostUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.ForumPostRepository)
	timeout := time.Duration(10)
	usecase := test.NewForumPostUsecase(mockRepo, timeout)

	forumPostID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, forumPostID).Return(nil)

	err := usecase.Delete(context.Background(), forumPostID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestForumPostUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.ForumPostRepository)
	timeout := time.Duration(10)
	usecase := test.NewForumPostUsecase(mockRepo, timeout)

	fetchedForumPosts := []domain.ForumPost{
		{
			ID:           bson.NewObjectID(),
			TopicID:      bson.NewObjectID(),
			CustomerID:   bson.NewObjectID(),
			Text:         "This is the first forum post.",
			IPAddress:    "192.168.1.1",
			CreatedOnUtc: time.Now().AddDate(0, 0, -10), // Created 10 days ago
			UpdatedOnUtc: time.Now().AddDate(0, 0, -5),  // Updated 5 days ago
			VoteCount:    3,
		},
		{
			ID:           bson.NewObjectID(),
			TopicID:      bson.NewObjectID(),
			CustomerID:   bson.NewObjectID(),
			Text:         "This is the second forum post.",
			IPAddress:    "192.168.1.2",
			CreatedOnUtc: time.Now().AddDate(0, 0, -7), // Created 7 days ago
			UpdatedOnUtc: time.Now(),
			VoteCount:    10,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedForumPosts, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedForumPosts, result)
	mockRepo.AssertExpectations(t)
}
