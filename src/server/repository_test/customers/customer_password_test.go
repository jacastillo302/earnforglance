package repository_test

import (
	"context"
	"errors"
	"testing"
	"time"

	domain "earnforglance/server/domain/customers"
	repository "earnforglance/server/repository/customers"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultCustomerPassword struct {
	mock.Mock
}

func (m *MockSingleResultCustomerPassword) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.CustomerPassword); ok {
		*v.(*domain.CustomerPassword) = *result
	}
	return args.Error(1)
}

var mockItemCustomerPassword = &domain.CustomerPassword{
	ID:               primitive.NewObjectID(),
	CustomerID:       primitive.NewObjectID(),
	Password:         "hashedpassword456",
	PasswordFormatID: 2, // Example password format ID (e.g., 2 for encrypted)
	PasswordSalt:     "anotherrandomsalt",
	CreatedOnUTC:     time.Now().UTC().Add(-24 * time.Hour), // Created 1 day ago
}

func TestCustomerPasswordRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionCustomerPassword

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultCustomerPassword{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemCustomerPassword, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewCustomerPasswordRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemCustomerPassword.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultCustomerPassword{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewCustomerPasswordRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemCustomerPassword.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestCustomerPasswordRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionCustomerPassword

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemCustomerPassword).Return(nil, nil).Once()

	repo := repository.NewCustomerPasswordRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemCustomerPassword)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestCustomerPasswordRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionCustomerPassword

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemCustomerPassword.ID}
	update := bson.M{"$set": mockItemCustomerPassword}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewCustomerPasswordRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemCustomerPassword)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
