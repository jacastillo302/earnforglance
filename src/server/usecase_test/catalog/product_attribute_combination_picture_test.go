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

func TestProductAttributeCombinationPictureUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.ProductAttributeCombinationPictureRepository)
	timeout := time.Duration(10)
	usecase := test.NewProductAttributeCombinationPictureUsecase(mockRepo, timeout)

	productAttributeCombinationPictureID := bson.NewObjectID().Hex()

	expectedProductAttributeCombinationPicture := domain.ProductAttributeCombinationPicture{
		ID:                            bson.NewObjectID(),
		ProductAttributeCombinationID: bson.NewObjectID(),
		PictureID:                     bson.NewObjectID(),
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
	usecase := test.NewProductAttributeCombinationPictureUsecase(mockRepo, timeout)

	newProductAttributeCombinationPicture := &domain.ProductAttributeCombinationPicture{
		ProductAttributeCombinationID: bson.NewObjectID(),
		PictureID:                     bson.NewObjectID(),
	}

	mockRepo.On("Create", mock.Anything, newProductAttributeCombinationPicture).Return(nil)

	err := usecase.Create(context.Background(), newProductAttributeCombinationPicture)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductAttributeCombinationPictureUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.ProductAttributeCombinationPictureRepository)
	timeout := time.Duration(10)
	usecase := test.NewProductAttributeCombinationPictureUsecase(mockRepo, timeout)

	updatedProductAttributeCombinationPicture := &domain.ProductAttributeCombinationPicture{
		ID:                            bson.NewObjectID(),
		ProductAttributeCombinationID: bson.NewObjectID(),
		PictureID:                     bson.NewObjectID(),
	}

	mockRepo.On("Update", mock.Anything, updatedProductAttributeCombinationPicture).Return(nil)

	err := usecase.Update(context.Background(), updatedProductAttributeCombinationPicture)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductAttributeCombinationPictureUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.ProductAttributeCombinationPictureRepository)
	timeout := time.Duration(10)
	usecase := test.NewProductAttributeCombinationPictureUsecase(mockRepo, timeout)

	productAttributeCombinationPictureID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, productAttributeCombinationPictureID).Return(nil)

	err := usecase.Delete(context.Background(), productAttributeCombinationPictureID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductAttributeCombinationPictureUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.ProductAttributeCombinationPictureRepository)
	timeout := time.Duration(10)
	usecase := test.NewProductAttributeCombinationPictureUsecase(mockRepo, timeout)

	expectedProductAttributeCombinationPictures := []domain.ProductAttributeCombinationPicture{
		{
			ID:                            bson.NewObjectID(),
			ProductAttributeCombinationID: bson.NewObjectID(),
			PictureID:                     bson.NewObjectID(),
		},
		{
			ID:                            bson.NewObjectID(),
			ProductAttributeCombinationID: bson.NewObjectID(),
			PictureID:                     bson.NewObjectID(),
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(expectedProductAttributeCombinationPictures, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, expectedProductAttributeCombinationPictures, result)
	mockRepo.AssertExpectations(t)
}
