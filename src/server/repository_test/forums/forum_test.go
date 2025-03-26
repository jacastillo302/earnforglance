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

type MockSingleResultForum struct {
	mock.Mock
}

func (m *MockSingleResultForum) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.Forum); ok {
		*v.(*domain.Forum) = *result
	}
	return args.Error(1)
}

var mockItemForum = &domain.Forum{
	ID:                 primitive.NewObjectID(), // Existing ID of the record to update
	ForumGroupID:       primitive.NewObjectID(),
	Name:               "Announcements",
	Description:        "Official announcements and updates.",
	NumTopics:          10,
	NumPosts:           50,
	LastTopicID:        primitive.NewObjectID(),
	LastPostID:         primitive.NewObjectID(),
	LastPostCustomerID: primitive.NewObjectID(),
	LastPostTime:       new(time.Time),
	DisplayOrder:       2,
	CreatedOnUtc:       time.Now().AddDate(0, 0, -30), // Created 30 days ago
	UpdatedOnUtc:       time.Now(),
}

func TestForumRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionForum

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultForum{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemForum, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewForumRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemForum.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultForum{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewForumRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemForum.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestForumRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionForum

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemForum).Return(nil, nil).Once()

	repo := repository.NewForumRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemForum)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestForumRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionForum

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemForum.ID}
	update := bson.M{"$set": mockItemForum}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewForumRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemForum)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
