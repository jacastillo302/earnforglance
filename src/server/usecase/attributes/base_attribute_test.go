package usecase

import (
	"context"
	domian "earnforglance/server/domain/attributes"
	mocks "earnforglance/server/domain/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestBaseAttributeUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.BaseAttributeRepository)
	time := time.Duration(10)
	usecase := NewBaseAttributeUsecase(mockRepo, time) // Assuming a constructor exists

	attributesID := primitive.NewObjectID().Hex()

	expectedBaseAttribute := domian.BaseAttribute{
		ID:                     primitive.NewObjectID(),
		Name:                   "Test Attribute",
		IsRequired:             true,
		AttributeControlTypeId: 1, // Example: TextBox
		DisplayOrder:           5,
	}

	mockRepo.On("FetchByID", mock.Anything, attributesID).Return(expectedBaseAttribute, nil)

	result, err := usecase.FetchByID(context.Background(), attributesID)

	assert.NoError(t, err)
	assert.Equal(t, expectedBaseAttribute, result)
	mockRepo.AssertExpectations(t)
}

func TestBaseAttributeUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.BaseAttributeRepository)
	time := time.Duration(10)
	usecase := NewBaseAttributeUsecase(mockRepo, time) // Assuming a constructor exists

	newBaseAttribute := &domian.BaseAttribute{
		ID:                     primitive.NewObjectID(),
		Name:                   "Test Attribute",
		IsRequired:             true,
		AttributeControlTypeId: 1, // Example: TextBox
		DisplayOrder:           5,
	}

	mockRepo.On("Create", mock.Anything, newBaseAttribute).Return(nil)

	err := usecase.Create(context.Background(), newBaseAttribute)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestBaseAttributeUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.BaseAttributeRepository)
	time := time.Duration(10)
	usecase := NewBaseAttributeUsecase(mockRepo, time)

	updatedBaseAttribute := &domian.BaseAttribute{
		ID:                     primitive.NewObjectID(),
		Name:                   "Test Attribute",
		IsRequired:             true,
		AttributeControlTypeId: 1, // Example: TextBox
		DisplayOrder:           5,
	}

	mockRepo.On("Update", mock.Anything, updatedBaseAttribute).Return(nil)

	err := usecase.Update(context.Background(), updatedBaseAttribute)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestBaseAttributeUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.BaseAttributeRepository)
	time := time.Duration(10)
	usecase := NewBaseAttributeUsecase(mockRepo, time)

	attributesID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, attributesID).Return(nil)

	err := usecase.Delete(context.Background(), attributesID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestBaseAttributeUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.BaseAttributeRepository)
	time := time.Duration(10)
	usecase := NewBaseAttributeUsecase(mockRepo, time)

	expectedBaseAttributes := []domian.BaseAttribute{
		{
			ID:                     primitive.NewObjectID(),
			Name:                   "Test Attribute",
			IsRequired:             true,
			AttributeControlTypeId: 1, // Example: TextBox
			DisplayOrder:           5,
		},
		{
			ID:                     primitive.NewObjectID(),
			Name:                   "Test Attribute",
			IsRequired:             true,
			AttributeControlTypeId: 1, // Example: TextBox
			DisplayOrder:           5,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(expectedBaseAttributes, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, expectedBaseAttributes, result)
	mockRepo.AssertExpectations(t)
}
