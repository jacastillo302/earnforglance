package usecase

import (
	"context"
	domain "earnforglance/server/domain/forums"
	mocks "earnforglance/server/domain/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestPrivateMessageUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.PrivateMessageRepository)
	timeout := time.Duration(10)
	usecase := NewPrivateMessageUsecase(mockRepo, timeout)

	privateMessageID := primitive.NewObjectID().Hex()

	updatedPrivateMessage := domain.PrivateMessage{
		ID:                   primitive.NewObjectID(), // Existing ID of the record to update
		StoreID:              primitive.NewObjectID(),
		FromCustomerID:       primitive.NewObjectID(),
		ToCustomerID:         primitive.NewObjectID(),
		Subject:              "Updated Subject",
		Text:                 "This is an updated private message.",
		IsRead:               true,
		IsDeletedByAuthor:    false,
		IsDeletedByRecipient: true,
		CreatedOnUtc:         time.Now().AddDate(0, 0, -7), // Created 7 days ago
	}

	mockRepo.On("FetchByID", mock.Anything, privateMessageID).Return(updatedPrivateMessage, nil)

	result, err := usecase.FetchByID(context.Background(), privateMessageID)

	assert.NoError(t, err)
	assert.Equal(t, updatedPrivateMessage, result)
	mockRepo.AssertExpectations(t)
}

func TestPrivateMessageUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.PrivateMessageRepository)
	timeout := time.Duration(10)
	usecase := NewPrivateMessageUsecase(mockRepo, timeout)

	newPrivateMessage := &domain.PrivateMessage{
		StoreID:              primitive.NewObjectID(),
		FromCustomerID:       primitive.NewObjectID(),
		ToCustomerID:         primitive.NewObjectID(),
		Subject:              "Welcome to the platform",
		Text:                 "Hello! Welcome to our platform. Let us know if you need any help.",
		IsRead:               false,
		IsDeletedByAuthor:    false,
		IsDeletedByRecipient: false,
		CreatedOnUtc:         time.Now(),
	}

	mockRepo.On("Create", mock.Anything, newPrivateMessage).Return(nil)

	err := usecase.Create(context.Background(), newPrivateMessage)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestPrivateMessageUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.PrivateMessageRepository)
	timeout := time.Duration(10)
	usecase := NewPrivateMessageUsecase(mockRepo, timeout)

	updatedPrivateMessage := &domain.PrivateMessage{
		ID:                   primitive.NewObjectID(), // Existing ID of the record to update
		StoreID:              primitive.NewObjectID(),
		FromCustomerID:       primitive.NewObjectID(),
		ToCustomerID:         primitive.NewObjectID(),
		Subject:              "Updated Subject",
		Text:                 "This is an updated private message.",
		IsRead:               true,
		IsDeletedByAuthor:    false,
		IsDeletedByRecipient: true,
		CreatedOnUtc:         time.Now().AddDate(0, 0, -7), // Created 7 days ago
	}
	mockRepo.On("Update", mock.Anything, updatedPrivateMessage).Return(nil)

	err := usecase.Update(context.Background(), updatedPrivateMessage)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestPrivateMessageUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.PrivateMessageRepository)
	timeout := time.Duration(10)
	usecase := NewPrivateMessageUsecase(mockRepo, timeout)

	privateMessageID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, privateMessageID).Return(nil)

	err := usecase.Delete(context.Background(), privateMessageID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestPrivateMessageUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.PrivateMessageRepository)
	timeout := time.Duration(10)
	usecase := NewPrivateMessageUsecase(mockRepo, timeout)

	fetchedPrivateMessages := []domain.PrivateMessage{
		{
			ID:                   primitive.NewObjectID(),
			StoreID:              primitive.NewObjectID(),
			FromCustomerID:       primitive.NewObjectID(),
			ToCustomerID:         primitive.NewObjectID(),
			Subject:              "Welcome to the platform",
			Text:                 "Hello! Welcome to our platform. Let us know if you need any help.",
			IsRead:               false,
			IsDeletedByAuthor:    false,
			IsDeletedByRecipient: false,
			CreatedOnUtc:         time.Now().AddDate(0, 0, -10), // Created 10 days ago
		},
		{
			ID:                   primitive.NewObjectID(),
			StoreID:              primitive.NewObjectID(),
			FromCustomerID:       primitive.NewObjectID(),
			ToCustomerID:         primitive.NewObjectID(),
			Subject:              "Follow-up",
			Text:                 "Just checking in to see if you need assistance.",
			IsRead:               true,
			IsDeletedByAuthor:    true,
			IsDeletedByRecipient: false,
			CreatedOnUtc:         time.Now().AddDate(0, 0, -5), // Created 5 days ago
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedPrivateMessages, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedPrivateMessages, result)
	mockRepo.AssertExpectations(t)
}
