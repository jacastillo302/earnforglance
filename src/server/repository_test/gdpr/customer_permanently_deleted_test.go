package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/gdpr"
	repository "earnforglance/server/repository/gdpr"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultCustomerPermanentlyDeleted struct {
	mock.Mock
}

func (m *MockSingleResultCustomerPermanentlyDeleted) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.CustomerPermanentlyDeleted); ok {
		*v.(*domain.CustomerPermanentlyDeleted) = *result
	}
	return args.Error(1)
}

var mockItemCustomerPermanentlyDeleted = &domain.CustomerPermanentlyDeleted{
	CustomerID: primitive.NewObjectID(), // Existing CustomerID to update
	Email:      "updated_deleted_customer@example.com",
}

func TestCustomerPermanentlyDeletedRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionCustomerPermanentlyDeleted

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultCustomerPermanentlyDeleted{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemCustomerPermanentlyDeleted, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewCustomerPermanentlyDeletedRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemCustomerPermanentlyDeleted.CustomerID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultCustomerPermanentlyDeleted{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewCustomerPermanentlyDeletedRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemCustomerPermanentlyDeleted.CustomerID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestCustomerPermanentlyDeletedRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionCustomerPermanentlyDeleted

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemCustomerPermanentlyDeleted).Return(nil, nil).Once()

	repo := repository.NewCustomerPermanentlyDeletedRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemCustomerPermanentlyDeleted)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestCustomerPermanentlyDeletedRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionCustomerPermanentlyDeleted

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemCustomerPermanentlyDeleted.CustomerID}
	update := bson.M{"$set": mockItemCustomerPermanentlyDeleted}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewCustomerPermanentlyDeletedRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemCustomerPermanentlyDeleted)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
