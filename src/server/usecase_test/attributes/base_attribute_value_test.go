package usecase_test

import (
	"context"
	domain "earnforglance/server/domain/attributes"
	mocks "earnforglance/server/domain/mocks"
	test "earnforglance/server/usecase/attributes"

	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestBaseAttributeValueUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.BaseAttributeValueRepository)
	timeout := time.Duration(10)
	usecase := test.NewBaseAttributeValueUsecase(mockRepo, timeout)

	baseAttributeValueID := primitive.NewObjectID().Hex()

	expectedBaseAttributeValue := domain.BaseAttributeValue{
		ID:            primitive.NewObjectID(), // This can be omitted for creation as MongoDB generates it
		Name:          "Color",
		IsPreSelected: false,
		DisplayOrder:  1,
		AttributeId:   primitive.NewObjectID(), // Reference to the related attribute
	}

	mockRepo.On("FetchByID", mock.Anything, baseAttributeValueID).Return(expectedBaseAttributeValue, nil)

	result, err := usecase.FetchByID(context.Background(), baseAttributeValueID)

	assert.NoError(t, err)
	assert.Equal(t, expectedBaseAttributeValue, result)
	mockRepo.AssertExpectations(t)
}

func TestBaseAttributeValueUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.BaseAttributeValueRepository)
	timeout := time.Duration(10)
	usecase := test.NewBaseAttributeValueUsecase(mockRepo, timeout)

	newBaseAttributeValue := &domain.BaseAttributeValue{
		ID:            primitive.NewObjectID(), // Existing ID of the record to update
		Name:          "Size",
		IsPreSelected: true,
		DisplayOrder:  2,
		AttributeId:   primitive.NewObjectID(), // Reference to the related attribute
	}

	mockRepo.On("Create", mock.Anything, newBaseAttributeValue).Return(nil)

	err := usecase.Create(context.Background(), newBaseAttributeValue)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestBaseAttributeValueUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.BaseAttributeValueRepository)
	timeout := time.Duration(10)
	usecase := test.NewBaseAttributeValueUsecase(mockRepo, timeout)

	updatedBaseAttributeValue := &domain.BaseAttributeValue{
		ID:            primitive.NewObjectID(), // Existing ID of the record to update
		Name:          "Size",
		IsPreSelected: true,
		DisplayOrder:  2,
		AttributeId:   primitive.NewObjectID(), // Reference to the related attribute
	}

	mockRepo.On("Update", mock.Anything, updatedBaseAttributeValue).Return(nil)

	err := usecase.Update(context.Background(), updatedBaseAttributeValue)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestBaseAttributeValueUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.BaseAttributeValueRepository)
	timeout := time.Duration(10)
	usecase := test.NewBaseAttributeValueUsecase(mockRepo, timeout)

	baseAttributeValueID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, baseAttributeValueID).Return(nil)

	err := usecase.Delete(context.Background(), baseAttributeValueID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestBaseAttributeValueUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.BaseAttributeValueRepository)
	timeout := time.Duration(10)
	usecase := test.NewBaseAttributeValueUsecase(mockRepo, timeout)

	expectedBaseAttributeValues := []domain.BaseAttributeValue{
		{
			ID:            primitive.NewObjectID(),
			Name:          "Material",
			IsPreSelected: false,
			DisplayOrder:  3,
			AttributeId:   primitive.NewObjectID(),
		},
		{
			ID:            primitive.NewObjectID(),
			Name:          "Brand",
			IsPreSelected: true,
			DisplayOrder:  4,
			AttributeId:   primitive.NewObjectID(),
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(expectedBaseAttributeValues, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, expectedBaseAttributeValues, result)
	mockRepo.AssertExpectations(t)
}
