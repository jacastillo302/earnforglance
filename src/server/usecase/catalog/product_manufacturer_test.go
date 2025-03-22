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

func TestProductManufacturerUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.ProductManufacturerRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewProductManufacturerUsecase(mockRepo, timeout)

	productManufacturerID := primitive.NewObjectID().Hex()

	updatedProductManufacturer := domain.ProductManufacturer{
		ID:                primitive.NewObjectID(), // Existing ID of the record to update
		ProductID:         primitive.NewObjectID(),
		ManufacturerID:    primitive.NewObjectID(),
		IsFeaturedProduct: false,
		DisplayOrder:      2,
	}

	mockRepo.On("FetchByID", mock.Anything, productManufacturerID).Return(updatedProductManufacturer, nil)

	result, err := usecase.FetchByID(context.Background(), productManufacturerID)

	assert.NoError(t, err)
	assert.Equal(t, updatedProductManufacturer, result)
	mockRepo.AssertExpectations(t)
}

func TestProductManufacturerUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.ProductManufacturerRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewProductManufacturerUsecase(mockRepo, timeout)

	newProductManufacturer := &domain.ProductManufacturer{
		ProductID:         primitive.NewObjectID(),
		ManufacturerID:    primitive.NewObjectID(),
		IsFeaturedProduct: true,
		DisplayOrder:      1,
	}

	mockRepo.On("Create", mock.Anything, newProductManufacturer).Return(nil)

	err := usecase.Create(context.Background(), newProductManufacturer)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductManufacturerUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.ProductManufacturerRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewProductManufacturerUsecase(mockRepo, timeout)

	updatedProductManufacturer := &domain.ProductManufacturer{
		ID:                primitive.NewObjectID(), // Existing ID of the record to update
		ProductID:         primitive.NewObjectID(),
		ManufacturerID:    primitive.NewObjectID(),
		IsFeaturedProduct: false,
		DisplayOrder:      2,
	}

	mockRepo.On("Update", mock.Anything, updatedProductManufacturer).Return(nil)

	err := usecase.Update(context.Background(), updatedProductManufacturer)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductManufacturerUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.ProductManufacturerRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewProductManufacturerUsecase(mockRepo, timeout)

	productManufacturerID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, productManufacturerID).Return(nil)

	err := usecase.Delete(context.Background(), productManufacturerID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductManufacturerUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.ProductManufacturerRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewProductManufacturerUsecase(mockRepo, timeout)

	fetchedProductManufacturers := []domain.ProductManufacturer{
		{
			ID:                primitive.NewObjectID(),
			ProductID:         primitive.NewObjectID(),
			ManufacturerID:    primitive.NewObjectID(),
			IsFeaturedProduct: true,
			DisplayOrder:      1,
		},
		{
			ID:                primitive.NewObjectID(),
			ProductID:         primitive.NewObjectID(),
			ManufacturerID:    primitive.NewObjectID(),
			IsFeaturedProduct: false,
			DisplayOrder:      2,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedProductManufacturers, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedProductManufacturers, result)
	mockRepo.AssertExpectations(t)
}
