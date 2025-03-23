package usecase_test

import (
	"context"
	"testing"
	"time"

	domain "earnforglance/server/domain/catalog"
	mocks "earnforglance/server/domain/mocks"
	test "earnforglance/server/usecase/catalog"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestProductTemplateUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.ProductTemplateRepository)
	timeout := time.Duration(10)
	usecase := test.NewProductTemplateUsecase(mockRepo, timeout)

	productTemplateID := primitive.NewObjectID().Hex()

	updatedProductTemplate := domain.ProductTemplate{
		ID:                  primitive.NewObjectID(), // Existing ID of the record to update
		Name:                "Updated Template",
		ViewPath:            "/Views/Product/Updated.cshtml",
		DisplayOrder:        2,
		IgnoredProductTypes: "Service",
	}

	mockRepo.On("FetchByID", mock.Anything, productTemplateID).Return(updatedProductTemplate, nil)

	result, err := usecase.FetchByID(context.Background(), productTemplateID)

	assert.NoError(t, err)
	assert.Equal(t, updatedProductTemplate, result)
	mockRepo.AssertExpectations(t)
}

func TestProductTemplateUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.ProductTemplateRepository)
	timeout := time.Duration(10)
	usecase := test.NewProductTemplateUsecase(mockRepo, timeout)

	newProductTemplate := &domain.ProductTemplate{
		Name:                "Default Template",
		ViewPath:            "/Views/Product/Default.cshtml",
		DisplayOrder:        1,
		IgnoredProductTypes: "Digital,Service",
	}

	mockRepo.On("Create", mock.Anything, newProductTemplate).Return(nil)

	err := usecase.Create(context.Background(), newProductTemplate)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductTemplateUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.ProductTemplateRepository)
	timeout := time.Duration(10)
	usecase := test.NewProductTemplateUsecase(mockRepo, timeout)

	updatedProductTemplate := &domain.ProductTemplate{
		ID:                  primitive.NewObjectID(), // Existing ID of the record to update
		Name:                "Updated Template",
		ViewPath:            "/Views/Product/Updated.cshtml",
		DisplayOrder:        2,
		IgnoredProductTypes: "Service",
	}

	mockRepo.On("Update", mock.Anything, updatedProductTemplate).Return(nil)

	err := usecase.Update(context.Background(), updatedProductTemplate)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductTemplateUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.ProductTemplateRepository)
	timeout := time.Duration(10)
	usecase := test.NewProductTemplateUsecase(mockRepo, timeout)

	productTemplateID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, productTemplateID).Return(nil)

	err := usecase.Delete(context.Background(), productTemplateID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductTemplateUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.ProductTemplateRepository)
	timeout := time.Duration(10)
	usecase := test.NewProductTemplateUsecase(mockRepo, timeout)

	fetchedProductTemplates := []domain.ProductTemplate{
		{
			ID:                  primitive.NewObjectID(),
			Name:                "Default Template",
			ViewPath:            "/Views/Product/Default.cshtml",
			DisplayOrder:        1,
			IgnoredProductTypes: "Digital,Service",
		},
		{
			ID:                  primitive.NewObjectID(),
			Name:                "Custom Template",
			ViewPath:            "/Views/Product/Custom.cshtml",
			DisplayOrder:        2,
			IgnoredProductTypes: "Physical",
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedProductTemplates, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedProductTemplates, result)
	mockRepo.AssertExpectations(t)
}
