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

type MockSingleResultForumPost struct {
	mock.Mock
}

func (m *MockSingleResultForumPost) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.ForumPost); ok {
		*v.(*domain.ForumPost) = *result
	}
	return args.Error(1)
}

var mockItemForumPost = &domain.ForumPost{
	ID:           primitive.NewObjectID(), // Existing ID of the record to update
	TopicID:      primitive.NewObjectID(),
	CustomerID:   primitive.NewObjectID(),
	Text:         "This is an updated forum post.",
	IPAddress:    "192.168.1.2",
	CreatedOnUtc: time.Now().AddDate(0, 0, -7), // Created 7 days ago
	UpdatedOnUtc: time.Now(),
	VoteCount:    5,
}

func TestForumPostRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionForumPost

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultForumPost{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemForumPost, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewForumPostRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemForumPost.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultForumPost{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewForumPostRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemForumPost.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestForumPostRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionForumPost

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemForumPost).Return(nil, nil).Once()

	repo := repository.NewForumPostRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemForumPost)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestForumPostRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionForumPost

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemForumPost.ID}
	update := bson.M{"$set": mockItemForumPost}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewForumPostRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemForumPost)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
