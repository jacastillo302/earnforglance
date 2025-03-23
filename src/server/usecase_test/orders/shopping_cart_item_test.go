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

func TestShoppingCartItemUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.ShoppingCartItemRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewShoppingCartItemUsecase(mockRepo, timeout)

	shoppingCartItemID := primitive.NewObjectID().Hex()

	updatedShoppingCartItem := domain.ShoppingCartItem{
		ID:                   primitive.NewObjectID(), // Existing ID of the record to update
		StoreID:              primitive.NewObjectID(),
		ShoppingCartTypeID:   primitive.NewObjectID(),
		CustomerID:           primitive.NewObjectID(),
		ProductID:            primitive.NewObjectID(),
		AttributesXml:        "<Attributes><Color>Blue</Color><Size>L</Size></Attributes>",
		CustomerEnteredPrice: 59.99,
		Quantity:             3,
		RentalStartDateUtc:   new(time.Time),
		RentalEndDateUtc:     new(time.Time),
		CreatedOnUtc:         time.Now().AddDate(0, 0, -7), // Created 7 days ago
		UpdatedOnUtc:         time.Now(),
		ShoppingCartType:     2,
	}

	mockRepo.On("FetchByID", mock.Anything, shoppingCartItemID).Return(updatedShoppingCartItem, nil)

	result, err := usecase.FetchByID(context.Background(), shoppingCartItemID)

	assert.NoError(t, err)
	assert.Equal(t, updatedShoppingCartItem, result)
	mockRepo.AssertExpectations(t)
}

func TestShoppingCartItemUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.ShoppingCartItemRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewShoppingCartItemUsecase(mockRepo, timeout)

	newShoppingCartItem := &domain.ShoppingCartItem{
		StoreID:              primitive.NewObjectID(),
		ShoppingCartTypeID:   primitive.NewObjectID(),
		CustomerID:           primitive.NewObjectID(),
		ProductID:            primitive.NewObjectID(),
		AttributesXml:        "<Attributes><Color>Red</Color><Size>M</Size></Attributes>",
		CustomerEnteredPrice: 49.99,
		Quantity:             2,
		RentalStartDateUtc:   nil,
		RentalEndDateUtc:     nil,
		CreatedOnUtc:         time.Now(),
		UpdatedOnUtc:         time.Now(),
		ShoppingCartType:     1,
	}

	mockRepo.On("Create", mock.Anything, newShoppingCartItem).Return(nil)

	err := usecase.Create(context.Background(), newShoppingCartItem)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestShoppingCartItemUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.ShoppingCartItemRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewShoppingCartItemUsecase(mockRepo, timeout)

	updatedShoppingCartItem := &domain.ShoppingCartItem{
		ID:                   primitive.NewObjectID(), // Existing ID of the record to update
		StoreID:              primitive.NewObjectID(),
		ShoppingCartTypeID:   primitive.NewObjectID(),
		CustomerID:           primitive.NewObjectID(),
		ProductID:            primitive.NewObjectID(),
		AttributesXml:        "<Attributes><Color>Blue</Color><Size>L</Size></Attributes>",
		CustomerEnteredPrice: 59.99,
		Quantity:             3,
		RentalStartDateUtc:   new(time.Time),
		RentalEndDateUtc:     new(time.Time),
		CreatedOnUtc:         time.Now().AddDate(0, 0, -7), // Created 7 days ago
		UpdatedOnUtc:         time.Now(),
		ShoppingCartType:     2,
	}
	*updatedShoppingCartItem.RentalStartDateUtc = time.Now().AddDate(0, 0, -1) // Rental started 1 day ago
	*updatedShoppingCartItem.RentalEndDateUtc = time.Now().AddDate(0, 0, 5)    // Rental ends in 5 days

	mockRepo.On("Update", mock.Anything, updatedShoppingCartItem).Return(nil)

	err := usecase.Update(context.Background(), updatedShoppingCartItem)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestShoppingCartItemUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.ShoppingCartItemRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewShoppingCartItemUsecase(mockRepo, timeout)

	shoppingCartItemID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, shoppingCartItemID).Return(nil)

	err := usecase.Delete(context.Background(), shoppingCartItemID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestShoppingCartItemUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.ShoppingCartItemRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewShoppingCartItemUsecase(mockRepo, timeout)

	fetchedShoppingCartItems := []domain.ShoppingCartItem{
		{
			ID:                   primitive.NewObjectID(),
			StoreID:              primitive.NewObjectID(),
			ShoppingCartTypeID:   primitive.NewObjectID(),
			CustomerID:           primitive.NewObjectID(),
			ProductID:            primitive.NewObjectID(),
			AttributesXml:        "<Attributes><Color>Red</Color><Size>M</Size></Attributes>",
			CustomerEnteredPrice: 49.99,
			Quantity:             2,
			RentalStartDateUtc:   nil,
			RentalEndDateUtc:     nil,
			CreatedOnUtc:         time.Now().AddDate(0, 0, -10), // Created 10 days ago
			UpdatedOnUtc:         time.Now().AddDate(0, 0, -5),  // Updated 5 days ago
			ShoppingCartType:     1,
		},
		{
			ID:                   primitive.NewObjectID(),
			StoreID:              primitive.NewObjectID(),
			ShoppingCartTypeID:   primitive.NewObjectID(),
			CustomerID:           primitive.NewObjectID(),
			ProductID:            primitive.NewObjectID(),
			AttributesXml:        "<Attributes><Color>Blue</Color><Size>L</Size></Attributes>",
			CustomerEnteredPrice: 59.99,
			Quantity:             3,
			RentalStartDateUtc:   new(time.Time),
			RentalEndDateUtc:     new(time.Time),
			CreatedOnUtc:         time.Now().AddDate(0, 0, -7), // Created 7 days ago
			UpdatedOnUtc:         time.Now(),
			ShoppingCartType:     2,
		},
	}
	*fetchedShoppingCartItems[1].RentalStartDateUtc = time.Now().AddDate(0, 0, -1) // Rental started 1 day ago
	*fetchedShoppingCartItems[1].RentalEndDateUtc = time.Now().AddDate(0, 0, 5)    // Rental ends in 5 days

	mockRepo.On("Fetch", mock.Anything).Return(fetchedShoppingCartItems, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedShoppingCartItems, result)
	mockRepo.AssertExpectations(t)
}
