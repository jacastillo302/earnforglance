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

func TestAdminAreaSettingsUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.AdminAreaSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewAdminAreaSettingsUsecase(mockRepo, timeout)

	adminAreaSettingsID := primitive.NewObjectID().Hex()

	updatedAdminAreaSettings := domain.AdminAreaSettings{
		ID:                              primitive.NewObjectID(), // Existing ID of the record to update
		DefaultGridPageSize:             30,
		ProductsBulkEditGridPageSize:    100,
		PopupGridPageSize:               15,
		GridPageSizes:                   "15,30,60,120",
		RichEditorAdditionalSettings:    "font-family, background-color",
		RichEditorAllowJavaScript:       true,
		RichEditorAllowStyleTag:         false,
		UseRichEditorForCustomerEmails:  false,
		UseRichEditorInMessageTemplates: false,
		HideAdvertisementsOnAdminArea:   false,
		CheckLicense:                    false,
		LastNewsTitleAdminArea:          "Updated News",
		UseIsoDateFormatInJsonResult:    false,
		ShowDocumentationReferenceLinks: true,
		UseStickyHeaderLayout:           false,
		MinimumDropdownItemsForSearch:   10,
	}

	mockRepo.On("FetchByID", mock.Anything, adminAreaSettingsID).Return(updatedAdminAreaSettings, nil)

	result, err := usecase.FetchByID(context.Background(), adminAreaSettingsID)

	assert.NoError(t, err)
	assert.Equal(t, updatedAdminAreaSettings, result)
	mockRepo.AssertExpectations(t)
}

func TestAdminAreaSettingsUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.AdminAreaSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewAdminAreaSettingsUsecase(mockRepo, timeout)

	newAdminAreaSettings := &domain.AdminAreaSettings{
		DefaultGridPageSize:             20,
		ProductsBulkEditGridPageSize:    50,
		PopupGridPageSize:               10,
		GridPageSizes:                   "10,20,50,100",
		RichEditorAdditionalSettings:    "font-size, color",
		RichEditorAllowJavaScript:       false,
		RichEditorAllowStyleTag:         true,
		UseRichEditorForCustomerEmails:  true,
		UseRichEditorInMessageTemplates: true,
		HideAdvertisementsOnAdminArea:   true,
		CheckLicense:                    true,
		LastNewsTitleAdminArea:          "Latest Updates",
		UseIsoDateFormatInJsonResult:    true,
		ShowDocumentationReferenceLinks: false,
		UseStickyHeaderLayout:           true,
		MinimumDropdownItemsForSearch:   5,
	}

	mockRepo.On("Create", mock.Anything, newAdminAreaSettings).Return(nil)

	err := usecase.Create(context.Background(), newAdminAreaSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestAdminAreaSettingsUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.AdminAreaSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewAdminAreaSettingsUsecase(mockRepo, timeout)

	updatedAdminAreaSettings := &domain.AdminAreaSettings{
		ID:                              primitive.NewObjectID(), // Existing ID of the record to update
		DefaultGridPageSize:             30,
		ProductsBulkEditGridPageSize:    100,
		PopupGridPageSize:               15,
		GridPageSizes:                   "15,30,60,120",
		RichEditorAdditionalSettings:    "font-family, background-color",
		RichEditorAllowJavaScript:       true,
		RichEditorAllowStyleTag:         false,
		UseRichEditorForCustomerEmails:  false,
		UseRichEditorInMessageTemplates: false,
		HideAdvertisementsOnAdminArea:   false,
		CheckLicense:                    false,
		LastNewsTitleAdminArea:          "Updated News",
		UseIsoDateFormatInJsonResult:    false,
		ShowDocumentationReferenceLinks: true,
		UseStickyHeaderLayout:           false,
		MinimumDropdownItemsForSearch:   10,
	}

	mockRepo.On("Update", mock.Anything, updatedAdminAreaSettings).Return(nil)

	err := usecase.Update(context.Background(), updatedAdminAreaSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestAdminAreaSettingsUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.AdminAreaSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewAdminAreaSettingsUsecase(mockRepo, timeout)

	adminAreaSettingsID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, adminAreaSettingsID).Return(nil)

	err := usecase.Delete(context.Background(), adminAreaSettingsID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestAdminAreaSettingsUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.AdminAreaSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewAdminAreaSettingsUsecase(mockRepo, timeout)

	fetchedAdminAreaSettings := []domain.AdminAreaSettings{
		{
			ID:                              primitive.NewObjectID(),
			DefaultGridPageSize:             20,
			ProductsBulkEditGridPageSize:    50,
			PopupGridPageSize:               10,
			GridPageSizes:                   "10,20,50,100",
			RichEditorAdditionalSettings:    "font-size, color",
			RichEditorAllowJavaScript:       false,
			RichEditorAllowStyleTag:         true,
			UseRichEditorForCustomerEmails:  true,
			UseRichEditorInMessageTemplates: true,
			HideAdvertisementsOnAdminArea:   true,
			CheckLicense:                    true,
			LastNewsTitleAdminArea:          "Latest Updates",
			UseIsoDateFormatInJsonResult:    true,
			ShowDocumentationReferenceLinks: false,
			UseStickyHeaderLayout:           true,
			MinimumDropdownItemsForSearch:   5,
		},
		{
			ID:                              primitive.NewObjectID(),
			DefaultGridPageSize:             30,
			ProductsBulkEditGridPageSize:    100,
			PopupGridPageSize:               15,
			GridPageSizes:                   "15,30,60,120",
			RichEditorAdditionalSettings:    "font-family, background-color",
			RichEditorAllowJavaScript:       true,
			RichEditorAllowStyleTag:         false,
			UseRichEditorForCustomerEmails:  false,
			UseRichEditorInMessageTemplates: false,
			HideAdvertisementsOnAdminArea:   false,
			CheckLicense:                    false,
			LastNewsTitleAdminArea:          "Updated News",
			UseIsoDateFormatInJsonResult:    false,
			ShowDocumentationReferenceLinks: true,
			UseStickyHeaderLayout:           false,
			MinimumDropdownItemsForSearch:   10,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedAdminAreaSettings, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedAdminAreaSettings, result)
	mockRepo.AssertExpectations(t)
}
