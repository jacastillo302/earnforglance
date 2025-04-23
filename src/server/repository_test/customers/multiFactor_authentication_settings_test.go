package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/customers"
	repository "earnforglance/server/repository/customers"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type MockSingleResultMultiFactorAuthenticationSettings struct {
	mock.Mock
}

func (m *MockSingleResultMultiFactorAuthenticationSettings) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.MultiFactorAuthenticationSettings); ok {
		*v.(*domain.MultiFactorAuthenticationSettings) = *result
	}
	return args.Error(1)
}

var mockItemMultiFactorAuthenticationSettings = &domain.MultiFactorAuthenticationSettings{
	ActiveAuthenticationMethodSystemNames: []string{"GoogleAuthenticator", "Authy"}, // Example active authentication methods
	ForceMultifactorAuthentication:        true,
}

func TestMultiFactorAuthenticationSettingsRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionMultiFactorAuthenticationSettings

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultMultiFactorAuthenticationSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemMultiFactorAuthenticationSettings, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewMultiFactorAuthenticationSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemMultiFactorAuthenticationSettings.ActiveAuthenticationMethodSystemNames[0])

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultMultiFactorAuthenticationSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewMultiFactorAuthenticationSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemMultiFactorAuthenticationSettings.ActiveAuthenticationMethodSystemNames[0])

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestMultiFactorAuthenticationSettingsRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionMultiFactorAuthenticationSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemMultiFactorAuthenticationSettings).Return(nil, nil).Once()

	repo := repository.NewMultiFactorAuthenticationSettingsRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemMultiFactorAuthenticationSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestMultiFactorAuthenticationSettingsRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionMultiFactorAuthenticationSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemMultiFactorAuthenticationSettings.ActiveAuthenticationMethodSystemNames}
	update := bson.M{"$set": mockItemMultiFactorAuthenticationSettings}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewMultiFactorAuthenticationSettingsRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemMultiFactorAuthenticationSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
