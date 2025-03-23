package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/customers"
	"earnforglance/server/mongo/mocks"
	repository "earnforglance/server/repository/customers"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultRewardPointsSettings struct {
	mock.Mock
}

func (m *MockSingleResultRewardPointsSettings) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.RewardPointsSettings); ok {
		*v.(*domain.RewardPointsSettings) = *result
	}
	return args.Error(1)
}

var mockItemRewardPointsSettings = &domain.RewardPointsSettings{
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

func TestRewardPointsSettingsRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionRewardPointsSettings

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultRewardPointsSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemRewardPointsSettings, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewRewardPointsSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemRewardPointsSettings.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultRewardPointsSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewRewardPointsSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemRewardPointsSettings.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestRewardPointsSettingsRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionRewardPointsSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemRewardPointsSettings).Return(nil, nil).Once()

	repo := repository.NewRewardPointsSettingsRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemRewardPointsSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestRewardPointsSettingsRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionRewardPointsSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemRewardPointsSettings.ID}
	update := bson.M{"$set": mockItemRewardPointsSettings}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewRewardPointsSettingsRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemRewardPointsSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
