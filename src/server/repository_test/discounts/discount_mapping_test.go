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
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultDiscountMapping struct {
	mock.Mock
}

func (m *MockSingleResultDiscountMapping) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.DiscountMapping); ok {
		*v.(*domain.DiscountMapping) = *result
	}
	return args.Error(1)
}

var mockItemDiscountMapping = &domain.DiscountMapping{
	ID:         primitive.NewObjectID(), // Existing ID of the record to update
	DiscountID: primitive.NewObjectID(),
	EntityID:   primitive.NewObjectID(),
}

func TestDiscountMappingRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionDiscountMapping

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultDiscountMapping{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemDiscountMapping, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewDiscountMappingRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemDiscountMapping.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultDiscountMapping{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewDiscountMappingRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemDiscountMapping.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestDiscountMappingRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionDiscountMapping

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemDiscountMapping).Return(nil, nil).Once()

	repo := repository.NewDiscountMappingRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemDiscountMapping)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestDiscountMappingRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionDiscountMapping

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemDiscountMapping.ID}
	update := bson.M{"$set": mockItemDiscountMapping}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewDiscountMappingRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemDiscountMapping)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
