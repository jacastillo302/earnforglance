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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestMeasureWeightUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.MeasureWeightRepository)
	timeout := time.Duration(10)
	usecase := test.NewMeasureWeightUsecase(mockRepo, timeout)

	measureWeightID := primitive.NewObjectID().Hex()

	updatedMeasureWeight := domain.MeasureWeight{
		ID:            primitive.NewObjectID(), // Existing ID of the record to update
		Name:          "Pound",
		SystemKeyword: "lb",
		Ratio:         2.20462,
		DisplayOrder:  2,
	}

	mockRepo.On("FetchByID", mock.Anything, measureWeightID).Return(updatedMeasureWeight, nil)

	result, err := usecase.FetchByID(context.Background(), measureWeightID)

	assert.NoError(t, err)
	assert.Equal(t, updatedMeasureWeight, result)
	mockRepo.AssertExpectations(t)
}

func TestMeasureWeightUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.MeasureWeightRepository)
	timeout := time.Duration(10)
	usecase := test.NewMeasureWeightUsecase(mockRepo, timeout)

	newMeasureWeight := &domain.MeasureWeight{
		Name:          "Kilogram",
		SystemKeyword: "kg",
		Ratio:         1.0,
		DisplayOrder:  1,
	}

	mockRepo.On("Create", mock.Anything, newMeasureWeight).Return(nil)

	err := usecase.Create(context.Background(), newMeasureWeight)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestMeasureWeightUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.MeasureWeightRepository)
	timeout := time.Duration(10)
	usecase := test.NewMeasureWeightUsecase(mockRepo, timeout)

	updatedMeasureWeight := &domain.MeasureWeight{
		ID:            primitive.NewObjectID(), // Existing ID of the record to update
		Name:          "Pound",
		SystemKeyword: "lb",
		Ratio:         2.20462,
		DisplayOrder:  2,
	}

	mockRepo.On("Update", mock.Anything, updatedMeasureWeight).Return(nil)

	err := usecase.Update(context.Background(), updatedMeasureWeight)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestMeasureWeightUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.MeasureWeightRepository)
	timeout := time.Duration(10)
	usecase := test.NewMeasureWeightUsecase(mockRepo, timeout)

	measureWeightID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, measureWeightID).Return(nil)

	err := usecase.Delete(context.Background(), measureWeightID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestMeasureWeightUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.MeasureWeightRepository)
	timeout := time.Duration(10)
	usecase := test.NewMeasureWeightUsecase(mockRepo, timeout)

	fetchedMeasureWeights := []domain.MeasureWeight{
		{
			ID:            primitive.NewObjectID(),
			Name:          "Kilogram",
			SystemKeyword: "kg",
			Ratio:         1.0,
			DisplayOrder:  1,
		},
		{
			ID:            primitive.NewObjectID(),
			Name:          "Pound",
			SystemKeyword: "lb",
			Ratio:         2.20462,
			DisplayOrder:  2,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedMeasureWeights, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedMeasureWeights, result)
	mockRepo.AssertExpectations(t)
}
