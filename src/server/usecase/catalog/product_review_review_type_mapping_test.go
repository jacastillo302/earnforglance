package usecase

import (
	"context"
	"testing"
	"time"

	domain "earnforglance/server/domain/catalog"
	mocks "earnforglance/server/domain/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestProductReviewReviewTypeMappingUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.ProductReviewReviewTypeMappingRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewProductReviewReviewTypeMappingUsecase(mockRepo, timeout)

	mappingID := primitive.NewObjectID().Hex()

	updatedProductReviewReviewTypeMapping := domain.ProductReviewReviewTypeMapping{
		ID:              primitive.NewObjectID(), // Existing ID of the record to update
		ProductReviewID: primitive.NewObjectID(),
		ReviewTypeID:    primitive.NewObjectID(),
		Rating:          4,
	}

	mockRepo.On("FetchByID", mock.Anything, mappingID).Return(updatedProductReviewReviewTypeMapping, nil)

	result, err := usecase.FetchByID(context.Background(), mappingID)

	assert.NoError(t, err)
	assert.Equal(t, updatedProductReviewReviewTypeMapping, result)
	mockRepo.AssertExpectations(t)
}

func TestProductReviewReviewTypeMappingUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.ProductReviewReviewTypeMappingRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewProductReviewReviewTypeMappingUsecase(mockRepo, timeout)

	newProductReviewReviewTypeMapping := &domain.ProductReviewReviewTypeMapping{
		ProductReviewID: primitive.NewObjectID(),
		ReviewTypeID:    primitive.NewObjectID(),
		Rating:          5,
	}

	mockRepo.On("Create", mock.Anything, newProductReviewReviewTypeMapping).Return(nil)

	err := usecase.Create(context.Background(), newProductReviewReviewTypeMapping)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductReviewReviewTypeMappingUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.ProductReviewReviewTypeMappingRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewProductReviewReviewTypeMappingUsecase(mockRepo, timeout)

	updatedProductReviewReviewTypeMapping := &domain.ProductReviewReviewTypeMapping{
		ID:              primitive.NewObjectID(), // Existing ID of the record to update
		ProductReviewID: primitive.NewObjectID(),
		ReviewTypeID:    primitive.NewObjectID(),
		Rating:          4,
	}

	mockRepo.On("Update", mock.Anything, updatedProductReviewReviewTypeMapping).Return(nil)

	err := usecase.Update(context.Background(), updatedProductReviewReviewTypeMapping)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductReviewReviewTypeMappingUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.ProductReviewReviewTypeMappingRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewProductReviewReviewTypeMappingUsecase(mockRepo, timeout)

	mappingID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, mappingID).Return(nil)

	err := usecase.Delete(context.Background(), mappingID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductReviewReviewTypeMappingUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.ProductReviewReviewTypeMappingRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewProductReviewReviewTypeMappingUsecase(mockRepo, timeout)

	fetchedProductReviewReviewTypeMappings := []domain.ProductReviewReviewTypeMapping{
		{
			ID:              primitive.NewObjectID(),
			ProductReviewID: primitive.NewObjectID(),
			ReviewTypeID:    primitive.NewObjectID(),
			Rating:          5,
		},
		{
			ID:              primitive.NewObjectID(),
			ProductReviewID: primitive.NewObjectID(),
			ReviewTypeID:    primitive.NewObjectID(),
			Rating:          3,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedProductReviewReviewTypeMappings, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedProductReviewReviewTypeMappings, result)
	mockRepo.AssertExpectations(t)
}
