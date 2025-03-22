package usecase

import (
	"context"
	domain "earnforglance/server/domain/cms"
	mocks "earnforglance/server/domain/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestWidgetSettingsUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.WidgetSettingsRepository)
	timeout := time.Duration(10)
	usecase := NewWidgetSettingsUsecase(mockRepo, timeout)

	widgetSettingsID := primitive.NewObjectID().Hex()

	updatedWidgetSettings := domain.WidgetSettings{
		ID:                      primitive.NewObjectID(), // Existing ID of the record to update
		ActiveWidgetSystemNames: []string{"WidgetX", "WidgetY"},
	}

	mockRepo.On("FetchByID", mock.Anything, widgetSettingsID).Return(updatedWidgetSettings, nil)

	result, err := usecase.FetchByID(context.Background(), widgetSettingsID)

	assert.NoError(t, err)
	assert.Equal(t, updatedWidgetSettings, result)
	mockRepo.AssertExpectations(t)
}

func TestWidgetSettingsUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.WidgetSettingsRepository)
	timeout := time.Duration(10)
	usecase := NewWidgetSettingsUsecase(mockRepo, timeout)

	newWidgetSettings := &domain.WidgetSettings{
		ActiveWidgetSystemNames: []string{"WidgetA", "WidgetB", "WidgetC"},
	}

	mockRepo.On("Create", mock.Anything, newWidgetSettings).Return(nil)

	err := usecase.Create(context.Background(), newWidgetSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestWidgetSettingsUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.WidgetSettingsRepository)
	timeout := time.Duration(10)
	usecase := NewWidgetSettingsUsecase(mockRepo, timeout)

	updatedWidgetSettings := &domain.WidgetSettings{
		ID:                      primitive.NewObjectID(), // Existing ID of the record to update
		ActiveWidgetSystemNames: []string{"WidgetX", "WidgetY"},
	}

	mockRepo.On("Update", mock.Anything, updatedWidgetSettings).Return(nil)

	err := usecase.Update(context.Background(), updatedWidgetSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestWidgetSettingsUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.WidgetSettingsRepository)
	timeout := time.Duration(10)
	usecase := NewWidgetSettingsUsecase(mockRepo, timeout)

	widgetSettingsID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, widgetSettingsID).Return(nil)

	err := usecase.Delete(context.Background(), widgetSettingsID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestWidgetSettingsUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.WidgetSettingsRepository)
	timeout := time.Duration(10)
	usecase := NewWidgetSettingsUsecase(mockRepo, timeout)

	fetchedWidgetSettings := []domain.WidgetSettings{
		{
			ID:                      primitive.NewObjectID(),
			ActiveWidgetSystemNames: []string{"WidgetA", "WidgetB", "WidgetC"},
		},
		{
			ID:                      primitive.NewObjectID(),
			ActiveWidgetSystemNames: []string{"WidgetX", "WidgetY"},
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedWidgetSettings, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedWidgetSettings, result)
	mockRepo.AssertExpectations(t)
}
