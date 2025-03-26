package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/catalog"
	repository "earnforglance/server/repository/catalog"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultProductAttributeCombination struct {
	mock.Mock
}

func (m *MockSingleResultProductAttributeCombination) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.ProductAttributeCombination); ok {
		*v.(*domain.ProductAttributeCombination) = *result
	}
	return args.Error(1)
}

func TestProductAttributeCombinationRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionProductAttributeCombination

	mockItem := domain.ProductAttributeCombination{
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
		PictureID:                   nil,
	}

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultProductAttributeCombination{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItem, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewProductAttributeCombinationRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItem.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultProductAttributeCombination{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewProductAttributeCombinationRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItem.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestProductAttributeCombinationRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionProductAttributeCombination

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)

	mockProductAttributeCombination := &domain.ProductAttributeCombination{
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
		PictureID:                   nil,
	}

	collectionHelper.On("InsertOne", mock.Anything, mockProductAttributeCombination).Return(nil, nil).Once()

	repo := repository.NewProductAttributeCombinationRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockProductAttributeCombination)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestProductAttributeCombinationRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionProductAttributeCombination

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)

	mockProductAttributeCombination := &domain.ProductAttributeCombination{
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
		PictureID:                   nil,
	}

	filter := bson.M{"_id": mockProductAttributeCombination.ID}
	update := bson.M{"$set": mockProductAttributeCombination}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewProductAttributeCombinationRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockProductAttributeCombination)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
