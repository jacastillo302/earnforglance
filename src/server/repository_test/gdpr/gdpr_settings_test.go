package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/gdpr"
	repository "earnforglance/server/repository/gdpr"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultGdprSettings struct {
	mock.Mock
}

func (m *MockSingleResultGdprSettings) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.GdprSettings); ok {
		*v.(*domain.GdprSettings) = *result
	}
	return args.Error(1)
}

var mockItemGdprSettings = &domain.GdprSettings{
	ID:                                 primitive.NewObjectID(), // Existing ID of the record to update
	GdprEnabled:                        false,
	LogPrivacyPolicyConsent:            false,
	LogNewsletterConsent:               false,
	LogUserProfileChanges:              false,
	DeleteInactiveCustomersAfterMonths: 6,
}

func TestGdprSettingsRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionGdprSettings

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultGdprSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemGdprSettings, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewGdprSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemGdprSettings.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultGdprSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewGdprSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemGdprSettings.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestGdprSettingsRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionGdprSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemGdprSettings).Return(nil, nil).Once()

	repo := repository.NewGdprSettingsRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemGdprSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestGdprSettingsRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionGdprSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemGdprSettings.ID}
	update := bson.M{"$set": mockItemGdprSettings}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewGdprSettingsRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemGdprSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
