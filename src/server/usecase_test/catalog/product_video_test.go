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

func TestProductVideoUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.ProductVideoRepository)
	timeout := time.Duration(10)
	usecase := test.NewProductVideoUsecase(mockRepo, timeout)

	productVideoID := primitive.NewObjectID().Hex()

	updatedProductVideo := domain.ProductVideo{
		ID:           primitive.NewObjectID(), // Existing ID of the record to update
		ProductID:    primitive.NewObjectID(),
		VideoID:      primitive.NewObjectID(),
		DisplayOrder: 2,
	}

	mockRepo.On("FetchByID", mock.Anything, productVideoID).Return(updatedProductVideo, nil)

	result, err := usecase.FetchByID(context.Background(), productVideoID)

	assert.NoError(t, err)
	assert.Equal(t, updatedProductVideo, result)
	mockRepo.AssertExpectations(t)
}

func TestProductVideoUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.ProductVideoRepository)
	timeout := time.Duration(10)
	usecase := test.NewProductVideoUsecase(mockRepo, timeout)

	newProductVideo := &domain.ProductVideo{
		ProductID:    primitive.NewObjectID(),
		VideoID:      primitive.NewObjectID(),
		DisplayOrder: 1,
	}

	mockRepo.On("Create", mock.Anything, newProductVideo).Return(nil)

	err := usecase.Create(context.Background(), newProductVideo)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductVideoUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.ProductVideoRepository)
	timeout := time.Duration(10)
	usecase := test.NewProductVideoUsecase(mockRepo, timeout)

	updatedProductVideo := &domain.ProductVideo{
		ID:           primitive.NewObjectID(), // Existing ID of the record to update
		ProductID:    primitive.NewObjectID(),
		VideoID:      primitive.NewObjectID(),
		DisplayOrder: 2,
	}

	mockRepo.On("Update", mock.Anything, updatedProductVideo).Return(nil)

	err := usecase.Update(context.Background(), updatedProductVideo)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductVideoUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.ProductVideoRepository)
	timeout := time.Duration(10)
	usecase := test.NewProductVideoUsecase(mockRepo, timeout)

	productVideoID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, productVideoID).Return(nil)

	err := usecase.Delete(context.Background(), productVideoID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductVideoUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.ProductVideoRepository)
	timeout := time.Duration(10)
	usecase := test.NewProductVideoUsecase(mockRepo, timeout)

	fetchedProductVideos := []domain.ProductVideo{
		{
			ID:           primitive.NewObjectID(),
			ProductID:    primitive.NewObjectID(),
			VideoID:      primitive.NewObjectID(),
			DisplayOrder: 1,
		},
		{
			ID:           primitive.NewObjectID(),
			ProductID:    primitive.NewObjectID(),
			VideoID:      primitive.NewObjectID(),
			DisplayOrder: 2,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedProductVideos, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedProductVideos, result)
	mockRepo.AssertExpectations(t)
}
