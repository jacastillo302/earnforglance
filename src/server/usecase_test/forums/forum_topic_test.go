package usecase_test

import (
	"context"
	domain "earnforglance/server/domain/forums"
	mocks "earnforglance/server/domain/mocks"
	test "earnforglance/server/usecase/forums"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestForumTopicUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.ForumTopicRepository)
	timeout := time.Duration(10)
	usecase := test.NewForumTopicUsecase(mockRepo, timeout)

	forumTopicID := primitive.NewObjectID().Hex()

	updatedForumTopic := domain.ForumTopic{
		ID:                 primitive.NewObjectID(), // Existing ID of the record to update
		ForumID:            primitive.NewObjectID(),
		CustomerID:         primitive.NewObjectID(),
		TopicTypeID:        2,
		Subject:            "Updated Forum Topic",
		NumPosts:           10,
		Views:              100,
		LastPostID:         primitive.NewObjectID(),
		LastPostCustomerID: primitive.NewObjectID(),
		LastPostTime:       new(time.Time),
		CreatedOnUtc:       time.Now().AddDate(0, 0, -7), // Created 7 days ago
		UpdatedOnUtc:       time.Now(),
	}

	mockRepo.On("FetchByID", mock.Anything, forumTopicID).Return(updatedForumTopic, nil)

	result, err := usecase.FetchByID(context.Background(), forumTopicID)

	assert.NoError(t, err)
	assert.Equal(t, updatedForumTopic, result)
	mockRepo.AssertExpectations(t)
}

func TestForumTopicUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.ForumTopicRepository)
	timeout := time.Duration(10)
	usecase := test.NewForumTopicUsecase(mockRepo, timeout)

	newForumTopic := &domain.ForumTopic{
		ForumID:            primitive.NewObjectID(),
		CustomerID:         primitive.NewObjectID(),
		TopicTypeID:        1,
		Subject:            "Welcome to the Forum",
		NumPosts:           0,
		Views:              0,
		LastPostID:         primitive.NilObjectID,
		LastPostCustomerID: primitive.NilObjectID,
		LastPostTime:       nil,
		CreatedOnUtc:       time.Now(),
		UpdatedOnUtc:       time.Now(),
	}
	mockRepo.On("Create", mock.Anything, newForumTopic).Return(nil)

	err := usecase.Create(context.Background(), newForumTopic)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestForumTopicUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.ForumTopicRepository)
	timeout := time.Duration(10)
	usecase := test.NewForumTopicUsecase(mockRepo, timeout)

	updatedForumTopic := &domain.ForumTopic{
		ID:                 primitive.NewObjectID(), // Existing ID of the record to update
		ForumID:            primitive.NewObjectID(),
		CustomerID:         primitive.NewObjectID(),
		TopicTypeID:        2,
		Subject:            "Updated Forum Topic",
		NumPosts:           10,
		Views:              100,
		LastPostID:         primitive.NewObjectID(),
		LastPostCustomerID: primitive.NewObjectID(),
		LastPostTime:       new(time.Time),
		CreatedOnUtc:       time.Now().AddDate(0, 0, -7), // Created 7 days ago
		UpdatedOnUtc:       time.Now(),
	}
	*updatedForumTopic.LastPostTime = time.Now().AddDate(0, 0, -1) // Last post 1 day ago

	mockRepo.On("Update", mock.Anything, updatedForumTopic).Return(nil)

	err := usecase.Update(context.Background(), updatedForumTopic)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestForumTopicUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.ForumTopicRepository)
	timeout := time.Duration(10)
	usecase := test.NewForumTopicUsecase(mockRepo, timeout)

	forumTopicID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, forumTopicID).Return(nil)

	err := usecase.Delete(context.Background(), forumTopicID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestForumTopicUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.ForumTopicRepository)
	timeout := time.Duration(10)
	usecase := test.NewForumTopicUsecase(mockRepo, timeout)

	fetchedForumTopics := []domain.ForumTopic{
		{
			ID:                 primitive.NewObjectID(),
			ForumID:            primitive.NewObjectID(),
			CustomerID:         primitive.NewObjectID(),
			TopicTypeID:        1,
			Subject:            "Welcome to the Forum",
			NumPosts:           5,
			Views:              50,
			LastPostID:         primitive.NewObjectID(),
			LastPostCustomerID: primitive.NewObjectID(),
			LastPostTime:       new(time.Time),
			CreatedOnUtc:       time.Now().AddDate(0, 0, -10), // Created 10 days ago
			UpdatedOnUtc:       time.Now().AddDate(0, 0, -5),  // Updated 5 days ago

		},
		{
			ID:                 primitive.NewObjectID(),
			ForumID:            primitive.NewObjectID(),
			CustomerID:         primitive.NewObjectID(),
			TopicTypeID:        2,
			Subject:            "Forum Rules",
			NumPosts:           20,
			Views:              200,
			LastPostID:         primitive.NewObjectID(),
			LastPostCustomerID: primitive.NewObjectID(),
			LastPostTime:       new(time.Time),
			CreatedOnUtc:       time.Now().AddDate(0, 0, -30), // Created 30 days ago
			UpdatedOnUtc:       time.Now(),
		},
	}
	*fetchedForumTopics[0].LastPostTime = time.Now().AddDate(0, 0, -3) // Last post 3 days ago
	*fetchedForumTopics[1].LastPostTime = time.Now().AddDate(0, 0, -1) // Last post 1 day ago

	mockRepo.On("Fetch", mock.Anything).Return(fetchedForumTopics, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedForumTopics, result)
	mockRepo.AssertExpectations(t)
}
