package usecase

import (
	"context"
	mocks "earnforglance/server/domain/mocks"
	domain "earnforglance/server/domain/orders"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestShoppingCartSettingsUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.ShoppingCartSettingsRepository)
	timeout := time.Duration(10)
	usecase := NewShoppingCartSettingsUsecase(mockRepo, timeout)

	shoppingCartSettingsID := primitive.NewObjectID().Hex()

	updatedShoppingCartSettings := domain.ShoppingCartSettings{
		ID:                                          primitive.NewObjectID(), // Existing ID of the record to update
		DisplayCartAfterAddingProduct:               false,
		DisplayWishlistAfterAddingProduct:           true,
		MaximumShoppingCartItems:                    100,
		MaximumWishlistItems:                        50,
		AllowOutOfStockItemsToBeAddedToWishlist:     false,
		MoveItemsFromWishlistToCart:                 false,
		CartsSharedBetweenStores:                    true,
		ShowProductImagesOnShoppingCart:             false,
		ShowProductImagesOnWishList:                 false,
		ShowDiscountBox:                             false,
		ShowGiftCardBox:                             true,
		CrossSellsNumber:                            10,
		EmailWishlistEnabled:                        false,
		AllowAnonymousUsersToEmailWishlist:          true,
		MiniShoppingCartEnabled:                     false,
		ShowProductImagesInMiniShoppingCart:         false,
		MiniShoppingCartProductNumber:               5,
		RoundPricesDuringCalculation:                false,
		GroupTierPricesForDistinctShoppingCartItems: true,
		AllowCartItemEditing:                        false,
		RenderAssociatedAttributeValueQuantity:      true,
	}

	mockRepo.On("FetchByID", mock.Anything, shoppingCartSettingsID).Return(updatedShoppingCartSettings, nil)

	result, err := usecase.FetchByID(context.Background(), shoppingCartSettingsID)

	assert.NoError(t, err)
	assert.Equal(t, updatedShoppingCartSettings, result)
	mockRepo.AssertExpectations(t)
}

func TestShoppingCartSettingsUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.ShoppingCartSettingsRepository)
	timeout := time.Duration(10)
	usecase := NewShoppingCartSettingsUsecase(mockRepo, timeout)

	newShoppingCartSettings := &domain.ShoppingCartSettings{
		DisplayCartAfterAddingProduct:               true,
		DisplayWishlistAfterAddingProduct:           false,
		MaximumShoppingCartItems:                    50,
		MaximumWishlistItems:                        20,
		AllowOutOfStockItemsToBeAddedToWishlist:     true,
		MoveItemsFromWishlistToCart:                 true,
		CartsSharedBetweenStores:                    false,
		ShowProductImagesOnShoppingCart:             true,
		ShowProductImagesOnWishList:                 true,
		ShowDiscountBox:                             true,
		ShowGiftCardBox:                             false,
		CrossSellsNumber:                            5,
		EmailWishlistEnabled:                        true,
		AllowAnonymousUsersToEmailWishlist:          false,
		MiniShoppingCartEnabled:                     true,
		ShowProductImagesInMiniShoppingCart:         true,
		MiniShoppingCartProductNumber:               3,
		RoundPricesDuringCalculation:                true,
		GroupTierPricesForDistinctShoppingCartItems: false,
		AllowCartItemEditing:                        true,
		RenderAssociatedAttributeValueQuantity:      false,
	}

	mockRepo.On("Create", mock.Anything, newShoppingCartSettings).Return(nil)

	err := usecase.Create(context.Background(), newShoppingCartSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestShoppingCartSettingsUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.ShoppingCartSettingsRepository)
	timeout := time.Duration(10)
	usecase := NewShoppingCartSettingsUsecase(mockRepo, timeout)

	updatedShoppingCartSettings := &domain.ShoppingCartSettings{
		ID:                                          primitive.NewObjectID(), // Existing ID of the record to update
		DisplayCartAfterAddingProduct:               false,
		DisplayWishlistAfterAddingProduct:           true,
		MaximumShoppingCartItems:                    100,
		MaximumWishlistItems:                        50,
		AllowOutOfStockItemsToBeAddedToWishlist:     false,
		MoveItemsFromWishlistToCart:                 false,
		CartsSharedBetweenStores:                    true,
		ShowProductImagesOnShoppingCart:             false,
		ShowProductImagesOnWishList:                 false,
		ShowDiscountBox:                             false,
		ShowGiftCardBox:                             true,
		CrossSellsNumber:                            10,
		EmailWishlistEnabled:                        false,
		AllowAnonymousUsersToEmailWishlist:          true,
		MiniShoppingCartEnabled:                     false,
		ShowProductImagesInMiniShoppingCart:         false,
		MiniShoppingCartProductNumber:               5,
		RoundPricesDuringCalculation:                false,
		GroupTierPricesForDistinctShoppingCartItems: true,
		AllowCartItemEditing:                        false,
		RenderAssociatedAttributeValueQuantity:      true,
	}

	mockRepo.On("Update", mock.Anything, updatedShoppingCartSettings).Return(nil)

	err := usecase.Update(context.Background(), updatedShoppingCartSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestShoppingCartSettingsUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.ShoppingCartSettingsRepository)
	timeout := time.Duration(10)
	usecase := NewShoppingCartSettingsUsecase(mockRepo, timeout)

	shoppingCartSettingsID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, shoppingCartSettingsID).Return(nil)

	err := usecase.Delete(context.Background(), shoppingCartSettingsID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestShoppingCartSettingsUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.ShoppingCartSettingsRepository)
	timeout := time.Duration(10)
	usecase := NewShoppingCartSettingsUsecase(mockRepo, timeout)

	fetchedShoppingCartSettings := []domain.ShoppingCartSettings{
		{
			ID:                                          primitive.NewObjectID(),
			DisplayCartAfterAddingProduct:               true,
			DisplayWishlistAfterAddingProduct:           false,
			MaximumShoppingCartItems:                    50,
			MaximumWishlistItems:                        20,
			AllowOutOfStockItemsToBeAddedToWishlist:     true,
			MoveItemsFromWishlistToCart:                 true,
			CartsSharedBetweenStores:                    false,
			ShowProductImagesOnShoppingCart:             true,
			ShowProductImagesOnWishList:                 true,
			ShowDiscountBox:                             true,
			ShowGiftCardBox:                             false,
			CrossSellsNumber:                            5,
			EmailWishlistEnabled:                        true,
			AllowAnonymousUsersToEmailWishlist:          false,
			MiniShoppingCartEnabled:                     true,
			ShowProductImagesInMiniShoppingCart:         true,
			MiniShoppingCartProductNumber:               3,
			RoundPricesDuringCalculation:                true,
			GroupTierPricesForDistinctShoppingCartItems: false,
			AllowCartItemEditing:                        true,
			RenderAssociatedAttributeValueQuantity:      false,
		},
		{
			ID:                                          primitive.NewObjectID(),
			DisplayCartAfterAddingProduct:               false,
			DisplayWishlistAfterAddingProduct:           true,
			MaximumShoppingCartItems:                    100,
			MaximumWishlistItems:                        50,
			AllowOutOfStockItemsToBeAddedToWishlist:     false,
			MoveItemsFromWishlistToCart:                 false,
			CartsSharedBetweenStores:                    true,
			ShowProductImagesOnShoppingCart:             false,
			ShowProductImagesOnWishList:                 false,
			ShowDiscountBox:                             false,
			ShowGiftCardBox:                             true,
			CrossSellsNumber:                            10,
			EmailWishlistEnabled:                        false,
			AllowAnonymousUsersToEmailWishlist:          true,
			MiniShoppingCartEnabled:                     false,
			ShowProductImagesInMiniShoppingCart:         false,
			MiniShoppingCartProductNumber:               5,
			RoundPricesDuringCalculation:                false,
			GroupTierPricesForDistinctShoppingCartItems: true,
			AllowCartItemEditing:                        false,
			RenderAssociatedAttributeValueQuantity:      true,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedShoppingCartSettings, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedShoppingCartSettings, result)
	mockRepo.AssertExpectations(t)
}
