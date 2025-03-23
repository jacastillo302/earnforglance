package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/topics"
	"earnforglance/server/mongo/mocks"
	repository "earnforglance/server/repository/topics"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultTopicTemplate struct {
	mock.Mock
}

func (m *MockSingleResultTopicTemplate) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.TopicTemplate); ok {
		*v.(*domain.TopicTemplate) = *result
	}
	return args.Error(1)
}

var mockItemTopicTemplate = &domain.TopicTemplate{
	ID:           primitive.NewObjectID(), // Existing ID of the record to update
	Name:         "Custom Template",
	ViewPath:     "/Views/Topics/Custom.cshtml",
	DisplayOrder: 2,
}

func TestTopicTemplateRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionTopicTemplate

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultTopicTemplate{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemTopicTemplate, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewTopicTemplateRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemTopicTemplate.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultTopicTemplate{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewTopicTemplateRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemTopicTemplate.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestTopicTemplateRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionTopicTemplate

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemTopicTemplate).Return(nil, nil).Once()

	repo := repository.NewTopicTemplateRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemTopicTemplate)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestTopicTemplateRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionTopicTemplate

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemTopicTemplate.ID}
	update := bson.M{"$set": mockItemTopicTemplate}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewTopicTemplateRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemTopicTemplate)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
