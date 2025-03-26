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

type MockSingleResultLog struct {
	mock.Mock
}

func (m *MockSingleResultLog) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.Log); ok {
		*v.(*domain.Log) = *result
	}
	return args.Error(1)
}

var mockItemLog = &domain.Log{
	ID:           primitive.NewObjectID(), // Existing ID of the record to update
	LogLevelID:   2,
	ShortMessage: "Application error",
	FullMessage:  "An error occurred while processing the request.",
	IpAddress:    "192.168.1.2",
	CustomerID:   new(primitive.ObjectID),
	PageUrl:      "/error",
	ReferrerUrl:  "/home",
	CreatedOnUtc: time.Now().AddDate(0, 0, -7),
}

func TestLogRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionLog

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultLog{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemLog, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewLogRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemLog.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultLog{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewLogRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemLog.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestLogRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionLog

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemLog).Return(nil, nil).Once()

	repo := repository.NewLogRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemLog)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestLogRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionLog

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemLog.ID}
	update := bson.M{"$set": mockItemLog}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewLogRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemLog)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
