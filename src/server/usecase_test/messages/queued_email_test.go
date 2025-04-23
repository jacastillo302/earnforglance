package usecase_test

import (
	"context"
	"testing"
	"time"

	domain "earnforglance/server/domain/messages"
	mocks "earnforglance/server/domain/mocks"
	test "earnforglance/server/usecase/messages"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestQueuedEmailUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.QueuedEmailRepository)
	timeout := time.Duration(10)
	usecase := test.NewQueuedEmailUsecase(mockRepo, timeout)

	queuedEmailID := bson.NewObjectID().Hex()

	updatedQueuedEmail := domain.QueuedEmail{
		ID:                    bson.NewObjectID(), // Existing ID of the record to update
		PriorityID:            2,
		From:                  "updated_sender@example.com",
		FromName:              "Updated Sender Name",
		To:                    "updated_recipient@example.com",
		ToName:                "Updated Recipient Name",
		ReplyTo:               "updated_replyto@example.com",
		ReplyToName:           "Updated Reply To Name",
		CC:                    "updated_cc@example.com",
		Bcc:                   "updated_bcc@example.com",
		Subject:               "Updated Test Email",
		Body:                  "This is an updated test email.",
		AttachmentFilePath:    "/path/to/updated_attachment.pdf",
		AttachmentFileName:    "updated_attachment.pdf",
		AttachedDownloadID:    bson.NewObjectID(),
		CreatedOnUtc:          time.Now().AddDate(0, 0, -7), // Created 7 days ago
		DontSendBeforeDateUtc: new(time.Time),
		SentTries:             1,
		SentOnUtc:             new(time.Time),
		EmailAccountID:        bson.NewObjectID(),
	}

	mockRepo.On("FetchByID", mock.Anything, queuedEmailID).Return(updatedQueuedEmail, nil)

	result, err := usecase.FetchByID(context.Background(), queuedEmailID)

	assert.NoError(t, err)
	assert.Equal(t, updatedQueuedEmail, result)
	mockRepo.AssertExpectations(t)
}

func TestQueuedEmailUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.QueuedEmailRepository)
	timeout := time.Duration(10)
	usecase := test.NewQueuedEmailUsecase(mockRepo, timeout)

	newQueuedEmail := &domain.QueuedEmail{
		PriorityID:            1,
		From:                  "sender@example.com",
		FromName:              "Sender Name",
		To:                    "recipient@example.com",
		ToName:                "Recipient Name",
		ReplyTo:               "replyto@example.com",
		ReplyToName:           "Reply To Name",
		CC:                    "cc@example.com",
		Bcc:                   "bcc@example.com",
		Subject:               "Test Email",
		Body:                  "This is a test email.",
		AttachmentFilePath:    "/path/to/attachment.pdf",
		AttachmentFileName:    "attachment.pdf",
		AttachedDownloadID:    bson.NewObjectID(),
		CreatedOnUtc:          time.Now(),
		DontSendBeforeDateUtc: nil,
		SentTries:             0,
		SentOnUtc:             nil,
		EmailAccountID:        bson.NewObjectID(),
	}

	mockRepo.On("Create", mock.Anything, newQueuedEmail).Return(nil)

	err := usecase.Create(context.Background(), newQueuedEmail)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestQueuedEmailUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.QueuedEmailRepository)
	timeout := time.Duration(10)
	usecase := test.NewQueuedEmailUsecase(mockRepo, timeout)

	updatedQueuedEmail := &domain.QueuedEmail{
		ID:                    bson.NewObjectID(), // Existing ID of the record to update
		PriorityID:            2,
		From:                  "updated_sender@example.com",
		FromName:              "Updated Sender Name",
		To:                    "updated_recipient@example.com",
		ToName:                "Updated Recipient Name",
		ReplyTo:               "updated_replyto@example.com",
		ReplyToName:           "Updated Reply To Name",
		CC:                    "updated_cc@example.com",
		Bcc:                   "updated_bcc@example.com",
		Subject:               "Updated Test Email",
		Body:                  "This is an updated test email.",
		AttachmentFilePath:    "/path/to/updated_attachment.pdf",
		AttachmentFileName:    "updated_attachment.pdf",
		AttachedDownloadID:    bson.NewObjectID(),
		CreatedOnUtc:          time.Now().AddDate(0, 0, -7), // Created 7 days ago
		DontSendBeforeDateUtc: new(time.Time),
		SentTries:             1,
		SentOnUtc:             new(time.Time),
		EmailAccountID:        bson.NewObjectID(),
	}
	*updatedQueuedEmail.DontSendBeforeDateUtc = time.Now().AddDate(0, 0, 1) // Don't send before tomorrow
	*updatedQueuedEmail.SentOnUtc = time.Now()

	mockRepo.On("Update", mock.Anything, updatedQueuedEmail).Return(nil)

	err := usecase.Update(context.Background(), updatedQueuedEmail)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestQueuedEmailUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.QueuedEmailRepository)
	timeout := time.Duration(10)
	usecase := test.NewQueuedEmailUsecase(mockRepo, timeout)

	queuedEmailID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, queuedEmailID).Return(nil)

	err := usecase.Delete(context.Background(), queuedEmailID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestQueuedEmailUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.QueuedEmailRepository)
	timeout := time.Duration(10)
	usecase := test.NewQueuedEmailUsecase(mockRepo, timeout)
	fetchedQueuedEmails := []domain.QueuedEmail{
		{
			ID:                    bson.NewObjectID(),
			PriorityID:            1,
			From:                  "sender1@example.com",
			FromName:              "Sender One",
			To:                    "recipient1@example.com",
			ToName:                "Recipient One",
			ReplyTo:               "replyto1@example.com",
			ReplyToName:           "Reply To One",
			CC:                    "cc1@example.com",
			Bcc:                   "bcc1@example.com",
			Subject:               "First Test Email",
			Body:                  "This is the first test email.",
			AttachmentFilePath:    "/path/to/attachment1.pdf",
			AttachmentFileName:    "attachment1.pdf",
			AttachedDownloadID:    bson.NewObjectID(),
			CreatedOnUtc:          time.Now().AddDate(0, 0, -10), // Created 10 days ago
			DontSendBeforeDateUtc: nil,
			SentTries:             0,
			SentOnUtc:             nil,
			EmailAccountID:        bson.NewObjectID(),
		},
		{
			ID:                    bson.NewObjectID(),
			PriorityID:            2,
			From:                  "sender2@example.com",
			FromName:              "Sender Two",
			To:                    "recipient2@example.com",
			ToName:                "Recipient Two",
			ReplyTo:               "replyto2@example.com",
			ReplyToName:           "Reply To Two",
			CC:                    "cc2@example.com",
			Bcc:                   "bcc2@example.com",
			Subject:               "Second Test Email",
			Body:                  "This is the second test email.",
			AttachmentFilePath:    "/path/to/attachment2.pdf",
			AttachmentFileName:    "attachment2.pdf",
			AttachedDownloadID:    bson.NewObjectID(),
			CreatedOnUtc:          time.Now().AddDate(0, 0, -5), // Created 5 days ago
			DontSendBeforeDateUtc: new(time.Time),
			SentTries:             1,
			SentOnUtc:             new(time.Time),
			EmailAccountID:        bson.NewObjectID(),
		},
	}
	*fetchedQueuedEmails[1].DontSendBeforeDateUtc = time.Now().AddDate(0, 0, 2) // Don't send before 2 days from now
	*fetchedQueuedEmails[1].SentOnUtc = time.Now().AddDate(0, 0, -1)            // Sent 1 day ago

	mockRepo.On("Fetch", mock.Anything).Return(fetchedQueuedEmails, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedQueuedEmails, result)
	mockRepo.AssertExpectations(t)
}
