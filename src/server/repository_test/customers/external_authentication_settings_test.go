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
	"go.mongodb.org/mongo-driver/bson"
)

type MockSingleResultExternalAuthenticationSettings struct {
	mock.Mock
}

func (m *MockSingleResultExternalAuthenticationSettings) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.ExternalAuthenticationSettings); ok {
		*v.(*domain.ExternalAuthenticationSettings) = *result
	}
	return args.Error(1)
}

var mockItemExternalAuthenticationSettings = &domain.ExternalAuthenticationSettings{
	RequireEmailValidation:                false,                // Email validation is not required
	LogErrors:                             false,                // Do not log errors
	AllowCustomersToRemoveAssociations:    false,                // Do not allow users to remove associations
	ActiveAuthenticationMethodSystemNames: []string{"LinkedIn"}, // Example active authentication method
}

func TestExternalAuthenticationSettingsRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionExternalAuthenticationSettings

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultExternalAuthenticationSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemExternalAuthenticationSettings, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewExternalAuthenticationSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemExternalAuthenticationSettings.ActiveAuthenticationMethodSystemNames[0])

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultExternalAuthenticationSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewExternalAuthenticationSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemExternalAuthenticationSettings.ActiveAuthenticationMethodSystemNames[0])

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestExternalAuthenticationSettingsRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionExternalAuthenticationSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemExternalAuthenticationSettings).Return(nil, nil).Once()

	repo := repository.NewExternalAuthenticationSettingsRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemExternalAuthenticationSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestExternalAuthenticationSettingsRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionExternalAuthenticationSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemExternalAuthenticationSettings.ActiveAuthenticationMethodSystemNames}
	update := bson.M{"$set": mockItemExternalAuthenticationSettings}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewExternalAuthenticationSettingsRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemExternalAuthenticationSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
