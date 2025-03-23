package usecase_test

import (
	"context"
	mocks "earnforglance/server/domain/mocks"
	domain "earnforglance/server/domain/shipping"
	test "earnforglance/server/usecase/shipping"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestProductAvailabilityRangeUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.ProductAvailabilityRangeRepository)
	timeout := time.Duration(10)
	usecase := test.NewProductAvailabilityRangeUsecase(mockRepo, timeout)

	productAvailabilityRangeID := primitive.NewObjectID().Hex()

	updatedProductAvailabilityRange := domain.ProductAvailabilityRange{
		ID:           primitive.NewObjectID(), // Existing ID of the record to update
		Name:         "Out of Stock",
		DisplayOrder: 2,
	}

	mockRepo.On("FetchByID", mock.Anything, productAvailabilityRangeID).Return(updatedProductAvailabilityRange, nil)

	result, err := usecase.FetchByID(context.Background(), productAvailabilityRangeID)

	assert.NoError(t, err)
	assert.Equal(t, updatedProductAvailabilityRange, result)
	mockRepo.AssertExpectations(t)
}

func TestProductAvailabilityRangeUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.ProductAvailabilityRangeRepository)
	timeout := time.Duration(10)
	usecase := test.NewProductAvailabilityRangeUsecase(mockRepo, timeout)

	newProductAvailabilityRange := &domain.ProductAvailabilityRange{
		Name:         "In Stock",
		DisplayOrder: 1,
	}

	mockRepo.On("Create", mock.Anything, newProductAvailabilityRange).Return(nil)

	err := usecase.Create(context.Background(), newProductAvailabilityRange)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductAvailabilityRangeUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.ProductAvailabilityRangeRepository)
	timeout := time.Duration(10)
	usecase := test.NewProductAvailabilityRangeUsecase(mockRepo, timeout)

	updatedProductAvailabilityRange := &domain.ProductAvailabilityRange{
		ID:           primitive.NewObjectID(), // Existing ID of the record to update
		Name:         "Out of Stock",
		DisplayOrder: 2,
	}

	mockRepo.On("Update", mock.Anything, updatedProductAvailabilityRange).Return(nil)

	err := usecase.Update(context.Background(), updatedProductAvailabilityRange)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductAvailabilityRangeUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.ProductAvailabilityRangeRepository)
	timeout := time.Duration(10)
	usecase := test.NewProductAvailabilityRangeUsecase(mockRepo, timeout)

	productAvailabilityRangeID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, productAvailabilityRangeID).Return(nil)

	err := usecase.Delete(context.Background(), productAvailabilityRangeID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductAvailabilityRangeUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.ProductAvailabilityRangeRepository)
	timeout := time.Duration(10)
	usecase := test.NewProductAvailabilityRangeUsecase(mockRepo, timeout)

	fetchedProductAvailabilityRanges := []domain.ProductAvailabilityRange{
		{
			ID:           primitive.NewObjectID(),
			Name:         "In Stock",
			DisplayOrder: 1,
		},
		{
			ID:           primitive.NewObjectID(),
			Name:         "Out of Stock",
			DisplayOrder: 2,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedProductAvailabilityRanges, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedProductAvailabilityRanges, result)
	mockRepo.AssertExpectations(t)
}
