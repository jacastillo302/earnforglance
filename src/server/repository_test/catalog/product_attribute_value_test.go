package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/catalog"
	"earnforglance/server/mongo/mocks"
	repository "earnforglance/server/repository/catalog"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultProductAttributeValue struct {
	mock.Mock
}

func (m *MockSingleResultProductAttributeValue) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.ProductAttributeValue); ok {
		*v.(*domain.ProductAttributeValue) = *result
	}
	return args.Error(1)
}

var mockItemProductAttributeValue = &domain.ProductAttributeValue{
	ID:                           primitive.NewObjectID(), // Existing ID of the record to update
	ProductAttributeMappingID:    primitive.NewObjectID(),
	AttributeValueTypeID:         2,
	AssociatedProductID:          primitive.NewObjectID(),
	Name:                         "Size - Large",
	ColorSquaresRgb:              "",
	ImageSquaresPictureID:        primitive.NewObjectID(),
	PriceAdjustment:              15.0,
	PriceAdjustmentUsePercentage: true,
	WeightAdjustment:             1.0,
	Cost:                         7.0,
	CustomerEntersQty:            true,
	Quantity:                     50,
	IsPreSelected:                false,
	DisplayOrder:                 2,
	AttributeValueType:           5,
	PictureID:                    nil, // Deprecated field
}

func TestProductAttributeValueRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionProductAttributeValue

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultProductAttributeValue{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemProductAttributeValue, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewProductAttributeValueRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemProductAttributeValue.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultProductAttributeValue{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewProductAttributeValueRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemProductAttributeValue.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestProductAttributeValueRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionProductAttributeValue

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemProductAttributeValue).Return(nil, nil).Once()

	repo := repository.NewProductAttributeValueRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemProductAttributeValue)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestProductAttributeValueRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionProductAttributeValue

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemProductAttributeValue.ID}
	update := bson.M{"$set": mockItemProductAttributeValue}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewProductAttributeValueRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemProductAttributeValue)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
