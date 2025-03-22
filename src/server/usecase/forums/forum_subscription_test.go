package usecase

import (
	"context"
	domain "earnforglance/server/domain/forums"
	mocks "earnforglance/server/domain/mocks"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestForumSubscriptionUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.ForumSubscriptionRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewForumSubscriptionUsecase(mockRepo, timeout)

	forumSubscriptionID := primitive.NewObjectID().Hex()

	updatedForumSubscription := domain.ForumSubscription{
		ID:               primitive.NewObjectID(), // Existing ID of the record to update
		SubscriptionGuid: uuid.New(),
		CustomerID:       2,
		ForumID:          20,
		TopicID:          200,
		CreatedOnUtc:     time.Now().AddDate(0, 0, -7), // Created 7 days ago
	}

	mockRepo.On("FetchByID", mock.Anything, forumSubscriptionID).Return(updatedForumSubscription, nil)

	result, err := usecase.FetchByID(context.Background(), forumSubscriptionID)

	assert.NoError(t, err)
	assert.Equal(t, updatedForumSubscription, result)
	mockRepo.AssertExpectations(t)
}

func TestForumSubscriptionUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.ForumSubscriptionRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewForumSubscriptionUsecase(mockRepo, timeout)

	newForumSubscription := &domain.ForumSubscription{
		SubscriptionGuid: uuid.New(),
		CustomerID:       1,
		ForumID:          10,
		TopicID:          100,
		CreatedOnUtc:     time.Now(),
	}

	mockRepo.On("Create", mock.Anything, newForumSubscription).Return(nil)

	err := usecase.Create(context.Background(), newForumSubscription)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestForumSubscriptionUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.ForumSubscriptionRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewForumSubscriptionUsecase(mockRepo, timeout)

	updatedForumSubscription := &domain.ForumSubscription{
		ID:               primitive.NewObjectID(), // Existing ID of the record to update
		SubscriptionGuid: uuid.New(),
		CustomerID:       2,
		ForumID:          20,
		TopicID:          200,
		CreatedOnUtc:     time.Now().AddDate(0, 0, -7), // Created 7 days ago
	}

	mockRepo.On("Update", mock.Anything, updatedForumSubscription).Return(nil)

	err := usecase.Update(context.Background(), updatedForumSubscription)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestForumSubscriptionUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.ForumSubscriptionRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewForumSubscriptionUsecase(mockRepo, timeout)

	forumSubscriptionID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, forumSubscriptionID).Return(nil)

	err := usecase.Delete(context.Background(), forumSubscriptionID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestForumSubscriptionUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.ForumSubscriptionRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewForumSubscriptionUsecase(mockRepo, timeout)

	fetchedForumSubscriptions := []domain.ForumSubscription{
		{
			ID:               primitive.NewObjectID(),
			SubscriptionGuid: uuid.New(),
			CustomerID:       1,
			ForumID:          10,
			TopicID:          100,
			CreatedOnUtc:     time.Now().AddDate(0, 0, -10), // Created 10 days ago
		},
		{
			ID:               primitive.NewObjectID(),
			SubscriptionGuid: uuid.New(),
			CustomerID:       2,
			ForumID:          20,
			TopicID:          200,
			CreatedOnUtc:     time.Now().AddDate(0, 0, -5), // Created 5 days ago
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedForumSubscriptions, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedForumSubscriptions, result)
	mockRepo.AssertExpectations(t)
}
