package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/vendors"
	repository "earnforglance/server/repository/vendors"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultVendorAttributeValue struct {
	mock.Mock
}

func (m *MockSingleResultVendorAttributeValue) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.VendorAttributeValue); ok {
		*v.(*domain.VendorAttributeValue) = *result
	}
	return args.Error(1)
}

var mockItemVendorAttributeValue = &domain.VendorAttributeValue{
	ID: primitive.NewObjectID(), // Existing ID of the record to update
}

func TestVendorAttributeValueRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionVendorAttributeValue

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultVendorAttributeValue{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemVendorAttributeValue, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewVendorAttributeValueRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemVendorAttributeValue.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultVendorAttributeValue{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewVendorAttributeValueRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemVendorAttributeValue.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestVendorAttributeValueRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionVendorAttributeValue

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemVendorAttributeValue).Return(nil, nil).Once()

	repo := repository.NewVendorAttributeValueRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemVendorAttributeValue)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestVendorAttributeValueRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionVendorAttributeValue

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemVendorAttributeValue.ID}
	update := bson.M{"$set": mockItemVendorAttributeValue}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewVendorAttributeValueRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemVendorAttributeValue)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
