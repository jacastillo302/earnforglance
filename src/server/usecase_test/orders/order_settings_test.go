package usecase_test

import (
	"context"
	mocks "earnforglance/server/domain/mocks"
	domain "earnforglance/server/domain/orders"
	test "earnforglance/server/usecase/orders"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestOrderSettingsUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.OrderSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewOrderSettingsUsecase(mockRepo, timeout)

	orderSettingsID := primitive.NewObjectID().Hex()

	updatedOrderSettings := domain.OrderSettings{
		ID:                                  primitive.NewObjectID(), // Existing ID of the record to update
		IsReOrderAllowed:                    false,
		MinOrderSubtotalAmount:              75.00,
		MinOrderSubtotalAmountIncludingTax:  false,
		MinOrderTotalAmount:                 150.00,
		AutoUpdateOrderTotalsOnEditingOrder: false,
		AnonymousCheckoutAllowed:            false,
		CheckoutDisabled:                    true,
		TermsOfServiceOnShoppingCartPage:    false,
		TermsOfServiceOnOrderConfirmPage:    true,
		OnePageCheckoutEnabled:              false,
		OnePageCheckoutDisplayOrderTotalsOnPaymentInfoTab: false,
		DisableBillingAddressCheckoutStep:                 true,
		DisableOrderCompletedPage:                         true,
		DisplayPickupInStoreOnShippingMethodPage:          false,
		AttachPdfInvoiceToOrderPlacedEmail:                false,
		AttachPdfInvoiceToOrderPaidEmail:                  true,
		AttachPdfInvoiceToOrderProcessingEmail:            false,
		AttachPdfInvoiceToOrderCompletedEmail:             false,
		GeneratePdfInvoiceInCustomerLanguage:              false,
		ReturnRequestsEnabled:                             false,
		ReturnRequestsAllowFiles:                          false,
		ReturnRequestsFileMaximumSize:                     1024,
		ReturnRequestNumberMask:                           "UPDATED-RR-{0}",
		NumberOfDaysReturnRequestAvailable:                15,
		ActivateGiftCardsAfterCompletingOrder:             false,
		DeactivateGiftCardsAfterCancellingOrder:           true,
		DeactivateGiftCardsAfterDeletingOrder:             false,
		MinimumOrderPlacementInterval:                     20,
		CompleteOrderWhenDelivered:                        false,
		CustomOrderNumberMask:                             "UPDATED-ORD-{0}",
		ExportWithProducts:                                false,
		AllowAdminsToBuyCallForPriceProducts:              true,
		ShowProductThumbnailInOrderDetailsPage:            false,
		DeleteGiftCardUsageHistory:                        true,
		DisplayCustomerCurrencyOnOrders:                   false,
		DisplayOrderSummary:                               false,
		PlaceOrderWithLock:                                true,
	}

	mockRepo.On("FetchByID", mock.Anything, orderSettingsID).Return(updatedOrderSettings, nil)

	result, err := usecase.FetchByID(context.Background(), orderSettingsID)

	assert.NoError(t, err)
	assert.Equal(t, updatedOrderSettings, result)
	mockRepo.AssertExpectations(t)
}

func TestOrderSettingsUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.OrderSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewOrderSettingsUsecase(mockRepo, timeout)

	newOrderSettings := &domain.OrderSettings{
		IsReOrderAllowed:                                  true,
		MinOrderSubtotalAmount:                            50.00,
		MinOrderSubtotalAmountIncludingTax:                true,
		MinOrderTotalAmount:                               100.00,
		AutoUpdateOrderTotalsOnEditingOrder:               true,
		AnonymousCheckoutAllowed:                          true,
		CheckoutDisabled:                                  false,
		TermsOfServiceOnShoppingCartPage:                  true,
		TermsOfServiceOnOrderConfirmPage:                  false,
		OnePageCheckoutEnabled:                            true,
		OnePageCheckoutDisplayOrderTotalsOnPaymentInfoTab: true,
		DisableBillingAddressCheckoutStep:                 false,
		DisableOrderCompletedPage:                         false,
		DisplayPickupInStoreOnShippingMethodPage:          true,
		AttachPdfInvoiceToOrderPlacedEmail:                true,
		AttachPdfInvoiceToOrderPaidEmail:                  false,
		AttachPdfInvoiceToOrderProcessingEmail:            true,
		AttachPdfInvoiceToOrderCompletedEmail:             true,
		GeneratePdfInvoiceInCustomerLanguage:              true,
		ReturnRequestsEnabled:                             true,
		ReturnRequestsAllowFiles:                          true,
		ReturnRequestsFileMaximumSize:                     2048,
		ReturnRequestNumberMask:                           "RR-{0}",
		NumberOfDaysReturnRequestAvailable:                30,
		ActivateGiftCardsAfterCompletingOrder:             true,
		DeactivateGiftCardsAfterCancellingOrder:           false,
		DeactivateGiftCardsAfterDeletingOrder:             true,
		MinimumOrderPlacementInterval:                     10,
		CompleteOrderWhenDelivered:                        true,
		CustomOrderNumberMask:                             "ORD-{0}",
		ExportWithProducts:                                true,
		AllowAdminsToBuyCallForPriceProducts:              false,
		ShowProductThumbnailInOrderDetailsPage:            true,
		DeleteGiftCardUsageHistory:                        false,
		DisplayCustomerCurrencyOnOrders:                   true,
		DisplayOrderSummary:                               true,
		PlaceOrderWithLock:                                false,
	}
	mockRepo.On("Create", mock.Anything, newOrderSettings).Return(nil)

	err := usecase.Create(context.Background(), newOrderSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestOrderSettingsUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.OrderSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewOrderSettingsUsecase(mockRepo, timeout)

	updatedOrderSettings := &domain.OrderSettings{
		ID:                                  primitive.NewObjectID(), // Existing ID of the record to update
		IsReOrderAllowed:                    false,
		MinOrderSubtotalAmount:              75.00,
		MinOrderSubtotalAmountIncludingTax:  false,
		MinOrderTotalAmount:                 150.00,
		AutoUpdateOrderTotalsOnEditingOrder: false,
		AnonymousCheckoutAllowed:            false,
		CheckoutDisabled:                    true,
		TermsOfServiceOnShoppingCartPage:    false,
		TermsOfServiceOnOrderConfirmPage:    true,
		OnePageCheckoutEnabled:              false,
		OnePageCheckoutDisplayOrderTotalsOnPaymentInfoTab: false,
		DisableBillingAddressCheckoutStep:                 true,
		DisableOrderCompletedPage:                         true,
		DisplayPickupInStoreOnShippingMethodPage:          false,
		AttachPdfInvoiceToOrderPlacedEmail:                false,
		AttachPdfInvoiceToOrderPaidEmail:                  true,
		AttachPdfInvoiceToOrderProcessingEmail:            false,
		AttachPdfInvoiceToOrderCompletedEmail:             false,
		GeneratePdfInvoiceInCustomerLanguage:              false,
		ReturnRequestsEnabled:                             false,
		ReturnRequestsAllowFiles:                          false,
		ReturnRequestsFileMaximumSize:                     1024,
		ReturnRequestNumberMask:                           "UPDATED-RR-{0}",
		NumberOfDaysReturnRequestAvailable:                15,
		ActivateGiftCardsAfterCompletingOrder:             false,
		DeactivateGiftCardsAfterCancellingOrder:           true,
		DeactivateGiftCardsAfterDeletingOrder:             false,
		MinimumOrderPlacementInterval:                     20,
		CompleteOrderWhenDelivered:                        false,
		CustomOrderNumberMask:                             "UPDATED-ORD-{0}",
		ExportWithProducts:                                false,
		AllowAdminsToBuyCallForPriceProducts:              true,
		ShowProductThumbnailInOrderDetailsPage:            false,
		DeleteGiftCardUsageHistory:                        true,
		DisplayCustomerCurrencyOnOrders:                   false,
		DisplayOrderSummary:                               false,
		PlaceOrderWithLock:                                true,
	}

	mockRepo.On("Update", mock.Anything, updatedOrderSettings).Return(nil)

	err := usecase.Update(context.Background(), updatedOrderSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestOrderSettingsUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.OrderSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewOrderSettingsUsecase(mockRepo, timeout)

	orderSettingsID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, orderSettingsID).Return(nil)

	err := usecase.Delete(context.Background(), orderSettingsID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestOrderSettingsUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.OrderSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewOrderSettingsUsecase(mockRepo, timeout)

	fetchedOrderSettings := []domain.OrderSettings{
		{
			ID:                                  primitive.NewObjectID(),
			IsReOrderAllowed:                    true,
			MinOrderSubtotalAmount:              50.00,
			MinOrderSubtotalAmountIncludingTax:  true,
			MinOrderTotalAmount:                 100.00,
			AutoUpdateOrderTotalsOnEditingOrder: true,
			AnonymousCheckoutAllowed:            true,
			CheckoutDisabled:                    false,
			TermsOfServiceOnShoppingCartPage:    true,
			TermsOfServiceOnOrderConfirmPage:    false,
			OnePageCheckoutEnabled:              true,
			OnePageCheckoutDisplayOrderTotalsOnPaymentInfoTab: true,
			DisableBillingAddressCheckoutStep:                 false,
			DisableOrderCompletedPage:                         false,
			DisplayPickupInStoreOnShippingMethodPage:          true,
			AttachPdfInvoiceToOrderPlacedEmail:                true,
			AttachPdfInvoiceToOrderPaidEmail:                  false,
			AttachPdfInvoiceToOrderProcessingEmail:            true,
			AttachPdfInvoiceToOrderCompletedEmail:             true,
			GeneratePdfInvoiceInCustomerLanguage:              true,
			ReturnRequestsEnabled:                             true,
			ReturnRequestsAllowFiles:                          true,
			ReturnRequestsFileMaximumSize:                     2048,
			ReturnRequestNumberMask:                           "RR-{0}",
			NumberOfDaysReturnRequestAvailable:                30,
			ActivateGiftCardsAfterCompletingOrder:             true,
			DeactivateGiftCardsAfterCancellingOrder:           false,
			DeactivateGiftCardsAfterDeletingOrder:             true,
			MinimumOrderPlacementInterval:                     10,
			CompleteOrderWhenDelivered:                        true,
			CustomOrderNumberMask:                             "ORD-{0}",
			ExportWithProducts:                                true,
			AllowAdminsToBuyCallForPriceProducts:              false,
			ShowProductThumbnailInOrderDetailsPage:            true,
			DeleteGiftCardUsageHistory:                        false,
			DisplayCustomerCurrencyOnOrders:                   true,
			DisplayOrderSummary:                               true,
			PlaceOrderWithLock:                                false,
		},
		{
			ID:                                  primitive.NewObjectID(),
			IsReOrderAllowed:                    false,
			MinOrderSubtotalAmount:              75.00,
			MinOrderSubtotalAmountIncludingTax:  false,
			MinOrderTotalAmount:                 150.00,
			AutoUpdateOrderTotalsOnEditingOrder: false,
			AnonymousCheckoutAllowed:            false,
			CheckoutDisabled:                    true,
			TermsOfServiceOnShoppingCartPage:    false,
			TermsOfServiceOnOrderConfirmPage:    true,
			OnePageCheckoutEnabled:              false,
			OnePageCheckoutDisplayOrderTotalsOnPaymentInfoTab: false,
			DisableBillingAddressCheckoutStep:                 true,
			DisableOrderCompletedPage:                         true,
			DisplayPickupInStoreOnShippingMethodPage:          false,
			AttachPdfInvoiceToOrderPlacedEmail:                false,
			AttachPdfInvoiceToOrderPaidEmail:                  true,
			AttachPdfInvoiceToOrderProcessingEmail:            false,
			AttachPdfInvoiceToOrderCompletedEmail:             false,
			GeneratePdfInvoiceInCustomerLanguage:              false,
			ReturnRequestsEnabled:                             false,
			ReturnRequestsAllowFiles:                          false,
			ReturnRequestsFileMaximumSize:                     1024,
			ReturnRequestNumberMask:                           "UPDATED-RR-{0}",
			NumberOfDaysReturnRequestAvailable:                15,
			ActivateGiftCardsAfterCompletingOrder:             false,
			DeactivateGiftCardsAfterCancellingOrder:           true,
			DeactivateGiftCardsAfterDeletingOrder:             false,
			MinimumOrderPlacementInterval:                     20,
			CompleteOrderWhenDelivered:                        false,
			CustomOrderNumberMask:                             "UPDATED-ORD-{0}",
			ExportWithProducts:                                false,
			AllowAdminsToBuyCallForPriceProducts:              true,
			ShowProductThumbnailInOrderDetailsPage:            false,
			DeleteGiftCardUsageHistory:                        true,
			DisplayCustomerCurrencyOnOrders:                   false,
			DisplayOrderSummary:                               false,
			PlaceOrderWithLock:                                true,
		},
	}
	mockRepo.On("Fetch", mock.Anything).Return(fetchedOrderSettings, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedOrderSettings, result)
	mockRepo.AssertExpectations(t)
}
