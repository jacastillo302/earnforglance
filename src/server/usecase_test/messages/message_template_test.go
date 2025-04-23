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

func TestMessageTemplateUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.MessageTemplateRepository)
	timeout := time.Duration(10)
	usecase := test.NewMessageTemplateUsecase(mockRepo, timeout)

	messageTemplateID := ""

	updatedMessageTemplate := domain.MessageTemplate{
		ID:                 bson.NewObjectID(), // Existing ID of the record to update
		Name:               "Updated Order Confirmation",
		BccEmailAddresses:  "updated_bcc@example.com",
		Subject:            "Updated Order Confirmation Subject",
		Body:               "Your order has been updated. Your new order number is #67890.",
		IsActive:           false,
		DelayBeforeSend:    new(int),
		DelayPeriodID:      2,
		AttachedDownloadID: "",
		AllowDirectReply:   false,
		EmailAccountID:     bson.NewObjectID(),
		LimitedToStores:    true,
	}

	mockRepo.On("FetchByID", mock.Anything, messageTemplateID).Return(updatedMessageTemplate, nil)

	result, err := usecase.FetchByID(context.Background(), messageTemplateID)

	assert.NoError(t, err)
	assert.Equal(t, updatedMessageTemplate, result)
	mockRepo.AssertExpectations(t)
}

func TestMessageTemplateUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.MessageTemplateRepository)
	timeout := time.Duration(10)
	usecase := test.NewMessageTemplateUsecase(mockRepo, timeout)

	newMessageTemplate := &domain.MessageTemplate{
		Name:               "Order Confirmation",
		BccEmailAddresses:  "bcc@example.com",
		Subject:            "Your Order Confirmation",
		Body:               "Thank you for your order. Your order number is #12345.",
		IsActive:           true,
		DelayBeforeSend:    new(int),
		DelayPeriodID:      1,
		AttachedDownloadID: "",
		AllowDirectReply:   true,
		EmailAccountID:     bson.NewObjectID(),
		LimitedToStores:    false,
	}
	*newMessageTemplate.DelayBeforeSend = 2

	mockRepo.On("Create", mock.Anything, newMessageTemplate).Return(nil)

	err := usecase.Create(context.Background(), newMessageTemplate)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestMessageTemplateUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.MessageTemplateRepository)
	timeout := time.Duration(10)
	usecase := test.NewMessageTemplateUsecase(mockRepo, timeout)

	updatedMessageTemplate := &domain.MessageTemplate{
		ID:                 bson.NewObjectID(), // Existing ID of the record to update
		Name:               "Updated Order Confirmation",
		BccEmailAddresses:  "updated_bcc@example.com",
		Subject:            "Updated Order Confirmation Subject",
		Body:               "Your order has been updated. Your new order number is #67890.",
		IsActive:           false,
		DelayBeforeSend:    new(int),
		DelayPeriodID:      2,
		AttachedDownloadID: "",
		AllowDirectReply:   false,
		EmailAccountID:     bson.NewObjectID(),
		LimitedToStores:    true,
	}
	*updatedMessageTemplate.DelayBeforeSend = 1

	mockRepo.On("Update", mock.Anything, updatedMessageTemplate).Return(nil)

	err := usecase.Update(context.Background(), updatedMessageTemplate)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestMessageTemplateUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.MessageTemplateRepository)
	timeout := time.Duration(10)
	usecase := test.NewMessageTemplateUsecase(mockRepo, timeout)

	messageTemplateID := ""

	mockRepo.On("Delete", mock.Anything, messageTemplateID).Return(nil)

	err := usecase.Delete(context.Background(), messageTemplateID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestMessageTemplateUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.MessageTemplateRepository)
	timeout := time.Duration(10)
	usecase := test.NewMessageTemplateUsecase(mockRepo, timeout)

	fetchedMessageTemplates := []domain.MessageTemplate{
		{
			ID:                 bson.NewObjectID(),
			Name:               "Order Confirmation",
			BccEmailAddresses:  "bcc@example.com",
			Subject:            "Your Order Confirmation",
			Body:               "Thank you for your order. Your order number is #12345.",
			IsActive:           true,
			DelayBeforeSend:    new(int),
			DelayPeriodID:      1,
			AttachedDownloadID: "",
			AllowDirectReply:   true,
			EmailAccountID:     bson.NewObjectID(),
			LimitedToStores:    false,
		},
		{
			ID:                 bson.NewObjectID(),
			Name:               "Shipping Notification",
			BccEmailAddresses:  "shipping_bcc@example.com",
			Subject:            "Your Order Has Shipped",
			Body:               "Your order has been shipped. Tracking number: 123456789.",
			IsActive:           true,
			DelayBeforeSend:    new(int),
			DelayPeriodID:      2,
			AttachedDownloadID: "",
			AllowDirectReply:   false,
			EmailAccountID:     bson.NewObjectID(),
			LimitedToStores:    true,
		},
	}
	*fetchedMessageTemplates[0].DelayBeforeSend = 2
	*fetchedMessageTemplates[1].DelayBeforeSend = 1

	mockRepo.On("Fetch", mock.Anything).Return(fetchedMessageTemplates, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedMessageTemplates, result)
	mockRepo.AssertExpectations(t)
}
