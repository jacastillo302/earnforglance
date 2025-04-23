package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/seo"
	repository "earnforglance/server/repository/seo"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type MockSingleResultSeoSettings struct {
	mock.Mock
}

func (m *MockSingleResultSeoSettings) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.SeoSettings); ok {
		*v.(*domain.SeoSettings) = *result
	}
	return args.Error(1)
}

var mockItemSeoSettings = &domain.SeoSettings{
	ID:                                bson.NewObjectID(), // Existing ID of the record to update
	PageTitleSeparator:                "|",
	PageTitleSeoAdjustmentID:          2,
	GenerateProductMetaDescription:    false,
	ConvertNonWesternChars:            true,
	AllowUnicodeCharsInUrls:           false,
	CanonicalUrlsEnabled:              false,
	QueryStringInCanonicalUrlsEnabled: true,
	WwwRequirementID:                  2,
	TwitterMetaTags:                   false,
	OpenGraphMetaTags:                 false,
	ReservedUrlRecordSlugs:            []string{"home", "checkout", "cart"},
	CustomHeadTags:                    "<meta name='description' content='Updated Example'>",
	MicrodataEnabled:                  false,
}

func TestSeoSettingsRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionSeoSettings

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultSeoSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemSeoSettings, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewSeoSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemSeoSettings.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultSeoSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewSeoSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemSeoSettings.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestSeoSettingsRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionSeoSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemSeoSettings).Return(nil, nil).Once()

	repo := repository.NewSeoSettingsRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemSeoSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestSeoSettingsRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionSeoSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemSeoSettings.ID}
	update := bson.M{"$set": mockItemSeoSettings}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewSeoSettingsRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemSeoSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
