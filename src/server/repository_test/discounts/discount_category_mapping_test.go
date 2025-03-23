package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/discounts"
	"earnforglance/server/mongo/mocks"
	repository "earnforglance/server/repository/discounts"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultDiscountCategoryMapping struct {
	mock.Mock
}

func (m *MockSingleResultDiscountCategoryMapping) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.DiscountCategoryMapping); ok {
		*v.(*domain.DiscountCategoryMapping) = *result
	}
	return args.Error(1)
}

var mockItemDiscountCategoryMapping = &domain.DiscountCategoryMapping{
	DiscountMapping: domain.DiscountMapping{
		DiscountID: primitive.NewObjectID(),
	},
	EntityID: primitive.NewObjectID(),
}

func TestDiscountCategoryMappingRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionDiscountCategoryMapping

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultDiscountCategoryMapping{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemDiscountCategoryMapping, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewDiscountCategoryMappingRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemDiscountCategoryMapping.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultDiscountCategoryMapping{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewDiscountCategoryMappingRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemDiscountCategoryMapping.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestDiscountCategoryMappingRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionDiscountCategoryMapping

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemDiscountCategoryMapping).Return(nil, nil).Once()

	repo := repository.NewDiscountCategoryMappingRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemDiscountCategoryMapping)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestDiscountCategoryMappingRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionDiscountCategoryMapping

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemDiscountCategoryMapping.ID}
	update := bson.M{"$set": mockItemDiscountCategoryMapping}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewDiscountCategoryMappingRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemDiscountCategoryMapping)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
