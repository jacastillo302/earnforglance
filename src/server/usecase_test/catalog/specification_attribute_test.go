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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestSpecificationAttributeUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.SpecificationAttributeRepository)
	timeout := time.Duration(10)
	usecase := test.NewSpecificationAttributeUsecase(mockRepo, timeout)

	specificationAttributeID := primitive.NewObjectID().Hex()
	specificationAttributeGroupID := primitive.NewObjectID() // Create the ObjectID
	specificationAttributeGroupIDPointer := &specificationAttributeGroupID

	updatedSpecificationAttribute := domain.SpecificationAttribute{
		ID:                            primitive.NewObjectID(), // Existing ID of the record to update
		Name:                          "Color",
		DisplayOrder:                  2,
		SpecificationAttributeGroupID: specificationAttributeGroupIDPointer,
	}

	mockRepo.On("FetchByID", mock.Anything, specificationAttributeID).Return(updatedSpecificationAttribute, nil)

	result, err := usecase.FetchByID(context.Background(), specificationAttributeID)

	assert.NoError(t, err)
	assert.Equal(t, updatedSpecificationAttribute, result)
	mockRepo.AssertExpectations(t)
}

func TestSpecificationAttributeUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.SpecificationAttributeRepository)
	timeout := time.Duration(10)
	usecase := test.NewSpecificationAttributeUsecase(mockRepo, timeout)

	newSpecificationAttribute := &domain.SpecificationAttribute{
		Name:                          "Material",
		DisplayOrder:                  1,
		SpecificationAttributeGroupID: nil,
	}

	mockRepo.On("Create", mock.Anything, newSpecificationAttribute).Return(nil)

	err := usecase.Create(context.Background(), newSpecificationAttribute)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestSpecificationAttributeUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.SpecificationAttributeRepository)
	timeout := time.Duration(10)
	usecase := test.NewSpecificationAttributeUsecase(mockRepo, timeout)

	updatedSpecificationAttribute := &domain.SpecificationAttribute{
		ID:                            primitive.NewObjectID(), // Existing ID of the record to update
		Name:                          "Color",
		DisplayOrder:                  2,
		SpecificationAttributeGroupID: nil,
	}

	mockRepo.On("Update", mock.Anything, updatedSpecificationAttribute).Return(nil)

	err := usecase.Update(context.Background(), updatedSpecificationAttribute)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestSpecificationAttributeUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.SpecificationAttributeRepository)
	timeout := time.Duration(10)
	usecase := test.NewSpecificationAttributeUsecase(mockRepo, timeout)

	specificationAttributeID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, specificationAttributeID).Return(nil)

	err := usecase.Delete(context.Background(), specificationAttributeID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestSpecificationAttributeUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.SpecificationAttributeRepository)
	timeout := time.Duration(10)
	usecase := test.NewSpecificationAttributeUsecase(mockRepo, timeout)

	fetchedSpecificationAttributes := []domain.SpecificationAttribute{
		{
			ID:                            primitive.NewObjectID(),
			Name:                          "Material",
			DisplayOrder:                  1,
			SpecificationAttributeGroupID: nil,
		},
		{
			ID:                            primitive.NewObjectID(),
			Name:                          "Color",
			DisplayOrder:                  2,
			SpecificationAttributeGroupID: nil,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedSpecificationAttributes, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedSpecificationAttributes, result)
	mockRepo.AssertExpectations(t)
}
