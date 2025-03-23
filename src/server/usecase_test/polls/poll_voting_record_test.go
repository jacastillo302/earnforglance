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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestPollVotingRecordUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.PollVotingRecordRepository)
	timeout := time.Duration(10)
	usecase := test.NewPollVotingRecordUsecase(mockRepo, timeout)

	pollVotingRecordID := primitive.NewObjectID().Hex()

	updatedPollVotingRecord := domain.PollVotingRecord{
		ID:           primitive.NewObjectID(), // Existing ID of the record to update
		PollAnswerID: primitive.NewObjectID(),
		CustomerID:   primitive.NewObjectID(),
		CreatedOnUtc: time.Now().AddDate(0, 0, -7), // Created 7 days ago
	}

	mockRepo.On("FetchByID", mock.Anything, pollVotingRecordID).Return(updatedPollVotingRecord, nil)

	result, err := usecase.FetchByID(context.Background(), pollVotingRecordID)

	assert.NoError(t, err)
	assert.Equal(t, updatedPollVotingRecord, result)
	mockRepo.AssertExpectations(t)
}

func TestPollVotingRecordUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.PollVotingRecordRepository)
	timeout := time.Duration(10)
	usecase := test.NewPollVotingRecordUsecase(mockRepo, timeout)

	newPollVotingRecord := &domain.PollVotingRecord{
		PollAnswerID: primitive.NewObjectID(),
		CustomerID:   primitive.NewObjectID(),
		CreatedOnUtc: time.Now(),
	}

	mockRepo.On("Create", mock.Anything, newPollVotingRecord).Return(nil)

	err := usecase.Create(context.Background(), newPollVotingRecord)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestPollVotingRecordUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.PollVotingRecordRepository)
	timeout := time.Duration(10)
	usecase := test.NewPollVotingRecordUsecase(mockRepo, timeout)

	updatedPollVotingRecord := &domain.PollVotingRecord{
		ID:           primitive.NewObjectID(), // Existing ID of the record to update
		PollAnswerID: primitive.NewObjectID(),
		CustomerID:   primitive.NewObjectID(),
		CreatedOnUtc: time.Now().AddDate(0, 0, -7), // Created 7 days ago
	}

	mockRepo.On("Update", mock.Anything, updatedPollVotingRecord).Return(nil)

	err := usecase.Update(context.Background(), updatedPollVotingRecord)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestPollVotingRecordUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.PollVotingRecordRepository)
	timeout := time.Duration(10)
	usecase := test.NewPollVotingRecordUsecase(mockRepo, timeout)

	pollVotingRecordID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, pollVotingRecordID).Return(nil)

	err := usecase.Delete(context.Background(), pollVotingRecordID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestPollVotingRecordUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.PollVotingRecordRepository)
	timeout := time.Duration(10)
	usecase := test.NewPollVotingRecordUsecase(mockRepo, timeout)

	fetchedPollVotingRecords := []domain.PollVotingRecord{
		{
			ID:           primitive.NewObjectID(),
			PollAnswerID: primitive.NewObjectID(),
			CustomerID:   primitive.NewObjectID(),
			CreatedOnUtc: time.Now().AddDate(0, 0, -10), // Created 10 days ago
		},
		{
			ID:           primitive.NewObjectID(),
			PollAnswerID: primitive.NewObjectID(),
			CustomerID:   primitive.NewObjectID(),
			CreatedOnUtc: time.Now().AddDate(0, 0, -5), // Created 5 days ago
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedPollVotingRecords, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedPollVotingRecords, result)
	mockRepo.AssertExpectations(t)
}
