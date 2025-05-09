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

func TestCategoryTemplateUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.CategoryTemplateRepository)
	timeout := time.Duration(10)
	usecase := test.NewCategoryTemplateUsecase(mockRepo, timeout)

	categoryTemplateID := bson.NewObjectID().Hex()

	expectedCategoryTemplate := domain.CategoryTemplate{
		ID:           bson.NewObjectID(), // Existing ID of the record to update
		Name:         "Updated Category Template",
		ViewPath:     "/Views/Category/Updated.cshtml",
		DisplayOrder: 2,
	}

	mockRepo.On("FetchByID", mock.Anything, categoryTemplateID).Return(expectedCategoryTemplate, nil)

	result, err := usecase.FetchByID(context.Background(), categoryTemplateID)

	assert.NoError(t, err)
	assert.Equal(t, expectedCategoryTemplate, result)
	mockRepo.AssertExpectations(t)
}

func TestCategoryTemplateUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.CategoryTemplateRepository)
	timeout := time.Duration(10)
	usecase := test.NewCategoryTemplateUsecase(mockRepo, timeout)

	newCategoryTemplate := &domain.CategoryTemplate{
		Name:         "Default Category Template",
		ViewPath:     "/Views/Category/Default.cshtml",
		DisplayOrder: 1,
	}

	mockRepo.On("Create", mock.Anything, newCategoryTemplate).Return(nil)

	err := usecase.Create(context.Background(), newCategoryTemplate)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCategoryTemplateUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.CategoryTemplateRepository)
	timeout := time.Duration(10)
	usecase := test.NewCategoryTemplateUsecase(mockRepo, timeout)

	updatedCategoryTemplate := &domain.CategoryTemplate{
		ID:           bson.NewObjectID(), // Existing ID of the record to update
		Name:         "Updated Category Template",
		ViewPath:     "/Views/Category/Updated.cshtml",
		DisplayOrder: 2,
	}

	mockRepo.On("Update", mock.Anything, updatedCategoryTemplate).Return(nil)

	err := usecase.Update(context.Background(), updatedCategoryTemplate)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCategoryTemplateUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.CategoryTemplateRepository)
	timeout := time.Duration(10)
	usecase := test.NewCategoryTemplateUsecase(mockRepo, timeout)

	categoryTemplateID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, categoryTemplateID).Return(nil)

	err := usecase.Delete(context.Background(), categoryTemplateID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCategoryTemplateUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.CategoryTemplateRepository)
	timeout := time.Duration(10)
	usecase := test.NewCategoryTemplateUsecase(mockRepo, timeout)

	expectedCategoryTemplates := []domain.CategoryTemplate{
		{
			ID:           bson.NewObjectID(),
			Name:         "Default Category Template",
			ViewPath:     "/Views/Category/Default.cshtml",
			DisplayOrder: 1,
		},
		{
			ID:           bson.NewObjectID(),
			Name:         "Custom Category Template",
			ViewPath:     "/Views/Category/Custom.cshtml",
			DisplayOrder: 2,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(expectedCategoryTemplates, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, expectedCategoryTemplates, result)
	mockRepo.AssertExpectations(t)
}
