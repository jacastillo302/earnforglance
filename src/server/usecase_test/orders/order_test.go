package usecase_test

import (
	"context"
	"testing"
	"time"

	mocks "earnforglance/server/domain/mocks"
	domain "earnforglance/server/domain/orders"
	test "earnforglance/server/usecase/orders"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestOrderUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.OrderRepository)
	timeout := time.Duration(10)
	usecase := test.NewOrderUsecase(mockRepo, timeout)

	orderID := bson.NewObjectID().Hex()

	updatedOrder := domain.Order{
		ID:                                      bson.NewObjectID(), // Existing ID of the record to update
		OrderGuid:                               uuid.New(),
		StoreID:                                 bson.NewObjectID(),
		CustomerID:                              bson.NewObjectID(),
		BillingAddressID:                        bson.NewObjectID(),
		ShippingAddressID:                       new(bson.ObjectID),
		PickupAddressID:                         nil,
		PickupInStore:                           true,
		OrderStatusID:                           2,
		ShippingStatusID:                        3,
		PaymentStatusID:                         1,
		PaymentMethodSystemName:                 "PayPal",
		CustomerCurrencyCode:                    "EUR",
		CurrencyRate:                            0.85,
		CustomerTaxDisplayTypeID:                2,
		VatNumber:                               "987654321",
		OrderSubtotalInclTax:                    200.00,
		OrderSubtotalExclTax:                    180.00,
		OrderSubTotalDiscountInclTax:            20.00,
		OrderSubTotalDiscountExclTax:            18.00,
		OrderShippingInclTax:                    25.00,
		OrderShippingExclTax:                    22.00,
		PaymentMethodAdditionalFeeInclTax:       10.00,
		PaymentMethodAdditionalFeeExclTax:       9.00,
		TaxRates:                                "15%",
		OrderTax:                                30.00,
		OrderDiscount:                           10.00,
		OrderTotal:                              250.00,
		RefundedAmount:                          0.00,
		RewardPointsHistoryEntryID:              new(bson.ObjectID),
		CheckoutAttributeDescription:            "Gift Wrap and Note",
		CheckoutAttributesXml:                   "<Attributes><GiftWrap>Yes</GiftWrap><Note>Happy Birthday</Note></Attributes>",
		CustomerLanguageID:                      bson.NewObjectID(),
		AffiliateID:                             bson.NewObjectID(),
		CustomerIp:                              "192.168.1.2",
		AllowStoringCreditCardNumber:            true,
		CardType:                                "MasterCard",
		CardName:                                "Jane Doe",
		CardNumber:                              "5555555555554444",
		MaskedCreditCardNumber:                  "************4444",
		CardCvv2:                                "456",
		CardExpirationMonth:                     "06",
		CardExpirationYear:                      "2026",
		AuthorizationTransactionID:              "AUTH456",
		AuthorizationTransactionCode:            "AUTHCODE456",
		AuthorizationTransactionResult:          "Approved",
		CaptureTransactionID:                    "CAPTURE456",
		CaptureTransactionResult:                "Captured",
		SubscriptionTransactionID:               "SUB456",
		PaidDateUtc:                             new(time.Time),
		ShippingMethod:                          "Express",
		ShippingRateComputationMethodSystemName: "WeightBased",
		CustomValuesXml:                         "<CustomValues><GiftMessage>Happy Birthday!</GiftMessage></CustomValues>",
		Deleted:                                 false,
		CreatedOnUtc:                            time.Now().AddDate(0, 0, -7), // Created 7 days ago
		CustomOrderNumber:                       "ORD67890",
		RedeemedRewardPointsEntryID:             new(bson.ObjectID),
	}

	mockRepo.On("FetchByID", mock.Anything, orderID).Return(updatedOrder, nil)

	result, err := usecase.FetchByID(context.Background(), orderID)

	assert.NoError(t, err)
	assert.Equal(t, updatedOrder, result)
	mockRepo.AssertExpectations(t)
}

func TestOrderUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.OrderRepository)
	timeout := time.Duration(10)
	usecase := test.NewOrderUsecase(mockRepo, timeout)

	newOrder := &domain.Order{
		OrderGuid:                               uuid.New(),
		StoreID:                                 bson.NewObjectID(),
		CustomerID:                              bson.NewObjectID(),
		BillingAddressID:                        bson.NewObjectID(),
		ShippingAddressID:                       nil,
		PickupAddressID:                         nil,
		PickupInStore:                           false,
		OrderStatusID:                           1,
		ShippingStatusID:                        2,
		PaymentStatusID:                         3,
		PaymentMethodSystemName:                 "CreditCard",
		CustomerCurrencyCode:                    "USD",
		CurrencyRate:                            1.0,
		CustomerTaxDisplayTypeID:                1,
		VatNumber:                               "123456789",
		OrderSubtotalInclTax:                    100.00,
		OrderSubtotalExclTax:                    90.00,
		OrderSubTotalDiscountInclTax:            10.00,
		OrderSubTotalDiscountExclTax:            9.00,
		OrderShippingInclTax:                    15.00,
		OrderShippingExclTax:                    12.00,
		PaymentMethodAdditionalFeeInclTax:       5.00,
		PaymentMethodAdditionalFeeExclTax:       4.50,
		TaxRates:                                "10%",
		OrderTax:                                10.00,
		OrderDiscount:                           5.00,
		OrderTotal:                              120.00,
		RefundedAmount:                          0.00,
		RewardPointsHistoryEntryID:              nil,
		CheckoutAttributeDescription:            "Gift Wrap",
		CheckoutAttributesXml:                   "<Attributes><GiftWrap>Yes</GiftWrap></Attributes>",
		CustomerLanguageID:                      bson.NewObjectID(),
		AffiliateID:                             bson.NewObjectID(),
		CustomerIp:                              "192.168.1.1",
		AllowStoringCreditCardNumber:            false,
		CardType:                                "Visa",
		CardName:                                "John Doe",
		CardNumber:                              "4111111111111111",
		MaskedCreditCardNumber:                  "************1111",
		CardCvv2:                                "123",
		CardExpirationMonth:                     "12",
		CardExpirationYear:                      "2025",
		AuthorizationTransactionID:              "AUTH123",
		AuthorizationTransactionCode:            "AUTHCODE123",
		AuthorizationTransactionResult:          "Approved",
		CaptureTransactionID:                    "CAPTURE123",
		CaptureTransactionResult:                "Captured",
		SubscriptionTransactionID:               "SUB123",
		PaidDateUtc:                             nil,
		ShippingMethod:                          "Standard",
		ShippingRateComputationMethodSystemName: "FlatRate",
		CustomValuesXml:                         "<CustomValues></CustomValues>",
		Deleted:                                 false,
		CreatedOnUtc:                            time.Now(),
		CustomOrderNumber:                       "ORD12345",
		RedeemedRewardPointsEntryID:             nil,
	}

	mockRepo.On("Create", mock.Anything, newOrder).Return(nil)

	err := usecase.Create(context.Background(), newOrder)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestOrderUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.OrderRepository)
	timeout := time.Duration(10)
	usecase := test.NewOrderUsecase(mockRepo, timeout)

	updatedOrder := &domain.Order{
		ID:                                      bson.NewObjectID(), // Existing ID of the record to update
		OrderGuid:                               uuid.New(),
		StoreID:                                 bson.NewObjectID(),
		CustomerID:                              bson.NewObjectID(),
		BillingAddressID:                        bson.NewObjectID(),
		ShippingAddressID:                       new(bson.ObjectID),
		PickupAddressID:                         nil,
		PickupInStore:                           true,
		OrderStatusID:                           2,
		ShippingStatusID:                        3,
		PaymentStatusID:                         1,
		PaymentMethodSystemName:                 "PayPal",
		CustomerCurrencyCode:                    "EUR",
		CurrencyRate:                            0.85,
		CustomerTaxDisplayTypeID:                2,
		VatNumber:                               "987654321",
		OrderSubtotalInclTax:                    200.00,
		OrderSubtotalExclTax:                    180.00,
		OrderSubTotalDiscountInclTax:            20.00,
		OrderSubTotalDiscountExclTax:            18.00,
		OrderShippingInclTax:                    25.00,
		OrderShippingExclTax:                    22.00,
		PaymentMethodAdditionalFeeInclTax:       10.00,
		PaymentMethodAdditionalFeeExclTax:       9.00,
		TaxRates:                                "15%",
		OrderTax:                                30.00,
		OrderDiscount:                           10.00,
		OrderTotal:                              250.00,
		RefundedAmount:                          0.00,
		RewardPointsHistoryEntryID:              new(bson.ObjectID),
		CheckoutAttributeDescription:            "Gift Wrap and Note",
		CheckoutAttributesXml:                   "<Attributes><GiftWrap>Yes</GiftWrap><Note>Happy Birthday</Note></Attributes>",
		CustomerLanguageID:                      bson.NewObjectID(),
		AffiliateID:                             bson.NewObjectID(),
		CustomerIp:                              "192.168.1.2",
		AllowStoringCreditCardNumber:            true,
		CardType:                                "MasterCard",
		CardName:                                "Jane Doe",
		CardNumber:                              "5555555555554444",
		MaskedCreditCardNumber:                  "************4444",
		CardCvv2:                                "456",
		CardExpirationMonth:                     "06",
		CardExpirationYear:                      "2026",
		AuthorizationTransactionID:              "AUTH456",
		AuthorizationTransactionCode:            "AUTHCODE456",
		AuthorizationTransactionResult:          "Approved",
		CaptureTransactionID:                    "CAPTURE456",
		CaptureTransactionResult:                "Captured",
		SubscriptionTransactionID:               "SUB456",
		PaidDateUtc:                             new(time.Time),
		ShippingMethod:                          "Express",
		ShippingRateComputationMethodSystemName: "WeightBased",
		CustomValuesXml:                         "<CustomValues><GiftMessage>Happy Birthday!</GiftMessage></CustomValues>",
		Deleted:                                 false,
		CreatedOnUtc:                            time.Now().AddDate(0, 0, -7), // Created 7 days ago
		CustomOrderNumber:                       "ORD67890",
		RedeemedRewardPointsEntryID:             new(bson.ObjectID),
	}
	*updatedOrder.ShippingAddressID = bson.NewObjectID()
	*updatedOrder.RewardPointsHistoryEntryID = bson.NewObjectID()
	*updatedOrder.PaidDateUtc = time.Now().AddDate(0, 0, -1) // Paid 1 day ago

	mockRepo.On("Update", mock.Anything, updatedOrder).Return(nil)

	err := usecase.Update(context.Background(), updatedOrder)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestOrderUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.OrderRepository)
	timeout := time.Duration(10)
	usecase := test.NewOrderUsecase(mockRepo, timeout)

	orderID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, orderID).Return(nil)

	err := usecase.Delete(context.Background(), orderID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestOrderUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.OrderRepository)
	timeout := time.Duration(10)
	usecase := test.NewOrderUsecase(mockRepo, timeout)

	fetchedOrders := []domain.Order{
		{
			ID:                                      bson.NewObjectID(),
			OrderGuid:                               uuid.New(),
			StoreID:                                 bson.NewObjectID(),
			CustomerID:                              bson.NewObjectID(),
			BillingAddressID:                        bson.NewObjectID(),
			ShippingAddressID:                       nil,
			PickupAddressID:                         nil,
			PickupInStore:                           false,
			OrderStatusID:                           1,
			ShippingStatusID:                        2,
			PaymentStatusID:                         3,
			PaymentMethodSystemName:                 "CreditCard",
			CustomerCurrencyCode:                    "USD",
			CurrencyRate:                            1.0,
			CustomerTaxDisplayTypeID:                1,
			VatNumber:                               "123456789",
			OrderSubtotalInclTax:                    100.00,
			OrderSubtotalExclTax:                    90.00,
			OrderSubTotalDiscountInclTax:            10.00,
			OrderSubTotalDiscountExclTax:            9.00,
			OrderShippingInclTax:                    15.00,
			OrderShippingExclTax:                    12.00,
			PaymentMethodAdditionalFeeInclTax:       5.00,
			PaymentMethodAdditionalFeeExclTax:       4.50,
			TaxRates:                                "10%",
			OrderTax:                                10.00,
			OrderDiscount:                           5.00,
			OrderTotal:                              120.00,
			RefundedAmount:                          0.00,
			RewardPointsHistoryEntryID:              nil,
			CheckoutAttributeDescription:            "Gift Wrap",
			CheckoutAttributesXml:                   "<Attributes><GiftWrap>Yes</GiftWrap></Attributes>",
			CustomerLanguageID:                      bson.NewObjectID(),
			AffiliateID:                             bson.NewObjectID(),
			CustomerIp:                              "192.168.1.1",
			AllowStoringCreditCardNumber:            false,
			CardType:                                "Visa",
			CardName:                                "John Doe",
			CardNumber:                              "4111111111111111",
			MaskedCreditCardNumber:                  "************1111",
			CardCvv2:                                "123",
			CardExpirationMonth:                     "12",
			CardExpirationYear:                      "2025",
			AuthorizationTransactionID:              "AUTH123",
			AuthorizationTransactionCode:            "AUTHCODE123",
			AuthorizationTransactionResult:          "Approved",
			CaptureTransactionID:                    "CAPTURE123",
			CaptureTransactionResult:                "Captured",
			SubscriptionTransactionID:               "SUB123",
			PaidDateUtc:                             nil,
			ShippingMethod:                          "Standard",
			ShippingRateComputationMethodSystemName: "FlatRate",
			CustomValuesXml:                         "<CustomValues></CustomValues>",
			Deleted:                                 false,
			CreatedOnUtc:                            time.Now().AddDate(0, 0, -10), // Created 10 days ago
			CustomOrderNumber:                       "ORD12345",
			RedeemedRewardPointsEntryID:             nil,
		},
		{
			ID:                                      bson.NewObjectID(),
			OrderGuid:                               uuid.New(),
			StoreID:                                 bson.NewObjectID(),
			CustomerID:                              bson.NewObjectID(),
			BillingAddressID:                        bson.NewObjectID(),
			ShippingAddressID:                       new(bson.ObjectID),
			PickupAddressID:                         nil,
			PickupInStore:                           true,
			OrderStatusID:                           2,
			ShippingStatusID:                        3,
			PaymentStatusID:                         1,
			PaymentMethodSystemName:                 "PayPal",
			CustomerCurrencyCode:                    "EUR",
			CurrencyRate:                            0.85,
			CustomerTaxDisplayTypeID:                2,
			VatNumber:                               "987654321",
			OrderSubtotalInclTax:                    200.00,
			OrderSubtotalExclTax:                    180.00,
			OrderSubTotalDiscountInclTax:            20.00,
			OrderSubTotalDiscountExclTax:            18.00,
			OrderShippingInclTax:                    25.00,
			OrderShippingExclTax:                    22.00,
			PaymentMethodAdditionalFeeInclTax:       10.00,
			PaymentMethodAdditionalFeeExclTax:       9.00,
			TaxRates:                                "15%",
			OrderTax:                                30.00,
			OrderDiscount:                           10.00,
			OrderTotal:                              250.00,
			RefundedAmount:                          0.00,
			RewardPointsHistoryEntryID:              new(bson.ObjectID),
			CheckoutAttributeDescription:            "Gift Wrap and Note",
			CheckoutAttributesXml:                   "<Attributes><GiftWrap>Yes</GiftWrap><Note>Happy Birthday</Note></Attributes>",
			CustomerLanguageID:                      bson.NewObjectID(),
			AffiliateID:                             bson.NewObjectID(),
			CustomerIp:                              "192.168.1.2",
			AllowStoringCreditCardNumber:            true,
			CardType:                                "MasterCard",
			CardName:                                "Jane Doe",
			CardNumber:                              "5555555555554444",
			MaskedCreditCardNumber:                  "************4444",
			CardCvv2:                                "456",
			CardExpirationMonth:                     "06",
			CardExpirationYear:                      "2026",
			AuthorizationTransactionID:              "AUTH456",
			AuthorizationTransactionCode:            "AUTHCODE456",
			AuthorizationTransactionResult:          "Approved",
			CaptureTransactionID:                    "CAPTURE456",
			CaptureTransactionResult:                "Captured",
			SubscriptionTransactionID:               "SUB456",
			PaidDateUtc:                             new(time.Time),
			ShippingMethod:                          "Express",
			ShippingRateComputationMethodSystemName: "WeightBased",
			CustomValuesXml:                         "<CustomValues><GiftMessage>Happy Birthday!</GiftMessage></CustomValues>",
			Deleted:                                 false,
			CreatedOnUtc:                            time.Now().AddDate(0, 0, -5), // Created 5 days ago
			CustomOrderNumber:                       "ORD67890",
			RedeemedRewardPointsEntryID:             new(bson.ObjectID),
		},
	}
	*fetchedOrders[1].ShippingAddressID = bson.NewObjectID()
	*fetchedOrders[1].RewardPointsHistoryEntryID = bson.NewObjectID()
	*fetchedOrders[1].PaidDateUtc = time.Now().AddDate(0, 0, -1) // Paid 1 day ago

	mockRepo.On("Fetch", mock.Anything).Return(fetchedOrders, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedOrders, result)
	mockRepo.AssertExpectations(t)
}
