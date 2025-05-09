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

func TestProductAttributeMappingUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.ProductAttributeMappingRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewProductAttributeMappingUsecase(mockRepo, timeout)

	productAttributeMappingID := bson.NewObjectID().Hex()

	expectedProductAttributeMapping := domain.ProductAttributeMapping{
		ProductID:                       bson.NewObjectID(),
		ProductAttributeID:              bson.NewObjectID(),
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
	usecase := test.NewProductAttributeMappingUsecase(mockRepo, timeout)

	newProductAttributeMapping := &domain.ProductAttributeMapping{
		ProductID:                       bson.NewObjectID(),
		ProductAttributeID:              bson.NewObjectID(),
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
	}

	mockRepo.On("Create", mock.Anything, newProductAttributeMapping).Return(nil)

	err := usecase.Create(context.Background(), newProductAttributeMapping)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductAttributeMappingUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.ProductAttributeMappingRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewProductAttributeMappingUsecase(mockRepo, timeout)

	updatedProductAttributeMapping := &domain.ProductAttributeMapping{
		ProductID:                       bson.NewObjectID(),
		ProductAttributeID:              bson.NewObjectID(),
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
	usecase := test.NewProductAttributeMappingUsecase(mockRepo, timeout)

	productAttributeMappingID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, productAttributeMappingID).Return(nil)

	err := usecase.Delete(context.Background(), productAttributeMappingID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductAttributeMappingUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.ProductAttributeMappingRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewProductAttributeMappingUsecase(mockRepo, timeout)

	fetchedProductAttributeMappings := []domain.ProductAttributeMapping{
		{
			ProductID:                       bson.NewObjectID(),
			ProductAttributeID:              bson.NewObjectID(),
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
		},
		{
			ProductID:                       bson.NewObjectID(),
			ProductAttributeID:              bson.NewObjectID(),
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
