package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/orders"
	"earnforglance/server/mongo/mocks"
	repository "earnforglance/server/repository/orders"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultReturnRequestReason struct {
	mock.Mock
}

func (m *MockSingleResultReturnRequestReason) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.ReturnRequestReason); ok {
		*v.(*domain.ReturnRequestReason) = *result
	}
	return args.Error(1)
}

var mockItemReturnRequestReason = &domain.ReturnRequestReason{
	ID:           primitive.NewObjectID(), // Existing ID of the record to update
	Name:         "Wrong Item Delivered",
	DisplayOrder: 2,
}

func TestReturnRequestReasonRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionReturnRequestReason

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultReturnRequestReason{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemReturnRequestReason, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewReturnRequestReasonRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemReturnRequestReason.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultReturnRequestReason{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewReturnRequestReasonRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemReturnRequestReason.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestReturnRequestReasonRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionReturnRequestReason

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemReturnRequestReason).Return(nil, nil).Once()

	repo := repository.NewReturnRequestReasonRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemReturnRequestReason)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestReturnRequestReasonRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionReturnRequestReason

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemReturnRequestReason.ID}
	update := bson.M{"$set": mockItemReturnRequestReason}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewReturnRequestReasonRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemReturnRequestReason)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
