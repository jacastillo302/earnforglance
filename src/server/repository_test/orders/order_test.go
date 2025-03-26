package repository_test

import (
	"context"
	"errors"
	"testing"
	"time"

	domain "earnforglance/server/domain/orders"
	repository "earnforglance/server/repository/orders"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultOrder struct {
	mock.Mock
}

func (m *MockSingleResultOrder) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.Order); ok {
		*v.(*domain.Order) = *result
	}
	return args.Error(1)
}

var mockItemOrder = &domain.Order{
	ID:                                      primitive.NewObjectID(), // Existing ID of the record to update
	OrderGuid:                               uuid.New(),
	StoreID:                                 primitive.NewObjectID(),
	CustomerID:                              primitive.NewObjectID(),
	BillingAddressID:                        primitive.NewObjectID(),
	ShippingAddressID:                       new(primitive.ObjectID),
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
	RewardPointsHistoryEntryID:              new(primitive.ObjectID),
	CheckoutAttributeDescription:            "Gift Wrap and Note",
	CheckoutAttributesXml:                   "<Attributes><GiftWrap>Yes</GiftWrap><Note>Happy Birthday</Note></Attributes>",
	CustomerLanguageID:                      primitive.NewObjectID(),
	AffiliateID:                             primitive.NewObjectID(),
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
	RedeemedRewardPointsEntryID:             new(primitive.ObjectID),
}

func TestOrderRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionOrder

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultOrder{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemOrder, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewOrderRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemOrder.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultOrder{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewOrderRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemOrder.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestOrderRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionOrder

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemOrder).Return(nil, nil).Once()

	repo := repository.NewOrderRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemOrder)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestOrderRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionOrder

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemOrder.ID}
	update := bson.M{"$set": mockItemOrder}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewOrderRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemOrder)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
