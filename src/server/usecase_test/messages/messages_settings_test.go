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

func TestMessagesSettingsUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.MessagesSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewMessagesSettingsUsecase(mockRepo, timeout)

	messagesSettingsID := bson.NewObjectID().Hex()

	updatedMessagesSettings := domain.MessagesSettings{
		ID:                    bson.NewObjectID(), // Existing ID of the record to update
		UsePopupNotifications: false,
		UseDefaultEmailAccountForSendStoreOwnerEmails: true,
	}

	mockRepo.On("FetchByID", mock.Anything, messagesSettingsID).Return(updatedMessagesSettings, nil)

	result, err := usecase.FetchByID(context.Background(), messagesSettingsID)

	assert.NoError(t, err)
	assert.Equal(t, updatedMessagesSettings, result)
	mockRepo.AssertExpectations(t)
}

func TestMessagesSettingsUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.MessagesSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewMessagesSettingsUsecase(mockRepo, timeout)

	newMessagesSettings := &domain.MessagesSettings{
		UsePopupNotifications:                         true,
		UseDefaultEmailAccountForSendStoreOwnerEmails: false,
	}

	mockRepo.On("Create", mock.Anything, newMessagesSettings).Return(nil)

	err := usecase.Create(context.Background(), newMessagesSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestMessagesSettingsUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.MessagesSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewMessagesSettingsUsecase(mockRepo, timeout)

	updatedMessagesSettings := &domain.MessagesSettings{
		ID:                    bson.NewObjectID(), // Existing ID of the record to update
		UsePopupNotifications: false,
		UseDefaultEmailAccountForSendStoreOwnerEmails: true,
	}

	mockRepo.On("Update", mock.Anything, updatedMessagesSettings).Return(nil)

	err := usecase.Update(context.Background(), updatedMessagesSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestMessagesSettingsUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.MessagesSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewMessagesSettingsUsecase(mockRepo, timeout)

	messagesSettingsID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, messagesSettingsID).Return(nil)

	err := usecase.Delete(context.Background(), messagesSettingsID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestMessagesSettingsUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.MessagesSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewMessagesSettingsUsecase(mockRepo, timeout)

	fetchedMessagesSettings := []domain.MessagesSettings{
		{
			ID:                    bson.NewObjectID(),
			UsePopupNotifications: true,
			UseDefaultEmailAccountForSendStoreOwnerEmails: false,
		},
		{
			ID:                    bson.NewObjectID(),
			UsePopupNotifications: false,
			UseDefaultEmailAccountForSendStoreOwnerEmails: true,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedMessagesSettings, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedMessagesSettings, result)
	mockRepo.AssertExpectations(t)
}
