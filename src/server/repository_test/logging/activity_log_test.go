package repository_test

import (
	"context"
	"errors"
	"testing"
	"time"

	domain "earnforglance/server/domain/logging"
	repository "earnforglance/server/repository/logging"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultActivityLog struct {
	mock.Mock
}

func (m *MockSingleResultActivityLog) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.ActivityLog); ok {
		*v.(*domain.ActivityLog) = *result
	}
	return args.Error(1)
}

var mockItemActivityLog = &domain.ActivityLog{
	ID:                primitive.NewObjectID(), // Existing ID of the record to update
	ActivityLogTypeID: 2,
	EntityID:          new(primitive.ObjectID),
	EntityName:        "Product",
	CustomerID:        primitive.NewObjectID(),
	Comment:           "Customer viewed a product.",
	CreatedOnUtc:      time.Now().AddDate(0, 0, -7), // Created 7 days ago
	IpAddress:         "192.168.1.2",
}

func TestActivityLogRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionActivityLog

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultActivityLog{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemActivityLog, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewActivityLogRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemActivityLog.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultActivityLog{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewActivityLogRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemActivityLog.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestActivityLogRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionActivityLog

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemActivityLog).Return(nil, nil).Once()

	repo := repository.NewActivityLogRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemActivityLog)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestActivityLogRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionActivityLog

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemActivityLog.ID}
	update := bson.M{"$set": mockItemActivityLog}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewActivityLogRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemActivityLog)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
