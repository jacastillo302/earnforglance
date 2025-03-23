package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/messages"
	"earnforglance/server/mongo/mocks"
	repository "earnforglance/server/repository/messages"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultMessagesSettings struct {
	mock.Mock
}

func (m *MockSingleResultMessagesSettings) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.MessagesSettings); ok {
		*v.(*domain.MessagesSettings) = *result
	}
	return args.Error(1)
}

var mockItemMessagesSettings = &domain.MessagesSettings{
	ID:                    primitive.NewObjectID(), // Existing ID of the record to update
	UsePopupNotifications: false,
	UseDefaultEmailAccountForSendStoreOwnerEmails: true,
}

func TestMessagesSettingsRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionMessagesSettings

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultMessagesSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemMessagesSettings, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewMessagesSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemMessagesSettings.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultMessagesSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewMessagesSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemMessagesSettings.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestMessagesSettingsRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionMessagesSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemMessagesSettings).Return(nil, nil).Once()

	repo := repository.NewMessagesSettingsRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemMessagesSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestMessagesSettingsRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionMessagesSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemMessagesSettings.ID}
	update := bson.M{"$set": mockItemMessagesSettings}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewMessagesSettingsRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemMessagesSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
