package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/messages"
	"earnforglance/server/mongo/mocks"
	repository "earnforglance/server/repository/messages"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultMessageTemplate struct {
	mock.Mock
}

func (m *MockSingleResultMessageTemplate) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.MessageTemplate); ok {
		*v.(*domain.MessageTemplate) = *result
	}
	return args.Error(1)
}

var mockItemMessageTemplate = &domain.MessageTemplate{
	ID:                 primitive.NewObjectID(), // Existing ID of the record to update
	Name:               "Updated Order Confirmation",
	BccEmailAddresses:  "updated_bcc@example.com",
	Subject:            "Updated Order Confirmation Subject",
	Body:               "Your order has been updated. Your new order number is #67890.",
	IsActive:           false,
	DelayBeforeSend:    new(int),
	DelayPeriodID:      2,
	AttachedDownloadID: primitive.NewObjectID(),
	AllowDirectReply:   false,
	EmailAccountID:     primitive.NewObjectID(),
	LimitedToStores:    true,
	DelayPeriod:        1,
}

func TestMessageTemplateRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionMessageTemplate

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultMessageTemplate{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemMessageTemplate, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewMessageTemplateRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemMessageTemplate.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultMessageTemplate{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewMessageTemplateRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemMessageTemplate.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestMessageTemplateRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionMessageTemplate

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemMessageTemplate).Return(nil, nil).Once()

	repo := repository.NewMessageTemplateRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemMessageTemplate)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestMessageTemplateRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionMessageTemplate

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemMessageTemplate.ID}
	update := bson.M{"$set": mockItemMessageTemplate}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewMessageTemplateRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemMessageTemplate)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
