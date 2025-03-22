package usecase

import (
	"context"
	domain "earnforglance/server/domain/catalog"
	mocks "earnforglance/server/domain/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestProductAttributeMappingUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.ProductAttributeMappingRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewProductAttributeMappingUsecase(mockRepo, timeout)

	productAttributeMappingID := primitive.NewObjectID().Hex()

	expectedProductAttributeMapping := domain.ProductAttributeMapping{
		ID:                              primitive.NewObjectID(), // Existing ID of the record to update
		ProductID:                       primitive.NewObjectID(),
		ProductAttributeID:              primitive.NewObjectID(),
		TextPrompt:                      "Select a size",
		IsRequired:                      false,
		AttributeControlTypeID:          2,
		DisplayOrder:                    2,
		ValidationMinLength:             new(int),
		ValidationMaxLength:             new(int),
		ValidationFileAllowedExtensions: ".jpg,.png",
		ValidationFileMaximumSize:       new(int),
		DefaultValue:                    "Medium",
		ConditionAttributeXml:           "<attributes><size>medium</size></attributes>",
		AttributeControlType:            6,
	}

	mockRepo.On("FetchByID", mock.Anything, productAttributeMappingID).Return(expectedProductAttributeMapping, nil)

	result, err := usecase.FetchByID(context.Background(), productAttributeMappingID)

	assert.NoError(t, err)
	assert.Equal(t, expectedProductAttributeMapping, result)
	mockRepo.AssertExpectations(t)
}

func TestProductAttributeMappingUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.ProductAttributeMappingRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewProductAttributeMappingUsecase(mockRepo, timeout)

	newProductAttributeMapping := &domain.ProductAttributeMapping{
		ProductID:                       primitive.NewObjectID(),
		ProductAttributeID:              primitive.NewObjectID(),
		TextPrompt:                      "Select a color",
		IsRequired:                      true,
		AttributeControlTypeID:          1,
		DisplayOrder:                    1,
		ValidationMinLength:             nil,
		ValidationMaxLength:             nil,
		ValidationFileAllowedExtensions: "",
		ValidationFileMaximumSize:       nil,
		DefaultValue:                    "Red",
		ConditionAttributeXml:           "<attributes><color>red</color></attributes>",
		AttributeControlType:            5,
	}

	mockRepo.On("Create", mock.Anything, newProductAttributeMapping).Return(nil)

	err := usecase.Create(context.Background(), newProductAttributeMapping)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductAttributeMappingUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.ProductAttributeMappingRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewProductAttributeMappingUsecase(mockRepo, timeout)

	updatedProductAttributeMapping := &domain.ProductAttributeMapping{
		ID:                              primitive.NewObjectID(), // Existing ID of the record to update
		ProductID:                       primitive.NewObjectID(),
		ProductAttributeID:              primitive.NewObjectID(),
		TextPrompt:                      "Select a size",
		IsRequired:                      false,
		AttributeControlTypeID:          2,
		DisplayOrder:                    2,
		ValidationMinLength:             new(int),
		ValidationMaxLength:             new(int),
		ValidationFileAllowedExtensions: ".jpg,.png",
		ValidationFileMaximumSize:       new(int),
		DefaultValue:                    "Medium",
		ConditionAttributeXml:           "<attributes><size>medium</size></attributes>",
		AttributeControlType:            5,
	}
	*updatedProductAttributeMapping.ValidationMinLength = 5
	*updatedProductAttributeMapping.ValidationMaxLength = 50
	*updatedProductAttributeMapping.ValidationFileMaximumSize = 1024

	mockRepo.On("Update", mock.Anything, updatedProductAttributeMapping).Return(nil)

	err := usecase.Update(context.Background(), updatedProductAttributeMapping)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductAttributeMappingUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.ProductAttributeMappingRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewProductAttributeMappingUsecase(mockRepo, timeout)

	productAttributeMappingID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, productAttributeMappingID).Return(nil)

	err := usecase.Delete(context.Background(), productAttributeMappingID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductAttributeMappingUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.ProductAttributeMappingRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewProductAttributeMappingUsecase(mockRepo, timeout)

	fetchedProductAttributeMappings := []domain.ProductAttributeMapping{
		{
			ID:                              primitive.NewObjectID(),
			ProductID:                       primitive.NewObjectID(),
			ProductAttributeID:              primitive.NewObjectID(),
			TextPrompt:                      "Select a color",
			IsRequired:                      true,
			AttributeControlTypeID:          1,
			DisplayOrder:                    1,
			ValidationMinLength:             nil,
			ValidationMaxLength:             nil,
			ValidationFileAllowedExtensions: "",
			ValidationFileMaximumSize:       nil,
			DefaultValue:                    "Red",
			ConditionAttributeXml:           "<attributes><color>red</color></attributes>",
			AttributeControlType:            5,
		},
		{
			ID:                              primitive.NewObjectID(),
			ProductID:                       primitive.NewObjectID(),
			ProductAttributeID:              primitive.NewObjectID(),
			TextPrompt:                      "Select a size",
			IsRequired:                      false,
			AttributeControlTypeID:          2,
			DisplayOrder:                    2,
			ValidationMinLength:             new(int),
			ValidationMaxLength:             new(int),
			ValidationFileAllowedExtensions: ".jpg,.png",
			ValidationFileMaximumSize:       new(int),
			DefaultValue:                    "Medium",
			ConditionAttributeXml:           "<attributes><size>medium</size></attributes>",
			AttributeControlType:            1,
		},
	}
	*fetchedProductAttributeMappings[1].ValidationMinLength = 5
	*fetchedProductAttributeMappings[1].ValidationMaxLength = 50
	*fetchedProductAttributeMappings[1].ValidationFileMaximumSize = 1024

	mockRepo.On("Fetch", mock.Anything).Return(fetchedProductAttributeMappings, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedProductAttributeMappings, result)
	mockRepo.AssertExpectations(t)
}
