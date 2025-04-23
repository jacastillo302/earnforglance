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

	"go.mongodb.org/mongo-driver/v2/bson"
)

type MockSingleResultProxySettings struct {
	mock.Mock
}

func (m *MockSingleResultProxySettings) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.ProxySettings); ok {
		*v.(*domain.ProxySettings) = *result
	}
	return args.Error(1)
}

var mockItemProxySettings = &domain.ProxySettings{
	ID:              bson.NewObjectID(), // Existing ID of the record to update
	Enabled:         false,
	Address:         "10.0.0.1",
	Port:            "3128",
	Username:        "updateduser",
	Password:        "updatedpass",
	BypassOnLocal:   false,
	PreAuthenticate: true,
}

func TestProxySettingsRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionProxySettings

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultProxySettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemProxySettings, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewProxySettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemProxySettings.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultProxySettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewProxySettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemProxySettings.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestProxySettingsRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionProxySettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemProxySettings).Return(nil, nil).Once()

	repo := repository.NewProxySettingsRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemProxySettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestProxySettingsRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionProxySettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemProxySettings.ID}
	update := bson.M{"$set": mockItemProxySettings}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewProxySettingsRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemProxySettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
