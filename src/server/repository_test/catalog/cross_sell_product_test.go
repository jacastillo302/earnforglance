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

	"go.mongodb.org/mongo-driver/v2/bson"
)

type MockSingleResultCrossSellProduct struct {
	mock.Mock
}

func (m *MockSingleResultCrossSellProduct) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.CrossSellProduct); ok {
		*v.(*domain.CrossSellProduct) = *result
	}
	return args.Error(1)
}

func TestCrossSellProductRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionCrossSellProduct

	mockItem := domain.CrossSellProduct{ID: bson.ObjectID{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, ProductID1: bson.ObjectID{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, ProductID2: bson.ObjectID{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}}

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultCrossSellProduct{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItem, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewCrossSellProductRepository(databaseHelper, collectionName)

		result, err := ur.FetchByID(context.Background(), mockItem.ID.Hex())

		assert.NoError(t, err)
		assert.Equal(t, mockItem, result)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultCrossSellProduct{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewCrossSellProductRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItem.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestCrossSellProductRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionCrossSellProduct

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)

	mockCrossSellProduct := &domain.CrossSellProduct{
		ID:         bson.NewObjectID(), // Existing ID of the record to update
		ProductID1: bson.NewObjectID(),
		ProductID2: bson.NewObjectID(),
	}

	collectionHelper.On("InsertOne", mock.Anything, mockCrossSellProduct).Return(nil, nil).Once()

	repo := repository.NewCrossSellProductRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockCrossSellProduct)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestCrossSellProductRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionCrossSellProduct

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)

	mockCrossSellProduct := &domain.CrossSellProduct{
		ID:         bson.NewObjectID(), // Existing ID of the record to update
		ProductID1: bson.NewObjectID(),
		ProductID2: bson.NewObjectID(),
	}

	filter := bson.M{"_id": mockCrossSellProduct.ID}
	update := bson.M{"$set": mockCrossSellProduct}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewCrossSellProductRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockCrossSellProduct)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
