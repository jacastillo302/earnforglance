package usecase_test

import (
	"context"
	domain "earnforglance/server/domain/directory"
	mocks "earnforglance/server/domain/mocks"
	test "earnforglance/server/usecase/directory"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestMeasureSettingsUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.MeasureSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewMeasureSettingsUsecase(mockRepo, timeout)

	directoryID := bson.NewObjectID().Hex()

	updatedMeasureSettings := domain.MeasureSettings{
		ID:              bson.NewObjectID(), // Existing ID of the record to update
		BaseDimensionID: 3,
		BaseWeightID:    4,
	}

	mockRepo.On("FetchByID", mock.Anything, directoryID).Return(updatedMeasureSettings, nil)

	result, err := usecase.FetchByID(context.Background(), directoryID)

	assert.NoError(t, err)
	assert.Equal(t, updatedMeasureSettings, result)
	mockRepo.AssertExpectations(t)
}

func TestMeasureSettingsUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.MeasureSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewMeasureSettingsUsecase(mockRepo, timeout)

	newMeasureSettings := &domain.MeasureSettings{
		BaseDimensionID: 1,
		BaseWeightID:    2,
	}

	mockRepo.On("Create", mock.Anything, newMeasureSettings).Return(nil)

	err := usecase.Create(context.Background(), newMeasureSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestMeasureSettingsUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.MeasureSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewMeasureSettingsUsecase(mockRepo, timeout)

	updatedMeasureSettings := &domain.MeasureSettings{
		ID:              bson.NewObjectID(), // Existing ID of the record to update
		BaseDimensionID: 3,
		BaseWeightID:    4,
	}

	mockRepo.On("Update", mock.Anything, updatedMeasureSettings).Return(nil)

	err := usecase.Update(context.Background(), updatedMeasureSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestMeasureSettingsUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.MeasureSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewMeasureSettingsUsecase(mockRepo, timeout)

	directoryID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, directoryID).Return(nil)

	err := usecase.Delete(context.Background(), directoryID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestMeasureSettingsUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.MeasureSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewMeasureSettingsUsecase(mockRepo, timeout)

	fetchedMeasureSettings := []domain.MeasureSettings{
		{
			ID:              bson.NewObjectID(),
			BaseDimensionID: 1,
			BaseWeightID:    2,
		},
		{
			ID:              bson.NewObjectID(),
			BaseDimensionID: 3,
			BaseWeightID:    4,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedMeasureSettings, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedMeasureSettings, result)
	mockRepo.AssertExpectations(t)
}
