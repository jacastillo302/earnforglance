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

type MockSingleResultCustomerCustomerRoleMapping struct {
	mock.Mock
}

func (m *MockSingleResultCustomerCustomerRoleMapping) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.CustomerCustomerRoleMapping); ok {
		*v.(*domain.CustomerCustomerRoleMapping) = *result
	}
	return args.Error(1)
}

var mockItemCustomerCustomerRoleMapping = &domain.CustomerCustomerRoleMapping{
	ID:             bson.NewObjectID(), // Existing ID of the record to update
	CustomerID:     bson.NewObjectID(),
	CustomerRoleID: bson.NewObjectID(),
}

func TestCustomerCustomerRoleMappingRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionCustomerCustomerRoleMapping

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultCustomerCustomerRoleMapping{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemCustomerCustomerRoleMapping, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewCustomerCustomerRoleMappingRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemCustomerCustomerRoleMapping.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultCustomerCustomerRoleMapping{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewCustomerCustomerRoleMappingRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemCustomerCustomerRoleMapping.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestCustomerCustomerRoleMappingRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionCustomerCustomerRoleMapping

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemCustomerCustomerRoleMapping).Return(nil, nil).Once()

	repo := repository.NewCustomerCustomerRoleMappingRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemCustomerCustomerRoleMapping)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestCustomerCustomerRoleMappingRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionCustomerCustomerRoleMapping

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemCustomerCustomerRoleMapping.ID}
	update := bson.M{"$set": mockItemCustomerCustomerRoleMapping}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewCustomerCustomerRoleMappingRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemCustomerCustomerRoleMapping)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
