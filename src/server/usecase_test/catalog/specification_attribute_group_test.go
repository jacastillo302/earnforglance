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

func TestSpecificationAttributeGroupUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.SpecificationAttributeGroupRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewSpecificationAttributeGroupUsecase(mockRepo, timeout)

	specificationAttributeGroupID := bson.NewObjectID().Hex()

	updatedSpecificationAttributeGroup := domain.SpecificationAttributeGroup{
		ID:           bson.NewObjectID(), // Existing ID of the record to update
		Name:         "Technical Specifications",
		DisplayOrder: 2,
	}

	mockRepo.On("FetchByID", mock.Anything, specificationAttributeGroupID).Return(updatedSpecificationAttributeGroup, nil)

	result, err := usecase.FetchByID(context.Background(), specificationAttributeGroupID)

	assert.NoError(t, err)
	assert.Equal(t, updatedSpecificationAttributeGroup, result)
	mockRepo.AssertExpectations(t)
}

func TestSpecificationAttributeGroupUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.SpecificationAttributeGroupRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewSpecificationAttributeGroupUsecase(mockRepo, timeout)

	newSpecificationAttributeGroup := &domain.SpecificationAttributeGroup{
		Name:         "General Specifications",
		DisplayOrder: 1,
	}

	mockRepo.On("Create", mock.Anything, newSpecificationAttributeGroup).Return(nil)

	err := usecase.Create(context.Background(), newSpecificationAttributeGroup)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestSpecificationAttributeGroupUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.SpecificationAttributeGroupRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewSpecificationAttributeGroupUsecase(mockRepo, timeout)

	updatedSpecificationAttributeGroup := &domain.SpecificationAttributeGroup{
		ID:           bson.NewObjectID(), // Existing ID of the record to update
		Name:         "Technical Specifications",
		DisplayOrder: 2,
	}

	mockRepo.On("Update", mock.Anything, updatedSpecificationAttributeGroup).Return(nil)

	err := usecase.Update(context.Background(), updatedSpecificationAttributeGroup)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestSpecificationAttributeGroupUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.SpecificationAttributeGroupRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewSpecificationAttributeGroupUsecase(mockRepo, timeout)

	specificationAttributeGroupID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, specificationAttributeGroupID).Return(nil)

	err := usecase.Delete(context.Background(), specificationAttributeGroupID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestSpecificationAttributeGroupUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.SpecificationAttributeGroupRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewSpecificationAttributeGroupUsecase(mockRepo, timeout)

	fetchedSpecificationAttributeGroups := []domain.SpecificationAttributeGroup{
		{
			ID:           bson.NewObjectID(),
			Name:         "General Specifications",
			DisplayOrder: 1,
		},
		{
			ID:           bson.NewObjectID(),
			Name:         "Technical Specifications",
			DisplayOrder: 2,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedSpecificationAttributeGroups, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedSpecificationAttributeGroups, result)
	mockRepo.AssertExpectations(t)
}
