package usecase_test

import (
	"context"
	domain "earnforglance/server/domain/cms"
	mocks "earnforglance/server/domain/mocks"
	test "earnforglance/server/usecase/cms"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestWidgetSettingsUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.WidgetSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewWidgetSettingsUsecase(mockRepo, timeout)

	widgetSettingsID := bson.NewObjectID().Hex()

	updatedWidgetSettings := domain.WidgetSettings{
		ID:                      bson.NewObjectID(), // Existing ID of the record to update
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
	usecase := test.NewWidgetSettingsUsecase(mockRepo, timeout)

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
	usecase := test.NewWidgetSettingsUsecase(mockRepo, timeout)

	updatedWidgetSettings := &domain.WidgetSettings{
		ID:                      bson.NewObjectID(), // Existing ID of the record to update
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
	usecase := test.NewWidgetSettingsUsecase(mockRepo, timeout)

	widgetSettingsID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, widgetSettingsID).Return(nil)

	err := usecase.Delete(context.Background(), widgetSettingsID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestWidgetSettingsUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.WidgetSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewWidgetSettingsUsecase(mockRepo, timeout)

	fetchedWidgetSettings := []domain.WidgetSettings{
		{
			ID:                      bson.NewObjectID(),
			ActiveWidgetSystemNames: []string{"WidgetA", "WidgetB", "WidgetC"},
		},
		{
			ID:                      bson.NewObjectID(),
			ActiveWidgetSystemNames: []string{"WidgetX", "WidgetY"},
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedWidgetSettings, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedWidgetSettings, result)
	mockRepo.AssertExpectations(t)
}
