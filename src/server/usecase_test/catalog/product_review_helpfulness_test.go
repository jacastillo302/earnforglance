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

func TestProductReviewHelpfulnessUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.ProductReviewHelpfulnessRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewProductReviewHelpfulnessUsecase(mockRepo, timeout)

	productReviewHelpfulnessID := primitive.NewObjectID().Hex()

	updatedProductReviewHelpfulness := domain.ProductReviewHelpfulness{
		ID:              primitive.NewObjectID(), // Existing ID of the record to update
		ProductReviewID: primitive.NewObjectID(),
		WasHelpful:      false,
		CustomerID:      primitive.NewObjectID(),
	}

	mockRepo.On("FetchByID", mock.Anything, productReviewHelpfulnessID).Return(updatedProductReviewHelpfulness, nil)

	result, err := usecase.FetchByID(context.Background(), productReviewHelpfulnessID)

	assert.NoError(t, err)
	assert.Equal(t, updatedProductReviewHelpfulness, result)
	mockRepo.AssertExpectations(t)
}

func TestProductReviewHelpfulnessUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.ProductReviewHelpfulnessRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewProductReviewHelpfulnessUsecase(mockRepo, timeout)

	newProductReviewHelpfulness := &domain.ProductReviewHelpfulness{
		ProductReviewID: primitive.NewObjectID(),
		WasHelpful:      true,
		CustomerID:      primitive.NewObjectID(),
	}

	mockRepo.On("Create", mock.Anything, newProductReviewHelpfulness).Return(nil)

	err := usecase.Create(context.Background(), newProductReviewHelpfulness)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductReviewHelpfulnessUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.ProductReviewHelpfulnessRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewProductReviewHelpfulnessUsecase(mockRepo, timeout)

	updatedProductReviewHelpfulness := &domain.ProductReviewHelpfulness{
		ID:              primitive.NewObjectID(), // Existing ID of the record to update
		ProductReviewID: primitive.NewObjectID(),
		WasHelpful:      false,
		CustomerID:      primitive.NewObjectID(),
	}

	mockRepo.On("Update", mock.Anything, updatedProductReviewHelpfulness).Return(nil)

	err := usecase.Update(context.Background(), updatedProductReviewHelpfulness)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductReviewHelpfulnessUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.ProductReviewHelpfulnessRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewProductReviewHelpfulnessUsecase(mockRepo, timeout)

	productReviewHelpfulnessID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, productReviewHelpfulnessID).Return(nil)

	err := usecase.Delete(context.Background(), productReviewHelpfulnessID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductReviewHelpfulnessUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.ProductReviewHelpfulnessRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewProductReviewHelpfulnessUsecase(mockRepo, timeout)

	fetchedProductReviewHelpfulness := []domain.ProductReviewHelpfulness{
		{
			ID:              primitive.NewObjectID(),
			ProductReviewID: primitive.NewObjectID(),
			WasHelpful:      true,
			CustomerID:      primitive.NewObjectID(),
		},
		{
			ID:              primitive.NewObjectID(),
			ProductReviewID: primitive.NewObjectID(),
			WasHelpful:      false,
			CustomerID:      primitive.NewObjectID(),
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedProductReviewHelpfulness, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedProductReviewHelpfulness, result)
	mockRepo.AssertExpectations(t)
}
