package usecase

import (
	"context"
	mocks "earnforglance/server/domain/mocks"
	domain "earnforglance/server/domain/vendors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestVendorSettingsUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.VendorSettingsRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewVendorSettingsUsecase(mockRepo, timeout)

	vendorSettingsID := primitive.NewObjectID().Hex()

	updatedVendorSettings := domain.VendorSettings{
		ID:                                           primitive.NewObjectID(), // Existing ID of the record to update
		DefaultVendorPageSizeOptions:                 "5,15,25",
		VendorsBlockItemsToDisplay:                   15,
		ShowVendorOnProductDetailsPage:               false,
		ShowVendorOnOrderDetailsPage:                 false,
		AllowCustomersToContactVendors:               false,
		AllowCustomersToApplyForVendorAccount:        false,
		TermsOfServiceEnabled:                        true,
		AllowSearchByVendor:                          false,
		AllowVendorsToEditInfo:                       false,
		NotifyStoreOwnerAboutVendorInformationChange: false,
		MaximumProductNumber:                         50,
		AllowVendorsToImportProducts:                 true,
		MaximumProductPicturesNumber:                 10,
	}

	mockRepo.On("FetchByID", mock.Anything, vendorSettingsID).Return(updatedVendorSettings, nil)

	result, err := usecase.FetchByID(context.Background(), vendorSettingsID)

	assert.NoError(t, err)
	assert.Equal(t, updatedVendorSettings, result)
	mockRepo.AssertExpectations(t)
}

func TestVendorSettingsUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.VendorSettingsRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewVendorSettingsUsecase(mockRepo, timeout)

	newVendorSettings := &domain.VendorSettings{
		DefaultVendorPageSizeOptions:                 "10,20,50",
		VendorsBlockItemsToDisplay:                   10,
		ShowVendorOnProductDetailsPage:               true,
		ShowVendorOnOrderDetailsPage:                 true,
		AllowCustomersToContactVendors:               true,
		AllowCustomersToApplyForVendorAccount:        true,
		TermsOfServiceEnabled:                        false,
		AllowSearchByVendor:                          true,
		AllowVendorsToEditInfo:                       true,
		NotifyStoreOwnerAboutVendorInformationChange: true,
		MaximumProductNumber:                         100,
		AllowVendorsToImportProducts:                 false,
		MaximumProductPicturesNumber:                 5,
	}

	mockRepo.On("Create", mock.Anything, newVendorSettings).Return(nil)

	err := usecase.Create(context.Background(), newVendorSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestVendorSettingsUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.VendorSettingsRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewVendorSettingsUsecase(mockRepo, timeout)

	updatedVendorSettings := &domain.VendorSettings{
		ID:                                           primitive.NewObjectID(), // Existing ID of the record to update
		DefaultVendorPageSizeOptions:                 "5,15,25",
		VendorsBlockItemsToDisplay:                   15,
		ShowVendorOnProductDetailsPage:               false,
		ShowVendorOnOrderDetailsPage:                 false,
		AllowCustomersToContactVendors:               false,
		AllowCustomersToApplyForVendorAccount:        false,
		TermsOfServiceEnabled:                        true,
		AllowSearchByVendor:                          false,
		AllowVendorsToEditInfo:                       false,
		NotifyStoreOwnerAboutVendorInformationChange: false,
		MaximumProductNumber:                         50,
		AllowVendorsToImportProducts:                 true,
		MaximumProductPicturesNumber:                 10,
	}

	mockRepo.On("Update", mock.Anything, updatedVendorSettings).Return(nil)

	err := usecase.Update(context.Background(), updatedVendorSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestVendorSettingsUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.VendorSettingsRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewVendorSettingsUsecase(mockRepo, timeout)

	vendorSettingsID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, vendorSettingsID).Return(nil)

	err := usecase.Delete(context.Background(), vendorSettingsID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestVendorSettingsUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.VendorSettingsRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewVendorSettingsUsecase(mockRepo, timeout)

	fetchedVendorSettings := []domain.VendorSettings{
		{
			ID:                                           primitive.NewObjectID(),
			DefaultVendorPageSizeOptions:                 "10,20,50",
			VendorsBlockItemsToDisplay:                   10,
			ShowVendorOnProductDetailsPage:               true,
			ShowVendorOnOrderDetailsPage:                 true,
			AllowCustomersToContactVendors:               true,
			AllowCustomersToApplyForVendorAccount:        true,
			TermsOfServiceEnabled:                        false,
			AllowSearchByVendor:                          true,
			AllowVendorsToEditInfo:                       true,
			NotifyStoreOwnerAboutVendorInformationChange: true,
			MaximumProductNumber:                         100,
			AllowVendorsToImportProducts:                 false,
			MaximumProductPicturesNumber:                 5,
		},
		{
			ID:                                           primitive.NewObjectID(),
			DefaultVendorPageSizeOptions:                 "5,15,25",
			VendorsBlockItemsToDisplay:                   15,
			ShowVendorOnProductDetailsPage:               false,
			ShowVendorOnOrderDetailsPage:                 false,
			AllowCustomersToContactVendors:               false,
			AllowCustomersToApplyForVendorAccount:        false,
			TermsOfServiceEnabled:                        true,
			AllowSearchByVendor:                          false,
			AllowVendorsToEditInfo:                       false,
			NotifyStoreOwnerAboutVendorInformationChange: false,
			MaximumProductNumber:                         50,
			AllowVendorsToImportProducts:                 true,
			MaximumProductPicturesNumber:                 10,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedVendorSettings, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedVendorSettings, result)
	mockRepo.AssertExpectations(t)
}
