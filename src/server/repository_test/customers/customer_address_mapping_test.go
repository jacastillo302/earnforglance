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

	"go.mongodb.org/mongo-driver/v2/bson"
)

type MockSingleResultCustomerAddressMapping struct {
	mock.Mock
}

func (m *MockSingleResultCustomerAddressMapping) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.CustomerAddressMapping); ok {
		*v.(*domain.CustomerAddressMapping) = *result
	}
	return args.Error(1)
}

var mockItemCustomerAddressMapping = &domain.CustomerAddressMapping{
	ID:         bson.NewObjectID(), // Existing ID of the record to update
	CustomerID: bson.NewObjectID(),
	AddressID:  bson.NewObjectID(),
}

func TestCustomerAddressMappingRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionCustomerAddressMapping

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultCustomerAddressMapping{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemCustomerAddressMapping, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewCustomerAddressMappingRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemCustomerAddressMapping.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultCustomerAddressMapping{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewCustomerAddressMappingRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemCustomerAddressMapping.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestCustomerAddressMappingRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionCustomerAddressMapping

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemCustomerAddressMapping).Return(nil, nil).Once()

	repo := repository.NewCustomerAddressMappingRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemCustomerAddressMapping)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestCustomerAddressMappingRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionCustomerAddressMapping

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemCustomerAddressMapping.ID}
	update := bson.M{"$set": mockItemCustomerAddressMapping}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewCustomerAddressMappingRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemCustomerAddressMapping)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
