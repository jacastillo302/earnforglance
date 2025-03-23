package usecase_test

import (
	"context"
	mocks "earnforglance/server/domain/mocks"
	domain "earnforglance/server/domain/news"
	test "earnforglance/server/usecase/news"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestNewsSettingsUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.NewsSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewNewsSettingsUsecase(mockRepo, timeout)

	newsID := primitive.NewObjectID().Hex()

	updatedNewsSettings := domain.NewsSettings{
		ID:                                     primitive.NewObjectID(), // Existing ID of the record to update
		Enabled:                                false,
		AllowNotRegisteredUsersToLeaveComments: true,
		NotifyAboutNewNewsComments:             false,
		ShowNewsOnMainPage:                     false,
		MainPageNewsCount:                      3,
		NewsArchivePageSize:                    15,
		ShowHeaderRssUrl:                       false,
		NewsCommentsMustBeApproved:             false,
		ShowNewsCommentsPerStore:               true,
	}

	mockRepo.On("FetchByID", mock.Anything, newsID).Return(updatedNewsSettings, nil)

	result, err := usecase.FetchByID(context.Background(), newsID)

	assert.NoError(t, err)
	assert.Equal(t, updatedNewsSettings, result)
	mockRepo.AssertExpectations(t)
}

func TestNewsSettingsUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.NewsSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewNewsSettingsUsecase(mockRepo, timeout)

	newNewsSettings := &domain.NewsSettings{
		Enabled:                                true,
		AllowNotRegisteredUsersToLeaveComments: false,
		NotifyAboutNewNewsComments:             true,
		ShowNewsOnMainPage:                     true,
		MainPageNewsCount:                      5,
		NewsArchivePageSize:                    10,
		ShowHeaderRssUrl:                       true,
		NewsCommentsMustBeApproved:             true,
		ShowNewsCommentsPerStore:               false,
	}

	mockRepo.On("Create", mock.Anything, newNewsSettings).Return(nil)

	err := usecase.Create(context.Background(), newNewsSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestNewsSettingsUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.NewsSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewNewsSettingsUsecase(mockRepo, timeout)

	updatedNewsSettings := &domain.NewsSettings{
		ID:                                     primitive.NewObjectID(), // Existing ID of the record to update
		Enabled:                                false,
		AllowNotRegisteredUsersToLeaveComments: true,
		NotifyAboutNewNewsComments:             false,
		ShowNewsOnMainPage:                     false,
		MainPageNewsCount:                      3,
		NewsArchivePageSize:                    15,
		ShowHeaderRssUrl:                       false,
		NewsCommentsMustBeApproved:             false,
		ShowNewsCommentsPerStore:               true,
	}

	mockRepo.On("Update", mock.Anything, updatedNewsSettings).Return(nil)

	err := usecase.Update(context.Background(), updatedNewsSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestNewsSettingsUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.NewsSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewNewsSettingsUsecase(mockRepo, timeout)

	newsID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, newsID).Return(nil)

	err := usecase.Delete(context.Background(), newsID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestNewsSettingsUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.NewsSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewNewsSettingsUsecase(mockRepo, timeout)

	fetchedNewsSettings := []domain.NewsSettings{
		{
			ID:                                     primitive.NewObjectID(),
			Enabled:                                true,
			AllowNotRegisteredUsersToLeaveComments: false,
			NotifyAboutNewNewsComments:             true,
			ShowNewsOnMainPage:                     true,
			MainPageNewsCount:                      5,
			NewsArchivePageSize:                    10,
			ShowHeaderRssUrl:                       true,
			NewsCommentsMustBeApproved:             true,
			ShowNewsCommentsPerStore:               false,
		},
		{
			ID:                                     primitive.NewObjectID(),
			Enabled:                                false,
			AllowNotRegisteredUsersToLeaveComments: true,
			NotifyAboutNewNewsComments:             false,
			ShowNewsOnMainPage:                     false,
			MainPageNewsCount:                      3,
			NewsArchivePageSize:                    15,
			ShowHeaderRssUrl:                       false,
			NewsCommentsMustBeApproved:             false,
			ShowNewsCommentsPerStore:               true,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedNewsSettings, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedNewsSettings, result)
	mockRepo.AssertExpectations(t)
}
