package usecase

import (
	"context"
	domain "earnforglance/server/domain/configuration"
	mocks "earnforglance/server/domain/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestSettingUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.SettingRepository)
	timeout := time.Duration(10)
	usecase := NewSettingUsecase(mockRepo, timeout)

	settingID := primitive.NewObjectID().Hex()

	updatedSetting := domain.Setting{
		ID:      primitive.NewObjectID(), // Existing ID of the record to update
		Name:    "SiteTitle",
		Value:   "Updated E-Commerce Store",
		StoreID: primitive.NewObjectID(),
	}

	mockRepo.On("FetchByID", mock.Anything, settingID).Return(updatedSetting, nil)

	result, err := usecase.FetchByID(context.Background(), settingID)

	assert.NoError(t, err)
	assert.Equal(t, updatedSetting, result)
	mockRepo.AssertExpectations(t)
}

func TestSettingUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.SettingRepository)
	timeout := time.Duration(10)
	usecase := NewSettingUsecase(mockRepo, timeout)

	newSetting := &domain.Setting{
		Name:    "SiteTitle",
		Value:   "My E-Commerce Store",
		StoreID: primitive.NewObjectID(),
	}

	mockRepo.On("Create", mock.Anything, newSetting).Return(nil)

	err := usecase.Create(context.Background(), newSetting)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestSettingUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.SettingRepository)
	timeout := time.Duration(10)
	usecase := NewSettingUsecase(mockRepo, timeout)

	updatedSetting := &domain.Setting{
		ID:      primitive.NewObjectID(), // Existing ID of the record to update
		Name:    "SiteTitle",
		Value:   "Updated E-Commerce Store",
		StoreID: primitive.NewObjectID(),
	}

	mockRepo.On("Update", mock.Anything, updatedSetting).Return(nil)

	err := usecase.Update(context.Background(), updatedSetting)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestSettingUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.SettingRepository)
	timeout := time.Duration(10)
	usecase := NewSettingUsecase(mockRepo, timeout)

	settingID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, settingID).Return(nil)

	err := usecase.Delete(context.Background(), settingID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestSettingUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.SettingRepository)
	timeout := time.Duration(10)
	usecase := NewSettingUsecase(mockRepo, timeout)

	fetchedSettings := []domain.Setting{
		{
			ID:      primitive.NewObjectID(),
			Name:    "SiteTitle",
			Value:   "My E-Commerce Store",
			StoreID: primitive.NewObjectID(),
		},
		{
			ID:      primitive.NewObjectID(),
			Name:    "Currency",
			Value:   "USD",
			StoreID: primitive.NewObjectID(),
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedSettings, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedSettings, result)
	mockRepo.AssertExpectations(t)
}
