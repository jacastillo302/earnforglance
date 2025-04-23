package usecase_test

import (
	"context"
	mocks "earnforglance/server/domain/mocks"
	domain "earnforglance/server/domain/security"
	test "earnforglance/server/usecase/security"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestSecuritySettingsUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.SecuritySettingsRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewSecuritySettingsUsecase(mockRepo, timeout)

	securityID := bson.NewObjectID().Hex()

	updatedSecuritySettings := domain.SecuritySettings{
		ID:                               bson.NewObjectID(), // Existing ID of the record to update
		EncryptionKey:                    "updatedEncryptionKey456",
		AdminAreaAllowedIpAddresses:      []string{"127.0.0.1", "192.168.0.1"},
		HoneypotEnabled:                  false,
		HoneypotInputName:                "updated_honeypot_field",
		LogHoneypotDetection:             false,
		AllowNonAsciiCharactersInHeaders: true,
		UseAesEncryptionAlgorithm:        false,
		AllowStoreOwnerExportImportCustomersWithHashedPassword: true,
	}

	mockRepo.On("FetchByID", mock.Anything, securityID).Return(updatedSecuritySettings, nil)

	result, err := usecase.FetchByID(context.Background(), securityID)

	assert.NoError(t, err)
	assert.Equal(t, updatedSecuritySettings, result)
	mockRepo.AssertExpectations(t)
}

func TestSecuritySettingsUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.SecuritySettingsRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewSecuritySettingsUsecase(mockRepo, timeout)

	newSecuritySettings := &domain.SecuritySettings{
		EncryptionKey:                    "randomEncryptionKey123",
		AdminAreaAllowedIpAddresses:      []string{"192.168.1.1", "10.0.0.1"},
		HoneypotEnabled:                  true,
		HoneypotInputName:                "honeypot_field",
		LogHoneypotDetection:             true,
		AllowNonAsciiCharactersInHeaders: false,
		UseAesEncryptionAlgorithm:        true,
		AllowStoreOwnerExportImportCustomersWithHashedPassword: false,
	}

	mockRepo.On("Create", mock.Anything, newSecuritySettings).Return(nil)

	err := usecase.Create(context.Background(), newSecuritySettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestSecuritySettingsUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.SecuritySettingsRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewSecuritySettingsUsecase(mockRepo, timeout)

	updatedSecuritySettings := &domain.SecuritySettings{
		ID:                               bson.NewObjectID(), // Existing ID of the record to update
		EncryptionKey:                    "updatedEncryptionKey456",
		AdminAreaAllowedIpAddresses:      []string{"127.0.0.1", "192.168.0.1"},
		HoneypotEnabled:                  false,
		HoneypotInputName:                "updated_honeypot_field",
		LogHoneypotDetection:             false,
		AllowNonAsciiCharactersInHeaders: true,
		UseAesEncryptionAlgorithm:        false,
		AllowStoreOwnerExportImportCustomersWithHashedPassword: true,
	}

	mockRepo.On("Update", mock.Anything, updatedSecuritySettings).Return(nil)

	err := usecase.Update(context.Background(), updatedSecuritySettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestSecuritySettingsUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.SecuritySettingsRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewSecuritySettingsUsecase(mockRepo, timeout)

	securityID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, securityID).Return(nil)

	err := usecase.Delete(context.Background(), securityID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestSecuritySettingsUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.SecuritySettingsRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewSecuritySettingsUsecase(mockRepo, timeout)

	fetchedSecuritySettings := []domain.SecuritySettings{
		{
			ID:                               bson.NewObjectID(),
			EncryptionKey:                    "randomEncryptionKey123",
			AdminAreaAllowedIpAddresses:      []string{"192.168.1.1", "10.0.0.1"},
			HoneypotEnabled:                  true,
			HoneypotInputName:                "honeypot_field",
			LogHoneypotDetection:             true,
			AllowNonAsciiCharactersInHeaders: false,
			UseAesEncryptionAlgorithm:        true,
			AllowStoreOwnerExportImportCustomersWithHashedPassword: false,
		},
		{
			ID:                               bson.NewObjectID(),
			EncryptionKey:                    "updatedEncryptionKey456",
			AdminAreaAllowedIpAddresses:      []string{"127.0.0.1", "192.168.0.1"},
			HoneypotEnabled:                  false,
			HoneypotInputName:                "updated_honeypot_field",
			LogHoneypotDetection:             false,
			AllowNonAsciiCharactersInHeaders: true,
			UseAesEncryptionAlgorithm:        false,
			AllowStoreOwnerExportImportCustomersWithHashedPassword: true,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedSecuritySettings, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedSecuritySettings, result)
	mockRepo.AssertExpectations(t)
}
