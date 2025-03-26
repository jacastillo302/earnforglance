package repository_test

import (
	"context"
	"errors"
	"testing"
	"time"

	domain "earnforglance/server/domain/topics"
	repository "earnforglance/server/repository/topics"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultTopic struct {
	mock.Mock
}

func (m *MockSingleResultTopic) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.Topic); ok {
		*v.(*domain.Topic) = *result
	}
	return args.Error(1)
}

var mockItemTopic = &domain.Topic{
	ID:                        primitive.NewObjectID(), // Existing ID of the record to update
	SystemName:                "contact-us",
	IncludeInSitemap:          true,
	IncludeInTopMenu:          false,
	IncludeInFooterColumn1:    true,
	IncludeInFooterColumn2:    false,
	IncludeInFooterColumn3:    true,
	DisplayOrder:              2,
	AccessibleWhenStoreClosed: false,
	IsPasswordProtected:       true,
	Password:                  "securepassword",
	Title:                     "Contact Us",
	Body:                      "This is the Contact Us page content.",
	Published:                 true,
	TopicTemplateID:           primitive.NewObjectID(),
	MetaKeywords:              "contact, support, help",
	MetaDescription:           "Get in touch with us for support.",
	MetaTitle:                 "Contact Us - Support",
	SubjectToAcl:              true,
	LimitedToStores:           true,
	AvailableStartDateTimeUtc: new(time.Time),
	AvailableEndDateTimeUtc:   new(time.Time),
}

func TestTopicRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionTopic

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultTopic{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemTopic, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewTopicRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemTopic.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultTopic{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewTopicRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemTopic.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestTopicRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionTopic

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemTopic).Return(nil, nil).Once()

	repo := repository.NewTopicRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemTopic)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestTopicRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionTopic

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemTopic.ID}
	update := bson.M{"$set": mockItemTopic}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewTopicRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemTopic)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
