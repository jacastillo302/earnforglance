package usecase

import (
	"context"
	domain "earnforglance/server/domain/localization"
	mocks "earnforglance/server/domain/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestLocaleStringResourceUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.LocaleStringResourceRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewLocaleStringResourceUsecase(mockRepo, timeout)

	localeStringResourceID := primitive.NewObjectID().Hex()

	updatedLocaleStringResource := domain.LocaleStringResource{
		ID:            primitive.NewObjectID(), // Existing ID of the record to update
		LanguageID:    primitive.NewObjectID(),
		ResourceName:  "WelcomeMessage",
		ResourceValue: "Welcome to our updated platform!",
	}

	mockRepo.On("FetchByID", mock.Anything, localeStringResourceID).Return(updatedLocaleStringResource, nil)

	result, err := usecase.FetchByID(context.Background(), localeStringResourceID)

	assert.NoError(t, err)
	assert.Equal(t, updatedLocaleStringResource, result)
	mockRepo.AssertExpectations(t)
}

func TestLocaleStringResourceUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.LocaleStringResourceRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewLocaleStringResourceUsecase(mockRepo, timeout)

	newLocaleStringResource := &domain.LocaleStringResource{
		LanguageID:    primitive.NewObjectID(),
		ResourceName:  "WelcomeMessage",
		ResourceValue: "Welcome to our platform!",
	}

	mockRepo.On("Create", mock.Anything, newLocaleStringResource).Return(nil)

	err := usecase.Create(context.Background(), newLocaleStringResource)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestLocaleStringResourceUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.LocaleStringResourceRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewLocaleStringResourceUsecase(mockRepo, timeout)

	updatedLocaleStringResource := &domain.LocaleStringResource{
		ID:            primitive.NewObjectID(), // Existing ID of the record to update
		LanguageID:    primitive.NewObjectID(),
		ResourceName:  "WelcomeMessage",
		ResourceValue: "Welcome to our updated platform!",
	}

	mockRepo.On("Update", mock.Anything, updatedLocaleStringResource).Return(nil)

	err := usecase.Update(context.Background(), updatedLocaleStringResource)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestLocaleStringResourceUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.LocaleStringResourceRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewLocaleStringResourceUsecase(mockRepo, timeout)

	localeStringResourceID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, localeStringResourceID).Return(nil)

	err := usecase.Delete(context.Background(), localeStringResourceID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestLocaleStringResourceUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.LocaleStringResourceRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewLocaleStringResourceUsecase(mockRepo, timeout)

	fetchedLocaleStringResources := []domain.LocaleStringResource{
		{
			ID:            primitive.NewObjectID(),
			LanguageID:    primitive.NewObjectID(),
			ResourceName:  "WelcomeMessage",
			ResourceValue: "Welcome to our platform!",
		},
		{
			ID:            primitive.NewObjectID(),
			LanguageID:    primitive.NewObjectID(),
			ResourceName:  "GoodbyeMessage",
			ResourceValue: "Thank you for visiting!",
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedLocaleStringResources, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedLocaleStringResources, result)
	mockRepo.AssertExpectations(t)
}
