package usecase_test

import (
	"context"
	mocks "earnforglance/server/domain/mocks"
	domain "earnforglance/server/domain/seo"
	test "earnforglance/server/usecase/seo"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestSeoSettingsUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.SeoSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewSeoSettingsUsecase(mockRepo, timeout)

	seoID := bson.NewObjectID().Hex()

	updatedSeoSettings := domain.SeoSettings{
		ID:                                bson.NewObjectID(), // Existing ID of the record to update
		PageTitleSeparator:                "|",
		PageTitleSeoAdjustmentID:          2,
		GenerateProductMetaDescription:    false,
		ConvertNonWesternChars:            true,
		AllowUnicodeCharsInUrls:           false,
		CanonicalUrlsEnabled:              false,
		QueryStringInCanonicalUrlsEnabled: true,
		WwwRequirementID:                  2,
		TwitterMetaTags:                   false,
		OpenGraphMetaTags:                 false,
		ReservedUrlRecordSlugs:            []string{"home", "checkout", "cart"},
		CustomHeadTags:                    "<meta name='description' content='Updated Example'>",
		MicrodataEnabled:                  false,
	}

	mockRepo.On("FetchByID", mock.Anything, seoID).Return(updatedSeoSettings, nil)

	result, err := usecase.FetchByID(context.Background(), seoID)

	assert.NoError(t, err)
	assert.Equal(t, updatedSeoSettings, result)
	mockRepo.AssertExpectations(t)
}

func TestSeoSettingsUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.SeoSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewSeoSettingsUsecase(mockRepo, timeout)

	newSeoSettings := &domain.SeoSettings{
		PageTitleSeparator:                "-",
		PageTitleSeoAdjustmentID:          1,
		GenerateProductMetaDescription:    true,
		ConvertNonWesternChars:            false,
		AllowUnicodeCharsInUrls:           true,
		CanonicalUrlsEnabled:              true,
		QueryStringInCanonicalUrlsEnabled: false,
		WwwRequirementID:                  1,
		TwitterMetaTags:                   true,
		OpenGraphMetaTags:                 true,
		ReservedUrlRecordSlugs:            []string{"admin", "login", "register"},
		CustomHeadTags:                    "<meta name='author' content='Example'>",
		MicrodataEnabled:                  true,
	}

	mockRepo.On("Create", mock.Anything, newSeoSettings).Return(nil)

	err := usecase.Create(context.Background(), newSeoSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestSeoSettingsUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.SeoSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewSeoSettingsUsecase(mockRepo, timeout)

	updatedSeoSettings := &domain.SeoSettings{
		ID:                                bson.NewObjectID(), // Existing ID of the record to update
		PageTitleSeparator:                "|",
		PageTitleSeoAdjustmentID:          2,
		GenerateProductMetaDescription:    false,
		ConvertNonWesternChars:            true,
		AllowUnicodeCharsInUrls:           false,
		CanonicalUrlsEnabled:              false,
		QueryStringInCanonicalUrlsEnabled: true,
		WwwRequirementID:                  2,
		TwitterMetaTags:                   false,
		OpenGraphMetaTags:                 false,
		ReservedUrlRecordSlugs:            []string{"home", "checkout", "cart"},
		CustomHeadTags:                    "<meta name='description' content='Updated Example'>",
		MicrodataEnabled:                  false,
	}

	mockRepo.On("Update", mock.Anything, updatedSeoSettings).Return(nil)

	err := usecase.Update(context.Background(), updatedSeoSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestSeoSettingsUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.SeoSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewSeoSettingsUsecase(mockRepo, timeout)

	seoID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, seoID).Return(nil)

	err := usecase.Delete(context.Background(), seoID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestSeoSettingsUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.SeoSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewSeoSettingsUsecase(mockRepo, timeout)

	fetchedSeoSettings := []domain.SeoSettings{
		{
			ID:                                bson.NewObjectID(),
			PageTitleSeparator:                "-",
			PageTitleSeoAdjustmentID:          1,
			GenerateProductMetaDescription:    true,
			ConvertNonWesternChars:            false,
			AllowUnicodeCharsInUrls:           true,
			CanonicalUrlsEnabled:              true,
			QueryStringInCanonicalUrlsEnabled: false,
			WwwRequirementID:                  2,
			TwitterMetaTags:                   true,
			OpenGraphMetaTags:                 true,
			ReservedUrlRecordSlugs:            []string{"admin", "login", "register"},
			CustomHeadTags:                    "<meta name='author' content='Example'>",
			MicrodataEnabled:                  true,
		},
		{
			ID:                                bson.NewObjectID(),
			PageTitleSeparator:                "|",
			PageTitleSeoAdjustmentID:          2,
			GenerateProductMetaDescription:    false,
			ConvertNonWesternChars:            true,
			AllowUnicodeCharsInUrls:           false,
			CanonicalUrlsEnabled:              false,
			QueryStringInCanonicalUrlsEnabled: true,
			WwwRequirementID:                  1,
			TwitterMetaTags:                   false,
			OpenGraphMetaTags:                 false,
			ReservedUrlRecordSlugs:            []string{"home", "checkout", "cart"},
			CustomHeadTags:                    "<meta name='description' content='Updated Example'>",
			MicrodataEnabled:                  false,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedSeoSettings, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedSeoSettings, result)
	mockRepo.AssertExpectations(t)
}
