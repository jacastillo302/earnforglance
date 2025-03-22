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

func TestProductPictureUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.ProductPictureRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewProductPictureUsecase(mockRepo, timeout)

	productPictureID := primitive.NewObjectID().Hex()

	updatedProductPicture := domain.ProductPicture{
		ID:           primitive.NewObjectID(), // Existing ID of the record to update
		ProductID:    primitive.NewObjectID(),
		PictureID:    102,
		DisplayOrder: 2,
	}

	mockRepo.On("FetchByID", mock.Anything, productPictureID).Return(updatedProductPicture, nil)

	result, err := usecase.FetchByID(context.Background(), productPictureID)

	assert.NoError(t, err)
	assert.Equal(t, updatedProductPicture, result)
	mockRepo.AssertExpectations(t)
}

func TestProductPictureUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.ProductPictureRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewProductPictureUsecase(mockRepo, timeout)

	newProductPicture := &domain.ProductPicture{
		ProductID:    primitive.NewObjectID(),
		PictureID:    101,
		DisplayOrder: 1,
	}

	mockRepo.On("Create", mock.Anything, newProductPicture).Return(nil)

	err := usecase.Create(context.Background(), newProductPicture)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductPictureUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.ProductPictureRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewProductPictureUsecase(mockRepo, timeout)

	updatedProductPicture := &domain.ProductPicture{
		ID:           primitive.NewObjectID(), // Existing ID of the record to update
		ProductID:    primitive.NewObjectID(),
		PictureID:    102,
		DisplayOrder: 2,
	}
	mockRepo.On("Update", mock.Anything, updatedProductPicture).Return(nil)

	err := usecase.Update(context.Background(), updatedProductPicture)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductPictureUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.ProductPictureRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewProductPictureUsecase(mockRepo, timeout)

	productPictureID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, productPictureID).Return(nil)

	err := usecase.Delete(context.Background(), productPictureID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductPictureUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.ProductPictureRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewProductPictureUsecase(mockRepo, timeout)

	fetchedProductPictures := []domain.ProductPicture{
		{
			ID:           primitive.NewObjectID(),
			ProductID:    primitive.NewObjectID(),
			PictureID:    101,
			DisplayOrder: 1,
		},
		{
			ID:           primitive.NewObjectID(),
			ProductID:    primitive.NewObjectID(),
			PictureID:    102,
			DisplayOrder: 2,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedProductPictures, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedProductPictures, result)
	mockRepo.AssertExpectations(t)
}
