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

func TestProductReviewUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.ProductReviewRepository)
	timeout := time.Duration(10)
	usecase := test.NewProductReviewUsecase(mockRepo, timeout)

	productReviewID := bson.NewObjectID().Hex()

	updatedProductReview := domain.ProductReview{
		ID:                      bson.NewObjectID(), // Existing ID of the record to update
		CustomerID:              bson.NewObjectID(),
		ProductID:               bson.NewObjectID(),
		StoreID:                 bson.NewObjectID(),
		IsApproved:              false,
		Title:                   "Updated Review Title",
		ReviewText:              "Updated review text with more details.",
		ReplyText:               "Thank you for your feedback!",
		CustomerNotifiedOfReply: true,
		Rating:                  4,
		HelpfulYesTotal:         15,
		HelpfulNoTotal:          3,
		CreatedOnUtc:            time.Now().AddDate(0, 0, -7), // Created 7 days ago
	}

	mockRepo.On("FetchByID", mock.Anything, productReviewID).Return(updatedProductReview, nil)

	result, err := usecase.FetchByID(context.Background(), productReviewID)

	assert.NoError(t, err)
	assert.Equal(t, updatedProductReview, result)
	mockRepo.AssertExpectations(t)
}

func TestProductReviewUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.ProductReviewRepository)
	timeout := time.Duration(10)
	usecase := test.NewProductReviewUsecase(mockRepo, timeout)

	newProductReview := &domain.ProductReview{
		CustomerID:              bson.NewObjectID(),
		ProductID:               bson.NewObjectID(),
		StoreID:                 bson.NewObjectID(),
		IsApproved:              true,
		Title:                   "Great Product!",
		ReviewText:              "This product exceeded my expectations.",
		ReplyText:               "",
		CustomerNotifiedOfReply: false,
		Rating:                  5,
		HelpfulYesTotal:         10,
		HelpfulNoTotal:          2,
		CreatedOnUtc:            time.Now(),
	}

	mockRepo.On("Create", mock.Anything, newProductReview).Return(nil)

	err := usecase.Create(context.Background(), newProductReview)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductReviewUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.ProductReviewRepository)
	timeout := time.Duration(10)
	usecase := test.NewProductReviewUsecase(mockRepo, timeout)

	updatedProductReview := &domain.ProductReview{
		ID:                      bson.NewObjectID(), // Existing ID of the record to update
		CustomerID:              bson.NewObjectID(),
		ProductID:               bson.NewObjectID(),
		StoreID:                 bson.NewObjectID(),
		IsApproved:              false,
		Title:                   "Updated Review Title",
		ReviewText:              "Updated review text with more details.",
		ReplyText:               "Thank you for your feedback!",
		CustomerNotifiedOfReply: true,
		Rating:                  4,
		HelpfulYesTotal:         15,
		HelpfulNoTotal:          3,
		CreatedOnUtc:            time.Now().AddDate(0, 0, -7), // Created 7 days ago
	}

	mockRepo.On("Update", mock.Anything, updatedProductReview).Return(nil)

	err := usecase.Update(context.Background(), updatedProductReview)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductReviewUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.ProductReviewRepository)
	timeout := time.Duration(10)
	usecase := test.NewProductReviewUsecase(mockRepo, timeout)

	productReviewID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, productReviewID).Return(nil)

	err := usecase.Delete(context.Background(), productReviewID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductReviewUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.ProductReviewRepository)
	timeout := time.Duration(10)
	usecase := test.NewProductReviewUsecase(mockRepo, timeout)

	fetchedProductReviews := []domain.ProductReview{
		{
			ID:                      bson.NewObjectID(),
			CustomerID:              bson.NewObjectID(),
			ProductID:               bson.NewObjectID(),
			StoreID:                 bson.NewObjectID(),
			IsApproved:              true,
			Title:                   "Excellent Product",
			ReviewText:              "I absolutely love this product!",
			ReplyText:               "",
			CustomerNotifiedOfReply: false,
			Rating:                  5,
			HelpfulYesTotal:         20,
			HelpfulNoTotal:          1,
			CreatedOnUtc:            time.Now().AddDate(0, 0, -10), // Created 10 days ago
		},
		{
			ID:                      bson.NewObjectID(),
			CustomerID:              bson.NewObjectID(),
			ProductID:               bson.NewObjectID(),
			StoreID:                 bson.NewObjectID(),
			IsApproved:              false,
			Title:                   "Not Satisfied",
			ReviewText:              "The product did not meet my expectations.",
			ReplyText:               "We are sorry to hear that. Please contact support.",
			CustomerNotifiedOfReply: true,
			Rating:                  2,
			HelpfulYesTotal:         5,
			HelpfulNoTotal:          8,
			CreatedOnUtc:            time.Now().AddDate(0, 0, -20), // Created 20 days ago
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedProductReviews, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedProductReviews, result)
	mockRepo.AssertExpectations(t)
}
