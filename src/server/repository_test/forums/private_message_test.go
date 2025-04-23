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

type MockSingleResultPrivateMessage struct {
	mock.Mock
}

func (m *MockSingleResultPrivateMessage) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.PrivateMessage); ok {
		*v.(*domain.PrivateMessage) = *result
	}
	return args.Error(1)
}

var mockItemPrivateMessage = &domain.PrivateMessage{
	ID:                   bson.NewObjectID(), // Existing ID of the record to update
	StoreID:              bson.NewObjectID(),
	FromCustomerID:       bson.NewObjectID(),
	ToCustomerID:         bson.NewObjectID(),
	Subject:              "Updated Subject",
	Text:                 "This is an updated private message.",
	IsRead:               true,
	IsDeletedByAuthor:    false,
	IsDeletedByRecipient: true,
	CreatedOnUtc:         time.Now().AddDate(0, 0, -7), // Created 7 days ago
}

func TestPrivateMessageRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionPrivateMessage

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultPrivateMessage{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemPrivateMessage, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewPrivateMessageRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemPrivateMessage.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultPrivateMessage{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewPrivateMessageRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemPrivateMessage.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestPrivateMessageRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionPrivateMessage

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemPrivateMessage).Return(nil, nil).Once()

	repo := repository.NewPrivateMessageRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemPrivateMessage)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestPrivateMessageRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionPrivateMessage

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemPrivateMessage.ID}
	update := bson.M{"$set": mockItemPrivateMessage}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewPrivateMessageRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemPrivateMessage)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
