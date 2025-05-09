package usecase_test

import (
	"context"
	mocks "earnforglance/server/domain/mocks"
	domain "earnforglance/server/domain/polls"
	test "earnforglance/server/usecase/polls"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestPollAnswerUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.PollAnswerRepository)
	timeout := time.Duration(10)
	usecase := test.NewPollAnswerUsecase(mockRepo, timeout)

	pollAnswerID := bson.NewObjectID().Hex()

	updatedPollAnswer := domain.PollAnswer{
		ID:            bson.NewObjectID(), // Existing ID of the record to update
		PollID:        bson.NewObjectID(),
		Name:          "Updated Option A",
		NumberOfVotes: 10,
		DisplayOrder:  2,
	}

	mockRepo.On("FetchByID", mock.Anything, pollAnswerID).Return(updatedPollAnswer, nil)

	result, err := usecase.FetchByID(context.Background(), pollAnswerID)

	assert.NoError(t, err)
	assert.Equal(t, updatedPollAnswer, result)
	mockRepo.AssertExpectations(t)
}

func TestPollAnswerUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.PollAnswerRepository)
	timeout := time.Duration(10)
	usecase := test.NewPollAnswerUsecase(mockRepo, timeout)

	newPollAnswer := &domain.PollAnswer{
		PollID:        bson.NewObjectID(),
		Name:          "Option A",
		NumberOfVotes: 0,
		DisplayOrder:  1,
	}

	mockRepo.On("Create", mock.Anything, newPollAnswer).Return(nil)

	err := usecase.Create(context.Background(), newPollAnswer)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestPollAnswerUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.PollAnswerRepository)
	timeout := time.Duration(10)
	usecase := test.NewPollAnswerUsecase(mockRepo, timeout)

	updatedPollAnswer := &domain.PollAnswer{
		ID:            bson.NewObjectID(), // Existing ID of the record to update
		PollID:        bson.NewObjectID(),
		Name:          "Updated Option A",
		NumberOfVotes: 10,
		DisplayOrder:  2,
	}

	mockRepo.On("Update", mock.Anything, updatedPollAnswer).Return(nil)

	err := usecase.Update(context.Background(), updatedPollAnswer)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestPollAnswerUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.PollAnswerRepository)
	timeout := time.Duration(10)
	usecase := test.NewPollAnswerUsecase(mockRepo, timeout)

	pollAnswerID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, pollAnswerID).Return(nil)

	err := usecase.Delete(context.Background(), pollAnswerID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestPollAnswerUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.PollAnswerRepository)
	timeout := time.Duration(10)
	usecase := test.NewPollAnswerUsecase(mockRepo, timeout)

	fetchedPollAnswers := []domain.PollAnswer{
		{
			ID:            bson.NewObjectID(),
			PollID:        bson.NewObjectID(),
			Name:          "Option A",
			NumberOfVotes: 5,
			DisplayOrder:  1,
		},
		{
			ID:            bson.NewObjectID(),
			PollID:        bson.NewObjectID(),
			Name:          "Option B",
			NumberOfVotes: 8,
			DisplayOrder:  2,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedPollAnswers, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedPollAnswers, result)
	mockRepo.AssertExpectations(t)
}
