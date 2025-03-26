package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/security"
	repository "earnforglance/server/repository/security"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultPermissionRecord struct {
	mock.Mock
}

func (m *MockSingleResultPermissionRecord) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.PermissionRecord); ok {
		*v.(*domain.PermissionRecord) = *result
	}
	return args.Error(1)
}

var mockItemPermissionRecord = &domain.PermissionRecord{
	ID:         primitive.NewObjectID(), // Existing ID of the record to update
	Name:       "Manage Categories",
	SystemName: "manage_categories",
	Category:   "Catalog",
}

func TestPermissionRecordRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionPermissionRecord

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultPermissionRecord{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemPermissionRecord, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewPermissionRecordRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemPermissionRecord.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultPermissionRecord{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewPermissionRecordRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemPermissionRecord.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestPermissionRecordRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionPermissionRecord

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemPermissionRecord).Return(nil, nil).Once()

	repo := repository.NewPermissionRecordRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemPermissionRecord)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestPermissionRecordRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionPermissionRecord

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemPermissionRecord.ID}
	update := bson.M{"$set": mockItemPermissionRecord}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewPermissionRecordRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemPermissionRecord)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
