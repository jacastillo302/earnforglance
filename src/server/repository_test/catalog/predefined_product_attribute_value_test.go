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

type MockSingleResultPredefinedProductAttributeValue struct {
	mock.Mock
}

func (m *MockSingleResultPredefinedProductAttributeValue) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.PredefinedProductAttributeValue); ok {
		*v.(*domain.PredefinedProductAttributeValue) = *result
	}
	return args.Error(1)
}

func TestPredefinedProductAttributeValueRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionPredefinedProductAttributeValue

	mockItem := domain.PredefinedProductAttributeValue{
		ID:                           primitive.NewObjectID(),
		ProductAttributeID:           primitive.NewObjectID(),
		Name:                         "Color - Blue",
		PriceAdjustment:              15.0,
		PriceAdjustmentUsePercentage: true,
		WeightAdjustment:             0.8,
		Cost:                         7.0,
		IsPreSelected:                false,
		DisplayOrder:                 2,
	}

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultPredefinedProductAttributeValue{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItem, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewPredefinedProductAttributeValueRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItem.ID.Hex())

		assert.NoError(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultPredefinedProductAttributeValue{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewPredefinedProductAttributeValueRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItem.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestPredefinedProductAttributeValueRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionPredefinedProductAttributeValue

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)

	mockPredefinedProductAttributeValue := &domain.PredefinedProductAttributeValue{
		ID:                           primitive.NewObjectID(),
		ProductAttributeID:           primitive.NewObjectID(),
		Name:                         "Color - Blue",
		PriceAdjustment:              15.0,
		PriceAdjustmentUsePercentage: true,
		WeightAdjustment:             0.8,
		Cost:                         7.0,
		IsPreSelected:                false,
		DisplayOrder:                 2,
	}
	collectionHelper.On("InsertOne", mock.Anything, mockPredefinedProductAttributeValue).Return(nil, nil).Once()

	repo := repository.NewPredefinedProductAttributeValueRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockPredefinedProductAttributeValue)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestPredefinedProductAttributeValueRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionPredefinedProductAttributeValue

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)

	mockPredefinedProductAttributeValue := &domain.PredefinedProductAttributeValue{
		ID:                           primitive.NewObjectID(),
		ProductAttributeID:           primitive.NewObjectID(),
		Name:                         "Color - Blue",
		PriceAdjustment:              15.0,
		PriceAdjustmentUsePercentage: true,
		WeightAdjustment:             0.8,
		Cost:                         7.0,
		IsPreSelected:                false,
		DisplayOrder:                 2,
	}

	filter := bson.M{"_id": mockPredefinedProductAttributeValue.ID}
	update := bson.M{"$set": mockPredefinedProductAttributeValue}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewPredefinedProductAttributeValueRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockPredefinedProductAttributeValue)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
