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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestForumPostVoteUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.ForumPostVoteRepository)
	timeout := time.Duration(10)
	usecase := test.NewForumPostVoteUsecase(mockRepo, timeout)

	forumPostVoteID := primitive.NewObjectID().Hex()

	updatedForumPostVote := domain.ForumPostVote{
		ID:           primitive.NewObjectID(), // Existing ID of the record to update
		ForumPostID:  primitive.NewObjectID(),
		CustomerID:   primitive.NewObjectID(),
		IsUp:         false,
		CreatedOnUtc: time.Now().AddDate(0, 0, -7), // Created 7 days ago
	}

	mockRepo.On("FetchByID", mock.Anything, forumPostVoteID).Return(updatedForumPostVote, nil)

	result, err := usecase.FetchByID(context.Background(), forumPostVoteID)

	assert.NoError(t, err)
	assert.Equal(t, updatedForumPostVote, result)
	mockRepo.AssertExpectations(t)
}

func TestForumPostVoteUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.ForumPostVoteRepository)
	timeout := time.Duration(10)
	usecase := test.NewForumPostVoteUsecase(mockRepo, timeout)

	newForumPostVote := &domain.ForumPostVote{
		ForumPostID:  primitive.NewObjectID(),
		CustomerID:   primitive.NewObjectID(),
		IsUp:         true,
		CreatedOnUtc: time.Now(),
	}

	mockRepo.On("Create", mock.Anything, newForumPostVote).Return(nil)

	err := usecase.Create(context.Background(), newForumPostVote)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestForumPostVoteUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.ForumPostVoteRepository)
	timeout := time.Duration(10)
	usecase := test.NewForumPostVoteUsecase(mockRepo, timeout)

	updatedForumPostVote := &domain.ForumPostVote{
		ID:           primitive.NewObjectID(), // Existing ID of the record to update
		ForumPostID:  primitive.NewObjectID(),
		CustomerID:   primitive.NewObjectID(),
		IsUp:         false,
		CreatedOnUtc: time.Now().AddDate(0, 0, -7), // Created 7 days ago
	}

	mockRepo.On("Update", mock.Anything, updatedForumPostVote).Return(nil)

	err := usecase.Update(context.Background(), updatedForumPostVote)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestForumPostVoteUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.ForumPostVoteRepository)
	timeout := time.Duration(10)
	usecase := test.NewForumPostVoteUsecase(mockRepo, timeout)

	forumPostVoteID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, forumPostVoteID).Return(nil)

	err := usecase.Delete(context.Background(), forumPostVoteID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestForumPostVoteUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.ForumPostVoteRepository)
	timeout := time.Duration(10)
	usecase := test.NewForumPostVoteUsecase(mockRepo, timeout)

	fetchedForumPostVotes := []domain.ForumPostVote{
		{
			ID:           primitive.NewObjectID(),
			ForumPostID:  primitive.NewObjectID(),
			CustomerID:   primitive.NewObjectID(),
			IsUp:         true,
			CreatedOnUtc: time.Now().AddDate(0, 0, -10), // Created 10 days ago
		},
		{
			ID:           primitive.NewObjectID(),
			ForumPostID:  primitive.NewObjectID(),
			CustomerID:   primitive.NewObjectID(),
			IsUp:         false,
			CreatedOnUtc: time.Now().AddDate(0, 0, -5), // Created 5 days ago
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedForumPostVotes, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedForumPostVotes, result)
	mockRepo.AssertExpectations(t)
}
