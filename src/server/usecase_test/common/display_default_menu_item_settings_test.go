package usecase_test

import (
	"context"
	domain "earnforglance/server/domain/common"
	mocks "earnforglance/server/domain/mocks"
	test "earnforglance/server/usecase/common"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestDisplayDefaultMenuItemSettingsUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.DisplayDefaultMenuItemSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewDisplayDefaultMenuItemSettingsUsecase(mockRepo, timeout)

	menuItemID := primitive.NewObjectID().Hex()

	updatedDisplayDefaultMenuItemSettings := domain.DisplayDefaultMenuItemSettings{
		ID:                           primitive.NewObjectID(), // Existing ID of the record to update
		DisplayHomepageMenuItem:      false,
		DisplayNewProductsMenuItem:   false,
		DisplayProductSearchMenuItem: true,
		DisplayCustomerInfoMenuItem:  true,
		DisplayBlogMenuItem:          false,
		DisplayForumsMenuItem:        true,
		DisplayContactUsMenuItem:     false,
	}

	mockRepo.On("FetchByID", mock.Anything, menuItemID).Return(updatedDisplayDefaultMenuItemSettings, nil)

	result, err := usecase.FetchByID(context.Background(), menuItemID)

	assert.NoError(t, err)
	assert.Equal(t, updatedDisplayDefaultMenuItemSettings, result)
	mockRepo.AssertExpectations(t)
}

func TestDisplayDefaultMenuItemSettingsUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.DisplayDefaultMenuItemSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewDisplayDefaultMenuItemSettingsUsecase(mockRepo, timeout)

	newDisplayDefaultMenuItemSettings := &domain.DisplayDefaultMenuItemSettings{
		DisplayHomepageMenuItem:      true,
		DisplayNewProductsMenuItem:   true,
		DisplayProductSearchMenuItem: true,
		DisplayCustomerInfoMenuItem:  false,
		DisplayBlogMenuItem:          true,
		DisplayForumsMenuItem:        false,
		DisplayContactUsMenuItem:     true,
	}

	mockRepo.On("Create", mock.Anything, newDisplayDefaultMenuItemSettings).Return(nil)

	err := usecase.Create(context.Background(), newDisplayDefaultMenuItemSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDisplayDefaultMenuItemSettingsUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.DisplayDefaultMenuItemSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewDisplayDefaultMenuItemSettingsUsecase(mockRepo, timeout)

	updatedDisplayDefaultMenuItemSettings := &domain.DisplayDefaultMenuItemSettings{
		ID:                           primitive.NewObjectID(), // Existing ID of the record to update
		DisplayHomepageMenuItem:      false,
		DisplayNewProductsMenuItem:   false,
		DisplayProductSearchMenuItem: true,
		DisplayCustomerInfoMenuItem:  true,
		DisplayBlogMenuItem:          false,
		DisplayForumsMenuItem:        true,
		DisplayContactUsMenuItem:     false,
	}

	mockRepo.On("Update", mock.Anything, updatedDisplayDefaultMenuItemSettings).Return(nil)

	err := usecase.Update(context.Background(), updatedDisplayDefaultMenuItemSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDisplayDefaultMenuItemSettingsUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.DisplayDefaultMenuItemSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewDisplayDefaultMenuItemSettingsUsecase(mockRepo, timeout)

	menuItemID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, menuItemID).Return(nil)

	err := usecase.Delete(context.Background(), menuItemID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDisplayDefaultMenuItemSettingsUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.DisplayDefaultMenuItemSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewDisplayDefaultMenuItemSettingsUsecase(mockRepo, timeout)

	fetchedDisplayDefaultMenuItemSettings := []domain.DisplayDefaultMenuItemSettings{
		{
			ID:                           primitive.NewObjectID(),
			DisplayHomepageMenuItem:      true,
			DisplayNewProductsMenuItem:   true,
			DisplayProductSearchMenuItem: true,
			DisplayCustomerInfoMenuItem:  false,
			DisplayBlogMenuItem:          true,
			DisplayForumsMenuItem:        false,
			DisplayContactUsMenuItem:     true,
		},
		{
			ID:                           primitive.NewObjectID(),
			DisplayHomepageMenuItem:      false,
			DisplayNewProductsMenuItem:   false,
			DisplayProductSearchMenuItem: true,
			DisplayCustomerInfoMenuItem:  true,
			DisplayBlogMenuItem:          false,
			DisplayForumsMenuItem:        true,
			DisplayContactUsMenuItem:     false,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedDisplayDefaultMenuItemSettings, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedDisplayDefaultMenuItemSettings, result)
	mockRepo.AssertExpectations(t)
}
