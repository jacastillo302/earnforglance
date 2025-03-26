package repository_test

import (
	"context"
	"errors"
	"testing"
	"time"

	domain "earnforglance/server/domain/polls"
	repository "earnforglance/server/repository/polls"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultPollVotingRecord struct {
	mock.Mock
}

func (m *MockSingleResultPollVotingRecord) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.PollVotingRecord); ok {
		*v.(*domain.PollVotingRecord) = *result
	}
	return args.Error(1)
}

var mockItemPollVotingRecord = &domain.PollVotingRecord{
	ID:           primitive.NewObjectID(), // Existing ID of the record to update
	PollAnswerID: primitive.NewObjectID(),
	CustomerID:   primitive.NewObjectID(),
	CreatedOnUtc: time.Now().AddDate(0, 0, -7), // Created 7 days ago
}

func TestPollVotingRecordRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionPollVotingRecord

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultPollVotingRecord{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemPollVotingRecord, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewPollVotingRecordRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemPollVotingRecord.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultPollVotingRecord{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewPollVotingRecordRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemPollVotingRecord.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestPollVotingRecordRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionPollVotingRecord

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemPollVotingRecord).Return(nil, nil).Once()

	repo := repository.NewPollVotingRecordRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemPollVotingRecord)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestPollVotingRecordRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionPollVotingRecord

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemPollVotingRecord.ID}
	update := bson.M{"$set": mockItemPollVotingRecord}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewPollVotingRecordRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemPollVotingRecord)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
