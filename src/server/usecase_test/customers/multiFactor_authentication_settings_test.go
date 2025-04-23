package usecase_test

import (
	"context"
	domain "earnforglance/server/domain/customers"
	mocks "earnforglance/server/domain/mocks"
	test "earnforglance/server/usecase/customers"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestMultiFactorAuthenticationSettingsUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.MultiFactorAuthenticationSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewMultiFactorAuthenticationSettingsUsecase(mockRepo, timeout)

	customersID := bson.NewObjectID().Hex()

	updatedMultiFactorAuthenticationSettings := domain.MultiFactorAuthenticationSettings{
		ActiveAuthenticationMethodSystemNames: []string{"GoogleAuthenticator", "Authy"}, // Example active authentication methods
		ForceMultifactorAuthentication:        true,
	}

	mockRepo.On("FetchByID", mock.Anything, customersID).Return(updatedMultiFactorAuthenticationSettings, nil)

	result, err := usecase.FetchByID(context.Background(), customersID)

	assert.NoError(t, err)
	assert.Equal(t, updatedMultiFactorAuthenticationSettings, result)
	mockRepo.AssertExpectations(t)
}

func TestMultiFactorAuthenticationSettingsUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.MultiFactorAuthenticationSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewMultiFactorAuthenticationSettingsUsecase(mockRepo, timeout)

	newMultiFactorAuthenticationSettings := &domain.MultiFactorAuthenticationSettings{
		ActiveAuthenticationMethodSystemNames: []string{"GoogleAuthenticator", "Authy"}, // Example active authentication methods
		ForceMultifactorAuthentication:        true,
	}

	mockRepo.On("Create", mock.Anything, newMultiFactorAuthenticationSettings).Return(nil)

	err := usecase.Create(context.Background(), newMultiFactorAuthenticationSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestMultiFactorAuthenticationSettingsUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.MultiFactorAuthenticationSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewMultiFactorAuthenticationSettingsUsecase(mockRepo, timeout)

	updatedMultiFactorAuthenticationSettings := &domain.MultiFactorAuthenticationSettings{
		ActiveAuthenticationMethodSystemNames: []string{"GoogleAuthenticator", "Authy"}, // Example active authentication methods
		ForceMultifactorAuthentication:        true,
	}

	mockRepo.On("Update", mock.Anything, updatedMultiFactorAuthenticationSettings).Return(nil)

	err := usecase.Update(context.Background(), updatedMultiFactorAuthenticationSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestMultiFactorAuthenticationSettingsUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.MultiFactorAuthenticationSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewMultiFactorAuthenticationSettingsUsecase(mockRepo, timeout)

	customersID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, customersID).Return(nil)

	err := usecase.Delete(context.Background(), customersID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestMultiFactorAuthenticationSettingsUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.MultiFactorAuthenticationSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewMultiFactorAuthenticationSettingsUsecase(mockRepo, timeout)

	fetchedMultiFactorAuthenticationSettings := []domain.MultiFactorAuthenticationSettings{
		{
			ActiveAuthenticationMethodSystemNames: []string{"MicrosoftAuthenticator"}, // Example active authentication method
			ForceMultifactorAuthentication:        false,                              // Do not force multi-factor authentication
		},
		{
			ActiveAuthenticationMethodSystemNames: []string{"GoogleAuthenticator", "Authy", "DuoSecurity"}, // Multiple active authentication methods
			ForceMultifactorAuthentication:        true,                                                    // Force multi-factor authentication
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedMultiFactorAuthenticationSettings, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedMultiFactorAuthenticationSettings, result)
	mockRepo.AssertExpectations(t)
}
