package usecase_test

import (
	"context"
	mocks "earnforglance/server/domain/mocks"
	domain "earnforglance/server/domain/topics"
	test "earnforglance/server/usecase/topics"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestTopicTemplateUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.TopicTemplateRepository)
	timeout := time.Duration(10)
	usecase := test.NewTopicTemplateUsecase(mockRepo, timeout)

	topicTemplateID := bson.NewObjectID().Hex()

	updatedTopicTemplate := domain.TopicTemplate{
		ID:           bson.NewObjectID(), // Existing ID of the record to update
		Name:         "Custom Template",
		ViewPath:     "/Views/Topics/Custom.cshtml",
		DisplayOrder: 2,
	}

	mockRepo.On("FetchByID", mock.Anything, topicTemplateID).Return(updatedTopicTemplate, nil)

	result, err := usecase.FetchByID(context.Background(), topicTemplateID)

	assert.NoError(t, err)
	assert.Equal(t, updatedTopicTemplate, result)
	mockRepo.AssertExpectations(t)
}

func TestTopicTemplateUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.TopicTemplateRepository)
	timeout := time.Duration(10)
	usecase := test.NewTopicTemplateUsecase(mockRepo, timeout)

	newTopicTemplate := &domain.TopicTemplate{
		Name:         "Default Template",
		ViewPath:     "/Views/Topics/Default.cshtml",
		DisplayOrder: 1,
	}

	mockRepo.On("Create", mock.Anything, newTopicTemplate).Return(nil)

	err := usecase.Create(context.Background(), newTopicTemplate)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestTopicTemplateUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.TopicTemplateRepository)
	timeout := time.Duration(10)
	usecase := test.NewTopicTemplateUsecase(mockRepo, timeout)

	updatedTopicTemplate := &domain.TopicTemplate{
		ID:           bson.NewObjectID(), // Existing ID of the record to update
		Name:         "Custom Template",
		ViewPath:     "/Views/Topics/Custom.cshtml",
		DisplayOrder: 2,
	}

	mockRepo.On("Update", mock.Anything, updatedTopicTemplate).Return(nil)

	err := usecase.Update(context.Background(), updatedTopicTemplate)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestTopicTemplateUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.TopicTemplateRepository)
	timeout := time.Duration(10)
	usecase := test.NewTopicTemplateUsecase(mockRepo, timeout)

	topicTemplateID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, topicTemplateID).Return(nil)

	err := usecase.Delete(context.Background(), topicTemplateID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestTopicTemplateUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.TopicTemplateRepository)
	timeout := time.Duration(10)
	usecase := test.NewTopicTemplateUsecase(mockRepo, timeout)

	fetchedTopicTemplates := []domain.TopicTemplate{
		{
			ID:           bson.NewObjectID(),
			Name:         "Default Template",
			ViewPath:     "/Views/Topics/Default.cshtml",
			DisplayOrder: 1,
		},
		{
			ID:           bson.NewObjectID(),
			Name:         "Custom Template",
			ViewPath:     "/Views/Topics/Custom.cshtml",
			DisplayOrder: 2,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedTopicTemplates, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedTopicTemplates, result)
	mockRepo.AssertExpectations(t)
}
