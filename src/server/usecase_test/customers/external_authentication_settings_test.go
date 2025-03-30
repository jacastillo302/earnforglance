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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestExternalAuthenticationSettingsUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.ExternalAuthenticationSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewExternalAuthenticationSettingsUsecase(mockRepo, timeout)

	customersID := primitive.NewObjectID().Hex()

	updatedExternalAuthenticationSettings := domian.ExternalAuthenticationSettings{
		RequireEmailValidation:                true,                                      // Email validation is required
		LogErrors:                             true,                                      // Log errors during authentication
		AllowCustomersToRemoveAssociations:    true,                                      // Allow users to remove external authentication associations
		ActiveAuthenticationMethodSystemNames: []string{"Google", "Facebook", "Twitter"}, // Example active authentication methods
	}

	mockRepo.On("FetchByID", mock.Anything, customersID).Return(updatedExternalAuthenticationSettings, nil)

	result, err := usecase.FetchByID(context.Background(), customersID)

	assert.NoError(t, err)
	assert.Equal(t, updatedExternalAuthenticationSettings, result)
	mockRepo.AssertExpectations(t)
}

func TestExternalAuthenticationSettingsUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.ExternalAuthenticationSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewExternalAuthenticationSettingsUsecase(mockRepo, timeout)

	newExternalAuthenticationSettings := &domian.ExternalAuthenticationSettings{
		RequireEmailValidation:                true,                                      // Email validation is required
		LogErrors:                             true,                                      // Log errors during authentication
		AllowCustomersToRemoveAssociations:    true,                                      // Allow users to remove external authentication associations
		ActiveAuthenticationMethodSystemNames: []string{"Google", "Facebook", "Twitter"}, // Example active authentication methods
	}

	mockRepo.On("Create", mock.Anything, newExternalAuthenticationSettings).Return(nil)

	err := usecase.Create(context.Background(), newExternalAuthenticationSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestExternalAuthenticationSettingsUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.ExternalAuthenticationSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewExternalAuthenticationSettingsUsecase(mockRepo, timeout)

	updatedExternalAuthenticationSettings := &domian.ExternalAuthenticationSettings{
		RequireEmailValidation:                true,                                      // Email validation is required
		LogErrors:                             true,                                      // Log errors during authentication
		AllowCustomersToRemoveAssociations:    true,                                      // Allow users to remove external authentication associations
		ActiveAuthenticationMethodSystemNames: []string{"Google", "Facebook", "Twitter"}, // Example active authentication methods
	}

	mockRepo.On("Update", mock.Anything, updatedExternalAuthenticationSettings).Return(nil)

	err := usecase.Update(context.Background(), updatedExternalAuthenticationSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestExternalAuthenticationSettingsUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.ExternalAuthenticationSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewExternalAuthenticationSettingsUsecase(mockRepo, timeout)

	customersID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, customersID).Return(nil)

	err := usecase.Delete(context.Background(), customersID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestExternalAuthenticationSettingsUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.ExternalAuthenticationSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewExternalAuthenticationSettingsUsecase(mockRepo, timeout)

	fetchedExternalAuthenticationSettings := []domian.ExternalAuthenticationSettings{
		{
			RequireEmailValidation:                false,                // Email validation is not required
			LogErrors:                             false,                // Do not log errors
			AllowCustomersToRemoveAssociations:    false,                // Do not allow users to remove associations
			ActiveAuthenticationMethodSystemNames: []string{"LinkedIn"}, // Example active authentication method
		},
		{
			RequireEmailValidation:                true,
			LogErrors:                             true,
			AllowCustomersToRemoveAssociations:    true,
			ActiveAuthenticationMethodSystemNames: []string{"GitHub", "Microsoft"}, // Example active authentication methods
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedExternalAuthenticationSettings, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedExternalAuthenticationSettings, result)
	mockRepo.AssertExpectations(t)
}
