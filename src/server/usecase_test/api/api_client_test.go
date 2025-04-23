package usecase_test

import (
	"context"
	domian "earnforglance/server/domain/api"
	mocks "earnforglance/server/domain/mocks"
	test "earnforglance/server/usecase/api"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestApiClientUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.ApiClientRepository)
	timeout := time.Duration(10)
	usecase := test.NewApiClientUsecase(mockRepo, timeout)

	apiID := bson.NewObjectID().Hex()

	updatedApiClient := domian.ApiClient{
		ID:                            bson.NewObjectID(), // Generate a new MongoDB ObjectID
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

	mockRepo.On("FetchByID", mock.Anything, apiID).Return(updatedApiClient, nil)

	result, err := usecase.FetchByID(context.Background(), apiID)

	assert.NoError(t, err)
	assert.Equal(t, updatedApiClient, result)
	mockRepo.AssertExpectations(t)
}

func TestApiClientUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.ApiClientRepository)
	timeout := time.Duration(10)
	usecase := test.NewApiClientUsecase(mockRepo, timeout)

	newApiClient := &domian.ApiClient{
		ID:                            bson.NewObjectID(), // Generate a new MongoDB ObjectID
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

	mockRepo.On("Create", mock.Anything, newApiClient).Return(nil)

	err := usecase.Create(context.Background(), newApiClient)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestApiClientUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.ApiClientRepository)
	timeout := time.Duration(10)
	usecase := test.NewApiClientUsecase(mockRepo, timeout)

	updatedApiClient := &domian.ApiClient{
		ID:                            bson.NewObjectID(), // Generate a new MongoDB ObjectID
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

	mockRepo.On("Update", mock.Anything, updatedApiClient).Return(nil)

	err := usecase.Update(context.Background(), updatedApiClient)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestApiClientUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.ApiClientRepository)
	time := time.Duration(10)
	usecase := test.NewApiClientUsecase(mockRepo, time)

	apiID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, apiID).Return(nil)

	err := usecase.Delete(context.Background(), apiID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestApiClientUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.ApiClientRepository)
	timeout := time.Duration(10)
	usecase := test.NewApiClientUsecase(mockRepo, timeout)

	fetchedApiClients := []domian.ApiClient{
		{
			ID:                            bson.NewObjectID(), // Generate a new MongoDB ObjectID
			Secret:                        "supersecretkey123",
			Enable:                        true,
			DateExpired:                   time.Now().AddDate(1, 0, 0), // Set expiration date to 1 year from now
			Name:                          "Example API Client",
			IdentityUrl:                   "https://identity.example.com",
			CallbackUrl:                   "https://callback.example.com",
			DefaultAccessTokenExpiration:  3600,       // 1 hour in seconds
			DefaultRefreshTokenExpiration: 86400,      // 1 day in seconds
			CreatedDate:                   time.Now(), // Current date and time

		},
		{
			ID:                            bson.NewObjectID(), // Generate a new MongoDB ObjectID
			Secret:                        "supersecretkey123",
			Enable:                        true,
			DateExpired:                   time.Now().AddDate(1, 0, 0), // Set expiration date to 1 year from now
			Name:                          "Example API Client",
			IdentityUrl:                   "https://identity.example.com",
			CallbackUrl:                   "https://callback.example.com",
			DefaultAccessTokenExpiration:  3600,       // 1 hour in seconds
			DefaultRefreshTokenExpiration: 86400,      // 1 day in seconds
			CreatedDate:                   time.Now(), // Current date and time

		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedApiClients, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedApiClients, result)
	mockRepo.AssertExpectations(t)
}
