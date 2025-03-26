package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/messages"
	repository "earnforglance/server/repository/messages"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultMessageTemplatesSettings struct {
	mock.Mock
}

func (m *MockSingleResultMessageTemplatesSettings) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.MessageTemplatesSettings); ok {
		*v.(*domain.MessageTemplatesSettings) = *result
	}
	return args.Error(1)
}

var mockItemMessageTemplatesSettings = &domain.MessageTemplatesSettings{
	ID:                       primitive.NewObjectID(), // Existing ID of the record to update
	CaseInvariantReplacement: false,
	Color1:                   "#FFFFFF",
	Color2:                   "#000000",
	Color3:                   "#CCCCCC",
}

func TestMessageTemplatesSettingsRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionMessageTemplatesSettings

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultMessageTemplatesSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemMessageTemplatesSettings, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewMessageTemplatesSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemMessageTemplatesSettings.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultMessageTemplatesSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewMessageTemplatesSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemMessageTemplatesSettings.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestMessageTemplatesSettingsRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionMessageTemplatesSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemMessageTemplatesSettings).Return(nil, nil).Once()

	repo := repository.NewMessageTemplatesSettingsRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemMessageTemplatesSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestMessageTemplatesSettingsRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionMessageTemplatesSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemMessageTemplatesSettings.ID}
	update := bson.M{"$set": mockItemMessageTemplatesSettings}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewMessageTemplatesSettingsRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemMessageTemplatesSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
