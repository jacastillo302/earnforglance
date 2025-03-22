package usecase

import (
	"context"
	domain "earnforglance/server/domain/directory"
	mocks "earnforglance/server/domain/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestMeasureDimensionUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.MeasureDimensionRepository)
	timeout := time.Duration(10)
	usecase := NewMeasureDimensionUsecase(mockRepo, timeout)

	directoryID := primitive.NewObjectID().Hex()

	updatedMeasureDimension := domain.MeasureDimension{
		ID:            primitive.NewObjectID(), // Existing ID of the record to update
		Name:          "Inch",
		SystemKeyword: "in",
		Ratio:         2.54,
		DisplayOrder:  2,
	}

	mockRepo.On("FetchByID", mock.Anything, directoryID).Return(updatedMeasureDimension, nil)

	result, err := usecase.FetchByID(context.Background(), directoryID)

	assert.NoError(t, err)
	assert.Equal(t, updatedMeasureDimension, result)
	mockRepo.AssertExpectations(t)
}

func TestMeasureDimensionUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.MeasureDimensionRepository)
	timeout := time.Duration(10)
	usecase := NewMeasureDimensionUsecase(mockRepo, timeout)

	newMeasureDimension := &domain.MeasureDimension{
		Name:          "Centimeter",
		SystemKeyword: "cm",
		Ratio:         1.0,
		DisplayOrder:  1,
	}

	mockRepo.On("Create", mock.Anything, newMeasureDimension).Return(nil)

	err := usecase.Create(context.Background(), newMeasureDimension)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestMeasureDimensionUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.MeasureDimensionRepository)
	timeout := time.Duration(10)
	usecase := NewMeasureDimensionUsecase(mockRepo, timeout)

	updatedMeasureDimension := &domain.MeasureDimension{
		ID:            primitive.NewObjectID(), // Existing ID of the record to update
		Name:          "Inch",
		SystemKeyword: "in",
		Ratio:         2.54,
		DisplayOrder:  2,
	}

	mockRepo.On("Update", mock.Anything, updatedMeasureDimension).Return(nil)

	err := usecase.Update(context.Background(), updatedMeasureDimension)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestMeasureDimensionUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.MeasureDimensionRepository)
	timeout := time.Duration(10)
	usecase := NewMeasureDimensionUsecase(mockRepo, timeout)

	directoryID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, directoryID).Return(nil)

	err := usecase.Delete(context.Background(), directoryID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestMeasureDimensionUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.MeasureDimensionRepository)
	timeout := time.Duration(10)
	usecase := NewMeasureDimensionUsecase(mockRepo, timeout)

	fetchedMeasureDimensions := []domain.MeasureDimension{
		{
			ID:            primitive.NewObjectID(),
			Name:          "Centimeter",
			SystemKeyword: "cm",
			Ratio:         1.0,
			DisplayOrder:  1,
		},
		{
			ID:            primitive.NewObjectID(),
			Name:          "Inch",
			SystemKeyword: "in",
			Ratio:         2.54,
			DisplayOrder:  2,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedMeasureDimensions, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedMeasureDimensions, result)
	mockRepo.AssertExpectations(t)
}
