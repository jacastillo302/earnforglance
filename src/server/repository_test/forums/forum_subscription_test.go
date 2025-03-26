package repository_test

import (
	"context"
	"errors"
	"testing"
	"time"

	domain "earnforglance/server/domain/forums"
	repository "earnforglance/server/repository/forums"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultForumSubscription struct {
	mock.Mock
}

func (m *MockSingleResultForumSubscription) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.ForumSubscription); ok {
		*v.(*domain.ForumSubscription) = *result
	}
	return args.Error(1)
}

var mockItemForumSubscription = &domain.ForumSubscription{
	ID:               primitive.NewObjectID(), // Existing ID of the record to update
	SubscriptionGuid: uuid.New(),
	CustomerID:       2,
	ForumID:          20,
	TopicID:          200,
	CreatedOnUtc:     time.Now().AddDate(0, 0, -7), // Created 7 days ago
}

func TestForumSubscriptionRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionForumSubscription

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultForumSubscription{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemForumSubscription, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewForumSubscriptionRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemForumSubscription.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultForumSubscription{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewForumSubscriptionRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemForumSubscription.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestForumSubscriptionRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionForumSubscription

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemForumSubscription).Return(nil, nil).Once()

	repo := repository.NewForumSubscriptionRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemForumSubscription)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestForumSubscriptionRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionForumSubscription

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemForumSubscription.ID}
	update := bson.M{"$set": mockItemForumSubscription}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewForumSubscriptionRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemForumSubscription)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
