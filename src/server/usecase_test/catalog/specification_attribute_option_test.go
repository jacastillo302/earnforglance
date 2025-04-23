package usecase_test

import (
	"context"
	domain "earnforglance/server/domain/catalog"
	mocks "earnforglance/server/domain/mocks"
	test "earnforglance/server/usecase/catalog"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestSpecificationAttributeOptionUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.SpecificationAttributeOptionRepository)
	timeout := time.Duration(10)
	usecase := test.NewSpecificationAttributeOptionUsecase(mockRepo, timeout)

	specificationAttributeOptionID := bson.NewObjectID().Hex()

	updatedSpecificationAttributeOption := domain.SpecificationAttributeOption{
		ID:                       bson.NewObjectID(), // Existing ID of the record to update
		SpecificationAttributeID: bson.NewObjectID(),
		Name:                     "Size",
		ColorSquaresRgb:          "",
		DisplayOrder:             2,
	}

	mockRepo.On("FetchByID", mock.Anything, specificationAttributeOptionID).Return(updatedSpecificationAttributeOption, nil)

	result, err := usecase.FetchByID(context.Background(), specificationAttributeOptionID)

	assert.NoError(t, err)
	assert.Equal(t, updatedSpecificationAttributeOption, result)
	mockRepo.AssertExpectations(t)
}

func TestSpecificationAttributeOptionUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.SpecificationAttributeOptionRepository)
	timeout := time.Duration(10)
	usecase := test.NewSpecificationAttributeOptionUsecase(mockRepo, timeout)

	newSpecificationAttributeOption := &domain.SpecificationAttributeOption{
		SpecificationAttributeID: bson.NewObjectID(),
		Name:                     "Color",
		ColorSquaresRgb:          "#FF0000",
		DisplayOrder:             1,
	}

	mockRepo.On("Create", mock.Anything, newSpecificationAttributeOption).Return(nil)

	err := usecase.Create(context.Background(), newSpecificationAttributeOption)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestSpecificationAttributeOptionUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.SpecificationAttributeOptionRepository)
	timeout := time.Duration(10)
	usecase := test.NewSpecificationAttributeOptionUsecase(mockRepo, timeout)

	updatedSpecificationAttributeOption := &domain.SpecificationAttributeOption{
		ID:                       bson.NewObjectID(), // Existing ID of the record to update
		SpecificationAttributeID: bson.NewObjectID(),
		Name:                     "Size",
		ColorSquaresRgb:          "",
		DisplayOrder:             2,
	}

	mockRepo.On("Update", mock.Anything, updatedSpecificationAttributeOption).Return(nil)

	err := usecase.Update(context.Background(), updatedSpecificationAttributeOption)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestSpecificationAttributeOptionUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.SpecificationAttributeOptionRepository)
	timeout := time.Duration(10)
	usecase := test.NewSpecificationAttributeOptionUsecase(mockRepo, timeout)

	specificationAttributeOptionID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, specificationAttributeOptionID).Return(nil)

	err := usecase.Delete(context.Background(), specificationAttributeOptionID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestSpecificationAttributeOptionUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.SpecificationAttributeOptionRepository)
	timeout := time.Duration(10)
	usecase := test.NewSpecificationAttributeOptionUsecase(mockRepo, timeout)

	fetchedSpecificationAttributeOptions := []domain.SpecificationAttributeOption{
		{
			ID:                       bson.NewObjectID(),
			SpecificationAttributeID: bson.NewObjectID(),
			Name:                     "Color",
			ColorSquaresRgb:          "#FF0000",
			DisplayOrder:             1,
		},
		{
			ID:                       bson.NewObjectID(),
			SpecificationAttributeID: bson.NewObjectID(),
			Name:                     "Size",
			ColorSquaresRgb:          "",
			DisplayOrder:             2,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedSpecificationAttributeOptions, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedSpecificationAttributeOptions, result)
	mockRepo.AssertExpectations(t)
}
