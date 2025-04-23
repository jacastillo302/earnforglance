package usecase_test

import (
	"context"
	mocks "earnforglance/server/domain/mocks"
	domain "earnforglance/server/domain/orders"
	test "earnforglance/server/usecase/orders"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestOrderItemUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.OrderItemRepository)
	timeout := time.Duration(10)
	usecase := test.NewOrderItemUsecase(mockRepo, timeout)

	orderItemID := bson.NewObjectID().Hex()

	updatedOrderItem := domain.OrderItem{
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

	mockRepo.On("FetchByID", mock.Anything, orderItemID).Return(updatedOrderItem, nil)

	result, err := usecase.FetchByID(context.Background(), orderItemID)

	assert.NoError(t, err)
	assert.Equal(t, updatedOrderItem, result)
	mockRepo.AssertExpectations(t)
}

func TestOrderItemUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.OrderItemRepository)
	timeout := time.Duration(10)
	usecase := test.NewOrderItemUsecase(mockRepo, timeout)

	newOrderItem := &domain.OrderItem{
		OrderItemGuid:         uuid.New(),
		OrderID:               bson.NewObjectID(),
		ProductID:             bson.NewObjectID(),
		Quantity:              2,
		UnitPriceInclTax:      50.00,
		UnitPriceExclTax:      45.00,
		PriceInclTax:          100.00,
		PriceExclTax:          90.00,
		DiscountAmountInclTax: 10.00,
		DiscountAmountExclTax: 9.00,
		OriginalProductCost:   40.00,
		AttributeDescription:  "Color: Red, Size: M",
		AttributesXml:         "<Attributes><Color>Red</Color><Size>M</Size></Attributes>",
		DownloadCount:         0,
		IsDownloadActivated:   false,
		LicenseDownloadID:     nil,
		ItemWeight:            new(float64),
		RentalStartDateUtc:    nil,
		RentalEndDateUtc:      nil,
	}
	*newOrderItem.ItemWeight = 1.5

	mockRepo.On("Create", mock.Anything, newOrderItem).Return(nil)

	err := usecase.Create(context.Background(), newOrderItem)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestOrderItemUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.OrderItemRepository)
	timeout := time.Duration(10)
	usecase := test.NewOrderItemUsecase(mockRepo, timeout)

	updatedOrderItem := &domain.OrderItem{
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
	*updatedOrderItem.LicenseDownloadID = bson.NewObjectID()
	*updatedOrderItem.ItemWeight = 2.0
	*updatedOrderItem.RentalStartDateUtc = time.Now().AddDate(0, 0, -1) // Started 1 day ago
	*updatedOrderItem.RentalEndDateUtc = time.Now().AddDate(0, 0, 5)    // Ends in 5 days

	mockRepo.On("Update", mock.Anything, updatedOrderItem).Return(nil)

	err := usecase.Update(context.Background(), updatedOrderItem)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestOrderItemUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.OrderItemRepository)
	timeout := time.Duration(10)
	usecase := test.NewOrderItemUsecase(mockRepo, timeout)

	orderItemID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, orderItemID).Return(nil)

	err := usecase.Delete(context.Background(), orderItemID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestOrderItemUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.OrderItemRepository)
	timeout := time.Duration(10)
	usecase := test.NewOrderItemUsecase(mockRepo, timeout)

	fetchedOrderItems := []domain.OrderItem{
		{
			ID:                    bson.NewObjectID(),
			OrderItemGuid:         uuid.New(),
			OrderID:               bson.NewObjectID(),
			ProductID:             bson.NewObjectID(),
			Quantity:              2,
			UnitPriceInclTax:      50.00,
			UnitPriceExclTax:      45.00,
			PriceInclTax:          100.00,
			PriceExclTax:          90.00,
			DiscountAmountInclTax: 10.00,
			DiscountAmountExclTax: 9.00,
			OriginalProductCost:   40.00,
			AttributeDescription:  "Color: Red, Size: M",
			AttributesXml:         "<Attributes><Color>Red</Color><Size>M</Size></Attributes>",
			DownloadCount:         0,
			IsDownloadActivated:   false,
			LicenseDownloadID:     nil,
			ItemWeight:            new(float64),
			RentalStartDateUtc:    nil,
			RentalEndDateUtc:      nil,
		},
		{
			ID:                    bson.NewObjectID(),
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
		},
	}
	*fetchedOrderItems[0].ItemWeight = 1.5
	*fetchedOrderItems[1].LicenseDownloadID = bson.NewObjectID()
	*fetchedOrderItems[1].ItemWeight = 2.0
	*fetchedOrderItems[1].RentalStartDateUtc = time.Now().AddDate(0, 0, -1) // Started 1 day ago
	*fetchedOrderItems[1].RentalEndDateUtc = time.Now().AddDate(0, 0, 5)    // Ends in 5 days

	mockRepo.On("Fetch", mock.Anything).Return(fetchedOrderItems, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedOrderItems, result)
	mockRepo.AssertExpectations(t)
}
