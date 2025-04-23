package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/discounts"
	repository "earnforglance/server/repository/discounts"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type MockSingleResultDiscountManufacturerMapping struct {
	mock.Mock
}

func (m *MockSingleResultDiscountManufacturerMapping) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.DiscountManufacturerMapping); ok {
		*v.(*domain.DiscountManufacturerMapping) = *result
	}
	return args.Error(1)
}

var mockItemDiscountManufacturerMapping = &domain.DiscountManufacturerMapping{
	DiscountMapping: domain.DiscountMapping{
		DiscountID: bson.NewObjectID(),
	},
	EntityID: bson.NewObjectID(),
}

func TestDiscountManufacturerMappingRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionDiscountManufacturerMapping

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultDiscountManufacturerMapping{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemDiscountManufacturerMapping, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewDiscountManufacturerMappingRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemDiscountManufacturerMapping.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultDiscountManufacturerMapping{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewDiscountManufacturerMappingRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemDiscountManufacturerMapping.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestDiscountManufacturerMappingRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionDiscountManufacturerMapping

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemDiscountManufacturerMapping).Return(nil, nil).Once()

	repo := repository.NewDiscountManufacturerMappingRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemDiscountManufacturerMapping)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestDiscountManufacturerMappingRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionDiscountManufacturerMapping

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemDiscountManufacturerMapping.ID}
	update := bson.M{"$set": mockItemDiscountManufacturerMapping}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewDiscountManufacturerMappingRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemDiscountManufacturerMapping)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
