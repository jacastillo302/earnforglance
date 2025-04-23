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

func TestProductProductTagMappingUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.ProductProductTagMappingRepository)
	timeout := time.Duration(10)
	usecase := test.NewProductProductTagMappingUsecase(mockRepo, timeout)

	productProductTagMappingID := bson.NewObjectID().Hex()

	updatedProductProductTagMapping := domain.ProductProductTagMapping{
		ID:           bson.NewObjectID(), // Existing ID of the record to update
		ProductID:    bson.NewObjectID(),
		ProductTagID: bson.NewObjectID(),
	}

	mockRepo.On("FetchByID", mock.Anything, productProductTagMappingID).Return(updatedProductProductTagMapping, nil)

	result, err := usecase.FetchByID(context.Background(), productProductTagMappingID)

	assert.NoError(t, err)
	assert.Equal(t, updatedProductProductTagMapping, result)
	mockRepo.AssertExpectations(t)
}

func TestProductProductTagMappingUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.ProductProductTagMappingRepository)
	timeout := time.Duration(10)
	usecase := test.NewProductProductTagMappingUsecase(mockRepo, timeout)

	newProductProductTagMapping := &domain.ProductProductTagMapping{
		ProductID:    bson.NewObjectID(),
		ProductTagID: bson.NewObjectID(),
	}

	mockRepo.On("Create", mock.Anything, newProductProductTagMapping).Return(nil)

	err := usecase.Create(context.Background(), newProductProductTagMapping)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductProductTagMappingUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.ProductProductTagMappingRepository)
	timeout := time.Duration(10)
	usecase := test.NewProductProductTagMappingUsecase(mockRepo, timeout)

	updatedProductProductTagMapping := &domain.ProductProductTagMapping{
		ID:           bson.NewObjectID(), // Existing ID of the record to update
		ProductID:    bson.NewObjectID(),
		ProductTagID: bson.NewObjectID(),
	}

	mockRepo.On("Update", mock.Anything, updatedProductProductTagMapping).Return(nil)

	err := usecase.Update(context.Background(), updatedProductProductTagMapping)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductProductTagMappingUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.ProductProductTagMappingRepository)
	timeout := time.Duration(10)
	usecase := test.NewProductProductTagMappingUsecase(mockRepo, timeout)

	productProductTagMappingID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, productProductTagMappingID).Return(nil)

	err := usecase.Delete(context.Background(), productProductTagMappingID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductProductTagMappingUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.ProductProductTagMappingRepository)
	timeout := time.Duration(10)
	usecase := test.NewProductProductTagMappingUsecase(mockRepo, timeout)

	fetchedProductProductTagMappings := []domain.ProductProductTagMapping{
		{
			ID:           bson.NewObjectID(),
			ProductID:    bson.NewObjectID(),
			ProductTagID: bson.NewObjectID(),
		},
		{
			ID:           bson.NewObjectID(),
			ProductID:    bson.NewObjectID(),
			ProductTagID: bson.NewObjectID(),
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedProductProductTagMappings, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedProductProductTagMappings, result)
	mockRepo.AssertExpectations(t)
}
