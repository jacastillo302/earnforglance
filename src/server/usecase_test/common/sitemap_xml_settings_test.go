package usecase_test

import (
	"context"
	domain "earnforglance/server/domain/common"
	mocks "earnforglance/server/domain/mocks"
	test "earnforglance/server/usecase/common"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestSitemapXmlSettingsUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.SitemapXmlSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewSitemapXmlSettingsUsecase(mockRepo, timeout)

	sitemapXmlSettingsID := primitive.NewObjectID().Hex()

	updatedSitemapXmlSettings := domain.SitemapXmlSettings{
		ID:                                   primitive.NewObjectID(), // Existing ID of the record to update
		SitemapXmlEnabled:                    false,
		SitemapXmlIncludeBlogPosts:           false,
		SitemapXmlIncludeCategories:          false,
		SitemapXmlIncludeCustomUrls:          false,
		SitemapXmlIncludeManufacturers:       true,
		SitemapXmlIncludeNews:                false,
		SitemapXmlIncludeProducts:            false,
		SitemapXmlIncludeSitemapXmlSettingss: true,
		SitemapXmlIncludeTopics:              false,
		SitemapCustomUrls:                    []string{"https://example.com/updated1", "https://example.com/updated2"},
		RebuildSitemapXmlAfterHours:          48,
		SitemapBuildOperationDelay:           10,
	}

	mockRepo.On("FetchByID", mock.Anything, sitemapXmlSettingsID).Return(updatedSitemapXmlSettings, nil)

	result, err := usecase.FetchByID(context.Background(), sitemapXmlSettingsID)

	assert.NoError(t, err)
	assert.Equal(t, updatedSitemapXmlSettings, result)
	mockRepo.AssertExpectations(t)
}

func TestSitemapXmlSettingsUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.SitemapXmlSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewSitemapXmlSettingsUsecase(mockRepo, timeout)

	newSitemapXmlSettings := &domain.SitemapXmlSettings{
		SitemapXmlEnabled:                    true,
		SitemapXmlIncludeBlogPosts:           true,
		SitemapXmlIncludeCategories:          true,
		SitemapXmlIncludeCustomUrls:          true,
		SitemapXmlIncludeManufacturers:       false,
		SitemapXmlIncludeNews:                true,
		SitemapXmlIncludeProducts:            true,
		SitemapXmlIncludeSitemapXmlSettingss: false,
		SitemapXmlIncludeTopics:              true,
		SitemapCustomUrls:                    []string{"https://example.com/custom1", "https://example.com/custom2"},
		RebuildSitemapXmlAfterHours:          24,
		SitemapBuildOperationDelay:           5,
	}

	mockRepo.On("Create", mock.Anything, newSitemapXmlSettings).Return(nil)

	err := usecase.Create(context.Background(), newSitemapXmlSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestSitemapXmlSettingsUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.SitemapXmlSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewSitemapXmlSettingsUsecase(mockRepo, timeout)

	updatedSitemapXmlSettings := &domain.SitemapXmlSettings{
		ID:                                   primitive.NewObjectID(), // Existing ID of the record to update
		SitemapXmlEnabled:                    false,
		SitemapXmlIncludeBlogPosts:           false,
		SitemapXmlIncludeCategories:          false,
		SitemapXmlIncludeCustomUrls:          false,
		SitemapXmlIncludeManufacturers:       true,
		SitemapXmlIncludeNews:                false,
		SitemapXmlIncludeProducts:            false,
		SitemapXmlIncludeSitemapXmlSettingss: true,
		SitemapXmlIncludeTopics:              false,
		SitemapCustomUrls:                    []string{"https://example.com/updated1", "https://example.com/updated2"},
		RebuildSitemapXmlAfterHours:          48,
		SitemapBuildOperationDelay:           10,
	}

	mockRepo.On("Update", mock.Anything, updatedSitemapXmlSettings).Return(nil)

	err := usecase.Update(context.Background(), updatedSitemapXmlSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestSitemapXmlSettingsUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.SitemapXmlSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewSitemapXmlSettingsUsecase(mockRepo, timeout)

	sitemapXmlSettingsID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, sitemapXmlSettingsID).Return(nil)

	err := usecase.Delete(context.Background(), sitemapXmlSettingsID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestSitemapXmlSettingsUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.SitemapXmlSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewSitemapXmlSettingsUsecase(mockRepo, timeout)

	fetchedSitemapXmlSettings := []domain.SitemapXmlSettings{
		{
			ID:                                   primitive.NewObjectID(),
			SitemapXmlEnabled:                    true,
			SitemapXmlIncludeBlogPosts:           true,
			SitemapXmlIncludeCategories:          true,
			SitemapXmlIncludeCustomUrls:          true,
			SitemapXmlIncludeManufacturers:       false,
			SitemapXmlIncludeNews:                true,
			SitemapXmlIncludeProducts:            true,
			SitemapXmlIncludeSitemapXmlSettingss: false,
			SitemapXmlIncludeTopics:              true,
			SitemapCustomUrls:                    []string{"https://example.com/custom1", "https://example.com/custom2"},
			RebuildSitemapXmlAfterHours:          24,
			SitemapBuildOperationDelay:           5,
		},
		{
			ID:                                   primitive.NewObjectID(),
			SitemapXmlEnabled:                    false,
			SitemapXmlIncludeBlogPosts:           false,
			SitemapXmlIncludeCategories:          false,
			SitemapXmlIncludeCustomUrls:          false,
			SitemapXmlIncludeManufacturers:       true,
			SitemapXmlIncludeNews:                false,
			SitemapXmlIncludeProducts:            false,
			SitemapXmlIncludeSitemapXmlSettingss: true,
			SitemapXmlIncludeTopics:              false,
			SitemapCustomUrls:                    []string{"https://example.com/updated1", "https://example.com/updated2"},
			RebuildSitemapXmlAfterHours:          48,
			SitemapBuildOperationDelay:           10,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedSitemapXmlSettings, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedSitemapXmlSettings, result)
	mockRepo.AssertExpectations(t)
}
