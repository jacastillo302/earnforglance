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

func TestProductAttributeCombinationUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.ProductAttributeCombinationRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewProductAttributeCombinationUsecase(mockRepo, timeout)

	productAttributeCombinationID := primitive.NewObjectID().Hex()

	expectedProductAttributeCombination := domain.ProductAttributeCombination{
		ID:                          primitive.NewObjectID(), // Existing ID of the record to update
		ProductID:                   primitive.NewObjectID(),
		AttributesXml:               "<attributes><color>blue</color><size>large</size></attributes>",
		StockQuantity:               30,
		AllowOutOfStockOrders:       true,
		Sku:                         "SKU54321",
		ManufacturerPartNumber:      "MPN54321",
		Gtin:                        "0987654321098",
		OverriddenPrice:             new(float64),
		NotifyAdminForQuantityBelow: 5,
		MinStockQuantity:            2,
	}

	mockRepo.On("FetchByID", mock.Anything, productAttributeCombinationID).Return(expectedProductAttributeCombination, nil)

	result, err := usecase.FetchByID(context.Background(), productAttributeCombinationID)

	assert.NoError(t, err)
	assert.Equal(t, expectedProductAttributeCombination, result)
	mockRepo.AssertExpectations(t)
}

func TestProductAttributeCombinationUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.ProductAttributeCombinationRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewProductAttributeCombinationUsecase(mockRepo, timeout)

	newProductAttributeCombination := &domain.ProductAttributeCombination{
		ProductID:                   primitive.NewObjectID(),
		AttributesXml:               "<attributes><color>red</color><size>medium</size></attributes>",
		StockQuantity:               50,
		AllowOutOfStockOrders:       false,
		Sku:                         "SKU12345",
		ManufacturerPartNumber:      "MPN67890",
		Gtin:                        "0123456789012",
		OverriddenPrice:             nil,
		NotifyAdminForQuantityBelow: 10,
		MinStockQuantity:            5,
	}

	mockRepo.On("Create", mock.Anything, newProductAttributeCombination).Return(nil)

	err := usecase.Create(context.Background(), newProductAttributeCombination)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductAttributeCombinationUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.ProductAttributeCombinationRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewProductAttributeCombinationUsecase(mockRepo, timeout)

	updatedProductAttributeCombination := &domain.ProductAttributeCombination{
		ID:                          primitive.NewObjectID(), // Existing ID of the record to update
		ProductID:                   primitive.NewObjectID(),
		AttributesXml:               "<attributes><color>blue</color><size>large</size></attributes>",
		StockQuantity:               30,
		AllowOutOfStockOrders:       true,
		Sku:                         "SKU54321",
		ManufacturerPartNumber:      "MPN54321",
		Gtin:                        "0987654321098",
		OverriddenPrice:             new(float64),
		NotifyAdminForQuantityBelow: 5,
		MinStockQuantity:            2,
	}

	mockRepo.On("Update", mock.Anything, updatedProductAttributeCombination).Return(nil)

	err := usecase.Update(context.Background(), updatedProductAttributeCombination)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductAttributeCombinationUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.ProductAttributeCombinationRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewProductAttributeCombinationUsecase(mockRepo, timeout)

	productAttributeCombinationID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, productAttributeCombinationID).Return(nil)

	err := usecase.Delete(context.Background(), productAttributeCombinationID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductAttributeCombinationUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.ProductAttributeCombinationRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewProductAttributeCombinationUsecase(mockRepo, timeout)

	fetchedProductAttributeCombinations := []domain.ProductAttributeCombination{
		{
			ID:                          primitive.NewObjectID(),
			ProductID:                   primitive.NewObjectID(),
			AttributesXml:               "<attributes><color>blue</color><size>large</size></attributes>",
			StockQuantity:               20,
			AllowOutOfStockOrders:       false,
			Sku:                         "SKU98765",
			ManufacturerPartNumber:      "MPN54321",
			Gtin:                        "0987654321098",
			OverriddenPrice:             nil,
			NotifyAdminForQuantityBelow: 5,
			MinStockQuantity:            2,
		},
		{
			ID:                          primitive.NewObjectID(),
			ProductID:                   primitive.NewObjectID(),
			AttributesXml:               "<attributes><color>green</color><size>small</size></attributes>",
			StockQuantity:               15,
			AllowOutOfStockOrders:       true,
			Sku:                         "SKU11223",
			ManufacturerPartNumber:      "MPN33445",
			Gtin:                        "1234567890123",
			OverriddenPrice:             new(float64),
			NotifyAdminForQuantityBelow: 3,
			MinStockQuantity:            1,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedProductAttributeCombinations, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedProductAttributeCombinations, result)
	mockRepo.AssertExpectations(t)
}
