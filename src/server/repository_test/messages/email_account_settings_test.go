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

type MockSingleResultEmailAccountSettings struct {
	mock.Mock
}

func (m *MockSingleResultEmailAccountSettings) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.EmailAccountSettings); ok {
		*v.(*domain.EmailAccountSettings) = *result
	}
	return args.Error(1)
}

var mockItemEmailAccountSettings = &domain.EmailAccountSettings{
	ID:                    primitive.NewObjectID(), // Existing ID of the record to update
	DefaultEmailAccountID: primitive.NewObjectID(),
}

func TestEmailAccountSettingsRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionEmailAccountSettings

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultEmailAccountSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemEmailAccountSettings, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewEmailAccountSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemEmailAccountSettings.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultEmailAccountSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewEmailAccountSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemEmailAccountSettings.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestEmailAccountSettingsRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionEmailAccountSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemEmailAccountSettings).Return(nil, nil).Once()

	repo := repository.NewEmailAccountSettingsRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemEmailAccountSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestEmailAccountSettingsRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionEmailAccountSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemEmailAccountSettings.ID}
	update := bson.M{"$set": mockItemEmailAccountSettings}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewEmailAccountSettingsRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemEmailAccountSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
