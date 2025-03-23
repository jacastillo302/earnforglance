package repository_test

import (
	"context"
	"errors"
	"testing"
	"time"

	domain "earnforglance/server/domain/messages"
	"earnforglance/server/mongo/mocks"
	repository "earnforglance/server/repository/messages"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultQueuedEmail struct {
	mock.Mock
}

func (m *MockSingleResultQueuedEmail) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.QueuedEmail); ok {
		*v.(*domain.QueuedEmail) = *result
	}
	return args.Error(1)
}

var mockItemQueuedEmail = &domain.QueuedEmail{
	ID:                    primitive.NewObjectID(), // Existing ID of the record to update
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
	AttachedDownloadID:    primitive.NewObjectID(),
	CreatedOnUtc:          time.Now().AddDate(0, 0, -7), // Created 7 days ago
	DontSendBeforeDateUtc: new(time.Time),
	SentTries:             1,
	SentOnUtc:             new(time.Time),
	EmailAccountID:        primitive.NewObjectID(),
	Priority:              3,
}

func TestQueuedEmailRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionQueuedEmail

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultQueuedEmail{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemQueuedEmail, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewQueuedEmailRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemQueuedEmail.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultQueuedEmail{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewQueuedEmailRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemQueuedEmail.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestQueuedEmailRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionQueuedEmail

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemQueuedEmail).Return(nil, nil).Once()

	repo := repository.NewQueuedEmailRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemQueuedEmail)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestQueuedEmailRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionQueuedEmail

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemQueuedEmail.ID}
	update := bson.M{"$set": mockItemQueuedEmail}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewQueuedEmailRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemQueuedEmail)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
