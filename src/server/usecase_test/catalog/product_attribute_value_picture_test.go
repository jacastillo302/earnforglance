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

func TestProductAttributeValuePictureUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.ProductAttributeValuePictureRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewProductAttributeValuePictureUsecase(mockRepo, timeout)

	productAttributeValuePictureID := primitive.NewObjectID().Hex()

	expectedProductAttributeValuePicture := domain.ProductAttributeValuePicture{
		ID:                      primitive.NewObjectID(), // Existing ID of the record to update
		ProductAttributeValueID: primitive.NewObjectID(),
		PictureID:               primitive.NewObjectID(),
	}

	mockRepo.On("FetchByID", mock.Anything, productAttributeValuePictureID).Return(expectedProductAttributeValuePicture, nil)

	result, err := usecase.FetchByID(context.Background(), productAttributeValuePictureID)

	assert.NoError(t, err)
	assert.Equal(t, expectedProductAttributeValuePicture, result)
	mockRepo.AssertExpectations(t)
}

func TestProductAttributeValuePictureUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.ProductAttributeValuePictureRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewProductAttributeValuePictureUsecase(mockRepo, timeout)

	newProductAttributeValuePicture := &domain.ProductAttributeValuePicture{
		ProductAttributeValueID: primitive.NewObjectID(),
		PictureID:               primitive.NewObjectID(),
	}

	mockRepo.On("Create", mock.Anything, newProductAttributeValuePicture).Return(nil)

	err := usecase.Create(context.Background(), newProductAttributeValuePicture)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductAttributeValuePictureUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.ProductAttributeValuePictureRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewProductAttributeValuePictureUsecase(mockRepo, timeout)

	updatedProductAttributeValuePicture := &domain.ProductAttributeValuePicture{
		ID:                      primitive.NewObjectID(), // Existing ID of the record to update
		ProductAttributeValueID: primitive.NewObjectID(),
		PictureID:               primitive.NewObjectID(),
	}

	mockRepo.On("Update", mock.Anything, updatedProductAttributeValuePicture).Return(nil)

	err := usecase.Update(context.Background(), updatedProductAttributeValuePicture)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductAttributeValuePictureUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.ProductAttributeValuePictureRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewProductAttributeValuePictureUsecase(mockRepo, timeout)

	productAttributeValuePictureID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, productAttributeValuePictureID).Return(nil)

	err := usecase.Delete(context.Background(), productAttributeValuePictureID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductAttributeValuePictureUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.ProductAttributeValuePictureRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewProductAttributeValuePictureUsecase(mockRepo, timeout)

	fetchedProductAttributeValuePictures := []domain.ProductAttributeValuePicture{
		{
			ID:                      primitive.NewObjectID(),
			ProductAttributeValueID: primitive.NewObjectID(),
			PictureID:               primitive.NewObjectID(),
		},
		{
			ID:                      primitive.NewObjectID(),
			ProductAttributeValueID: primitive.NewObjectID(),
			PictureID:               primitive.NewObjectID(),
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedProductAttributeValuePictures, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedProductAttributeValuePictures, result)
	mockRepo.AssertExpectations(t)
}
