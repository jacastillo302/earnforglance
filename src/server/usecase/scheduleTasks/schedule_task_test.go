package usecase

import (
	"context"
	mocks "earnforglance/server/domain/mocks"
	domain "earnforglance/server/domain/scheduleTasks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestScheduleTaskUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.ScheduleTaskRepository)
	timeout := time.Duration(10)
	usecase := NewScheduleTaskUsecase(mockRepo, timeout)

	scheduleTaskID := primitive.NewObjectID().Hex()
	updatedScheduleTask := domain.ScheduleTask{
		ID:             primitive.NewObjectID(), // Existing ID of the record to update
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

	mockRepo.On("FetchByID", mock.Anything, scheduleTaskID).Return(updatedScheduleTask, nil)

	result, err := usecase.FetchByID(context.Background(), scheduleTaskID)

	assert.NoError(t, err)
	assert.Equal(t, updatedScheduleTask, result)
	mockRepo.AssertExpectations(t)
}

func TestScheduleTaskUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.ScheduleTaskRepository)
	timeout := time.Duration(10)
	usecase := NewScheduleTaskUsecase(mockRepo, timeout)

	newScheduleTask := &domain.ScheduleTask{
		Name:           "Daily Data Backup",
		Seconds:        86400,
		Type:           "BackupTask",
		LastEnabledUtc: nil,
		Enabled:        true,
		StopOnError:    false,
		LastStartUtc:   nil,
		LastEndUtc:     nil,
		LastSuccessUtc: nil,
	}

	mockRepo.On("Create", mock.Anything, newScheduleTask).Return(nil)

	err := usecase.Create(context.Background(), newScheduleTask)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestScheduleTaskUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.ScheduleTaskRepository)
	timeout := time.Duration(10)
	usecase := NewScheduleTaskUsecase(mockRepo, timeout)

	updatedScheduleTask := &domain.ScheduleTask{
		ID:             primitive.NewObjectID(), // Existing ID of the record to update
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
	*updatedScheduleTask.LastEnabledUtc = time.Now().AddDate(0, 0, -7) // Enabled 7 days ago
	*updatedScheduleTask.LastStartUtc = time.Now().AddDate(0, 0, -1)   // Started 1 day ago
	*updatedScheduleTask.LastEndUtc = time.Now().AddDate(0, 0, -1)     // Ended 1 day ago
	*updatedScheduleTask.LastSuccessUtc = time.Now().AddDate(0, 0, -1) // Succeeded 1 day ago

	mockRepo.On("Update", mock.Anything, updatedScheduleTask).Return(nil)

	err := usecase.Update(context.Background(), updatedScheduleTask)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestScheduleTaskUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.ScheduleTaskRepository)
	timeout := time.Duration(10)
	usecase := NewScheduleTaskUsecase(mockRepo, timeout)

	scheduleTaskID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, scheduleTaskID).Return(nil)

	err := usecase.Delete(context.Background(), scheduleTaskID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestScheduleTaskUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.ScheduleTaskRepository)
	timeout := time.Duration(10)
	usecase := NewScheduleTaskUsecase(mockRepo, timeout)
	fetchedScheduleTasks := []domain.ScheduleTask{
		{
			ID:             primitive.NewObjectID(),
			Name:           "Daily Data Backup",
			Seconds:        86400,
			Type:           "BackupTask",
			LastEnabledUtc: nil,
			Enabled:        true,
			StopOnError:    false,
			LastStartUtc:   nil,
			LastEndUtc:     nil,
			LastSuccessUtc: nil,
		},
		{
			ID:             primitive.NewObjectID(),
			Name:           "Weekly Data Cleanup",
			Seconds:        604800,
			Type:           "CleanupTask",
			LastEnabledUtc: new(time.Time),
			Enabled:        false,
			StopOnError:    true,
			LastStartUtc:   new(time.Time),
			LastEndUtc:     new(time.Time),
			LastSuccessUtc: new(time.Time),
		},
	}
	*fetchedScheduleTasks[1].LastEnabledUtc = time.Now().AddDate(0, 0, -7) // Enabled 7 days ago
	*fetchedScheduleTasks[1].LastStartUtc = time.Now().AddDate(0, 0, -1)   // Started 1 day ago
	*fetchedScheduleTasks[1].LastEndUtc = time.Now().AddDate(0, 0, -1)     // Ended 1 day ago
	*fetchedScheduleTasks[1].LastSuccessUtc = time.Now().AddDate(0, 0, -1) // Succeeded 1 day ago

	mockRepo.On("Fetch", mock.Anything).Return(fetchedScheduleTasks, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedScheduleTasks, result)
	mockRepo.AssertExpectations(t)
}
