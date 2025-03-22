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

func TestProductTagUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.ProductTagRepository)
	timeout := time.Duration(10)
	usecase := NewProductTagUsecase(mockRepo, timeout)

	productTagID := primitive.NewObjectID().Hex()

	updatedProductTag := domain.ProductTag{
		ID:              primitive.NewObjectID(), // Existing ID of the record to update
		Name:            "Home Appliances",
		MetaDescription: "Tags related to home appliances",
		MetaKeywords:    "appliances, home, kitchen",
		MetaTitle:       "Home Appliances Tag",
	}

	mockRepo.On("FetchByID", mock.Anything, productTagID).Return(updatedProductTag, nil)

	result, err := usecase.FetchByID(context.Background(), productTagID)

	assert.NoError(t, err)
	assert.Equal(t, updatedProductTag, result)
	mockRepo.AssertExpectations(t)
}

func TestProductTagUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.ProductTagRepository)
	timeout := time.Duration(10)
	usecase := NewProductTagUsecase(mockRepo, timeout)

	newProductTag := &domain.ProductTag{
		Name:            "Electronics",
		MetaDescription: "Tags related to electronic products",
		MetaKeywords:    "electronics, gadgets, devices",
		MetaTitle:       "Electronics Tag",
	}

	mockRepo.On("Create", mock.Anything, newProductTag).Return(nil)

	err := usecase.Create(context.Background(), newProductTag)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductTagUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.ProductTagRepository)
	timeout := time.Duration(10)
	usecase := NewProductTagUsecase(mockRepo, timeout)

	updatedProductTag := &domain.ProductTag{
		ID:              primitive.NewObjectID(), // Existing ID of the record to update
		Name:            "Home Appliances",
		MetaDescription: "Tags related to home appliances",
		MetaKeywords:    "appliances, home, kitchen",
		MetaTitle:       "Home Appliances Tag",
	}

	mockRepo.On("Update", mock.Anything, updatedProductTag).Return(nil)

	err := usecase.Update(context.Background(), updatedProductTag)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductTagUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.ProductTagRepository)
	timeout := time.Duration(10)
	usecase := NewProductTagUsecase(mockRepo, timeout)

	productTagID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, productTagID).Return(nil)

	err := usecase.Delete(context.Background(), productTagID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductTagUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.ProductTagRepository)
	timeout := time.Duration(10)
	usecase := NewProductTagUsecase(mockRepo, timeout)

	fetchedProductTags := []domain.ProductTag{
		{
			ID:              primitive.NewObjectID(),
			Name:            "Electronics",
			MetaDescription: "Tags related to electronic products",
			MetaKeywords:    "electronics, gadgets, devices",
			MetaTitle:       "Electronics Tag",
		},
		{
			ID:              primitive.NewObjectID(),
			Name:            "Furniture",
			MetaDescription: "Tags related to furniture products",
			MetaKeywords:    "furniture, home, decor",
			MetaTitle:       "Furniture Tag",
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedProductTags, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedProductTags, result)
	mockRepo.AssertExpectations(t)
}
