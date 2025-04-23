package repository_test

import (
	"context"
	"errors"
	"testing"
	"time"

	domain "earnforglance/server/domain/scheduleTasks"
	repository "earnforglance/server/repository/scheduleTasks"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type MockSingleResultScheduleTask struct {
	mock.Mock
}

func (m *MockSingleResultScheduleTask) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.ScheduleTask); ok {
		*v.(*domain.ScheduleTask) = *result
	}
	return args.Error(1)
}

var mockItemScheduleTask = &domain.ScheduleTask{
	ID:             bson.NewObjectID(), // Existing ID of the record to update
	Name:           "Weekly Data Cleanup",
	Seconds:        604800,
	Type:           "CleanupTask",
	LastEnabledUtc: new(time.Time),
	Enabled:        false,
	StopOnError:    true,
	LastStartUtc:   new(time.Time),
	LastEndUtc:     new(time.Time),
	LastSuccessUtc: new(time.Time),
}

func TestScheduleTaskRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionScheduleTask

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultScheduleTask{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemScheduleTask, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewScheduleTaskRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemScheduleTask.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultScheduleTask{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewScheduleTaskRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemScheduleTask.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestScheduleTaskRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionScheduleTask

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemScheduleTask).Return(nil, nil).Once()

	repo := repository.NewScheduleTaskRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemScheduleTask)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestScheduleTaskRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionScheduleTask

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemScheduleTask.ID}
	update := bson.M{"$set": mockItemScheduleTask}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewScheduleTaskRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemScheduleTask)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
