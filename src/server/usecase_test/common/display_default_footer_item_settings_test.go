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

func TestDisplayDefaultFooterItemSettingsUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.DisplayDefaultFooterItemSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewDisplayDefaultFooterItemSettingsUsecase(mockRepo, timeout)

	footerItemID := primitive.NewObjectID().Hex()

	updatedDisplayDefaultFooterItemSettings := domain.DisplayDefaultFooterItemSettings{
		ID:                                      primitive.NewObjectID(), // Existing ID of the record to update
		DisplaySitemapFooterItem:                false,
		DisplayContactUsFooterItem:              false,
		DisplayProductSearchFooterItem:          true,
		DisplayNewsFooterItem:                   true,
		DisplayBlogFooterItem:                   false,
		DisplayForumsFooterItem:                 true,
		DisplayRecentlyViewedProductsFooterItem: false,
		DisplayCompareProductsFooterItem:        true,
		DisplayNewProductsFooterItem:            false,
		DisplayCustomerInfoFooterItem:           false,
		DisplayCustomerOrdersFooterItem:         true,
		DisplayCustomerAddressesFooterItem:      true,
		DisplayShoppingCartFooterItem:           false,
		DisplayWishlistFooterItem:               false,
		DisplayApplyVendorAccountFooterItem:     true,
	}

	mockRepo.On("FetchByID", mock.Anything, footerItemID).Return(updatedDisplayDefaultFooterItemSettings, nil)

	result, err := usecase.FetchByID(context.Background(), footerItemID)

	assert.NoError(t, err)
	assert.Equal(t, updatedDisplayDefaultFooterItemSettings, result)
	mockRepo.AssertExpectations(t)
}

func TestDisplayDefaultFooterItemSettingsUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.DisplayDefaultFooterItemSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewDisplayDefaultFooterItemSettingsUsecase(mockRepo, timeout)

	newDisplayDefaultFooterItemSettings := &domain.DisplayDefaultFooterItemSettings{
		DisplaySitemapFooterItem:                true,
		DisplayContactUsFooterItem:              true,
		DisplayProductSearchFooterItem:          true,
		DisplayNewsFooterItem:                   false,
		DisplayBlogFooterItem:                   true,
		DisplayForumsFooterItem:                 false,
		DisplayRecentlyViewedProductsFooterItem: true,
		DisplayCompareProductsFooterItem:        false,
		DisplayNewProductsFooterItem:            true,
		DisplayCustomerInfoFooterItem:           true,
		DisplayCustomerOrdersFooterItem:         true,
		DisplayCustomerAddressesFooterItem:      false,
		DisplayShoppingCartFooterItem:           true,
		DisplayWishlistFooterItem:               true,
		DisplayApplyVendorAccountFooterItem:     false,
	}

	mockRepo.On("Create", mock.Anything, newDisplayDefaultFooterItemSettings).Return(nil)

	err := usecase.Create(context.Background(), newDisplayDefaultFooterItemSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDisplayDefaultFooterItemSettingsUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.DisplayDefaultFooterItemSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewDisplayDefaultFooterItemSettingsUsecase(mockRepo, timeout)

	updatedDisplayDefaultFooterItemSettings := &domain.DisplayDefaultFooterItemSettings{
		ID:                                      primitive.NewObjectID(), // Existing ID of the record to update
		DisplaySitemapFooterItem:                false,
		DisplayContactUsFooterItem:              false,
		DisplayProductSearchFooterItem:          true,
		DisplayNewsFooterItem:                   true,
		DisplayBlogFooterItem:                   false,
		DisplayForumsFooterItem:                 true,
		DisplayRecentlyViewedProductsFooterItem: false,
		DisplayCompareProductsFooterItem:        true,
		DisplayNewProductsFooterItem:            false,
		DisplayCustomerInfoFooterItem:           false,
		DisplayCustomerOrdersFooterItem:         true,
		DisplayCustomerAddressesFooterItem:      true,
		DisplayShoppingCartFooterItem:           false,
		DisplayWishlistFooterItem:               false,
		DisplayApplyVendorAccountFooterItem:     true,
	}

	mockRepo.On("Update", mock.Anything, updatedDisplayDefaultFooterItemSettings).Return(nil)

	err := usecase.Update(context.Background(), updatedDisplayDefaultFooterItemSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDisplayDefaultFooterItemSettingsUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.DisplayDefaultFooterItemSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewDisplayDefaultFooterItemSettingsUsecase(mockRepo, timeout)

	footerItemID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, footerItemID).Return(nil)

	err := usecase.Delete(context.Background(), footerItemID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDisplayDefaultFooterItemSettingsUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.DisplayDefaultFooterItemSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewDisplayDefaultFooterItemSettingsUsecase(mockRepo, timeout)

	fetchedDisplayDefaultFooterItemSettings := []domain.DisplayDefaultFooterItemSettings{
		{
			ID:                                      primitive.NewObjectID(),
			DisplaySitemapFooterItem:                true,
			DisplayContactUsFooterItem:              true,
			DisplayProductSearchFooterItem:          true,
			DisplayNewsFooterItem:                   false,
			DisplayBlogFooterItem:                   true,
			DisplayForumsFooterItem:                 false,
			DisplayRecentlyViewedProductsFooterItem: true,
			DisplayCompareProductsFooterItem:        false,
			DisplayNewProductsFooterItem:            true,
			DisplayCustomerInfoFooterItem:           true,
			DisplayCustomerOrdersFooterItem:         true,
			DisplayCustomerAddressesFooterItem:      false,
			DisplayShoppingCartFooterItem:           true,
			DisplayWishlistFooterItem:               true,
			DisplayApplyVendorAccountFooterItem:     false,
		},
		{
			ID:                                      primitive.NewObjectID(),
			DisplaySitemapFooterItem:                false,
			DisplayContactUsFooterItem:              false,
			DisplayProductSearchFooterItem:          true,
			DisplayNewsFooterItem:                   true,
			DisplayBlogFooterItem:                   false,
			DisplayForumsFooterItem:                 true,
			DisplayRecentlyViewedProductsFooterItem: false,
			DisplayCompareProductsFooterItem:        true,
			DisplayNewProductsFooterItem:            false,
			DisplayCustomerInfoFooterItem:           false,
			DisplayCustomerOrdersFooterItem:         true,
			DisplayCustomerAddressesFooterItem:      true,
			DisplayShoppingCartFooterItem:           false,
			DisplayWishlistFooterItem:               false,
			DisplayApplyVendorAccountFooterItem:     true,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedDisplayDefaultFooterItemSettings, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedDisplayDefaultFooterItemSettings, result)
	mockRepo.AssertExpectations(t)
}
