package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/orders"
	"earnforglance/server/mongo/mocks"
	repository "earnforglance/server/repository/orders"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultOrderSettings struct {
	mock.Mock
}

func (m *MockSingleResultOrderSettings) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.OrderSettings); ok {
		*v.(*domain.OrderSettings) = *result
	}
	return args.Error(1)
}

var mockItemOrderSettings = &domain.OrderSettings{
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

func TestOrderSettingsRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionOrderSettings

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultOrderSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemOrderSettings, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewOrderSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemOrderSettings.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultOrderSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewOrderSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemOrderSettings.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestOrderSettingsRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionOrderSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemOrderSettings).Return(nil, nil).Once()

	repo := repository.NewOrderSettingsRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemOrderSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestOrderSettingsRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionOrderSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemOrderSettings.ID}
	update := bson.M{"$set": mockItemOrderSettings}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewOrderSettingsRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemOrderSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
