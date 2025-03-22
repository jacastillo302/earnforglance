package usecase

import (
	"context"
	domain "earnforglance/server/domain/catalog"
	mocks "earnforglance/server/domain/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestProductEditorSettingsUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.ProductEditorSettingsRepository)
	timeout := time.Duration(10)
	usecase := NewProductEditorSettingsUsecase(mockRepo, timeout)

	productEditorSettingsID := primitive.NewObjectID().Hex()

	updatedProductEditorSettings := domain.ProductEditorSettings{
		ID:                              primitive.NewObjectID(), // Existing ID of the record to update
		ProductType:                     false,
		VisibleIndividually:             false,
		ProductTemplate:                 true,
		AdminComment:                    true,
		Vendor:                          false,
		Stores:                          false,
		ACL:                             true,
		ShowOnHomepage:                  false,
		AllowCustomerReviews:            false,
		ProductTags:                     false,
		ManufacturerPartNumber:          true,
		GTIN:                            false,
		ProductCost:                     false,
		TierPrices:                      true,
		Discounts:                       false,
		DisableBuyButton:                true,
		DisableWishlistButton:           true,
		AvailableForPreOrder:            false,
		CallForPrice:                    true,
		OldPrice:                        false,
		CustomerEntersPrice:             true,
		PAngV:                           true,
		RequireOtherProductsAddedToCart: true,
		IsGiftCard:                      true,
		DownloadableProduct:             false,
		RecurringProduct:                true,
		IsRental:                        true,
		FreeShipping:                    false,
		ShipSeparately:                  true,
		AdditionalShippingCharge:        false,
		DeliveryDate:                    false,
		ProductAvailabilityRange:        true,
		UseMultipleWarehouses:           false,
		Warehouse:                       false,
		DisplayStockAvailability:        false,
		MinimumStockQuantity:            false,
		LowStockActivity:                true,
		NotifyAdminForQuantityBelow:     false,
		Backorders:                      false,
		AllowBackInStockSubscriptions:   false,
		MinimumCartQuantity:             false,
		MaximumCartQuantity:             false,
		AllowedQuantities:               false,
		AllowAddingOnlyExistingAttributeCombinations: true,
		NotReturnable:           true,
		Weight:                  false,
		Dimensions:              false,
		AvailableStartDate:      false,
		AvailableEndDate:        false,
		MarkAsNew:               false,
		Published:               false,
		RelatedProducts:         false,
		CrossSellsProducts:      false,
		Seo:                     false,
		PurchasedWithOrders:     true,
		ProductAttributes:       false,
		SpecificationAttributes: false,
		Manufacturers:           false,
		StockQuantityChange:     false,
		AgeVerification:         true,
	}

	mockRepo.On("FetchByID", mock.Anything, productEditorSettingsID).Return(updatedProductEditorSettings, nil)

	result, err := usecase.FetchByID(context.Background(), productEditorSettingsID)

	assert.NoError(t, err)
	assert.Equal(t, updatedProductEditorSettings, result)
	mockRepo.AssertExpectations(t)
}

func TestProductEditorSettingsUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.ProductEditorSettingsRepository)
	timeout := time.Duration(10)
	usecase := NewProductEditorSettingsUsecase(mockRepo, timeout)

	newProductEditorSettings := &domain.ProductEditorSettings{
		ProductType:                     true,
		VisibleIndividually:             true,
		ProductTemplate:                 true,
		AdminComment:                    false,
		Vendor:                          true,
		Stores:                          true,
		ACL:                             false,
		ShowOnHomepage:                  true,
		AllowCustomerReviews:            true,
		ProductTags:                     true,
		ManufacturerPartNumber:          false,
		GTIN:                            true,
		ProductCost:                     true,
		TierPrices:                      false,
		Discounts:                       true,
		DisableBuyButton:                false,
		DisableWishlistButton:           false,
		AvailableForPreOrder:            true,
		CallForPrice:                    false,
		OldPrice:                        true,
		CustomerEntersPrice:             false,
		PAngV:                           false,
		RequireOtherProductsAddedToCart: false,
		IsGiftCard:                      false,
		DownloadableProduct:             true,
		RecurringProduct:                false,
		IsRental:                        false,
		FreeShipping:                    true,
		ShipSeparately:                  false,
		AdditionalShippingCharge:        true,
		DeliveryDate:                    true,
		ProductAvailabilityRange:        false,
		UseMultipleWarehouses:           true,
		Warehouse:                       true,
		DisplayStockAvailability:        true,
		MinimumStockQuantity:            true,
		LowStockActivity:                false,
		NotifyAdminForQuantityBelow:     true,
		Backorders:                      true,
		AllowBackInStockSubscriptions:   true,
		MinimumCartQuantity:             true,
		MaximumCartQuantity:             true,
		AllowedQuantities:               true,
		AllowAddingOnlyExistingAttributeCombinations: false,
		NotReturnable:           false,
		Weight:                  true,
		Dimensions:              true,
		AvailableStartDate:      true,
		AvailableEndDate:        true,
		MarkAsNew:               true,
		Published:               true,
		RelatedProducts:         true,
		CrossSellsProducts:      true,
		Seo:                     true,
		PurchasedWithOrders:     false,
		ProductAttributes:       true,
		SpecificationAttributes: true,
		Manufacturers:           true,
		StockQuantityChange:     true,
		AgeVerification:         false,
	}

	mockRepo.On("Create", mock.Anything, newProductEditorSettings).Return(nil)

	err := usecase.Create(context.Background(), newProductEditorSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductEditorSettingsUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.ProductEditorSettingsRepository)
	timeout := time.Duration(10)
	usecase := NewProductEditorSettingsUsecase(mockRepo, timeout)

	updatedProductEditorSettings := &domain.ProductEditorSettings{
		ID:                              primitive.NewObjectID(), // Existing ID of the record to update
		ProductType:                     false,
		VisibleIndividually:             false,
		ProductTemplate:                 true,
		AdminComment:                    true,
		Vendor:                          false,
		Stores:                          false,
		ACL:                             true,
		ShowOnHomepage:                  false,
		AllowCustomerReviews:            false,
		ProductTags:                     false,
		ManufacturerPartNumber:          true,
		GTIN:                            false,
		ProductCost:                     false,
		TierPrices:                      true,
		Discounts:                       false,
		DisableBuyButton:                true,
		DisableWishlistButton:           true,
		AvailableForPreOrder:            false,
		CallForPrice:                    true,
		OldPrice:                        false,
		CustomerEntersPrice:             true,
		PAngV:                           true,
		RequireOtherProductsAddedToCart: true,
		IsGiftCard:                      true,
		DownloadableProduct:             false,
		RecurringProduct:                true,
		IsRental:                        true,
		FreeShipping:                    false,
		ShipSeparately:                  true,
		AdditionalShippingCharge:        false,
		DeliveryDate:                    false,
		ProductAvailabilityRange:        true,
		UseMultipleWarehouses:           false,
		Warehouse:                       false,
		DisplayStockAvailability:        false,
		MinimumStockQuantity:            false,
		LowStockActivity:                true,
		NotifyAdminForQuantityBelow:     false,
		Backorders:                      false,
		AllowBackInStockSubscriptions:   false,
		MinimumCartQuantity:             false,
		MaximumCartQuantity:             false,
		AllowedQuantities:               false,
		AllowAddingOnlyExistingAttributeCombinations: true,
		NotReturnable:           true,
		Weight:                  false,
		Dimensions:              false,
		AvailableStartDate:      false,
		AvailableEndDate:        false,
		MarkAsNew:               false,
		Published:               false,
		RelatedProducts:         false,
		CrossSellsProducts:      false,
		Seo:                     false,
		PurchasedWithOrders:     true,
		ProductAttributes:       false,
		SpecificationAttributes: false,
		Manufacturers:           false,
		StockQuantityChange:     false,
		AgeVerification:         true,
	}

	mockRepo.On("Update", mock.Anything, updatedProductEditorSettings).Return(nil)

	err := usecase.Update(context.Background(), updatedProductEditorSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductEditorSettingsUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.ProductEditorSettingsRepository)
	timeout := time.Duration(10)
	usecase := NewProductEditorSettingsUsecase(mockRepo, timeout)

	productEditorSettingsID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, productEditorSettingsID).Return(nil)

	err := usecase.Delete(context.Background(), productEditorSettingsID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductEditorSettingsUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.ProductEditorSettingsRepository)
	timeout := time.Duration(10)
	usecase := NewProductEditorSettingsUsecase(mockRepo, timeout)

	fetchedProductEditorSettings := []domain.ProductEditorSettings{
		{
			ID:                              primitive.NewObjectID(),
			ProductType:                     true,
			VisibleIndividually:             true,
			ProductTemplate:                 true,
			AdminComment:                    false,
			Vendor:                          true,
			Stores:                          true,
			ACL:                             false,
			ShowOnHomepage:                  true,
			AllowCustomerReviews:            true,
			ProductTags:                     true,
			ManufacturerPartNumber:          false,
			GTIN:                            true,
			ProductCost:                     true,
			TierPrices:                      false,
			Discounts:                       true,
			DisableBuyButton:                false,
			DisableWishlistButton:           false,
			AvailableForPreOrder:            true,
			CallForPrice:                    false,
			OldPrice:                        true,
			CustomerEntersPrice:             false,
			PAngV:                           false,
			RequireOtherProductsAddedToCart: false,
			IsGiftCard:                      false,
			DownloadableProduct:             true,
			RecurringProduct:                false,
			IsRental:                        false,
			FreeShipping:                    true,
			ShipSeparately:                  false,
			AdditionalShippingCharge:        true,
			DeliveryDate:                    true,
			ProductAvailabilityRange:        false,
			UseMultipleWarehouses:           true,
			Warehouse:                       true,
			DisplayStockAvailability:        true,
			MinimumStockQuantity:            true,
			LowStockActivity:                false,
			NotifyAdminForQuantityBelow:     true,
			Backorders:                      true,
			AllowBackInStockSubscriptions:   true,
			MinimumCartQuantity:             true,
			MaximumCartQuantity:             true,
			AllowedQuantities:               true,
			AllowAddingOnlyExistingAttributeCombinations: false,
			NotReturnable:           false,
			Weight:                  true,
			Dimensions:              true,
			AvailableStartDate:      true,
			AvailableEndDate:        true,
			MarkAsNew:               true,
			Published:               true,
			RelatedProducts:         true,
			CrossSellsProducts:      true,
			Seo:                     true,
			PurchasedWithOrders:     false,
			ProductAttributes:       true,
			SpecificationAttributes: true,
			Manufacturers:           true,
			StockQuantityChange:     true,
			AgeVerification:         false,
		},
		{
			ID:                              primitive.NewObjectID(),
			ProductType:                     false,
			VisibleIndividually:             false,
			ProductTemplate:                 false,
			AdminComment:                    true,
			Vendor:                          false,
			Stores:                          false,
			ACL:                             true,
			ShowOnHomepage:                  false,
			AllowCustomerReviews:            false,
			ProductTags:                     false,
			ManufacturerPartNumber:          true,
			GTIN:                            false,
			ProductCost:                     false,
			TierPrices:                      true,
			Discounts:                       false,
			DisableBuyButton:                true,
			DisableWishlistButton:           true,
			AvailableForPreOrder:            false,
			CallForPrice:                    true,
			OldPrice:                        false,
			CustomerEntersPrice:             true,
			PAngV:                           true,
			RequireOtherProductsAddedToCart: true,
			IsGiftCard:                      true,
			DownloadableProduct:             false,
			RecurringProduct:                true,
			IsRental:                        true,
			FreeShipping:                    false,
			ShipSeparately:                  true,
			AdditionalShippingCharge:        false,
			DeliveryDate:                    false,
			ProductAvailabilityRange:        true,
			UseMultipleWarehouses:           false,
			Warehouse:                       false,
			DisplayStockAvailability:        false,
			MinimumStockQuantity:            false,
			LowStockActivity:                true,
			NotifyAdminForQuantityBelow:     false,
			Backorders:                      false,
			AllowBackInStockSubscriptions:   false,
			MinimumCartQuantity:             false,
			MaximumCartQuantity:             false,
			AllowedQuantities:               false,
			AllowAddingOnlyExistingAttributeCombinations: true,
			NotReturnable:           true,
			Weight:                  false,
			Dimensions:              false,
			AvailableStartDate:      false,
			AvailableEndDate:        false,
			MarkAsNew:               false,
			Published:               false,
			RelatedProducts:         false,
			CrossSellsProducts:      false,
			Seo:                     false,
			PurchasedWithOrders:     true,
			ProductAttributes:       false,
			SpecificationAttributes: false,
			Manufacturers:           false,
			StockQuantityChange:     false,
			AgeVerification:         true,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedProductEditorSettings, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedProductEditorSettings, result)
	mockRepo.AssertExpectations(t)
}
