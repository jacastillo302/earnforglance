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

type MockSingleResultDiscountProductMapping struct {
	mock.Mock
}

func (m *MockSingleResultDiscountProductMapping) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.DiscountProductMapping); ok {
		*v.(*domain.DiscountProductMapping) = *result
	}
	return args.Error(1)
}

var mockItemDiscountProductMapping = &domain.DiscountProductMapping{
	DiscountMapping: domain.DiscountMapping{
		DiscountID: bson.NewObjectID(),
	},
	EntityID: bson.NewObjectID(),
}

func TestDiscountProductMappingRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionDiscountProductMapping

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultDiscountProductMapping{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemDiscountProductMapping, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewDiscountProductMappingRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemDiscountProductMapping.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultDiscountProductMapping{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewDiscountProductMappingRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemDiscountProductMapping.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestDiscountProductMappingRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionDiscountProductMapping

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemDiscountProductMapping).Return(nil, nil).Once()

	repo := repository.NewDiscountProductMappingRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemDiscountProductMapping)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestDiscountProductMappingRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionDiscountProductMapping

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemDiscountProductMapping.ID}
	update := bson.M{"$set": mockItemDiscountProductMapping}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewDiscountProductMappingRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemDiscountProductMapping)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
