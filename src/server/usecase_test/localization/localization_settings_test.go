package usecase_test

import (
	"context"
	domain "earnforglance/server/domain/localization"
	mocks "earnforglance/server/domain/mocks"
	test "earnforglance/server/usecase/localization"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestLocalizationSettingsUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.LocalizationSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewLocalizationSettingsUsecase(mockRepo, timeout)

	localizationID := bson.NewObjectID().Hex()

	updatedLocalizationSettings := domain.LocalizationSettings{
		ID:                                  bson.NewObjectID(), // Existing ID of the record to update
		DefaultAdminLanguageID:              bson.NewObjectID(),
		UseImagesForLanguageSelection:       false,
		SeoFriendlyUrlsForLanguagesEnabled:  false,
		AutomaticallyDetectLanguage:         true,
		LoadAllLocaleRecordsOnStartup:       false,
		LoadAllLocalizedPropertiesOnStartup: true,
		LoadAllUrlRecordsOnStartup:          false,
		IgnoreRtlPropertyForAdminArea:       true,
	}

	mockRepo.On("FetchByID", mock.Anything, localizationID).Return(updatedLocalizationSettings, nil)

	result, err := usecase.FetchByID(context.Background(), localizationID)

	assert.NoError(t, err)
	assert.Equal(t, updatedLocalizationSettings, result)
	mockRepo.AssertExpectations(t)
}

func TestLocalizationSettingsUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.LocalizationSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewLocalizationSettingsUsecase(mockRepo, timeout)

	newLocalizationSettings := &domain.LocalizationSettings{
		DefaultAdminLanguageID:              bson.NewObjectID(),
		UseImagesForLanguageSelection:       true,
		SeoFriendlyUrlsForLanguagesEnabled:  true,
		AutomaticallyDetectLanguage:         false,
		LoadAllLocaleRecordsOnStartup:       true,
		LoadAllLocalizedPropertiesOnStartup: false,
		LoadAllUrlRecordsOnStartup:          true,
		IgnoreRtlPropertyForAdminArea:       false,
	}
	mockRepo.On("Create", mock.Anything, newLocalizationSettings).Return(nil)

	err := usecase.Create(context.Background(), newLocalizationSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestLocalizationSettingsUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.LocalizationSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewLocalizationSettingsUsecase(mockRepo, timeout)
	updatedLocalizationSettings := &domain.LocalizationSettings{
		ID:                                  bson.NewObjectID(), // Existing ID of the record to update
		DefaultAdminLanguageID:              bson.NewObjectID(),
		UseImagesForLanguageSelection:       false,
		SeoFriendlyUrlsForLanguagesEnabled:  false,
		AutomaticallyDetectLanguage:         true,
		LoadAllLocaleRecordsOnStartup:       false,
		LoadAllLocalizedPropertiesOnStartup: true,
		LoadAllUrlRecordsOnStartup:          false,
		IgnoreRtlPropertyForAdminArea:       true,
	}

	mockRepo.On("Update", mock.Anything, updatedLocalizationSettings).Return(nil)

	err := usecase.Update(context.Background(), updatedLocalizationSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestLocalizationSettingsUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.LocalizationSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewLocalizationSettingsUsecase(mockRepo, timeout)

	localizationID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, localizationID).Return(nil)

	err := usecase.Delete(context.Background(), localizationID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestLocalizationSettingsUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.LocalizationSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewLocalizationSettingsUsecase(mockRepo, timeout)

	fetchedLocalizationSettings := []domain.LocalizationSettings{
		{
			ID:                                  bson.NewObjectID(),
			DefaultAdminLanguageID:              bson.NewObjectID(),
			UseImagesForLanguageSelection:       true,
			SeoFriendlyUrlsForLanguagesEnabled:  true,
			AutomaticallyDetectLanguage:         false,
			LoadAllLocaleRecordsOnStartup:       true,
			LoadAllLocalizedPropertiesOnStartup: false,
			LoadAllUrlRecordsOnStartup:          true,
			IgnoreRtlPropertyForAdminArea:       false,
		},
		{
			ID:                                  bson.NewObjectID(),
			DefaultAdminLanguageID:              bson.NewObjectID(),
			UseImagesForLanguageSelection:       false,
			SeoFriendlyUrlsForLanguagesEnabled:  false,
			AutomaticallyDetectLanguage:         true,
			LoadAllLocaleRecordsOnStartup:       false,
			LoadAllLocalizedPropertiesOnStartup: true,
			LoadAllUrlRecordsOnStartup:          false,
			IgnoreRtlPropertyForAdminArea:       true,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedLocalizationSettings, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedLocalizationSettings, result)
	mockRepo.AssertExpectations(t)
}
