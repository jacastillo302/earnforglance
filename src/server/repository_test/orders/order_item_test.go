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

	"go.mongodb.org/mongo-driver/v2/bson"
)

type MockSingleResultOrderItem struct {
	mock.Mock
}

func (m *MockSingleResultOrderItem) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.OrderItem); ok {
		*v.(*domain.OrderItem) = *result
	}
	return args.Error(1)
}

var mockItemOrderItem = &domain.OrderItem{
	ID:                    bson.NewObjectID(), // Existing ID of the record to update
	OrderItemGuid:         uuid.New(),
	OrderID:               bson.NewObjectID(),
	ProductID:             bson.NewObjectID(),
	Quantity:              3,
	UnitPriceInclTax:      60.00,
	UnitPriceExclTax:      55.00,
	PriceInclTax:          180.00,
	PriceExclTax:          165.00,
	DiscountAmountInclTax: 15.00,
	DiscountAmountExclTax: 13.50,
	OriginalProductCost:   50.00,
	AttributeDescription:  "Color: Blue, Size: L",
	AttributesXml:         "<Attributes><Color>Blue</Color><Size>L</Size></Attributes>",
	DownloadCount:         1,
	IsDownloadActivated:   true,
	LicenseDownloadID:     new(bson.ObjectID),
	ItemWeight:            new(float64),
	RentalStartDateUtc:    new(time.Time),
	RentalEndDateUtc:      new(time.Time),
}

func TestOrderItemRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionOrderItem

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultOrderItem{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemOrderItem, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewOrderItemRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemOrderItem.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultOrderItem{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewOrderItemRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemOrderItem.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestOrderItemRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionOrderItem

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemOrderItem).Return(nil, nil).Once()

	repo := repository.NewOrderItemRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemOrderItem)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestOrderItemRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionOrderItem

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemOrderItem.ID}
	update := bson.M{"$set": mockItemOrderItem}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewOrderItemRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemOrderItem)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
