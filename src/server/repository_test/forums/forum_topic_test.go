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

	"go.mongodb.org/mongo-driver/v2/bson"
)

type MockSingleResultForumTopic struct {
	mock.Mock
}

func (m *MockSingleResultForumTopic) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.ForumTopic); ok {
		*v.(*domain.ForumTopic) = *result
	}
	return args.Error(1)
}

var mockItemForumTopic = &domain.ForumTopic{
	ID:                 bson.NewObjectID(), // Existing ID of the record to update
	ForumID:            bson.NewObjectID(),
	CustomerID:         bson.NewObjectID(),
	TopicTypeID:        2,
	Subject:            "Updated Forum Topic",
	NumPosts:           10,
	Views:              100,
	LastPostID:         bson.NewObjectID(),
	LastPostCustomerID: bson.NewObjectID(),
	LastPostTime:       new(time.Time),
	CreatedOnUtc:       time.Now().AddDate(0, 0, -7), // Created 7 days ago
	UpdatedOnUtc:       time.Now(),
}

func TestForumTopicRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionForumTopic

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultForumTopic{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemForumTopic, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewForumTopicRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemForumTopic.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultForumTopic{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewForumTopicRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemForumTopic.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestForumTopicRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionForumTopic

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemForumTopic).Return(nil, nil).Once()

	repo := repository.NewForumTopicRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemForumTopic)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestForumTopicRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionForumTopic

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemForumTopic.ID}
	update := bson.M{"$set": mockItemForumTopic}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewForumTopicRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemForumTopic)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
