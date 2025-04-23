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
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestSitemapSettingsUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.SitemapSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewSitemapSettingsUsecase(mockRepo, timeout)

	sitemapSettingsID := bson.NewObjectID().Hex()

	updatedSitemapSettings := domain.SitemapSettings{
		ID:                             bson.NewObjectID(), // Existing ID of the record to update
		SitemapEnabled:                 false,
		SitemapPageSize:                50,
		SitemapIncludeBlogPosts:        false,
		SitemapIncludeCategories:       false,
		SitemapIncludeManufacturers:    true,
		SitemapIncludeNews:             false,
		SitemapIncludeProducts:         false,
		SitemapIncludeSitemapSettingss: true,
		SitemapIncludeTopics:           false,
	}

	mockRepo.On("FetchByID", mock.Anything, sitemapSettingsID).Return(updatedSitemapSettings, nil)

	result, err := usecase.FetchByID(context.Background(), sitemapSettingsID)

	assert.NoError(t, err)
	assert.Equal(t, updatedSitemapSettings, result)
	mockRepo.AssertExpectations(t)
}

func TestSitemapSettingsUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.SitemapSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewSitemapSettingsUsecase(mockRepo, timeout)

	newSitemapSettings := &domain.SitemapSettings{
		SitemapEnabled:                 true,
		SitemapPageSize:                100,
		SitemapIncludeBlogPosts:        true,
		SitemapIncludeCategories:       true,
		SitemapIncludeManufacturers:    false,
		SitemapIncludeNews:             true,
		SitemapIncludeProducts:         true,
		SitemapIncludeSitemapSettingss: false,
		SitemapIncludeTopics:           true,
	}

	mockRepo.On("Create", mock.Anything, newSitemapSettings).Return(nil)

	err := usecase.Create(context.Background(), newSitemapSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestSitemapSettingsUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.SitemapSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewSitemapSettingsUsecase(mockRepo, timeout)

	updatedSitemapSettings := &domain.SitemapSettings{
		ID:                             bson.NewObjectID(), // Existing ID of the record to update
		SitemapEnabled:                 false,
		SitemapPageSize:                50,
		SitemapIncludeBlogPosts:        false,
		SitemapIncludeCategories:       false,
		SitemapIncludeManufacturers:    true,
		SitemapIncludeNews:             false,
		SitemapIncludeProducts:         false,
		SitemapIncludeSitemapSettingss: true,
		SitemapIncludeTopics:           false,
	}

	mockRepo.On("Update", mock.Anything, updatedSitemapSettings).Return(nil)

	err := usecase.Update(context.Background(), updatedSitemapSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestSitemapSettingsUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.SitemapSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewSitemapSettingsUsecase(mockRepo, timeout)

	sitemapSettingsID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, sitemapSettingsID).Return(nil)

	err := usecase.Delete(context.Background(), sitemapSettingsID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestSitemapSettingsUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.SitemapSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewSitemapSettingsUsecase(mockRepo, timeout)

	fetchedSitemapSettings := []domain.SitemapSettings{
		{
			ID:                             bson.NewObjectID(),
			SitemapEnabled:                 true,
			SitemapPageSize:                100,
			SitemapIncludeBlogPosts:        true,
			SitemapIncludeCategories:       true,
			SitemapIncludeManufacturers:    false,
			SitemapIncludeNews:             true,
			SitemapIncludeProducts:         true,
			SitemapIncludeSitemapSettingss: false,
			SitemapIncludeTopics:           true,
		},
		{
			ID:                             bson.NewObjectID(),
			SitemapEnabled:                 false,
			SitemapPageSize:                50,
			SitemapIncludeBlogPosts:        false,
			SitemapIncludeCategories:       false,
			SitemapIncludeManufacturers:    true,
			SitemapIncludeNews:             false,
			SitemapIncludeProducts:         false,
			SitemapIncludeSitemapSettingss: true,
			SitemapIncludeTopics:           false,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedSitemapSettings, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedSitemapSettings, result)
	mockRepo.AssertExpectations(t)
}
