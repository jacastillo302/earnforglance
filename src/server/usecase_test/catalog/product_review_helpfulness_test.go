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

func TestProductReviewHelpfulnessUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.ProductReviewHelpfulnessRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewProductReviewHelpfulnessUsecase(mockRepo, timeout)

	productReviewHelpfulnessID := bson.NewObjectID().Hex()

	updatedProductReviewHelpfulness := domain.ProductReviewHelpfulness{
		ID:              bson.NewObjectID(), // Existing ID of the record to update
		ProductReviewID: bson.NewObjectID(),
		WasHelpful:      false,
		CustomerID:      bson.NewObjectID(),
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
		ProductReviewID: bson.NewObjectID(),
		WasHelpful:      true,
		CustomerID:      bson.NewObjectID(),
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
		ID:              bson.NewObjectID(), // Existing ID of the record to update
		ProductReviewID: bson.NewObjectID(),
		WasHelpful:      false,
		CustomerID:      bson.NewObjectID(),
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

	productReviewHelpfulnessID := bson.NewObjectID().Hex()

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
			ID:              bson.NewObjectID(),
			ProductReviewID: bson.NewObjectID(),
			WasHelpful:      true,
			CustomerID:      bson.NewObjectID(),
		},
		{
			ID:              bson.NewObjectID(),
			ProductReviewID: bson.NewObjectID(),
			WasHelpful:      false,
			CustomerID:      bson.NewObjectID(),
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedProductReviewHelpfulness, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedProductReviewHelpfulness, result)
	mockRepo.AssertExpectations(t)
}
