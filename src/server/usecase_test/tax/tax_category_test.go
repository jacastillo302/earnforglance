package usecase_test

import (
	"context"
	mocks "earnforglance/server/domain/mocks"
	domain "earnforglance/server/domain/tax"
	test "earnforglance/server/usecase/tax"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestTaxCategoryUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.TaxCategoryRepository)
	timeout := time.Duration(10)
	usecase := test.NewTaxCategoryUsecase(mockRepo, timeout)

	taxCategoryID := primitive.NewObjectID().Hex()

	updatedTaxCategory := domain.TaxCategory{
		ID:           primitive.NewObjectID(), // Existing ID of the record to update
		Name:         "Reduced Rate",
		DisplayOrder: 2,
	}

	mockRepo.On("FetchByID", mock.Anything, taxCategoryID).Return(updatedTaxCategory, nil)

	result, err := usecase.FetchByID(context.Background(), taxCategoryID)

	assert.NoError(t, err)
	assert.Equal(t, updatedTaxCategory, result)
	mockRepo.AssertExpectations(t)
}

func TestTaxCategoryUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.TaxCategoryRepository)
	timeout := time.Duration(10)
	usecase := test.NewTaxCategoryUsecase(mockRepo, timeout)

	newTaxCategory := &domain.TaxCategory{
		Name:         "Standard Rate",
		DisplayOrder: 1,
	}

	mockRepo.On("Create", mock.Anything, newTaxCategory).Return(nil)

	err := usecase.Create(context.Background(), newTaxCategory)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestTaxCategoryUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.TaxCategoryRepository)
	timeout := time.Duration(10)
	usecase := test.NewTaxCategoryUsecase(mockRepo, timeout)

	updatedTaxCategory := &domain.TaxCategory{
		ID:           primitive.NewObjectID(), // Existing ID of the record to update
		Name:         "Reduced Rate",
		DisplayOrder: 2,
	}

	mockRepo.On("Update", mock.Anything, updatedTaxCategory).Return(nil)

	err := usecase.Update(context.Background(), updatedTaxCategory)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestTaxCategoryUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.TaxCategoryRepository)
	timeout := time.Duration(10)
	usecase := test.NewTaxCategoryUsecase(mockRepo, timeout)

	taxCategoryID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, taxCategoryID).Return(nil)

	err := usecase.Delete(context.Background(), taxCategoryID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestTaxCategoryUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.TaxCategoryRepository)
	timeout := time.Duration(10)
	usecase := test.NewTaxCategoryUsecase(mockRepo, timeout)

	fetchedTaxCategories := []domain.TaxCategory{
		{
			ID:           primitive.NewObjectID(),
			Name:         "Standard Rate",
			DisplayOrder: 1,
		},
		{
			ID:           primitive.NewObjectID(),
			Name:         "Reduced Rate",
			DisplayOrder: 2,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedTaxCategories, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedTaxCategories, result)
	mockRepo.AssertExpectations(t)
}
