package usecase_test

import (
	"context"
	domain "earnforglance/server/domain/gdpr"
	mocks "earnforglance/server/domain/mocks"
	test "earnforglance/server/usecase/gdpr"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGdprConsentUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.GdprConsentRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewGdprConsentUsecase(mockRepo, timeout)

	gdprID := primitive.NewObjectID().Hex()

	updatedGdprConsent := domain.GdprConsent{
		ID:                        primitive.NewObjectID(), // Existing ID of the record to update
		Message:                   "Updated GDPR consent message.",
		IsRequired:                false,
		RequiredMessage:           "Consent is optional.",
		DisplayDuringRegistration: false,
		DisplayOnCustomerInfoPage: true,
		DisplayOrder:              2,
	}

	mockRepo.On("FetchByID", mock.Anything, gdprID).Return(updatedGdprConsent, nil)

	result, err := usecase.FetchByID(context.Background(), gdprID)

	assert.NoError(t, err)
	assert.Equal(t, updatedGdprConsent, result)
	mockRepo.AssertExpectations(t)
}

func TestGdprConsentUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.GdprConsentRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewGdprConsentUsecase(mockRepo, timeout)

	newGdprConsent := &domain.GdprConsent{
		Message:                   "We need your consent to process your data.",
		IsRequired:                true,
		RequiredMessage:           "Consent is mandatory to proceed.",
		DisplayDuringRegistration: true,
		DisplayOnCustomerInfoPage: true,
		DisplayOrder:              1,
	}

	mockRepo.On("Create", mock.Anything, newGdprConsent).Return(nil)

	err := usecase.Create(context.Background(), newGdprConsent)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGdprConsentUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.GdprConsentRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewGdprConsentUsecase(mockRepo, timeout)

	updatedGdprConsent := &domain.GdprConsent{
		ID:                        primitive.NewObjectID(), // Existing ID of the record to update
		Message:                   "Updated GDPR consent message.",
		IsRequired:                false,
		RequiredMessage:           "Consent is optional.",
		DisplayDuringRegistration: false,
		DisplayOnCustomerInfoPage: true,
		DisplayOrder:              2,
	}

	mockRepo.On("Update", mock.Anything, updatedGdprConsent).Return(nil)

	err := usecase.Update(context.Background(), updatedGdprConsent)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGdprConsentUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.GdprConsentRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewGdprConsentUsecase(mockRepo, timeout)

	gdprID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, gdprID).Return(nil)

	err := usecase.Delete(context.Background(), gdprID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGdprConsentUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.GdprConsentRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewGdprConsentUsecase(mockRepo, timeout)

	fetchedGdprConsents := []domain.GdprConsent{
		{
			ID:                        primitive.NewObjectID(),
			Message:                   "We need your consent to process your data.",
			IsRequired:                true,
			RequiredMessage:           "Consent is mandatory to proceed.",
			DisplayDuringRegistration: true,
			DisplayOnCustomerInfoPage: true,
			DisplayOrder:              1,
		},
		{
			ID:                        primitive.NewObjectID(),
			Message:                   "Updated GDPR consent message.",
			IsRequired:                false,
			RequiredMessage:           "Consent is optional.",
			DisplayDuringRegistration: false,
			DisplayOnCustomerInfoPage: true,
			DisplayOrder:              2,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedGdprConsents, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedGdprConsents, result)
	mockRepo.AssertExpectations(t)
}
