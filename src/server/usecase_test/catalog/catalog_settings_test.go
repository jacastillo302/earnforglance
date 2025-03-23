package usecase_test

import (
	"context"
	domain "earnforglance/server/domain/catalog"
	mocks "earnforglance/server/domain/mocks"
	test "earnforglance/server/usecase/catalog"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCatalogSettingsUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.CatalogSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewCatalogSettingsUsecase(mockRepo, timeout)

	catalogID := primitive.NewObjectID().Hex()

	expectedCatalogSettings := domain.CatalogSettings{
		ID:                              primitive.NewObjectID(), // Existing ID of the record to update
		AllowViewUnpublishedProductPage: false,
		DisplayDiscontinuedMessageForUnpublishedProducts: true,
		PublishBackProductWhenCancellingOrders:           false,
		ShowSkuOnProductDetailsPage:                      false,
		ShowSkuOnCatalogPages:                            true,
		ShowManufacturerPartNumber:                       false,
		ShowGtin:                                         true,
		ShowFreeShippingNotification:                     false,
		ShowShortDescriptionOnCatalogPages:               false,
		AllowProductSorting:                              false,
		DefaultViewMode:                                  "list",
		RecentlyViewedProductsEnabled:                    false,
		RecentlyViewedProductsNumber:                     3,
		NewProductsEnabled:                               false,
		NewProductsPageSize:                              20,
		CompareProductsEnabled:                           false,
		CompareProductsNumber:                            2,
		ProductSearchEnabled:                             false,
		ProductSearchAutoCompleteEnabled:                 false,
		ProductSearchAutoCompleteNumberOfProducts:        0,
		ShowProductImagesInSearchAutoComplete:            false,
		ShowBestsellersOnHomepage:                        false,
		NumberOfBestsellersOnHomepage:                    0,
	}

	mockRepo.On("FetchByID", mock.Anything, catalogID).Return(expectedCatalogSettings, nil)

	result, err := usecase.FetchByID(context.Background(), catalogID)

	assert.NoError(t, err)
	assert.Equal(t, expectedCatalogSettings, result)
	mockRepo.AssertExpectations(t)
}

func TestCatalogSettingsUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.CatalogSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewCatalogSettingsUsecase(mockRepo, timeout)

	newCatalogSettings := &domain.CatalogSettings{
		AllowViewUnpublishedProductPage:                  true,
		DisplayDiscontinuedMessageForUnpublishedProducts: false,
		PublishBackProductWhenCancellingOrders:           true,
		ShowSkuOnProductDetailsPage:                      true,
		ShowSkuOnCatalogPages:                            false,
		ShowManufacturerPartNumber:                       true,
		ShowGtin:                                         false,
		ShowFreeShippingNotification:                     true,
		ShowShortDescriptionOnCatalogPages:               true,
		AllowProductSorting:                              true,
		DefaultViewMode:                                  "grid",
		RecentlyViewedProductsEnabled:                    true,
		RecentlyViewedProductsNumber:                     5,
		NewProductsEnabled:                               true,
		NewProductsPageSize:                              10,
		CompareProductsEnabled:                           true,
		CompareProductsNumber:                            4,
		ProductSearchEnabled:                             true,
		ProductSearchAutoCompleteEnabled:                 true,
		ProductSearchAutoCompleteNumberOfProducts:        5,
		ShowProductImagesInSearchAutoComplete:            true,
		ShowBestsellersOnHomepage:                        true,
		NumberOfBestsellersOnHomepage:                    10,
	}

	mockRepo.On("Create", mock.Anything, newCatalogSettings).Return(nil)

	err := usecase.Create(context.Background(), newCatalogSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCatalogSettingsUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.CatalogSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewCatalogSettingsUsecase(mockRepo, timeout)

	updatedCatalogSettings := &domain.CatalogSettings{
		ID:                              primitive.NewObjectID(), // Existing ID of the record to update
		AllowViewUnpublishedProductPage: false,
		DisplayDiscontinuedMessageForUnpublishedProducts: true,
		PublishBackProductWhenCancellingOrders:           false,
		ShowSkuOnProductDetailsPage:                      false,
		ShowSkuOnCatalogPages:                            true,
		ShowManufacturerPartNumber:                       false,
		ShowGtin:                                         true,
		ShowFreeShippingNotification:                     false,
		ShowShortDescriptionOnCatalogPages:               false,
		AllowProductSorting:                              false,
		DefaultViewMode:                                  "list",
		RecentlyViewedProductsEnabled:                    false,
		RecentlyViewedProductsNumber:                     3,
		NewProductsEnabled:                               false,
		NewProductsPageSize:                              20,
		CompareProductsEnabled:                           false,
		CompareProductsNumber:                            2,
		ProductSearchEnabled:                             false,
		ProductSearchAutoCompleteEnabled:                 false,
		ProductSearchAutoCompleteNumberOfProducts:        0,
		ShowProductImagesInSearchAutoComplete:            false,
		ShowBestsellersOnHomepage:                        false,
		NumberOfBestsellersOnHomepage:                    0,
	}

	mockRepo.On("Update", mock.Anything, updatedCatalogSettings).Return(nil)

	err := usecase.Update(context.Background(), updatedCatalogSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCatalogSettingsUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.CatalogSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewCatalogSettingsUsecase(mockRepo, timeout)

	catalogID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, catalogID).Return(nil)

	err := usecase.Delete(context.Background(), catalogID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCatalogSettingsUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.CatalogSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewCatalogSettingsUsecase(mockRepo, timeout)

	expectedCatalogSettings := []domain.CatalogSettings{
		{
			ID:                              primitive.NewObjectID(),
			AllowViewUnpublishedProductPage: true,
			DisplayDiscontinuedMessageForUnpublishedProducts: false,
			PublishBackProductWhenCancellingOrders:           true,
			ShowSkuOnProductDetailsPage:                      true,
			ShowSkuOnCatalogPages:                            false,
			ShowManufacturerPartNumber:                       true,
			ShowGtin:                                         false,
			ShowFreeShippingNotification:                     true,
			ShowShortDescriptionOnCatalogPages:               true,
			AllowProductSorting:                              true,
			DefaultViewMode:                                  "grid",
			RecentlyViewedProductsEnabled:                    true,
			RecentlyViewedProductsNumber:                     5,
			NewProductsEnabled:                               true,
			NewProductsPageSize:                              10,
			CompareProductsEnabled:                           true,
			CompareProductsNumber:                            4,
			ProductSearchEnabled:                             true,
			ProductSearchAutoCompleteEnabled:                 true,
			ProductSearchAutoCompleteNumberOfProducts:        5,
			ShowProductImagesInSearchAutoComplete:            true,
			ShowBestsellersOnHomepage:                        true,
			NumberOfBestsellersOnHomepage:                    10,
		},
		{
			ID:                              primitive.NewObjectID(),
			AllowViewUnpublishedProductPage: false,
			DisplayDiscontinuedMessageForUnpublishedProducts: true,
			PublishBackProductWhenCancellingOrders:           false,
			ShowSkuOnProductDetailsPage:                      false,
			ShowSkuOnCatalogPages:                            true,
			ShowManufacturerPartNumber:                       false,
			ShowGtin:                                         true,
			ShowFreeShippingNotification:                     false,
			ShowShortDescriptionOnCatalogPages:               false,
			AllowProductSorting:                              false,
			DefaultViewMode:                                  "list",
			RecentlyViewedProductsEnabled:                    false,
			RecentlyViewedProductsNumber:                     3,
			NewProductsEnabled:                               false,
			NewProductsPageSize:                              20,
			CompareProductsEnabled:                           false,
			CompareProductsNumber:                            2,
			ProductSearchEnabled:                             false,
			ProductSearchAutoCompleteEnabled:                 false,
			ProductSearchAutoCompleteNumberOfProducts:        0,
			ShowProductImagesInSearchAutoComplete:            false,
			ShowBestsellersOnHomepage:                        false,
			NumberOfBestsellersOnHomepage:                    0,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(expectedCatalogSettings, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, expectedCatalogSettings, result)
	mockRepo.AssertExpectations(t)
}
