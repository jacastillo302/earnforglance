package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/orders"
	repository "earnforglance/server/repository/orders"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultReturnRequestAction struct {
	mock.Mock
}

func (m *MockSingleResultReturnRequestAction) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.ReturnRequestAction); ok {
		*v.(*domain.ReturnRequestAction) = *result
	}
	return args.Error(1)
}

var mockItemReturnRequestAction = &domain.ReturnRequestAction{
	ID:           primitive.NewObjectID(), // Existing ID of the record to update
	Name:         "Refund Item",
	DisplayOrder: 2,
}

func TestReturnRequestActionRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionReturnRequestAction

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultReturnRequestAction{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemReturnRequestAction, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewReturnRequestActionRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemReturnRequestAction.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultReturnRequestAction{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewReturnRequestActionRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemReturnRequestAction.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestReturnRequestActionRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionReturnRequestAction

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemReturnRequestAction).Return(nil, nil).Once()

	repo := repository.NewReturnRequestActionRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemReturnRequestAction)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestReturnRequestActionRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionReturnRequestAction

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemReturnRequestAction.ID}
	update := bson.M{"$set": mockItemReturnRequestAction}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewReturnRequestActionRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemReturnRequestAction)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
