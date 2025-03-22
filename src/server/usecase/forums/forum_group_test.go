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

func TestForumGroupUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.ForumGroupRepository)
	timeout := time.Duration(10)
	usecase := NewForumGroupUsecase(mockRepo, timeout)

	forumGroupID := primitive.NewObjectID().Hex()

	updatedForumGroup := domain.ForumGroup{
		ID:           primitive.NewObjectID(), // Existing ID of the record to update
		Name:         "Announcements",
		DisplayOrder: 2,
		CreatedOnUtc: time.Now().AddDate(0, 0, -30), // Created 30 days ago
		UpdatedOnUtc: time.Now(),
	}

	mockRepo.On("FetchByID", mock.Anything, forumGroupID).Return(updatedForumGroup, nil)

	result, err := usecase.FetchByID(context.Background(), forumGroupID)

	assert.NoError(t, err)
	assert.Equal(t, updatedForumGroup, result)
	mockRepo.AssertExpectations(t)
}

func TestForumGroupUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.ForumGroupRepository)
	timeout := time.Duration(10)
	usecase := NewForumGroupUsecase(mockRepo, timeout)

	newForumGroup := &domain.ForumGroup{
		Name:         "General Discussion",
		DisplayOrder: 1,
		CreatedOnUtc: time.Now(),
		UpdatedOnUtc: time.Now(),
	}

	mockRepo.On("Create", mock.Anything, newForumGroup).Return(nil)

	err := usecase.Create(context.Background(), newForumGroup)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestForumGroupUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.ForumGroupRepository)
	timeout := time.Duration(10)
	usecase := NewForumGroupUsecase(mockRepo, timeout)

	updatedForumGroup := &domain.ForumGroup{
		ID:           primitive.NewObjectID(), // Existing ID of the record to update
		Name:         "Announcements",
		DisplayOrder: 2,
		CreatedOnUtc: time.Now().AddDate(0, 0, -30), // Created 30 days ago
		UpdatedOnUtc: time.Now(),
	}

	mockRepo.On("Update", mock.Anything, updatedForumGroup).Return(nil)

	err := usecase.Update(context.Background(), updatedForumGroup)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestForumGroupUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.ForumGroupRepository)
	timeout := time.Duration(10)
	usecase := NewForumGroupUsecase(mockRepo, timeout)

	forumGroupID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, forumGroupID).Return(nil)

	err := usecase.Delete(context.Background(), forumGroupID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestForumGroupUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.ForumGroupRepository)
	timeout := time.Duration(10)
	usecase := NewForumGroupUsecase(mockRepo, timeout)

	fetchedForumGroups := []domain.ForumGroup{
		{
			ID:           primitive.NewObjectID(),
			Name:         "General Discussion",
			DisplayOrder: 1,
			CreatedOnUtc: time.Now().AddDate(0, 0, -10), // Created 10 days ago
			UpdatedOnUtc: time.Now(),
		},
		{
			ID:           primitive.NewObjectID(),
			Name:         "Announcements",
			DisplayOrder: 2,
			CreatedOnUtc: time.Now().AddDate(0, 0, -30), // Created 30 days ago
			UpdatedOnUtc: time.Now(),
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedForumGroups, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedForumGroups, result)
	mockRepo.AssertExpectations(t)
}
