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
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestMessageTemplatesSettingsUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.MessageTemplatesSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewMessageTemplatesSettingsUsecase(mockRepo, timeout)

	messageID := bson.NewObjectID().Hex()

	updatedMessageTemplatesSettings := domain.MessageTemplatesSettings{
		ID:                       bson.NewObjectID(), // Existing ID of the record to update
		CaseInvariantReplacement: false,
		Color1:                   "#FFFFFF",
		Color2:                   "#000000",
		Color3:                   "#CCCCCC",
	}

	mockRepo.On("FetchByID", mock.Anything, messageID).Return(updatedMessageTemplatesSettings, nil)

	result, err := usecase.FetchByID(context.Background(), messageID)

	assert.NoError(t, err)
	assert.Equal(t, updatedMessageTemplatesSettings, result)
	mockRepo.AssertExpectations(t)
}

func TestMessageTemplatesSettingsUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.MessageTemplatesSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewMessageTemplatesSettingsUsecase(mockRepo, timeout)

	newMessageTemplatesSettings := &domain.MessageTemplatesSettings{
		CaseInvariantReplacement: true,
		Color1:                   "#FF5733",
		Color2:                   "#33FF57",
		Color3:                   "#3357FF",
	}

	mockRepo.On("Create", mock.Anything, newMessageTemplatesSettings).Return(nil)

	err := usecase.Create(context.Background(), newMessageTemplatesSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestMessageTemplatesSettingsUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.MessageTemplatesSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewMessageTemplatesSettingsUsecase(mockRepo, timeout)

	updatedMessageTemplatesSettings := &domain.MessageTemplatesSettings{
		ID:                       bson.NewObjectID(), // Existing ID of the record to update
		CaseInvariantReplacement: false,
		Color1:                   "#FFFFFF",
		Color2:                   "#000000",
		Color3:                   "#CCCCCC",
	}

	mockRepo.On("Update", mock.Anything, updatedMessageTemplatesSettings).Return(nil)

	err := usecase.Update(context.Background(), updatedMessageTemplatesSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestMessageTemplatesSettingsUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.MessageTemplatesSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewMessageTemplatesSettingsUsecase(mockRepo, timeout)

	messageID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, messageID).Return(nil)

	err := usecase.Delete(context.Background(), messageID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestMessageTemplatesSettingsUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.MessageTemplatesSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewMessageTemplatesSettingsUsecase(mockRepo, timeout)

	fetchedMessageTemplatesSettings := []domain.MessageTemplatesSettings{
		{
			ID:                       bson.NewObjectID(),
			CaseInvariantReplacement: true,
			Color1:                   "#FF5733",
			Color2:                   "#33FF57",
			Color3:                   "#3357FF",
		},
		{
			ID:                       bson.NewObjectID(),
			CaseInvariantReplacement: false,
			Color1:                   "#FFFFFF",
			Color2:                   "#000000",
			Color3:                   "#CCCCCC",
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedMessageTemplatesSettings, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedMessageTemplatesSettings, result)
	mockRepo.AssertExpectations(t)
}
