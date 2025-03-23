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

func TestCommonSettingsUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.CommonSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewCommonSettingsUsecase(mockRepo, timeout)

	commonID := primitive.NewObjectID().Hex()

	updatedCommonSettings := domain.CommonSettings{
		ID:                               primitive.NewObjectID(), // Existing ID of the record to update
		SubjectFieldOnContactUsForm:      false,
		UseSystemEmailForContactUsForm:   true,
		DisplayJavaScriptDisabledWarning: false,
		Log404Errors:                     false,
		BreadcrumbDelimiter:              " / ",
		IgnoreLogWordlist:                []string{"info", "trace"},
		ClearLogOlderThanDays:            60,
		BbcodeEditorOpenLinksInNewWindow: false,
		PopupForTermsOfServiceLinks:      true,
		JqueryMigrateScriptLoggingActive: false,
		UseResponseCompression:           false,
		FaviconAndAppIconsHeadCode:       "<link rel='icon' href='/newfavicon.ico'>",
		EnableHtmlMinification:           true,
		RestartTimeout:                   new(int),
		HeaderCustomHtml:                 "<header>Updated Header</header>",
		FooterCustomHtml:                 "<footer>Updated Footer</footer>",
	}

	mockRepo.On("FetchByID", mock.Anything, commonID).Return(updatedCommonSettings, nil)

	result, err := usecase.FetchByID(context.Background(), commonID)

	assert.NoError(t, err)
	assert.Equal(t, updatedCommonSettings, result)
	mockRepo.AssertExpectations(t)
}

func TestCommonSettingsUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.CommonSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewCommonSettingsUsecase(mockRepo, timeout)

	newCommonSettings := &domain.CommonSettings{
		SubjectFieldOnContactUsForm:      true,
		UseSystemEmailForContactUsForm:   false,
		DisplayJavaScriptDisabledWarning: true,
		Log404Errors:                     true,
		BreadcrumbDelimiter:              " > ",
		IgnoreLogWordlist:                []string{"error", "debug"},
		ClearLogOlderThanDays:            30,
		BbcodeEditorOpenLinksInNewWindow: true,
		PopupForTermsOfServiceLinks:      false,
		JqueryMigrateScriptLoggingActive: true,
		UseResponseCompression:           true,
		FaviconAndAppIconsHeadCode:       "<link rel='icon' href='/favicon.ico'>",
		EnableHtmlMinification:           false,
		RestartTimeout:                   nil,
		HeaderCustomHtml:                 "<header>Custom Header</header>",
		FooterCustomHtml:                 "<footer>Custom Footer</footer>",
	}

	mockRepo.On("Create", mock.Anything, newCommonSettings).Return(nil)

	err := usecase.Create(context.Background(), newCommonSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCommonSettingsUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.CommonSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewCommonSettingsUsecase(mockRepo, timeout)

	updatedCommonSettings := &domain.CommonSettings{
		ID:                               primitive.NewObjectID(), // Existing ID of the record to update
		SubjectFieldOnContactUsForm:      false,
		UseSystemEmailForContactUsForm:   true,
		DisplayJavaScriptDisabledWarning: false,
		Log404Errors:                     false,
		BreadcrumbDelimiter:              " / ",
		IgnoreLogWordlist:                []string{"info", "trace"},
		ClearLogOlderThanDays:            60,
		BbcodeEditorOpenLinksInNewWindow: false,
		PopupForTermsOfServiceLinks:      true,
		JqueryMigrateScriptLoggingActive: false,
		UseResponseCompression:           false,
		FaviconAndAppIconsHeadCode:       "<link rel='icon' href='/newfavicon.ico'>",
		EnableHtmlMinification:           true,
		RestartTimeout:                   new(int),
		HeaderCustomHtml:                 "<header>Updated Header</header>",
		FooterCustomHtml:                 "<footer>Updated Footer</footer>",
	}
	*updatedCommonSettings.RestartTimeout = 120

	mockRepo.On("Update", mock.Anything, updatedCommonSettings).Return(nil)

	err := usecase.Update(context.Background(), updatedCommonSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCommonSettingsUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.CommonSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewCommonSettingsUsecase(mockRepo, timeout)

	commonID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, commonID).Return(nil)

	err := usecase.Delete(context.Background(), commonID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCommonSettingsUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.CommonSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewCommonSettingsUsecase(mockRepo, timeout)

	fetchedCommonSettings := []domain.CommonSettings{
		{
			ID:                               primitive.NewObjectID(),
			SubjectFieldOnContactUsForm:      true,
			UseSystemEmailForContactUsForm:   false,
			DisplayJavaScriptDisabledWarning: true,
			Log404Errors:                     true,
			BreadcrumbDelimiter:              " > ",
			IgnoreLogWordlist:                []string{"error", "debug"},
			ClearLogOlderThanDays:            30,
			BbcodeEditorOpenLinksInNewWindow: true,
			PopupForTermsOfServiceLinks:      false,
			JqueryMigrateScriptLoggingActive: true,
			UseResponseCompression:           true,
			FaviconAndAppIconsHeadCode:       "<link rel='icon' href='/favicon.ico'>",
			EnableHtmlMinification:           false,
			RestartTimeout:                   nil,
			HeaderCustomHtml:                 "<header>Custom Header</header>",
			FooterCustomHtml:                 "<footer>Custom Footer</footer>",
		},
		{
			ID:                               primitive.NewObjectID(),
			SubjectFieldOnContactUsForm:      false,
			UseSystemEmailForContactUsForm:   true,
			DisplayJavaScriptDisabledWarning: false,
			Log404Errors:                     false,
			BreadcrumbDelimiter:              " / ",
			IgnoreLogWordlist:                []string{"info", "trace"},
			ClearLogOlderThanDays:            60,
			BbcodeEditorOpenLinksInNewWindow: false,
			PopupForTermsOfServiceLinks:      true,
			JqueryMigrateScriptLoggingActive: false,
			UseResponseCompression:           false,
			FaviconAndAppIconsHeadCode:       "<link rel='icon' href='/newfavicon.ico'>",
			EnableHtmlMinification:           true,
			RestartTimeout:                   new(int),
			HeaderCustomHtml:                 "<header>Updated Header</header>",
			FooterCustomHtml:                 "<footer>Updated Footer</footer>",
		},
	}
	*fetchedCommonSettings[1].RestartTimeout = 120

	mockRepo.On("Fetch", mock.Anything).Return(fetchedCommonSettings, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedCommonSettings, result)
	mockRepo.AssertExpectations(t)
}
