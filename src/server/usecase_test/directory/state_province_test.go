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

func TestStateProvinceUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.StateProvinceRepository)
	timeout := time.Duration(10)
	usecase := test.NewStateProvinceUsecase(mockRepo, timeout)

	stateprovinceID := bson.NewObjectID().Hex()

	updatedStateProvince := domain.StateProvince{
		ID:           bson.NewObjectID(), // Existing ID of the record to update
		CountryID:    bson.NewObjectID(),
		Name:         "Ontario",
		Abbreviation: "ON",
		Published:    false,
		DisplayOrder: 2,
	}

	mockRepo.On("FetchByID", mock.Anything, stateprovinceID).Return(updatedStateProvince, nil)

	result, err := usecase.FetchByID(context.Background(), stateprovinceID)

	assert.NoError(t, err)
	assert.Equal(t, updatedStateProvince, result)
	mockRepo.AssertExpectations(t)
}

func TestStateProvinceUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.StateProvinceRepository)
	timeout := time.Duration(10)
	usecase := test.NewStateProvinceUsecase(mockRepo, timeout)

	newStateProvince := &domain.StateProvince{
		CountryID:    bson.NewObjectID(),
		Name:         "California",
		Abbreviation: "CA",
		Published:    true,
		DisplayOrder: 1,
	}

	mockRepo.On("Create", mock.Anything, newStateProvince).Return(nil)

	err := usecase.Create(context.Background(), newStateProvince)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestStateProvinceUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.StateProvinceRepository)
	timeout := time.Duration(10)
	usecase := test.NewStateProvinceUsecase(mockRepo, timeout)

	updatedStateProvince := &domain.StateProvince{
		ID:           bson.NewObjectID(), // Existing ID of the record to update
		CountryID:    bson.NewObjectID(),
		Name:         "Ontario",
		Abbreviation: "ON",
		Published:    false,
		DisplayOrder: 2,
	}

	mockRepo.On("Update", mock.Anything, updatedStateProvince).Return(nil)

	err := usecase.Update(context.Background(), updatedStateProvince)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestStateProvinceUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.StateProvinceRepository)
	timeout := time.Duration(10)
	usecase := test.NewStateProvinceUsecase(mockRepo, timeout)

	stateprovinceID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, stateprovinceID).Return(nil)

	err := usecase.Delete(context.Background(), stateprovinceID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestStateProvinceUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.StateProvinceRepository)
	timeout := time.Duration(10)
	usecase := test.NewStateProvinceUsecase(mockRepo, timeout)

	fetchedStateProvinces := []domain.StateProvince{
		{
			ID:           bson.NewObjectID(),
			CountryID:    bson.NewObjectID(),
			Name:         "California",
			Abbreviation: "CA",
			Published:    true,
			DisplayOrder: 1,
		},
		{
			ID:           bson.NewObjectID(),
			CountryID:    bson.NewObjectID(),
			Name:         "Ontario",
			Abbreviation: "ON",
			Published:    false,
			DisplayOrder: 2,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedStateProvinces, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedStateProvinces, result)
	mockRepo.AssertExpectations(t)
}
