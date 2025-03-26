package repository_test

import (
	"context"
	"errors"
	"testing"
	"time"

	domain "earnforglance/server/domain/gdpr"
	repository "earnforglance/server/repository/gdpr"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultGdprLog struct {
	mock.Mock
}

func (m *MockSingleResultGdprLog) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.GdprLog); ok {
		*v.(*domain.GdprLog) = *result
	}
	return args.Error(1)
}

var mockItemGdprLog = &domain.GdprLog{
	ID:             primitive.NewObjectID(), // Existing ID of the record to update
	CustomerID:     primitive.NewObjectID(),
	ConsentID:      primitive.NewObjectID(),
	CustomerInfo:   "Jane Doe, jane.doe@example.com",
	RequestTypeID:  2,
	RequestDetails: "Request to export personal data.",
	CreatedOnUtc:   time.Now().AddDate(0, 0, -7), // Created 7 days ago

}

func TestGdprLogRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionGdprLog

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultGdprLog{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemGdprLog, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewGdprLogRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemGdprLog.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultGdprLog{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewGdprLogRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemGdprLog.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestGdprLogRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionGdprLog

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemGdprLog).Return(nil, nil).Once()

	repo := repository.NewGdprLogRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemGdprLog)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestGdprLogRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionGdprLog

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemGdprLog.ID}
	update := bson.M{"$set": mockItemGdprLog}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewGdprLogRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemGdprLog)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
