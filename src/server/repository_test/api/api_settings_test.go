package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/api"
	repository "earnforglance/server/repository/api"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultApiSettings struct {
	mock.Mock
}

func (m *MockSingleResultApiSettings) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.ApiSettings); ok {
		*v.(*domain.ApiSettings) = *result
	}
	return args.Error(1)
}

var mockItemApiSettings = &domain.ApiSettings{
	ID:                       primitive.NewObjectID(), // Generate a new MongoDB ObjectID
	EnableApi:                true,                    // API is enabled
	AllowRequestsFromSwagger: true,                    // Allow requests from Swagger
	EnableLogging:            true,                    // Enable logging for API requests
	TokenKey:                 "example-token-key",     // Example token key for authentication

}

func TestApiSettingsRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionApiSettings

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultApiSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemApiSettings, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewApiSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemApiSettings.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultApiSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewApiSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemApiSettings.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestApiSettingsRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionApiSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemApiSettings).Return(nil, nil).Once()

	repo := repository.NewApiSettingsRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemApiSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestApiSettingsRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionApiSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemApiSettings.ID}
	update := bson.M{"$set": mockItemApiSettings}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewApiSettingsRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemApiSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
