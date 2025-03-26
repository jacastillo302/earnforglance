package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/common"
	repository "earnforglance/server/repository/common"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultAddressAttributeValue struct {
	mock.Mock
}

func (m *MockSingleResultAddressAttributeValue) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.AddressAttributeValue); ok {
		*v.(*domain.AddressAttributeValue) = *result
	}
	return args.Error(1)
}

var mockItemAddressAttributeValue = &domain.AddressAttributeValue{
	ID:                 primitive.NewObjectID(), // Existing ID of the record to update
	AddressAttributeID: primitive.NewObjectID(),
	Name:               "State",
	IsPreSelected:      false,
	DisplayOrder:       2,
}

func TestAddressAttributeValueRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionAddressAttributeValue

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultAddressAttributeValue{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemAddressAttributeValue, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewAddressAttributeValueRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemAddressAttributeValue.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultAddressAttributeValue{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewAddressAttributeValueRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemAddressAttributeValue.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestAddressAttributeValueRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionAddressAttributeValue

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemAddressAttributeValue).Return(nil, nil).Once()

	repo := repository.NewAddressAttributeValueRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemAddressAttributeValue)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestAddressAttributeValueRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionAddressAttributeValue

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemAddressAttributeValue.ID}
	update := bson.M{"$set": mockItemAddressAttributeValue}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewAddressAttributeValueRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemAddressAttributeValue)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
