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

type MockSingleResultProductManufacturer struct {
	mock.Mock
}

func (m *MockSingleResultProductManufacturer) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.ProductManufacturer); ok {
		*v.(*domain.ProductManufacturer) = *result
	}
	return args.Error(1)
}

var mockItemProductManufacturer = &domain.ProductManufacturer{
	ID:                bson.NewObjectID(), // Existing ID of the record to update
	ProductID:         bson.NewObjectID(),
	ManufacturerID:    bson.NewObjectID(),
	IsFeaturedProduct: false,
	DisplayOrder:      2,
}

func TestProductManufacturerRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionProductManufacturer

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultProductManufacturer{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemProductManufacturer, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewProductManufacturerRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemProductManufacturer.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultProductManufacturer{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewProductManufacturerRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemProductManufacturer.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestProductManufacturerRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionProductManufacturer

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemProductManufacturer).Return(nil, nil).Once()

	repo := repository.NewProductManufacturerRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemProductManufacturer)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestProductManufacturerRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionProductManufacturer

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemProductManufacturer.ID}
	update := bson.M{"$set": mockItemProductManufacturer}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewProductManufacturerRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemProductManufacturer)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
