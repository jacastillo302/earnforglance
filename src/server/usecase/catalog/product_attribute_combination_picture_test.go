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

func TestProductAttributeCombinationPictureUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.ProductAttributeCombinationPictureRepository)
	timeout := time.Duration(10)
	usecase := NewProductAttributeCombinationPictureUsecase(mockRepo, timeout)

	productAttributeCombinationPictureID := primitive.NewObjectID().Hex()

	expectedProductAttributeCombinationPicture := domain.ProductAttributeCombinationPicture{
		ID:                            primitive.NewObjectID(),
		ProductAttributeCombinationID: primitive.NewObjectID(),
		PictureID:                     primitive.NewObjectID(),
	}

	mockRepo.On("FetchByID", mock.Anything, productAttributeCombinationPictureID).Return(expectedProductAttributeCombinationPicture, nil)

	result, err := usecase.FetchByID(context.Background(), productAttributeCombinationPictureID)

	assert.NoError(t, err)
	assert.Equal(t, expectedProductAttributeCombinationPicture, result)
	mockRepo.AssertExpectations(t)
}

func TestProductAttributeCombinationPictureUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.ProductAttributeCombinationPictureRepository)
	timeout := time.Duration(10)
	usecase := NewProductAttributeCombinationPictureUsecase(mockRepo, timeout)

	newProductAttributeCombinationPicture := &domain.ProductAttributeCombinationPicture{
		ProductAttributeCombinationID: primitive.NewObjectID(),
		PictureID:                     primitive.NewObjectID(),
	}

	mockRepo.On("Create", mock.Anything, newProductAttributeCombinationPicture).Return(nil)

	err := usecase.Create(context.Background(), newProductAttributeCombinationPicture)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductAttributeCombinationPictureUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.ProductAttributeCombinationPictureRepository)
	timeout := time.Duration(10)
	usecase := NewProductAttributeCombinationPictureUsecase(mockRepo, timeout)

	updatedProductAttributeCombinationPicture := &domain.ProductAttributeCombinationPicture{
		ID:                            primitive.NewObjectID(),
		ProductAttributeCombinationID: primitive.NewObjectID(),
		PictureID:                     primitive.NewObjectID(),
	}

	mockRepo.On("Update", mock.Anything, updatedProductAttributeCombinationPicture).Return(nil)

	err := usecase.Update(context.Background(), updatedProductAttributeCombinationPicture)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductAttributeCombinationPictureUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.ProductAttributeCombinationPictureRepository)
	timeout := time.Duration(10)
	usecase := NewProductAttributeCombinationPictureUsecase(mockRepo, timeout)

	productAttributeCombinationPictureID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, productAttributeCombinationPictureID).Return(nil)

	err := usecase.Delete(context.Background(), productAttributeCombinationPictureID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductAttributeCombinationPictureUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.ProductAttributeCombinationPictureRepository)
	timeout := time.Duration(10)
	usecase := NewProductAttributeCombinationPictureUsecase(mockRepo, timeout)

	expectedProductAttributeCombinationPictures := []domain.ProductAttributeCombinationPicture{
		{
			ID:                            primitive.NewObjectID(),
			ProductAttributeCombinationID: primitive.NewObjectID(),
			PictureID:                     primitive.NewObjectID(),
		},
		{
			ID:                            primitive.NewObjectID(),
			ProductAttributeCombinationID: primitive.NewObjectID(),
			PictureID:                     primitive.NewObjectID(),
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(expectedProductAttributeCombinationPictures, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, expectedProductAttributeCombinationPictures, result)
	mockRepo.AssertExpectations(t)
}
