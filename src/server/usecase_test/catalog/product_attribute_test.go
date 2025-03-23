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

func TestProductAttributeUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.ProductAttributeRepository)
	timeout := time.Duration(10)
	usecase := test.NewProductAttributeUsecase(mockRepo, timeout)

	productAttributeID := primitive.NewObjectID().Hex()

	updatedProductAttribute := domain.ProductAttribute{
		ID:          primitive.NewObjectID(),
		Name:        "Size",
		Description: "Defines the size of the product",
	}

	mockRepo.On("FetchByID", mock.Anything, productAttributeID).Return(updatedProductAttribute, nil)

	result, err := usecase.FetchByID(context.Background(), productAttributeID)

	assert.NoError(t, err)
	assert.Equal(t, updatedProductAttribute, result)
	mockRepo.AssertExpectations(t)
}

func TestProductAttributeUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.ProductAttributeRepository)
	timeout := time.Duration(10)
	usecase := test.NewProductAttributeUsecase(mockRepo, timeout)

	newProductAttribute := &domain.ProductAttribute{
		Name:        "Color",
		Description: "Defines the color of the product",
	}

	mockRepo.On("Create", mock.Anything, newProductAttribute).Return(nil)

	err := usecase.Create(context.Background(), newProductAttribute)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductAttributeUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.ProductAttributeRepository)
	timeout := time.Duration(10)
	usecase := test.NewProductAttributeUsecase(mockRepo, timeout)

	updatedProductAttribute := &domain.ProductAttribute{
		ID:          primitive.NewObjectID(),
		Name:        "Size",
		Description: "Defines the size of the product",
	}

	mockRepo.On("Update", mock.Anything, updatedProductAttribute).Return(nil)

	err := usecase.Update(context.Background(), updatedProductAttribute)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductAttributeUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.ProductAttributeRepository)
	timeout := time.Duration(10)
	usecase := test.NewProductAttributeUsecase(mockRepo, timeout)

	productAttributeID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, productAttributeID).Return(nil)

	err := usecase.Delete(context.Background(), productAttributeID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductAttributeUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.ProductAttributeRepository)
	timeout := time.Duration(10)
	usecase := test.NewProductAttributeUsecase(mockRepo, timeout)

	fetchedProductAttributes := []domain.ProductAttribute{
		{
			ID:          primitive.NewObjectID(),
			Name:        "Material",
			Description: "Defines the material of the product",
		},
		{
			ID:          primitive.NewObjectID(),
			Name:        "Brand",
			Description: "Defines the brand of the product",
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedProductAttributes, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedProductAttributes, result)
	mockRepo.AssertExpectations(t)
}
