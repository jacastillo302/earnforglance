package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/catalog"
	repository "earnforglance/server/repository/catalog"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type MockSingleResultCatalogSettings struct {
	mock.Mock
}

func (m *MockSingleResultCatalogSettings) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.CatalogSettings); ok {
		*v.(*domain.CatalogSettings) = *result
	}
	return args.Error(1)
}

func TestCatalogSettingsRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionCatalogSettings

	mockItem := domain.CatalogSettings{ID: bson.ObjectID{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, AllowViewUnpublishedProductPage: false, DisplayDiscontinuedMessageForUnpublishedProducts: false, PublishBackProductWhenCancellingOrders: false, ShowSkuOnProductDetailsPage: false, ShowSkuOnCatalogPages: false, ShowManufacturerPartNumber: false, ShowGtin: false, ShowFreeShippingNotification: false, ShowShortDescriptionOnCatalogPages: false, AllowProductSorting: false, AllowProductViewModeChanging: false, DefaultViewMode: "", ShowProductsFromSubcategories: false, ShowCategoryProductNumber: false, ShowCategoryProductNumberIncludingSubcategories: false, CategoryBreadcrumbEnabled: false, ShowShareButton: false, PageShareCode: "", ProductReviewsMustBeApproved: false, OneReviewPerProductFromCustomer: false, DefaultProductRatingValue: 0, AllowAnonymousUsersToReviewProduct: false, ProductReviewPossibleOnlyAfterPurchasing: false, NotifyStoreOwnerAboutNewProductReviews: false, NotifyCustomerAboutProductReviewReply: false, ShowProductReviewsPerStore: false, ShowProductReviewsTabOnAccountPage: false, ProductReviewsPageSizeOnAccountPage: 0, ProductReviewsSortByCreatedDateAscending: false, EmailAFriendEnabled: false, AllowAnonymousUsersToEmailAFriend: false, RecentlyViewedProductsNumber: 0, RecentlyViewedProductsEnabled: false, NewProductsEnabled: false, NewProductsPageSize: 0, NewProductsAllowCustomersToSelectPageSize: false, NewProductsPageSizeOptions: "", CompareProductsEnabled: false, CompareProductsNumber: 0, ProductSearchAutoCompleteEnabled: false, ProductSearchEnabled: false, ProductSearchAutoCompleteNumberOfProducts: 0, ShowProductImagesInSearchAutoComplete: false, ShowLinkToAllResultInSearchAutoComplete: false, ProductSearchTermMinimumLength: 0, ShowBestsellersOnHomepage: false, NumberOfBestsellersOnHomepage: 0, ShowSearchBoxCategories: false, SearchPageProductsPerPage: 0, SearchPageAllowCustomersToSelectPageSize: false, SearchPagePageSizeOptions: "", SearchPagePriceRangeFiltering: false, SearchPagePriceFrom: 0, SearchPagePriceTo: 0, SearchPageManuallyPriceRange: false, ProductsAlsoPurchasedEnabled: false, ProductsAlsoPurchasedNumber: 0, AjaxProcessAttributeChange: false, NumberOfProductTags: 0, ProductsByTagPageSize: 0, ProductsByTagAllowCustomersToSelectPageSize: false, ProductsByTagPageSizeOptions: "", ProductsByTagPriceRangeFiltering: false, ProductsByTagPriceFrom: 0, ProductsByTagPriceTo: 0, ProductsByTagManuallyPriceRange: false, IncludeShortDescriptionInCompareProducts: false, IncludeFullDescriptionInCompareProducts: false, IncludeFeaturedProductsInNormalLists: false, UseLinksInRequiredProductWarnings: false, DisplayTierPricesWithDiscounts: false, IgnoreDiscounts: false, IgnoreFeaturedProducts: false, IgnoreAcl: false, IgnoreStoreLimitations: false, CacheProductPrices: false, MaximumBackInStockSubscriptions: 0, ManufacturersBlockItemsToDisplay: 0, DisplayTaxShippingInfoFooter: false, DisplayTaxShippingInfoProductDetailsPage: false, DisplayTaxShippingInfoProductBoxes: false, DisplayTaxShippingInfoShoppingCart: false, DisplayTaxShippingInfoWishlist: false, DisplayTaxShippingInfoOrderDetailsPage: false, DefaultCategoryPageSizeOptions: "", DefaultCategoryPageSize: 0, DefaultManufacturerPageSizeOptions: "", DefaultManufacturerPageSize: 0, ProductSortingEnumDisabled: []int(nil), ProductSortingEnumDisplayOrder: map[int]int(nil), ExportImportProductAttributes: false, ExportImportProductUseLimitedToStores: false, ExportImportCategoryUseLimitedToStores: false, ExportImportProductSpecificationAttributes: false, ExportImportTierPrices: false, ExportImportUseDropdownlistsForAssociatedEntities: false, ExportImportProductCategoryBreadcrumb: false, ExportImportCategoriesUsingCategoryName: false, ExportImportAllowDownloadImages: false, ExportImportSplitProductsFile: false, ExportImportProductsCountInOneFile: 0, RemoveRequiredProducts: false, ExportImportRelatedEntitiesByName: false, CountDisplayedYearsDatePicker: 0, DisplayDatePreOrderAvailability: false, UseAjaxLoadMenu: false, UseAjaxCatalogProductsLoading: false, EnableManufacturerFiltering: false, EnablePriceRangeFiltering: false, EnableSpecificationAttributeFiltering: false, DisplayFromPrices: false, AttributeValueOutOfStockDisplayTypeID: false, AllowCustomersToSearchWithManufacturerName: false, AllowCustomersToSearchWithCategoryName: false, DisplayAllPicturesOnCatalogPages: false, ProductUrlStructureTypeID: 0, ActiveSearchProviderSystemName: "", UseStandardSearchWhenSearchProviderThrowsException: false, VendorProductReviewsPageSize: 0}

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultCatalogSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItem, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewCatalogSettingsRepository(databaseHelper, collectionName)

		result, err := ur.FetchByID(context.Background(), mockItem.ID.Hex())

		assert.NoError(t, err)
		assert.Equal(t, mockItem, result)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultCatalogSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewCatalogSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItem.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestCatalogSettingsRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionCatalogSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)

	mockCatalogSettings := &domain.CatalogSettings{
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

	collectionHelper.On("InsertOne", mock.Anything, mockCatalogSettings).Return(nil, nil).Once()

	repo := repository.NewCatalogSettingsRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockCatalogSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestCatalogSettingsRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionCatalogSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)

	mockCatalogSettings := &domain.CatalogSettings{
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

	filter := bson.M{"_id": mockCatalogSettings.ID}
	update := bson.M{"$set": mockCatalogSettings}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewCatalogSettingsRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockCatalogSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
