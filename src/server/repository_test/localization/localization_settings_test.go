package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/localization"
	"earnforglance/server/mongo/mocks"
	repository "earnforglance/server/repository/localization"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultLocalizationSettings struct {
	mock.Mock
}

func (m *MockSingleResultLocalizationSettings) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.LocalizationSettings); ok {
		*v.(*domain.LocalizationSettings) = *result
	}
	return args.Error(1)
}

var mockItemLocalizationSettings = &domain.LocalizationSettings{
	ID:                                  primitive.NewObjectID(), // Existing ID of the record to update
	DefaultAdminLanguageID:              primitive.NewObjectID(),
	UseImagesForLanguageSelection:       false,
	SeoFriendlyUrlsForLanguagesEnabled:  false,
	AutomaticallyDetectLanguage:         true,
	LoadAllLocaleRecordsOnStartup:       false,
	LoadAllLocalizedPropertiesOnStartup: true,
	LoadAllUrlRecordsOnStartup:          false,
	IgnoreRtlPropertyForAdminArea:       true,
}

func TestLocalizationSettingsRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionLocalizationSettings

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultLocalizationSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemLocalizationSettings, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewLocalizationSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemLocalizationSettings.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultLocalizationSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewLocalizationSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemLocalizationSettings.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestLocalizationSettingsRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionLocalizationSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemLocalizationSettings).Return(nil, nil).Once()

	repo := repository.NewLocalizationSettingsRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemLocalizationSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestLocalizationSettingsRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionLocalizationSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemLocalizationSettings.ID}
	update := bson.M{"$set": mockItemLocalizationSettings}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewLocalizationSettingsRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemLocalizationSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
