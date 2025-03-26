package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/stores"
	repository "earnforglance/server/repository/stores"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultStore struct {
	mock.Mock
}

func (m *MockSingleResultStore) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.Store); ok {
		*v.(*domain.Store) = *result
	}
	return args.Error(1)
}

var mockItemStore = &domain.Store{
	ID:                     primitive.NewObjectID(), // Existing ID of the record to update
	Name:                   "Updated Store",
	DefaultMetaKeywords:    "updated, store, gadgets",
	DefaultMetaDescription: "An updated description for the store.",
	DefaultTitle:           "Updated Store - Gadgets & More",
	HomepageTitle:          "Welcome to Updated Store",
	HomepageDescription:    "Discover updated gadgets and more.",
	Url:                    "https://www.updatedstore.com",
	SslEnabled:             false,
	Hosts:                  "www.updatedstore.com",
	DefaultLanguageID:      primitive.NewObjectID(),
	DisplayOrder:           2,
	CompanyName:            "Updated Store LLC",
	CompanyAddress:         "456 Updated Avenue, Townsville",
	CompanyPhoneNumber:     "+1-800-987-6543",
	CompanyVat:             "VAT654321",
	Deleted:                true,
}

func TestStoreRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionStore

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultStore{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemStore, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewStoreRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemStore.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultStore{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewStoreRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemStore.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestStoreRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionStore

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemStore).Return(nil, nil).Once()

	repo := repository.NewStoreRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemStore)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestStoreRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionStore

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemStore.ID}
	update := bson.M{"$set": mockItemStore}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewStoreRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemStore)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
