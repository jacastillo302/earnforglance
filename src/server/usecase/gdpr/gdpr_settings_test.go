package usecase

import (
	"context"
	domain "earnforglance/server/domain/gdpr"
	mocks "earnforglance/server/domain/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGdprSettingsUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.GdprSettingsRepository)
	timeout := time.Duration(10)
	usecase := NewGdprSettingsUsecase(mockRepo, timeout)

	gdprID := primitive.NewObjectID().Hex()

	updatedGdprSettings := domain.GdprSettings{
		ID:                                 primitive.NewObjectID(), // Existing ID of the record to update
		GdprEnabled:                        false,
		LogPrivacyPolicyConsent:            false,
		LogNewsletterConsent:               false,
		LogUserProfileChanges:              false,
		DeleteInactiveCustomersAfterMonths: 6,
	}

	mockRepo.On("FetchByID", mock.Anything, gdprID).Return(updatedGdprSettings, nil)

	result, err := usecase.FetchByID(context.Background(), gdprID)

	assert.NoError(t, err)
	assert.Equal(t, updatedGdprSettings, result)
	mockRepo.AssertExpectations(t)
}

func TestGdprSettingsUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.GdprSettingsRepository)
	timeout := time.Duration(10)
	usecase := NewGdprSettingsUsecase(mockRepo, timeout)

	newGdprSettings := &domain.GdprSettings{
		GdprEnabled:                        true,
		LogPrivacyPolicyConsent:            true,
		LogNewsletterConsent:               true,
		LogUserProfileChanges:              true,
		DeleteInactiveCustomersAfterMonths: 12,
	}
	mockRepo.On("Create", mock.Anything, newGdprSettings).Return(nil)

	err := usecase.Create(context.Background(), newGdprSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGdprSettingsUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.GdprSettingsRepository)
	timeout := time.Duration(10)
	usecase := NewGdprSettingsUsecase(mockRepo, timeout)

	updatedGdprSettings := &domain.GdprSettings{
		ID:                                 primitive.NewObjectID(), // Existing ID of the record to update
		GdprEnabled:                        false,
		LogPrivacyPolicyConsent:            false,
		LogNewsletterConsent:               false,
		LogUserProfileChanges:              false,
		DeleteInactiveCustomersAfterMonths: 6,
	}

	mockRepo.On("Update", mock.Anything, updatedGdprSettings).Return(nil)

	err := usecase.Update(context.Background(), updatedGdprSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGdprSettingsUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.GdprSettingsRepository)
	timeout := time.Duration(10)
	usecase := NewGdprSettingsUsecase(mockRepo, timeout)

	gdprID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, gdprID).Return(nil)

	err := usecase.Delete(context.Background(), gdprID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGdprSettingsUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.GdprSettingsRepository)
	timeout := time.Duration(10)
	usecase := NewGdprSettingsUsecase(mockRepo, timeout)

	fetchedGdprSettings := []domain.GdprSettings{
		{
			ID:                                 primitive.NewObjectID(),
			GdprEnabled:                        true,
			LogPrivacyPolicyConsent:            true,
			LogNewsletterConsent:               true,
			LogUserProfileChanges:              true,
			DeleteInactiveCustomersAfterMonths: 12,
		},
		{
			ID:                                 primitive.NewObjectID(),
			GdprEnabled:                        false,
			LogPrivacyPolicyConsent:            false,
			LogNewsletterConsent:               false,
			LogUserProfileChanges:              false,
			DeleteInactiveCustomersAfterMonths: 6,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedGdprSettings, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedGdprSettings, result)
	mockRepo.AssertExpectations(t)
}
