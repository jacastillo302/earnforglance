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

func TestProductUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.ProductRepository)
	timeout := time.Duration(10)
	usecase := test.NewProductUsecase(mockRepo, timeout)

	productID := primitive.NewObjectID().Hex()

	updatedProduct := domain.Product{
		ID:                               primitive.NewObjectID(), // Existing ID of the record to update
		ProductTypeID:                    2,
		ParentGroupedID:                  primitive.NewObjectID(),
		VisibleIndividually:              false,
		Name:                             "Updated Product",
		ShortDescription:                 "Updated short description.",
		FullDescription:                  "Updated full description of the product.",
		AdminComment:                     "Updated admin comment.",
		ProductTemplateID:                primitive.NewObjectID(),
		VendorID:                         primitive.NewObjectID(),
		ShowOnHomepage:                   false,
		MetaKeywords:                     "updated, product, ecommerce",
		MetaDescription:                  "Updated meta description for the product.",
		MetaTitle:                        "Updated Product Title",
		AllowCustomerReviews:             false,
		ApprovedRatingSum:                30,
		NotApprovedRatingSum:             10,
		ApprovedTotalReviews:             5,
		NotApprovedTotalReviews:          3,
		SubjectToAcl:                     true,
		LimitedToStores:                  true,
		Sku:                              "SKU54321",
		ManufacturerPartNumber:           "MPN09876",
		Gtin:                             "0987654321098",
		IsGiftCard:                       true,
		GiftCardTypeID:                   1,
		OverriddenGiftCardAmount:         new(float64),
		RequireOtherProducts:             true,
		RequiredIDs:                      "123,456",
		AutomaticallyAddRequiredProducts: true,
		IsDownload:                       true,
		DownloadID:                       1,
		UnlimitedDownloads:               true,
		MaxNumberOfDownloads:             5,
		DownloadExpirationDays:           new(int),
		DownloadActivationTypeID:         1,
		HasSampleDownload:                true,
		SampleDownloadID:                 2,
		HasUserAgreement:                 true,
		UserAgreementText:                "Updated user agreement text.",
		IsRecurring:                      true,
		RecurringCycleLength:             30,
		RecurringProductCyclePeriodID:    2,
		RecurringTotalCycles:             12,
		IsRental:                         true,
		RentalPriceLength:                7,
		RentalPricePeriodID:              1,
		IsShipEnabled:                    false,
		IsFreeShipping:                   true,
		ShipSeparately:                   true,
		AdditionalShippingCharge:         0.0,
		DeliveryDateID:                   primitive.NewObjectID(),
		IsTaxExempt:                      true,
		TaxCategoryID:                    primitive.NewObjectID(),
		ManageInventoryMethodID:          2,
		ProductAvailabilityRangeID:       primitive.NewObjectID(),
		UseMultipleWarehouses:            true,
		WarehouseID:                      primitive.NewObjectID(),
		StockQuantity:                    50,
		DisplayStockAvailability:         false,
		DisplayStockQuantity:             false,
		MinStockQuantity:                 2,
		LowStockActivityID:               2,
		NotifyAdminForQuantityBelow:      5,
		BackorderModeID:                  1,
		AllowBackInStockSubscriptions:    false,
		OrderMinimumQuantity:             2,
		OrderMaximumQuantity:             5,
		AllowedQuantities:                "2,3,4,5",
		AllowAddingOnlyExistingAttributeCombinations: true,
		DisplayAttributeCombinationImagesOnly:        true,
		NotReturnable:                                true,
		DisableBuyButton:                             true,
		DisableWishlistButton:                        true,
		AvailableForPreOrder:                         true,
		PreOrderAvailabilityStartDateTimeUtc:         new(time.Time),
		CallForPrice:                                 true,
		Price:                                        79.99,
		OldPrice:                                     100.00,
		ProductCost:                                  40.00,
		CustomerEntersPrice:                          true,
		MinimumCustomerEnteredPrice:                  10.0,
		MaximumCustomerEnteredPrice:                  50.0,
		BasepriceEnabled:                             true,
		BasepriceAmount:                              1.0,
		BasepriceUnitID:                              1,
		BasepriceBaseAmount:                          10.0,
		BasepriceBaseUnitID:                          2,
		MarkAsNew:                                    false,
		MarkAsNewStartDateTimeUtc:                    new(time.Time),
		MarkAsNewEndDateTimeUtc:                      new(time.Time),
		Weight:                                       2.0,
		Length:                                       15.0,
		Width:                                        7.0,
		Height:                                       4.0,
		AvailableStartDateTimeUtc:                    new(time.Time),
		AvailableEndDateTimeUtc:                      new(time.Time),
		DisplayOrder:                                 2,
		Published:                                    false,
		Deleted:                                      true,
		CreatedOnUtc:                                 time.Now().AddDate(0, 0, -30), // Created 30 days ago
		UpdatedOnUtc:                                 time.Now(),
		AgeVerification:                              true,
		MinimumAgeToPurchase:                         18,
	}

	mockRepo.On("FetchByID", mock.Anything, productID).Return(updatedProduct, nil)

	result, err := usecase.FetchByID(context.Background(), productID)

	assert.NoError(t, err)
	assert.Equal(t, updatedProduct, result)
	mockRepo.AssertExpectations(t)
}

func TestProductUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.ProductRepository)
	timeout := time.Duration(10)
	usecase := test.NewProductUsecase(mockRepo, timeout)

	newProduct := &domain.Product{
		ProductTypeID:                    1,
		ParentGroupedID:                  primitive.NewObjectID(),
		VisibleIndividually:              true,
		Name:                             "Sample Product",
		ShortDescription:                 "This is a short description.",
		FullDescription:                  "This is a full description of the product.",
		AdminComment:                     "Admin comment here.",
		ProductTemplateID:                primitive.NewObjectID(),
		VendorID:                         primitive.NewObjectID(),
		ShowOnHomepage:                   true,
		MetaKeywords:                     "sample, product, ecommerce",
		MetaDescription:                  "Meta description for the sample product.",
		MetaTitle:                        "Sample Product Title",
		AllowCustomerReviews:             true,
		ApprovedRatingSum:                50,
		NotApprovedRatingSum:             5,
		ApprovedTotalReviews:             10,
		NotApprovedTotalReviews:          2,
		SubjectToAcl:                     false,
		LimitedToStores:                  false,
		Sku:                              "SKU12345",
		ManufacturerPartNumber:           "MPN67890",
		Gtin:                             "0123456789012",
		IsGiftCard:                       false,
		GiftCardTypeID:                   0,
		OverriddenGiftCardAmount:         nil,
		RequireOtherProducts:             false,
		RequiredIDs:                      "",
		AutomaticallyAddRequiredProducts: false,
		IsDownload:                       false,
		DownloadID:                       0,
		UnlimitedDownloads:               false,
		MaxNumberOfDownloads:             0,
		DownloadExpirationDays:           nil,
		DownloadActivationTypeID:         0,
		HasSampleDownload:                false,
		SampleDownloadID:                 0,
		HasUserAgreement:                 false,
		UserAgreementText:                "",
		IsRecurring:                      false,
		RecurringCycleLength:             0,
		RecurringProductCyclePeriodID:    0,
		RecurringTotalCycles:             0,
		IsRental:                         false,
		RentalPriceLength:                0,
		RentalPricePeriodID:              0,
		IsShipEnabled:                    true,
		IsFreeShipping:                   false,
		ShipSeparately:                   false,
		AdditionalShippingCharge:         5.0,
		DeliveryDateID:                   primitive.NewObjectID(),
		IsTaxExempt:                      false,
		TaxCategoryID:                    primitive.NewObjectID(),
		ManageInventoryMethodID:          1,
		ProductAvailabilityRangeID:       primitive.NewObjectID(),
		UseMultipleWarehouses:            false,
		WarehouseID:                      primitive.NewObjectID(),
		StockQuantity:                    100,
		DisplayStockAvailability:         true,
		DisplayStockQuantity:             true,
		MinStockQuantity:                 5,
		LowStockActivityID:               1,
		NotifyAdminForQuantityBelow:      10,
		BackorderModeID:                  0,
		AllowBackInStockSubscriptions:    true,
		OrderMinimumQuantity:             1,
		OrderMaximumQuantity:             10,
		AllowedQuantities:                "1,2,3,4,5",
		AllowAddingOnlyExistingAttributeCombinations: false,
		DisplayAttributeCombinationImagesOnly:        false,
		NotReturnable:                                false,
		DisableBuyButton:                             false,
		DisableWishlistButton:                        false,
		AvailableForPreOrder:                         false,
		PreOrderAvailabilityStartDateTimeUtc:         nil,
		CallForPrice:                                 false,
		Price:                                        99.99,
		OldPrice:                                     120.00,
		ProductCost:                                  50.00,
		CustomerEntersPrice:                          false,
		MinimumCustomerEnteredPrice:                  0.0,
		MaximumCustomerEnteredPrice:                  0.0,
		BasepriceEnabled:                             false,
		BasepriceAmount:                              0.0,
		BasepriceUnitID:                              0,
		BasepriceBaseAmount:                          0.0,
		BasepriceBaseUnitID:                          0,
		MarkAsNew:                                    true,
		MarkAsNewStartDateTimeUtc:                    nil,
		MarkAsNewEndDateTimeUtc:                      nil,
		Weight:                                       1.5,
		Length:                                       10.0,
		Width:                                        5.0,
		Height:                                       3.0,
		AvailableStartDateTimeUtc:                    nil,
		AvailableEndDateTimeUtc:                      nil,
		DisplayOrder:                                 1,
		Published:                                    true,
		Deleted:                                      false,
		CreatedOnUtc:                                 time.Now(),
		UpdatedOnUtc:                                 time.Now(),
		AgeVerification:                              false,
		MinimumAgeToPurchase:                         0,
	}

	mockRepo.On("Create", mock.Anything, newProduct).Return(nil)

	err := usecase.Create(context.Background(), newProduct)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.ProductRepository)
	timeout := time.Duration(10)
	usecase := test.NewProductUsecase(mockRepo, timeout)

	updatedProduct := &domain.Product{
		ID:                               primitive.NewObjectID(), // Existing ID of the record to update
		ProductTypeID:                    2,
		ParentGroupedID:                  primitive.NewObjectID(),
		VisibleIndividually:              false,
		Name:                             "Updated Product",
		ShortDescription:                 "Updated short description.",
		FullDescription:                  "Updated full description of the product.",
		AdminComment:                     "Updated admin comment.",
		ProductTemplateID:                primitive.NewObjectID(),
		VendorID:                         primitive.NewObjectID(),
		ShowOnHomepage:                   false,
		MetaKeywords:                     "updated, product, ecommerce",
		MetaDescription:                  "Updated meta description for the product.",
		MetaTitle:                        "Updated Product Title",
		AllowCustomerReviews:             false,
		ApprovedRatingSum:                30,
		NotApprovedRatingSum:             10,
		ApprovedTotalReviews:             5,
		NotApprovedTotalReviews:          3,
		SubjectToAcl:                     true,
		LimitedToStores:                  true,
		Sku:                              "SKU54321",
		ManufacturerPartNumber:           "MPN09876",
		Gtin:                             "0987654321098",
		IsGiftCard:                       true,
		GiftCardTypeID:                   1,
		OverriddenGiftCardAmount:         new(float64),
		RequireOtherProducts:             true,
		RequiredIDs:                      "123,456",
		AutomaticallyAddRequiredProducts: true,
		IsDownload:                       true,
		DownloadID:                       1,
		UnlimitedDownloads:               true,
		MaxNumberOfDownloads:             5,
		DownloadExpirationDays:           new(int),
		DownloadActivationTypeID:         1,
		HasSampleDownload:                true,
		SampleDownloadID:                 2,
		HasUserAgreement:                 true,
		UserAgreementText:                "Updated user agreement text.",
		IsRecurring:                      true,
		RecurringCycleLength:             30,
		RecurringProductCyclePeriodID:    2,
		RecurringTotalCycles:             12,
		IsRental:                         true,
		RentalPriceLength:                7,
		RentalPricePeriodID:              1,
		IsShipEnabled:                    false,
		IsFreeShipping:                   true,
		ShipSeparately:                   true,
		AdditionalShippingCharge:         0.0,
		DeliveryDateID:                   primitive.NewObjectID(),
		IsTaxExempt:                      true,
		TaxCategoryID:                    primitive.NewObjectID(),
		ManageInventoryMethodID:          2,
		ProductAvailabilityRangeID:       primitive.NewObjectID(),
		UseMultipleWarehouses:            true,
		WarehouseID:                      primitive.NewObjectID(),
		StockQuantity:                    50,
		DisplayStockAvailability:         false,
		DisplayStockQuantity:             false,
		MinStockQuantity:                 2,
		LowStockActivityID:               2,
		NotifyAdminForQuantityBelow:      5,
		BackorderModeID:                  1,
		AllowBackInStockSubscriptions:    false,
		OrderMinimumQuantity:             2,
		OrderMaximumQuantity:             5,
		AllowedQuantities:                "2,3,4,5",
		AllowAddingOnlyExistingAttributeCombinations: true,
		DisplayAttributeCombinationImagesOnly:        true,
		NotReturnable:                                true,
		DisableBuyButton:                             true,
		DisableWishlistButton:                        true,
		AvailableForPreOrder:                         true,
		PreOrderAvailabilityStartDateTimeUtc:         new(time.Time),
		CallForPrice:                                 true,
		Price:                                        79.99,
		OldPrice:                                     100.00,
		ProductCost:                                  40.00,
		CustomerEntersPrice:                          true,
		MinimumCustomerEnteredPrice:                  10.0,
		MaximumCustomerEnteredPrice:                  50.0,
		BasepriceEnabled:                             true,
		BasepriceAmount:                              1.0,
		BasepriceUnitID:                              1,
		BasepriceBaseAmount:                          10.0,
		BasepriceBaseUnitID:                          2,
		MarkAsNew:                                    false,
		MarkAsNewStartDateTimeUtc:                    new(time.Time),
		MarkAsNewEndDateTimeUtc:                      new(time.Time),
		Weight:                                       2.0,
		Length:                                       15.0,
		Width:                                        7.0,
		Height:                                       4.0,
		AvailableStartDateTimeUtc:                    new(time.Time),
		AvailableEndDateTimeUtc:                      new(time.Time),
		DisplayOrder:                                 2,
		Published:                                    false,
		Deleted:                                      true,
		CreatedOnUtc:                                 time.Now().AddDate(0, 0, -30), // Created 30 days ago
		UpdatedOnUtc:                                 time.Now(),
		AgeVerification:                              true,
		MinimumAgeToPurchase:                         18,
	}
	*updatedProduct.OverriddenGiftCardAmount = 25.00
	*updatedProduct.DownloadExpirationDays = 30
	*updatedProduct.PreOrderAvailabilityStartDateTimeUtc = time.Now().AddDate(0, 0, 7) // 7 days from now
	*updatedProduct.MarkAsNewStartDateTimeUtc = time.Now().AddDate(0, 0, -7)           // 7 days ago
	*updatedProduct.MarkAsNewEndDateTimeUtc = time.Now().AddDate(0, 0, 7)              // 7 days from now

	mockRepo.On("Update", mock.Anything, updatedProduct).Return(nil)

	err := usecase.Update(context.Background(), updatedProduct)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.ProductRepository)
	timeout := time.Duration(10)
	usecase := test.NewProductUsecase(mockRepo, timeout)

	productID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, productID).Return(nil)

	err := usecase.Delete(context.Background(), productID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.ProductRepository)
	timeout := time.Duration(10)
	usecase := test.NewProductUsecase(mockRepo, timeout)

	fetchedProducts := []domain.Product{
		{
			ID:                  primitive.NewObjectID(),
			ProductTypeID:       1,
			ParentGroupedID:     primitive.NewObjectID(),
			VisibleIndividually: true,
			Name:                "Product 1",
			ShortDescription:    "Short description for product 1.",
			FullDescription:     "Full description for product 1.",
			Price:               49.99,
			StockQuantity:       100,
			Published:           true,
			CreatedOnUtc:        time.Now().AddDate(0, 0, -10), // Created 10 days ago
			UpdatedOnUtc:        time.Now(),
		},
		{
			ID:                  primitive.NewObjectID(),
			ProductTypeID:       2,
			ParentGroupedID:     primitive.NewObjectID(),
			VisibleIndividually: false,
			Name:                "Product 2",
			ShortDescription:    "Short description for product 2.",
			FullDescription:     "Full description for product 2.",
			Price:               79.99,
			StockQuantity:       50,
			Published:           false,
			CreatedOnUtc:        time.Now().AddDate(0, 0, -20), // Created 20 days ago
			UpdatedOnUtc:        time.Now(),
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedProducts, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedProducts, result)
	mockRepo.AssertExpectations(t)
}
