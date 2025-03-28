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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestApiSettingsUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.ApiSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewApiSettingsUsecase(mockRepo, timeout)

	apiID := primitive.NewObjectID().Hex()

	updatedApiSettings := domian.ApiSettings{
		ID:                       primitive.NewObjectID(), // Generate a new MongoDB ObjectID
		EnableApi:                true,                    // API is enabled
		AllowRequestsFromSwagger: true,                    // Allow requests from Swagger
		EnableLogging:            true,                    // Enable logging for API requests
		TokenKey:                 "example-token-key",     // Example token key for authentication
	}

	mockRepo.On("FetchByID", mock.Anything, apiID).Return(updatedApiSettings, nil)

	result, err := usecase.FetchByID(context.Background(), apiID)

	assert.NoError(t, err)
	assert.Equal(t, updatedApiSettings, result)
	mockRepo.AssertExpectations(t)
}

func TestApiSettingsUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.ApiSettingsRepository)
	time := time.Duration(10)
	usecase := test.NewApiSettingsUsecase(mockRepo, time)

	newApiSettings := &domian.ApiSettings{
		ID:                       primitive.NewObjectID(), // Generate a new MongoDB ObjectID
		EnableApi:                true,                    // API is enabled
		AllowRequestsFromSwagger: true,                    // Allow requests from Swagger
		EnableLogging:            true,                    // Enable logging for API requests
		TokenKey:                 "example-token-key",     // Example token key for authentication
	}

	mockRepo.On("Create", mock.Anything, newApiSettings).Return(nil)

	err := usecase.Create(context.Background(), newApiSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestApiSettingsUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.ApiSettingsRepository)
	time := time.Duration(10)
	usecase := test.NewApiSettingsUsecase(mockRepo, time)

	updatedApiSettings := &domian.ApiSettings{
		ID:                       primitive.NewObjectID(), // Generate a new MongoDB ObjectID
		EnableApi:                true,                    // API is enabled
		AllowRequestsFromSwagger: true,                    // Allow requests from Swagger
		EnableLogging:            true,                    // Enable logging for API requests
		TokenKey:                 "example-token-key",     // Example token key for authentication
	}

	mockRepo.On("Update", mock.Anything, updatedApiSettings).Return(nil)

	err := usecase.Update(context.Background(), updatedApiSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestApiSettingsUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.ApiSettingsRepository)
	time := time.Duration(10)
	usecase := test.NewApiSettingsUsecase(mockRepo, time)

	apiID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, apiID).Return(nil)

	err := usecase.Delete(context.Background(), apiID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestApiSettingsUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.ApiSettingsRepository)
	time := time.Duration(10)
	usecase := test.NewApiSettingsUsecase(mockRepo, time)

	fetchedApiSettings := []domian.ApiSettings{
		{

			ID:                       primitive.NewObjectID(), // Generate a new MongoDB ObjectID
			EnableApi:                true,                    // API is enabled
			AllowRequestsFromSwagger: true,                    // Allow requests from Swagger
			EnableLogging:            true,                    // Enable logging for API requests
			TokenKey:                 "example-token-key",     // Example token key for authentication

		},
		{

			ID:                       primitive.NewObjectID(), // Generate a new MongoDB ObjectID
			EnableApi:                true,                    // API is enabled
			AllowRequestsFromSwagger: true,                    // Allow requests from Swagger
			EnableLogging:            true,                    // Enable logging for API requests
			TokenKey:                 "example-token-key",     // Example token key for authentication

		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedApiSettings, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedApiSettings, result)
	mockRepo.AssertExpectations(t)
}
