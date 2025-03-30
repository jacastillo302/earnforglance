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

type MockSingleResultCustomerRole struct {
	mock.Mock
}

func (m *MockSingleResultCustomerRole) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.CustomerRole); ok {
		*v.(*domain.CustomerRole) = *result
	}
	return args.Error(1)
}

var mockItemCustomerRole = &domain.CustomerRole{
	ID:                      primitive.NewObjectID(),
	Name:                    "Registered",
	FreeShipping:            false,
	TaxExempt:               false,
	Active:                  true,
	IsSystemRole:            true,
	SystemName:              "Registered",
	EnablePasswordLifetime:  false,
	OverrideTaxDisplayType:  false,
	DefaultTaxDisplayTypeID: 2,
	PurchasedWithProductId:  0,
}

func TestCustomerRoleRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionCustomerRole

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultCustomerRole{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemCustomerRole, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewCustomerRoleRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemCustomerRole.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultCustomerRole{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewCustomerRoleRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemCustomerRole.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestCustomerRoleRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionCustomerRole

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemCustomerRole).Return(nil, nil).Once()

	repo := repository.NewCustomerRoleRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemCustomerRole)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestCustomerRoleRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionCustomerRole

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemCustomerRole.ID}
	update := bson.M{"$set": mockItemCustomerRole}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewCustomerRoleRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemCustomerRole)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
