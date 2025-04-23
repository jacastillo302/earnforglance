package usecase_test

import (
	"context"
	mocks "earnforglance/server/domain/mocks"
	domain "earnforglance/server/domain/security"
	test "earnforglance/server/usecase/security"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestProxySettingsUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.ProxySettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewProxySettingsUsecase(mockRepo, timeout)

	securityID := bson.NewObjectID().Hex()

	updatedProxySettings := domain.ProxySettings{
		ID:              bson.NewObjectID(), // Existing ID of the record to update
		Enabled:         false,
		Address:         "10.0.0.1",
		Port:            "3128",
		Username:        "updateduser",
		Password:        "updatedpass",
		BypassOnLocal:   false,
		PreAuthenticate: true,
	}
	mockRepo.On("FetchByID", mock.Anything, securityID).Return(updatedProxySettings, nil)

	result, err := usecase.FetchByID(context.Background(), securityID)

	assert.NoError(t, err)
	assert.Equal(t, updatedProxySettings, result)
	mockRepo.AssertExpectations(t)
}

func TestProxySettingsUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.ProxySettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewProxySettingsUsecase(mockRepo, timeout)

	newProxySettings := &domain.ProxySettings{
		Enabled:         true,
		Address:         "192.168.1.1",
		Port:            "8080",
		Username:        "proxyuser",
		Password:        "proxypass",
		BypassOnLocal:   true,
		PreAuthenticate: false,
	}

	mockRepo.On("Create", mock.Anything, newProxySettings).Return(nil)

	err := usecase.Create(context.Background(), newProxySettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProxySettingsUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.ProxySettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewProxySettingsUsecase(mockRepo, timeout)

	updatedProxySettings := &domain.ProxySettings{
		ID:              bson.NewObjectID(), // Existing ID of the record to update
		Enabled:         false,
		Address:         "10.0.0.1",
		Port:            "3128",
		Username:        "updateduser",
		Password:        "updatedpass",
		BypassOnLocal:   false,
		PreAuthenticate: true,
	}

	mockRepo.On("Update", mock.Anything, updatedProxySettings).Return(nil)

	err := usecase.Update(context.Background(), updatedProxySettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProxySettingsUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.ProxySettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewProxySettingsUsecase(mockRepo, timeout)

	securityID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, securityID).Return(nil)

	err := usecase.Delete(context.Background(), securityID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProxySettingsUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.ProxySettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewProxySettingsUsecase(mockRepo, timeout)

	fetchedProxySettings := []domain.ProxySettings{
		{
			ID:              bson.NewObjectID(),
			Enabled:         true,
			Address:         "192.168.1.1",
			Port:            "8080",
			Username:        "proxyuser",
			Password:        "proxypass",
			BypassOnLocal:   true,
			PreAuthenticate: false,
		},
		{
			ID:              bson.NewObjectID(),
			Enabled:         false,
			Address:         "10.0.0.1",
			Port:            "3128",
			Username:        "updateduser",
			Password:        "updatedpass",
			BypassOnLocal:   false,
			PreAuthenticate: true,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedProxySettings, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedProxySettings, result)
	mockRepo.AssertExpectations(t)
}
