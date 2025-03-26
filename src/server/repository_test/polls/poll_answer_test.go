package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/polls"
	repository "earnforglance/server/repository/polls"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultPollAnswer struct {
	mock.Mock
}

func (m *MockSingleResultPollAnswer) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.PollAnswer); ok {
		*v.(*domain.PollAnswer) = *result
	}
	return args.Error(1)
}

var mockItemPollAnswer = &domain.PollAnswer{
	ID:            primitive.NewObjectID(), // Existing ID of the record to update
	PollID:        primitive.NewObjectID(),
	Name:          "Updated Option A",
	NumberOfVotes: 10,
	DisplayOrder:  2,
}

func TestPollAnswerRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionPollAnswer

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultPollAnswer{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemPollAnswer, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewPollAnswerRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemPollAnswer.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultPollAnswer{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewPollAnswerRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemPollAnswer.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestPollAnswerRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionPollAnswer

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemPollAnswer).Return(nil, nil).Once()

	repo := repository.NewPollAnswerRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemPollAnswer)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestPollAnswerRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionPollAnswer

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemPollAnswer.ID}
	update := bson.M{"$set": mockItemPollAnswer}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewPollAnswerRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemPollAnswer)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
