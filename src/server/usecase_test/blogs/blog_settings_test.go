package usecase_test

import (
	"context"
	domain "earnforglance/server/domain/blogs"
	mocks "earnforglance/server/domain/mocks"
	test "earnforglance/server/usecase/blogs"

	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestBlogSettingsUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.BlogSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewBlogSettingsUsecase(mockRepo, timeout)

	blogID := primitive.NewObjectID().Hex()

	expectedBlogSettings := domain.BlogSettings{
		Enabled:                                true,
		PostsPageSize:                          10,
		AllowNotRegisteredUsersToLeaveComments: true,
		NotifyAboutNewBlogComments:             false,
		NumberOfTags:                           5,
		ShowHeaderRssUrl:                       true,
		BlogCommentsMustBeApproved:             false,
		ShowBlogCommentsPerStore:               true,
	}

	mockRepo.On("FetchByID", mock.Anything, blogID).Return(expectedBlogSettings, nil)

	result, err := usecase.FetchByID(context.Background(), blogID)

	assert.NoError(t, err)
	assert.Equal(t, expectedBlogSettings, result)
	mockRepo.AssertExpectations(t)
}

func TestBlogSettingsUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.BlogSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewBlogSettingsUsecase(mockRepo, timeout)

	newBlogSettings := &domain.BlogSettings{
		Enabled:                                true,
		PostsPageSize:                          10,
		AllowNotRegisteredUsersToLeaveComments: true,
		NotifyAboutNewBlogComments:             false,
		NumberOfTags:                           5,
		ShowHeaderRssUrl:                       true,
		BlogCommentsMustBeApproved:             false,
		ShowBlogCommentsPerStore:               true,
	}

	mockRepo.On("Create", mock.Anything, newBlogSettings).Return(nil)

	err := usecase.Create(context.Background(), newBlogSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestBlogSettingsUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.BlogSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewBlogSettingsUsecase(mockRepo, timeout)

	updatedBlogSettings := &domain.BlogSettings{
		ID:                                     primitive.NewObjectID(), // Existing ID of the record to update
		Enabled:                                false,
		PostsPageSize:                          20,
		AllowNotRegisteredUsersToLeaveComments: false,
		NotifyAboutNewBlogComments:             true,
		NumberOfTags:                           10,
		ShowHeaderRssUrl:                       false,
		BlogCommentsMustBeApproved:             true,
		ShowBlogCommentsPerStore:               false,
	}

	mockRepo.On("Update", mock.Anything, updatedBlogSettings).Return(nil)

	err := usecase.Update(context.Background(), updatedBlogSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestBlogSettingsUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.BlogSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewBlogSettingsUsecase(mockRepo, timeout)

	blogID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, blogID).Return(nil)

	err := usecase.Delete(context.Background(), blogID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestBlogSettingsUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.BlogSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewBlogSettingsUsecase(mockRepo, timeout)

	expectedBlogSettings := []domain.BlogSettings{
		{
			ID:                                     primitive.NewObjectID(),
			Enabled:                                true,
			PostsPageSize:                          15,
			AllowNotRegisteredUsersToLeaveComments: true,
			NotifyAboutNewBlogComments:             false,
			NumberOfTags:                           8,
			ShowHeaderRssUrl:                       true,
			BlogCommentsMustBeApproved:             false,
			ShowBlogCommentsPerStore:               true,
		},
		{
			ID:                                     primitive.NewObjectID(),
			Enabled:                                false,
			PostsPageSize:                          25,
			AllowNotRegisteredUsersToLeaveComments: false,
			NotifyAboutNewBlogComments:             true,
			NumberOfTags:                           12,
			ShowHeaderRssUrl:                       false,
			BlogCommentsMustBeApproved:             true,
			ShowBlogCommentsPerStore:               false,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(expectedBlogSettings, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, expectedBlogSettings, result)
	mockRepo.AssertExpectations(t)
}
