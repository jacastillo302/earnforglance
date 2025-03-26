package repository_test

import (
	"context"
	"errors"
	"testing"
	"time"

	domain "earnforglance/server/domain/forums"
	repository "earnforglance/server/repository/forums"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultForumPostVote struct {
	mock.Mock
}

func (m *MockSingleResultForumPostVote) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.ForumPostVote); ok {
		*v.(*domain.ForumPostVote) = *result
	}
	return args.Error(1)
}

var mockItemForumPostVote = &domain.ForumPostVote{
	ID:           primitive.NewObjectID(), // Existing ID of the record to update
	ForumPostID:  primitive.NewObjectID(),
	CustomerID:   primitive.NewObjectID(),
	IsUp:         false,
	CreatedOnUtc: time.Now().AddDate(0, 0, -7), // Created 7 days ago
}

func TestForumPostVoteRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionForumPostVote

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultForumPostVote{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemForumPostVote, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewForumPostVoteRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemForumPostVote.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultForumPostVote{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewForumPostVoteRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemForumPostVote.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestForumPostVoteRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionForumPostVote

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemForumPostVote).Return(nil, nil).Once()

	repo := repository.NewForumPostVoteRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemForumPostVote)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestForumPostVoteRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionForumPostVote

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemForumPostVote.ID}
	update := bson.M{"$set": mockItemForumPostVote}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewForumPostVoteRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemForumPostVote)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
