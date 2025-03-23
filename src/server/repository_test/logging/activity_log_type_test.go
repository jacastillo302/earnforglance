package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/logging"
	"earnforglance/server/mongo/mocks"
	repository "earnforglance/server/repository/logging"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultActivityLogType struct {
	mock.Mock
}

func (m *MockSingleResultActivityLogType) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.ActivityLogType); ok {
		*v.(*domain.ActivityLogType) = *result
	}
	return args.Error(1)
}

var mockItemActivityLogType = &domain.ActivityLogType{
	ID:            primitive.NewObjectID(), // Existing ID of the record to update
	SystemKeyword: "customer_registration",
	Name:          "Customer Registration",
	Enabled:       false,
}

func TestActivityLogTypeRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionActivityLogType

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultActivityLogType{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemActivityLogType, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewActivityLogTypeRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemActivityLogType.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultActivityLogType{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewActivityLogTypeRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemActivityLogType.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestActivityLogTypeRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionActivityLogType

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemActivityLogType).Return(nil, nil).Once()

	repo := repository.NewActivityLogTypeRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemActivityLogType)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestActivityLogTypeRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionActivityLogType

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemActivityLogType.ID}
	update := bson.M{"$set": mockItemActivityLogType}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewActivityLogTypeRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemActivityLogType)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
