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

func TestForumUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.ForumRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewForumUsecase(mockRepo, timeout)

	forumID := primitive.NewObjectID().Hex()

	updatedForum := domain.Forum{
		ID:                 primitive.NewObjectID(), // Existing ID of the record to update
		ForumGroupID:       primitive.NewObjectID(),
		Name:               "Announcements",
		Description:        "Official announcements and updates.",
		NumTopics:          10,
		NumPosts:           50,
		LastTopicID:        primitive.NewObjectID(),
		LastPostID:         primitive.NewObjectID(),
		LastPostCustomerID: primitive.NewObjectID(),
		LastPostTime:       new(time.Time),
		DisplayOrder:       2,
		CreatedOnUtc:       time.Now().AddDate(0, 0, -30), // Created 30 days ago
		UpdatedOnUtc:       time.Now(),
	}
	*updatedForum.LastPostTime = time.Now().AddDate(0, 0, -1) // Last post 1 day ago

	mockRepo.On("FetchByID", mock.Anything, forumID).Return(updatedForum, nil)

	result, err := usecase.FetchByID(context.Background(), forumID)

	assert.NoError(t, err)
	assert.Equal(t, updatedForum, result)
	mockRepo.AssertExpectations(t)
}

func TestForumUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.ForumRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewForumUsecase(mockRepo, timeout)

	newForum := &domain.Forum{
		ForumGroupID:       primitive.NewObjectID(),
		Name:               "General Discussion",
		Description:        "A place for general topics and discussions.",
		NumTopics:          0,
		NumPosts:           0,
		LastTopicID:        primitive.NilObjectID,
		LastPostID:         primitive.NilObjectID,
		LastPostCustomerID: primitive.NilObjectID,
		LastPostTime:       nil,
		DisplayOrder:       1,
		CreatedOnUtc:       time.Now(),
		UpdatedOnUtc:       time.Now(),
	}

	mockRepo.On("Create", mock.Anything, newForum).Return(nil)

	err := usecase.Create(context.Background(), newForum)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestForumUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.ForumRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewForumUsecase(mockRepo, timeout)

	updatedForum := &domain.Forum{
		ID:                 primitive.NewObjectID(), // Existing ID of the record to update
		ForumGroupID:       primitive.NewObjectID(),
		Name:               "Announcements",
		Description:        "Official announcements and updates.",
		NumTopics:          10,
		NumPosts:           50,
		LastTopicID:        primitive.NewObjectID(),
		LastPostID:         primitive.NewObjectID(),
		LastPostCustomerID: primitive.NewObjectID(),
		LastPostTime:       new(time.Time),
		DisplayOrder:       2,
		CreatedOnUtc:       time.Now().AddDate(0, 0, -30), // Created 30 days ago
		UpdatedOnUtc:       time.Now(),
	}
	*updatedForum.LastPostTime = time.Now().AddDate(0, 0, -1) // Last post 1 day ago

	mockRepo.On("Update", mock.Anything, updatedForum).Return(nil)

	err := usecase.Update(context.Background(), updatedForum)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestForumUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.ForumRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewForumUsecase(mockRepo, timeout)

	forumID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, forumID).Return(nil)

	err := usecase.Delete(context.Background(), forumID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestForumUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.ForumRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewForumUsecase(mockRepo, timeout)

	fetchedForums := []domain.Forum{
		{
			ID:                 primitive.NewObjectID(),
			ForumGroupID:       primitive.NewObjectID(),
			Name:               "General Discussion",
			Description:        "A place for general topics and discussions.",
			NumTopics:          5,
			NumPosts:           20,
			LastTopicID:        primitive.NewObjectID(),
			LastPostID:         primitive.NewObjectID(),
			LastPostCustomerID: primitive.NewObjectID(),
			LastPostTime:       new(time.Time),
			DisplayOrder:       1,
			CreatedOnUtc:       time.Now().AddDate(0, 0, -10), // Created 10 days ago
			UpdatedOnUtc:       time.Now().AddDate(0, 0, -5),  // Updated 5 days ago
		},
		{
			ID:                 primitive.NewObjectID(),
			ForumGroupID:       primitive.NewObjectID(),
			Name:               "Announcements",
			Description:        "Official announcements and updates.",
			NumTopics:          10,
			NumPosts:           50,
			LastTopicID:        primitive.NewObjectID(),
			LastPostID:         primitive.NewObjectID(),
			LastPostCustomerID: primitive.NewObjectID(),
			LastPostTime:       new(time.Time),
			DisplayOrder:       2,
			CreatedOnUtc:       time.Now().AddDate(0, 0, -30), // Created 30 days ago
			UpdatedOnUtc:       time.Now(),
		},
	}
	*fetchedForums[0].LastPostTime = time.Now().AddDate(0, 0, -3) // Last post 3 days ago
	*fetchedForums[1].LastPostTime = time.Now().AddDate(0, 0, -1) // Last post 1 day ago

	mockRepo.On("Fetch", mock.Anything).Return(fetchedForums, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedForums, result)
	mockRepo.AssertExpectations(t)
}
