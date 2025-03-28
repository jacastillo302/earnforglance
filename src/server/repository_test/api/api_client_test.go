package repository_test

import (
	"context"
	domain "earnforglance/server/domain/api"
	repository "earnforglance/server/repository/api"
	"earnforglance/server/service/data/mongo/mocks"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultApiClient struct {
	mock.Mock
}

func (m *MockSingleResultApiClient) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.ApiClient); ok {
		*v.(*domain.ApiClient) = *result
	}
	return args.Error(1)
}

var mockItemApiClient = &domain.ApiClient{
	ID:                            primitive.NewObjectID(), // Generate a new MongoDB ObjectID
	Secret:                        "supersecretkey123",
	Enable:                        true,
	DateExpired:                   time.Now().AddDate(1, 0, 0), // Set expiration date to 1 year from now
	Name:                          "Example API Client",
	IdentityUrl:                   "https://identity.example.com",
	CallbackUrl:                   "https://callback.example.com",
	DefaultAccessTokenExpiration:  3600,       // 1 hour in seconds
	DefaultRefreshTokenExpiration: 86400,      // 1 day in seconds
	CreatedDate:                   time.Now(), // Current date and time

}

func TestApiClientRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionApiClient

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultApiClient{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemApiClient, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewApiClientRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemApiClient.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultApiClient{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewApiClientRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemApiClient.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestApiClientRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionApiClient

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemApiClient).Return(nil, nil).Once()

	repo := repository.NewApiClientRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemApiClient)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestApiClientRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionApiClient

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemApiClient.ID}
	update := bson.M{"$set": mockItemApiClient}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewApiClientRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemApiClient)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
