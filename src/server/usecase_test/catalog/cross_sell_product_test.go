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

func TestCrossSellProductUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.CrossSellProductRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewCrossSellProductUsecase(mockRepo, timeout)

	crossSellProductID := bson.NewObjectID().Hex()

	expectedCrossSellProduct := domain.CrossSellProduct{
		ID:         bson.NewObjectID(), // Existing ID of the record to update
		ProductID1: bson.NewObjectID(),
		ProductID2: bson.NewObjectID(),
	}

	mockRepo.On("FetchByID", mock.Anything, crossSellProductID).Return(expectedCrossSellProduct, nil)

	result, err := usecase.FetchByID(context.Background(), crossSellProductID)

	assert.NoError(t, err)
	assert.Equal(t, expectedCrossSellProduct, result)
	mockRepo.AssertExpectations(t)
}

func TestCrossSellProductUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.CrossSellProductRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewCrossSellProductUsecase(mockRepo, timeout)

	newCrossSellProduct := &domain.CrossSellProduct{
		ID:         bson.NewObjectID(), // Existing ID of the record to update
		ProductID1: bson.NewObjectID(),
		ProductID2: bson.NewObjectID(),
	}

	mockRepo.On("Create", mock.Anything, newCrossSellProduct).Return(nil)

	err := usecase.Create(context.Background(), newCrossSellProduct)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCrossSellProductUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.CrossSellProductRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewCrossSellProductUsecase(mockRepo, timeout)

	updatedCrossSellProduct := &domain.CrossSellProduct{
		ID:         bson.NewObjectID(), // Existing ID of the record to update
		ProductID1: bson.NewObjectID(),
		ProductID2: bson.NewObjectID(),
	}

	mockRepo.On("Update", mock.Anything, updatedCrossSellProduct).Return(nil)

	err := usecase.Update(context.Background(), updatedCrossSellProduct)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCrossSellProductUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.CrossSellProductRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewCrossSellProductUsecase(mockRepo, timeout)

	crossSellProductID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, crossSellProductID).Return(nil)

	err := usecase.Delete(context.Background(), crossSellProductID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCrossSellProductUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.CrossSellProductRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewCrossSellProductUsecase(mockRepo, timeout)

	expectedCrossSellProducts := []domain.CrossSellProduct{
		{
			ID:         bson.NewObjectID(),
			ProductID1: bson.NewObjectID(),
			ProductID2: bson.NewObjectID(),
		},
		{
			ID:         bson.NewObjectID(),
			ProductID1: bson.NewObjectID(),
			ProductID2: bson.NewObjectID(),
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(expectedCrossSellProducts, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, expectedCrossSellProducts, result)
	mockRepo.AssertExpectations(t)
}
