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

func TestManufacturerTemplateUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.ManufacturerTemplateRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewManufacturerTemplateUsecase(mockRepo, timeout)

	manufacturerTemplateID := bson.NewObjectID().Hex()

	expectedManufacturerTemplate := domain.ManufacturerTemplate{
		ID:           bson.NewObjectID(),
		Name:         "Updated Manufacturer Template",
		ViewPath:     "/Views/Manufacturer/Updated.cshtml",
		DisplayOrder: 2,
	}

	mockRepo.On("FetchByID", mock.Anything, manufacturerTemplateID).Return(expectedManufacturerTemplate, nil)

	result, err := usecase.FetchByID(context.Background(), manufacturerTemplateID)

	assert.NoError(t, err)
	assert.Equal(t, expectedManufacturerTemplate, result)
	mockRepo.AssertExpectations(t)
}

func TestManufacturerTemplateUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.ManufacturerTemplateRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewManufacturerTemplateUsecase(mockRepo, timeout)

	newManufacturerTemplate := &domain.ManufacturerTemplate{
		Name:         "Default Manufacturer Template",
		ViewPath:     "/Views/Manufacturer/Default.cshtml",
		DisplayOrder: 1,
	}
	mockRepo.On("Create", mock.Anything, newManufacturerTemplate).Return(nil)

	err := usecase.Create(context.Background(), newManufacturerTemplate)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestManufacturerTemplateUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.ManufacturerTemplateRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewManufacturerTemplateUsecase(mockRepo, timeout)

	updatedManufacturerTemplate := &domain.ManufacturerTemplate{
		ID:           bson.NewObjectID(),
		Name:         "Updated Manufacturer Template",
		ViewPath:     "/Views/Manufacturer/Updated.cshtml",
		DisplayOrder: 2,
	}

	mockRepo.On("Update", mock.Anything, updatedManufacturerTemplate).Return(nil)

	err := usecase.Update(context.Background(), updatedManufacturerTemplate)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestManufacturerTemplateUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.ManufacturerTemplateRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewManufacturerTemplateUsecase(mockRepo, timeout)

	manufacturerTemplateID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, manufacturerTemplateID).Return(nil)

	err := usecase.Delete(context.Background(), manufacturerTemplateID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestManufacturerTemplateUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.ManufacturerTemplateRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewManufacturerTemplateUsecase(mockRepo, timeout)

	expectedManufacturerTemplates := []domain.ManufacturerTemplate{
		{
			ID:           bson.NewObjectID(),
			Name:         "Default Manufacturer Template",
			ViewPath:     "/Views/Manufacturer/Default.cshtml",
			DisplayOrder: 1,
		},
		{
			ID:           bson.NewObjectID(),
			Name:         "Custom Manufacturer Template",
			ViewPath:     "/Views/Manufacturer/Custom.cshtml",
			DisplayOrder: 2,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(expectedManufacturerTemplates, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, expectedManufacturerTemplates, result)
	mockRepo.AssertExpectations(t)
}
