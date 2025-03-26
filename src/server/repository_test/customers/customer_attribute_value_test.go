package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/customers"
	repository "earnforglance/server/repository/customers"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultCustomerAttributeValue struct {
	mock.Mock
}

func (m *MockSingleResultCustomerAttributeValue) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.CustomerAttributeValue); ok {
		*v.(*domain.CustomerAttributeValue) = *result
	}
	return args.Error(1)
}

var mockItemCustomerAttributeValue = &domain.CustomerAttributeValue{
	ID:                  primitive.NewObjectID(), // Existing ID of the record to update
	CustomerAttributeID: primitive.NewObjectID(),
	Name:                "Preferred Currency",
	IsPreSelected:       false,
	DisplayOrder:        2,
}

func TestCustomerAttributeValueRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionCustomerAttributeValue

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultCustomerAttributeValue{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemCustomerAttributeValue, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewCustomerAttributeValueRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemCustomerAttributeValue.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultCustomerAttributeValue{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewCustomerAttributeValueRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemCustomerAttributeValue.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestCustomerAttributeValueRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionCustomerAttributeValue

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemCustomerAttributeValue).Return(nil, nil).Once()

	repo := repository.NewCustomerAttributeValueRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemCustomerAttributeValue)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestCustomerAttributeValueRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionCustomerAttributeValue

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemCustomerAttributeValue.ID}
	update := bson.M{"$set": mockItemCustomerAttributeValue}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewCustomerAttributeValueRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemCustomerAttributeValue)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
