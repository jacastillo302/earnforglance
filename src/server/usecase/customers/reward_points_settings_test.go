package usecase

import (
	"context"
	domain "earnforglance/server/domain/customers"
	mocks "earnforglance/server/domain/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestRewardPointsSettingsUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.RewardPointsSettingsRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewRewardPointsSettingsUsecase(mockRepo, timeout)

	rewardPointsSettingsID := primitive.NewObjectID().Hex()

	updatedRewardPointsSettings := domain.RewardPointsSettings{
		ID:                               primitive.NewObjectID(), // Existing ID of the record to update
		Enabled:                          false,
		ExchangeRate:                     0.02,
		MinimumRewardPointsToUse:         200,
		MaximumRewardPointsToUsePerOrder: 500,
		MaximumRedeemedRate:              0.3,
		PointsForRegistration:            100,
		RegistrationPointsValidity:       new(int),
		PointsForPurchasesAmount:         200.0,
		PointsForPurchasesPoints:         20,
		PurchasesPointsValidity:          new(int),
		MinOrderTotalToAwardPoints:       100.0,
		ActivationDelay:                  14,
		ActivationDelayPeriodID:          2,
		DisplayHowMuchWillBeEarned:       false,
		PointsAccumulatedForAllStores:    true,
		PageSize:                         50,
	}

	mockRepo.On("FetchByID", mock.Anything, rewardPointsSettingsID).Return(updatedRewardPointsSettings, nil)

	result, err := usecase.FetchByID(context.Background(), rewardPointsSettingsID)

	assert.NoError(t, err)
	assert.Equal(t, updatedRewardPointsSettings, result)
	mockRepo.AssertExpectations(t)
}

func TestRewardPointsSettingsUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.RewardPointsSettingsRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewRewardPointsSettingsUsecase(mockRepo, timeout)

	newRewardPointsSettings := &domain.RewardPointsSettings{
		Enabled:                          true,
		ExchangeRate:                     0.01,
		MinimumRewardPointsToUse:         100,
		MaximumRewardPointsToUsePerOrder: 1000,
		MaximumRedeemedRate:              0.5,
		PointsForRegistration:            50,
		RegistrationPointsValidity:       new(int),
		PointsForPurchasesAmount:         100.0,
		PointsForPurchasesPoints:         10,
		PurchasesPointsValidity:          new(int),
		MinOrderTotalToAwardPoints:       50.0,
		ActivationDelay:                  7,
		ActivationDelayPeriodID:          1,
		DisplayHowMuchWillBeEarned:       true,
		PointsAccumulatedForAllStores:    false,
		PageSize:                         20,
	}
	*newRewardPointsSettings.RegistrationPointsValidity = 365
	*newRewardPointsSettings.PurchasesPointsValidity = 180

	mockRepo.On("Create", mock.Anything, newRewardPointsSettings).Return(nil)

	err := usecase.Create(context.Background(), newRewardPointsSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestRewardPointsSettingsUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.RewardPointsSettingsRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewRewardPointsSettingsUsecase(mockRepo, timeout)

	updatedRewardPointsSettings := &domain.RewardPointsSettings{
		ID:                               primitive.NewObjectID(), // Existing ID of the record to update
		Enabled:                          false,
		ExchangeRate:                     0.02,
		MinimumRewardPointsToUse:         200,
		MaximumRewardPointsToUsePerOrder: 500,
		MaximumRedeemedRate:              0.3,
		PointsForRegistration:            100,
		RegistrationPointsValidity:       new(int),
		PointsForPurchasesAmount:         200.0,
		PointsForPurchasesPoints:         20,
		PurchasesPointsValidity:          new(int),
		MinOrderTotalToAwardPoints:       100.0,
		ActivationDelay:                  14,
		ActivationDelayPeriodID:          2,
		DisplayHowMuchWillBeEarned:       false,
		PointsAccumulatedForAllStores:    true,
		PageSize:                         50,
	}
	*updatedRewardPointsSettings.RegistrationPointsValidity = 180
	*updatedRewardPointsSettings.PurchasesPointsValidity = 90

	mockRepo.On("Update", mock.Anything, updatedRewardPointsSettings).Return(nil)

	err := usecase.Update(context.Background(), updatedRewardPointsSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestRewardPointsSettingsUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.RewardPointsSettingsRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewRewardPointsSettingsUsecase(mockRepo, timeout)

	rewardPointsSettingsID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, rewardPointsSettingsID).Return(nil)

	err := usecase.Delete(context.Background(), rewardPointsSettingsID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestRewardPointsSettingsUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.RewardPointsSettingsRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewRewardPointsSettingsUsecase(mockRepo, timeout)

	fetchedRewardPointsSettings := []domain.RewardPointsSettings{
		{
			ID:                               primitive.NewObjectID(),
			Enabled:                          true,
			ExchangeRate:                     0.01,
			MinimumRewardPointsToUse:         100,
			MaximumRewardPointsToUsePerOrder: 1000,
			MaximumRedeemedRate:              0.5,
			PointsForRegistration:            50,
			RegistrationPointsValidity:       new(int),
			PointsForPurchasesAmount:         100.0,
			PointsForPurchasesPoints:         10,
			PurchasesPointsValidity:          new(int),
			MinOrderTotalToAwardPoints:       50.0,
			ActivationDelay:                  7,
			ActivationDelayPeriodID:          1,
			DisplayHowMuchWillBeEarned:       true,
			PointsAccumulatedForAllStores:    false,
			PageSize:                         20,
		},
		{
			ID:                               primitive.NewObjectID(),
			Enabled:                          false,
			ExchangeRate:                     0.02,
			MinimumRewardPointsToUse:         200,
			MaximumRewardPointsToUsePerOrder: 500,
			MaximumRedeemedRate:              0.3,
			PointsForRegistration:            100,
			RegistrationPointsValidity:       new(int),
			PointsForPurchasesAmount:         200.0,
			PointsForPurchasesPoints:         20,
			PurchasesPointsValidity:          new(int),
			MinOrderTotalToAwardPoints:       100.0,
			ActivationDelay:                  14,
			ActivationDelayPeriodID:          2,
			DisplayHowMuchWillBeEarned:       false,
			PointsAccumulatedForAllStores:    true,
			PageSize:                         50,
		},
	}
	*fetchedRewardPointsSettings[0].RegistrationPointsValidity = 365
	*fetchedRewardPointsSettings[0].PurchasesPointsValidity = 180
	*fetchedRewardPointsSettings[1].RegistrationPointsValidity = 180
	*fetchedRewardPointsSettings[1].PurchasesPointsValidity = 90

	mockRepo.On("Fetch", mock.Anything).Return(fetchedRewardPointsSettings, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedRewardPointsSettings, result)
	mockRepo.AssertExpectations(t)
}
