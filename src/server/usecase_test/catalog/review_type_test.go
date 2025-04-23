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

func TestReviewTypeUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.ReviewTypeRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewReviewTypeUsecase(mockRepo, timeout)

	reviewTypeID := bson.NewObjectID().Hex()

	updatedReviewType := domain.ReviewType{
		ID:                    bson.NewObjectID(), // Existing ID of the record to update
		Name:                  "Durability",
		Description:           "Review the durability of the product.",
		DisplayOrder:          2,
		VisibleToAllCustomers: false,
		IsRequired:            false,
	}

	mockRepo.On("FetchByID", mock.Anything, reviewTypeID).Return(updatedReviewType, nil)

	result, err := usecase.FetchByID(context.Background(), reviewTypeID)

	assert.NoError(t, err)
	assert.Equal(t, updatedReviewType, result)
	mockRepo.AssertExpectations(t)
}

func TestReviewTypeUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.ReviewTypeRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewReviewTypeUsecase(mockRepo, timeout)

	newReviewType := &domain.ReviewType{
		Name:                  "Quality",
		Description:           "Review the quality of the product.",
		DisplayOrder:          1,
		VisibleToAllCustomers: true,
		IsRequired:            true,
	}

	mockRepo.On("Create", mock.Anything, newReviewType).Return(nil)

	err := usecase.Create(context.Background(), newReviewType)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestReviewTypeUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.ReviewTypeRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewReviewTypeUsecase(mockRepo, timeout)

	updatedReviewType := &domain.ReviewType{
		ID:                    bson.NewObjectID(), // Existing ID of the record to update
		Name:                  "Durability",
		Description:           "Review the durability of the product.",
		DisplayOrder:          2,
		VisibleToAllCustomers: false,
		IsRequired:            false,
	}

	mockRepo.On("Update", mock.Anything, updatedReviewType).Return(nil)

	err := usecase.Update(context.Background(), updatedReviewType)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestReviewTypeUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.ReviewTypeRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewReviewTypeUsecase(mockRepo, timeout)

	reviewTypeID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, reviewTypeID).Return(nil)

	err := usecase.Delete(context.Background(), reviewTypeID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestReviewTypeUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.ReviewTypeRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewReviewTypeUsecase(mockRepo, timeout)

	fetchedReviewTypes := []domain.ReviewType{
		{
			ID:                    bson.NewObjectID(),
			Name:                  "Quality",
			Description:           "Review the quality of the product.",
			DisplayOrder:          1,
			VisibleToAllCustomers: true,
			IsRequired:            true,
		},
		{
			ID:                    bson.NewObjectID(),
			Name:                  "Durability",
			Description:           "Review the durability of the product.",
			DisplayOrder:          2,
			VisibleToAllCustomers: false,
			IsRequired:            false,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedReviewTypes, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedReviewTypes, result)
	mockRepo.AssertExpectations(t)
}
