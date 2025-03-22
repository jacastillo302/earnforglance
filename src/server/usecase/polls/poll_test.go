package usecase

import (
	"context"
	mocks "earnforglance/server/domain/mocks"
	domain "earnforglance/server/domain/polls"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestPollUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.PollRepository)
	timeout := time.Duration(10)
	usecase := NewPollUsecase(mockRepo, timeout)

	pollID := primitive.NewObjectID().Hex()

	updatedPoll := domain.Poll{
		ID:                primitive.NewObjectID(), // Existing ID of the record to update
		LanguageID:        primitive.NewObjectID(),
		Name:              "Updated Poll Name",
		SystemKeyword:     "updated_poll_keyword",
		Published:         false,
		ShowOnHomepage:    false,
		AllowGuestsToVote: false,
		DisplayOrder:      2,
		LimitedToStores:   true,
		StartDateUtc:      new(time.Time),
		EndDateUtc:        new(time.Time),
	}
	*updatedPoll.StartDateUtc = time.Now().AddDate(0, 0, -7) // Started 7 days ago
	*updatedPoll.EndDateUtc = time.Now().AddDate(0, 0, 7)    // Ends in 7 days

	mockRepo.On("FetchByID", mock.Anything, pollID).Return(updatedPoll, nil)

	result, err := usecase.FetchByID(context.Background(), pollID)

	assert.NoError(t, err)
	assert.Equal(t, updatedPoll, result)
	mockRepo.AssertExpectations(t)
}

func TestPollUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.PollRepository)
	timeout := time.Duration(10)
	usecase := NewPollUsecase(mockRepo, timeout)

	newPoll := &domain.Poll{
		LanguageID:        primitive.NewObjectID(),
		Name:              "Favorite Programming Language",
		SystemKeyword:     "favorite_programming_language",
		Published:         true,
		ShowOnHomepage:    true,
		AllowGuestsToVote: true,
		DisplayOrder:      1,
		LimitedToStores:   false,
		StartDateUtc:      nil,
		EndDateUtc:        nil,
	}

	mockRepo.On("Create", mock.Anything, newPoll).Return(nil)

	err := usecase.Create(context.Background(), newPoll)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestPollUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.PollRepository)
	timeout := time.Duration(10)
	usecase := NewPollUsecase(mockRepo, timeout)

	updatedPoll := &domain.Poll{
		ID:                primitive.NewObjectID(), // Existing ID of the record to update
		LanguageID:        primitive.NewObjectID(),
		Name:              "Updated Poll Name",
		SystemKeyword:     "updated_poll_keyword",
		Published:         false,
		ShowOnHomepage:    false,
		AllowGuestsToVote: false,
		DisplayOrder:      2,
		LimitedToStores:   true,
		StartDateUtc:      new(time.Time),
		EndDateUtc:        new(time.Time),
	}
	*updatedPoll.StartDateUtc = time.Now().AddDate(0, 0, -7) // Started 7 days ago
	*updatedPoll.EndDateUtc = time.Now().AddDate(0, 0, 7)    // Ends in 7 days

	mockRepo.On("Update", mock.Anything, updatedPoll).Return(nil)

	err := usecase.Update(context.Background(), updatedPoll)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestPollUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.PollRepository)
	timeout := time.Duration(10)
	usecase := NewPollUsecase(mockRepo, timeout)

	pollID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, pollID).Return(nil)

	err := usecase.Delete(context.Background(), pollID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestPollUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.PollRepository)
	timeout := time.Duration(10)
	usecase := NewPollUsecase(mockRepo, timeout)

	fetchedPolls := []domain.Poll{
		{
			ID:                primitive.NewObjectID(),
			LanguageID:        primitive.NewObjectID(),
			Name:              "Favorite Programming Language",
			SystemKeyword:     "favorite_programming_language",
			Published:         true,
			ShowOnHomepage:    true,
			AllowGuestsToVote: true,
			DisplayOrder:      1,
			LimitedToStores:   false,
			StartDateUtc:      nil,
			EndDateUtc:        nil,
		},
		{
			ID:                primitive.NewObjectID(),
			LanguageID:        primitive.NewObjectID(),
			Name:              "Updated Poll Name",
			SystemKeyword:     "updated_poll_keyword",
			Published:         false,
			ShowOnHomepage:    false,
			AllowGuestsToVote: false,
			DisplayOrder:      2,
			LimitedToStores:   true,
			StartDateUtc:      new(time.Time),
			EndDateUtc:        new(time.Time),
		},
	}
	*fetchedPolls[1].StartDateUtc = time.Now().AddDate(0, 0, -7) // Started 7 days ago
	*fetchedPolls[1].EndDateUtc = time.Now().AddDate(0, 0, 7)    // Ends in 7 days

	mockRepo.On("Fetch", mock.Anything).Return(fetchedPolls, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedPolls, result)
	mockRepo.AssertExpectations(t)
}
