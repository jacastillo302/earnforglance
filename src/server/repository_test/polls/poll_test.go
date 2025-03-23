package repository_test

import (
	"context"
	"errors"
	"testing"
	"time"

	domain "earnforglance/server/domain/polls"
	"earnforglance/server/mongo/mocks"
	repository "earnforglance/server/repository/polls"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultPoll struct {
	mock.Mock
}

func (m *MockSingleResultPoll) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.Poll); ok {
		*v.(*domain.Poll) = *result
	}
	return args.Error(1)
}

var mockItemPoll = &domain.Poll{
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

func TestPollRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionPoll

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultPoll{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemPoll, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewPollRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemPoll.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultPoll{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewPollRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemPoll.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestPollRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionPoll

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemPoll).Return(nil, nil).Once()

	repo := repository.NewPollRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemPoll)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestPollRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionPoll

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemPoll.ID}
	update := bson.M{"$set": mockItemPoll}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewPollRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemPoll)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
