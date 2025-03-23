package usecase_test

import (
	"context"
	domain "earnforglance/server/domain/messages"
	mocks "earnforglance/server/domain/mocks"
	test "earnforglance/server/usecase/messages"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestEmailAccountSettingsUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.EmailAccountSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewEmailAccountSettingsUsecase(mockRepo, timeout)

	emailAccountSettingsID := primitive.NewObjectID().Hex()
	updatedEmailAccountSettings := domain.EmailAccountSettings{
		ID:                    primitive.NewObjectID(), // Existing ID of the record to update
		DefaultEmailAccountID: primitive.NewObjectID(),
	}

	mockRepo.On("FetchByID", mock.Anything, emailAccountSettingsID).Return(updatedEmailAccountSettings, nil)

	result, err := usecase.FetchByID(context.Background(), emailAccountSettingsID)

	assert.NoError(t, err)
	assert.Equal(t, updatedEmailAccountSettings, result)
	mockRepo.AssertExpectations(t)
}

func TestEmailAccountSettingsUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.EmailAccountSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewEmailAccountSettingsUsecase(mockRepo, timeout)

	newEmailAccountSettings := &domain.EmailAccountSettings{
		DefaultEmailAccountID: primitive.NewObjectID(),
	}

	mockRepo.On("Create", mock.Anything, newEmailAccountSettings).Return(nil)

	err := usecase.Create(context.Background(), newEmailAccountSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestEmailAccountSettingsUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.EmailAccountSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewEmailAccountSettingsUsecase(mockRepo, timeout)

	updatedEmailAccountSettings := &domain.EmailAccountSettings{
		ID:                    primitive.NewObjectID(), // Existing ID of the record to update
		DefaultEmailAccountID: primitive.NewObjectID(),
	}

	mockRepo.On("Update", mock.Anything, updatedEmailAccountSettings).Return(nil)

	err := usecase.Update(context.Background(), updatedEmailAccountSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestEmailAccountSettingsUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.EmailAccountSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewEmailAccountSettingsUsecase(mockRepo, timeout)

	emailAccountSettingsID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, emailAccountSettingsID).Return(nil)

	err := usecase.Delete(context.Background(), emailAccountSettingsID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestEmailAccountSettingsUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.EmailAccountSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewEmailAccountSettingsUsecase(mockRepo, timeout)

	fetchedEmailAccountSettings := []domain.EmailAccountSettings{
		{
			ID:                    primitive.NewObjectID(),
			DefaultEmailAccountID: primitive.NewObjectID(),
		},
		{
			ID:                    primitive.NewObjectID(),
			DefaultEmailAccountID: primitive.NewObjectID(),
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedEmailAccountSettings, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedEmailAccountSettings, result)
	mockRepo.AssertExpectations(t)
}
