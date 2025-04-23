package usecase_test

import (
	"context"
	test "earnforglance/server/usecase/catalog"
	"testing"
	"time"

	domain "earnforglance/server/domain/catalog"
	mocks "earnforglance/server/domain/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestProductSpecificationAttributeUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.ProductSpecificationAttributeRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewProductSpecificationAttributeUsecase(mockRepo, timeout)

	productSpecificationAttributeID := bson.NewObjectID().Hex()

	updatedProductSpecificationAttribute := domain.ProductSpecificationAttribute{
		ProductID:                      bson.NewObjectID(),
		SpecificationAttributeTypeID:   10,
		SpecificationAttributeOptionID: bson.NewObjectID(),
		CustomValue:                    "Updated Custom Value",
		AllowFiltering:                 false,
		ShowOnProductPage:              false,
		DisplayOrder:                   2,
	}

	mockRepo.On("FetchByID", mock.Anything, productSpecificationAttributeID).Return(updatedProductSpecificationAttribute, nil)

	result, err := usecase.FetchByID(context.Background(), productSpecificationAttributeID)

	assert.NoError(t, err)
	assert.Equal(t, updatedProductSpecificationAttribute, result)
	mockRepo.AssertExpectations(t)
}

func TestProductSpecificationAttributeUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.ProductSpecificationAttributeRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewProductSpecificationAttributeUsecase(mockRepo, timeout)

	newProductSpecificationAttribute := &domain.ProductSpecificationAttribute{
		ProductID:                      bson.NewObjectID(),
		SpecificationAttributeTypeID:   20,
		SpecificationAttributeOptionID: bson.NewObjectID(),
		CustomValue:                    "Custom Value Example",
		AllowFiltering:                 true,
		ShowOnProductPage:              true,
		DisplayOrder:                   1,
	}

	mockRepo.On("Create", mock.Anything, newProductSpecificationAttribute).Return(nil)

	err := usecase.Create(context.Background(), newProductSpecificationAttribute)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductSpecificationAttributeUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.ProductSpecificationAttributeRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewProductSpecificationAttributeUsecase(mockRepo, timeout)

	updatedProductSpecificationAttribute := &domain.ProductSpecificationAttribute{
		ProductID:                      bson.NewObjectID(),
		SpecificationAttributeTypeID:   30,
		SpecificationAttributeOptionID: bson.NewObjectID(),
		CustomValue:                    "Updated Custom Value",
		AllowFiltering:                 false,
		ShowOnProductPage:              false,
		DisplayOrder:                   2,
	}

	mockRepo.On("Update", mock.Anything, updatedProductSpecificationAttribute).Return(nil)

	err := usecase.Update(context.Background(), updatedProductSpecificationAttribute)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductSpecificationAttributeUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.ProductSpecificationAttributeRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewProductSpecificationAttributeUsecase(mockRepo, timeout)

	productSpecificationAttributeID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, productSpecificationAttributeID).Return(nil)

	err := usecase.Delete(context.Background(), productSpecificationAttributeID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductSpecificationAttributeUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.ProductSpecificationAttributeRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewProductSpecificationAttributeUsecase(mockRepo, timeout)

	fetchedProductSpecificationAttributes := []domain.ProductSpecificationAttribute{
		{
			ProductID:                      bson.NewObjectID(),
			SpecificationAttributeTypeID:   0,
			SpecificationAttributeOptionID: bson.NewObjectID(),
			CustomValue:                    "Custom Value 1",
			AllowFiltering:                 true,
			ShowOnProductPage:              true,
			DisplayOrder:                   1,
		},
		{
			ProductID:                      bson.NewObjectID(),
			SpecificationAttributeTypeID:   20,
			SpecificationAttributeOptionID: bson.NewObjectID(),
			CustomValue:                    "Custom Value 2",
			AllowFiltering:                 false,
			ShowOnProductPage:              false,
			DisplayOrder:                   2,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedProductSpecificationAttributes, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedProductSpecificationAttributes, result)
	mockRepo.AssertExpectations(t)
}
