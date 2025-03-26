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

type MockSingleResultVendorAttribute struct {
	mock.Mock
}

func (m *MockSingleResultVendorAttribute) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.VendorAttribute); ok {
		*v.(*domain.VendorAttribute) = *result
	}
	return args.Error(1)
}

var mockItemVendorAttribute = &domain.VendorAttribute{
	ID: primitive.NewObjectID(), // Existing ID of the record to update
}

func TestVendorAttributeRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionVendorAttribute

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultVendorAttribute{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemVendorAttribute, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewVendorAttributeRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemVendorAttribute.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultVendorAttribute{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewVendorAttributeRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemVendorAttribute.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestVendorAttributeRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionVendorAttribute

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemVendorAttribute).Return(nil, nil).Once()

	repo := repository.NewVendorAttributeRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemVendorAttribute)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestVendorAttributeRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionVendorAttribute

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemVendorAttribute.ID}
	update := bson.M{"$set": mockItemVendorAttribute}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewVendorAttributeRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemVendorAttribute)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
