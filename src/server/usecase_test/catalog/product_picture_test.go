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

func TestProductPictureUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.ProductPictureRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewProductPictureUsecase(mockRepo, timeout)

	productPictureID := bson.NewObjectID().Hex()

	updatedProductPicture := domain.ProductPicture{
		ID:           bson.NewObjectID(), // Existing ID of the record to update
		ProductID:    bson.NewObjectID(),
		PictureID:    bson.NewObjectID(),
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
	usecase := test.NewProductPictureUsecase(mockRepo, timeout)

	newProductPicture := &domain.ProductPicture{
		ProductID:    bson.NewObjectID(),
		PictureID:    bson.NewObjectID(),
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
	usecase := test.NewProductPictureUsecase(mockRepo, timeout)

	updatedProductPicture := &domain.ProductPicture{
		ID:           bson.NewObjectID(), // Existing ID of the record to update
		ProductID:    bson.NewObjectID(),
		PictureID:    bson.NewObjectID(),
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
	usecase := test.NewProductPictureUsecase(mockRepo, timeout)

	productPictureID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, productPictureID).Return(nil)

	err := usecase.Delete(context.Background(), productPictureID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductPictureUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.ProductPictureRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewProductPictureUsecase(mockRepo, timeout)

	fetchedProductPictures := []domain.ProductPicture{
		{
			ID:           bson.NewObjectID(),
			ProductID:    bson.NewObjectID(),
			PictureID:    bson.NewObjectID(),
			DisplayOrder: 1,
		},
		{
			ID:           bson.NewObjectID(),
			ProductID:    bson.NewObjectID(),
			PictureID:    bson.NewObjectID(),
			DisplayOrder: 2,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedProductPictures, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedProductPictures, result)
	mockRepo.AssertExpectations(t)
}
