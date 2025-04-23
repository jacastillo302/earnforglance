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

func TestRelatedProductUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.RelatedProductRepository)
	timeout := time.Duration(10)
	usecase := test.NewRelatedProductUsecase(mockRepo, timeout)

	catalogID := bson.NewObjectID().Hex()

	updatedRelatedProduct := domain.RelatedProduct{
		ID:           bson.NewObjectID(), // Existing ID of the record to update
		ProductID1:   bson.NewObjectID(),
		ProductID2:   bson.NewObjectID(),
		DisplayOrder: 2,
	}

	mockRepo.On("FetchByID", mock.Anything, catalogID).Return(updatedRelatedProduct, nil)

	result, err := usecase.FetchByID(context.Background(), catalogID)

	assert.NoError(t, err)
	assert.Equal(t, updatedRelatedProduct, result)
	mockRepo.AssertExpectations(t)
}

func TestRelatedProductUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.RelatedProductRepository)
	timeout := time.Duration(10)
	usecase := test.NewRelatedProductUsecase(mockRepo, timeout)

	newRelatedProduct := &domain.RelatedProduct{
		ProductID1:   bson.NewObjectID(),
		ProductID2:   bson.NewObjectID(),
		DisplayOrder: 1,
	}

	mockRepo.On("Create", mock.Anything, newRelatedProduct).Return(nil)

	err := usecase.Create(context.Background(), newRelatedProduct)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestRelatedProductUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.RelatedProductRepository)
	timeout := time.Duration(10)
	usecase := test.NewRelatedProductUsecase(mockRepo, timeout)
	updatedRelatedProduct := &domain.RelatedProduct{
		ID:           bson.NewObjectID(), // Existing ID of the record to update
		ProductID1:   bson.NewObjectID(),
		ProductID2:   bson.NewObjectID(),
		DisplayOrder: 2,
	}

	mockRepo.On("Update", mock.Anything, updatedRelatedProduct).Return(nil)

	err := usecase.Update(context.Background(), updatedRelatedProduct)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestRelatedProductUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.RelatedProductRepository)
	timeout := time.Duration(10)
	usecase := test.NewRelatedProductUsecase(mockRepo, timeout)

	catalogID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, catalogID).Return(nil)

	err := usecase.Delete(context.Background(), catalogID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestRelatedProductUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.RelatedProductRepository)
	timeout := time.Duration(10)
	usecase := test.NewRelatedProductUsecase(mockRepo, timeout)

	fetchedRelatedProducts := []domain.RelatedProduct{
		{
			ID:           bson.NewObjectID(),
			ProductID1:   bson.NewObjectID(),
			ProductID2:   bson.NewObjectID(),
			DisplayOrder: 1,
		},
		{
			ID:           bson.NewObjectID(),
			ProductID1:   bson.NewObjectID(),
			ProductID2:   bson.NewObjectID(),
			DisplayOrder: 2,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedRelatedProducts, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedRelatedProducts, result)
	mockRepo.AssertExpectations(t)
}
