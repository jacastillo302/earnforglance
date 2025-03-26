package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/security"
	repository "earnforglance/server/repository/security"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultSecuritySettings struct {
	mock.Mock
}

func (m *MockSingleResultSecuritySettings) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.SecuritySettings); ok {
		*v.(*domain.SecuritySettings) = *result
	}
	return args.Error(1)
}

var mockItemSecuritySettings = &domain.SecuritySettings{
	ID:                               primitive.NewObjectID(), // Existing ID of the record to update
	EncryptionKey:                    "updatedEncryptionKey456",
	AdminAreaAllowedIpAddresses:      []string{"127.0.0.1", "192.168.0.1"},
	HoneypotEnabled:                  false,
	HoneypotInputName:                "updated_honeypot_field",
	LogHoneypotDetection:             false,
	AllowNonAsciiCharactersInHeaders: true,
	UseAesEncryptionAlgorithm:        false,
	AllowStoreOwnerExportImportCustomersWithHashedPassword: true,
}

func TestSecuritySettingsRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionSecuritySettings

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultSecuritySettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemSecuritySettings, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewSecuritySettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemSecuritySettings.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultSecuritySettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewSecuritySettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemSecuritySettings.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestSecuritySettingsRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionSecuritySettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemSecuritySettings).Return(nil, nil).Once()

	repo := repository.NewSecuritySettingsRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemSecuritySettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestSecuritySettingsRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionSecuritySettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemSecuritySettings.ID}
	update := bson.M{"$set": mockItemSecuritySettings}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewSecuritySettingsRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemSecuritySettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
