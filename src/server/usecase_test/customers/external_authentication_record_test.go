package usecase_test

import (
	"context"
	domian "earnforglance/server/domain/customers"
	mocks "earnforglance/server/domain/mocks"
	test "earnforglance/server/usecase/customers"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestExternalAuthenticationRecordUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.ExternalAuthenticationRecordRepository)
	timeout := time.Duration(10)
	usecase := test.NewExternalAuthenticationRecordUsecase(mockRepo, timeout)

	customersID := bson.NewObjectID().Hex()

	updatedExternalAuthenticationRecord := domian.ExternalAuthenticationRecord{
		CustomerID:                bson.NewObjectID(),       // Example Customer ID
		Email:                     "example@example.com",    // Example external email
		ExternalIdentifier:        "external-id-12345",      // Example external identifier
		ExternalDisplayIdentifier: "ExternalUser123",        // Example external display identifier
		OAuthToken:                "oauth-token-abc123",     // Example OAuth token
		OAuthAccessToken:          "oauth-access-token-xyz", // Example OAuth access token
		ProviderSystemName:        "Google",
	}

	mockRepo.On("FetchByID", mock.Anything, customersID).Return(updatedExternalAuthenticationRecord, nil)

	result, err := usecase.FetchByID(context.Background(), customersID)

	assert.NoError(t, err)
	assert.Equal(t, updatedExternalAuthenticationRecord, result)
	mockRepo.AssertExpectations(t)
}

func TestExternalAuthenticationRecordUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.ExternalAuthenticationRecordRepository)
	timeout := time.Duration(10)
	usecase := test.NewExternalAuthenticationRecordUsecase(mockRepo, timeout)

	newExternalAuthenticationRecord := &domian.ExternalAuthenticationRecord{
		ID:                        bson.NewObjectID(),       // Generate a new MongoDB ObjectID
		CustomerID:                bson.NewObjectID(),       // Example Customer ID
		Email:                     "example@example.com",    // Example external email
		ExternalIdentifier:        "external-id-12345",      // Example external identifier
		ExternalDisplayIdentifier: "ExternalUser123",        // Example external display identifier
		OAuthToken:                "oauth-token-abc123",     // Example OAuth token
		OAuthAccessToken:          "oauth-access-token-xyz", // Example OAuth access token
		ProviderSystemName:        "Google",
	}

	mockRepo.On("Create", mock.Anything, newExternalAuthenticationRecord).Return(nil)

	err := usecase.Create(context.Background(), newExternalAuthenticationRecord)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestExternalAuthenticationRecordUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.ExternalAuthenticationRecordRepository)
	timeout := time.Duration(10)
	usecase := test.NewExternalAuthenticationRecordUsecase(mockRepo, timeout)

	updatedExternalAuthenticationRecord := &domian.ExternalAuthenticationRecord{
		ID:                        bson.NewObjectID(),       // Generate a new MongoDB ObjectID
		CustomerID:                bson.NewObjectID(),       // Example Customer ID
		Email:                     "example@example.com",    // Example external email
		ExternalIdentifier:        "external-id-12345",      // Example external identifier
		ExternalDisplayIdentifier: "ExternalUser123",        // Example external display identifier
		OAuthToken:                "oauth-token-abc123",     // Example OAuth token
		OAuthAccessToken:          "oauth-access-token-xyz", // Example OAuth access token
		ProviderSystemName:        "Google",
	}

	mockRepo.On("Update", mock.Anything, updatedExternalAuthenticationRecord).Return(nil)

	err := usecase.Update(context.Background(), updatedExternalAuthenticationRecord)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestExternalAuthenticationRecordUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.ExternalAuthenticationRecordRepository)
	timeout := time.Duration(10)
	usecase := test.NewExternalAuthenticationRecordUsecase(mockRepo, timeout)

	customersID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, customersID).Return(nil)

	err := usecase.Delete(context.Background(), customersID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestExternalAuthenticationRecordUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.ExternalAuthenticationRecordRepository)
	timeout := time.Duration(10)
	usecase := test.NewExternalAuthenticationRecordUsecase(mockRepo, timeout)

	fetchedExternalAuthenticationRecords := []domian.ExternalAuthenticationRecord{
		{
			ID:                        bson.NewObjectID(),
			CustomerID:                bson.NewObjectID(),
			Email:                     "user1@example.com",
			ExternalIdentifier:        "external-id-67890",
			ExternalDisplayIdentifier: "ExternalUser1",
			OAuthToken:                "oauth-token-def456",
			OAuthAccessToken:          "oauth-access-token-uvw",
			ProviderSystemName:        "Facebook",
		},
		{
			ID:                        bson.NewObjectID(),
			CustomerID:                bson.NewObjectID(),
			Email:                     "user2@example.com",
			ExternalIdentifier:        "external-id-11223",
			ExternalDisplayIdentifier: "ExternalUser2",
			OAuthToken:                "oauth-token-ghi789",
			OAuthAccessToken:          "oauth-access-token-klm",
			ProviderSystemName:        "Twitter",
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedExternalAuthenticationRecords, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedExternalAuthenticationRecords, result)
	mockRepo.AssertExpectations(t)
}
