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

func TestProductCategoryUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.ProductCategoryRepository)
	timeout := time.Duration(10)
	usecase := test.NewProductCategoryUsecase(mockRepo, timeout)

	productCategoryID := primitive.NewObjectID().Hex()

	updatedProductCategory := domain.ProductCategory{
		ID:                primitive.NewObjectID(),
		ProductID:         primitive.NewObjectID(),
		CategoryID:        primitive.NewObjectID(),
		IsFeaturedProduct: false,
		DisplayOrder:      2,
	}

	mockRepo.On("FetchByID", mock.Anything, productCategoryID).Return(updatedProductCategory, nil)

	result, err := usecase.FetchByID(context.Background(), productCategoryID)

	assert.NoError(t, err)
	assert.Equal(t, updatedProductCategory, result)
	mockRepo.AssertExpectations(t)
}

func TestProductCategoryUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.ProductCategoryRepository)
	timeout := time.Duration(10)
	usecase := test.NewProductCategoryUsecase(mockRepo, timeout)

	newProductCategory := &domain.ProductCategory{
		ProductID:         primitive.NewObjectID(),
		CategoryID:        primitive.NewObjectID(),
		IsFeaturedProduct: true,
		DisplayOrder:      1,
	}

	mockRepo.On("Create", mock.Anything, newProductCategory).Return(nil)

	err := usecase.Create(context.Background(), newProductCategory)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductCategoryUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.ProductCategoryRepository)
	timeout := time.Duration(10)
	usecase := test.NewProductCategoryUsecase(mockRepo, timeout)

	updatedProductCategory := &domain.ProductCategory{
		ID:                primitive.NewObjectID(),
		ProductID:         primitive.NewObjectID(),
		CategoryID:        primitive.NewObjectID(),
		IsFeaturedProduct: false,
		DisplayOrder:      2,
	}

	mockRepo.On("Update", mock.Anything, updatedProductCategory).Return(nil)

	err := usecase.Update(context.Background(), updatedProductCategory)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductCategoryUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.ProductCategoryRepository)
	timeout := time.Duration(10)
	usecase := test.NewProductCategoryUsecase(mockRepo, timeout)

	productCategoryID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, productCategoryID).Return(nil)

	err := usecase.Delete(context.Background(), productCategoryID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductCategoryUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.ProductCategoryRepository)
	timeout := time.Duration(10)
	usecase := test.NewProductCategoryUsecase(mockRepo, timeout)

	fetchedProductCategories := []domain.ProductCategory{
		{
			ID:                primitive.NewObjectID(),
			ProductID:         primitive.NewObjectID(),
			CategoryID:        primitive.NewObjectID(),
			IsFeaturedProduct: true,
			DisplayOrder:      1,
		},
		{
			ID:                primitive.NewObjectID(),
			ProductID:         primitive.NewObjectID(),
			CategoryID:        primitive.NewObjectID(),
			IsFeaturedProduct: false,
			DisplayOrder:      2,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedProductCategories, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedProductCategories, result)
	mockRepo.AssertExpectations(t)
}
