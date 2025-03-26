package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/common"
	repository "earnforglance/server/repository/common"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultCommonSettings struct {
	mock.Mock
}

func (m *MockSingleResultCommonSettings) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.CommonSettings); ok {
		*v.(*domain.CommonSettings) = *result
	}
	return args.Error(1)
}

var mockItemCommonSettings = &domain.CommonSettings{
	ID:                               primitive.NewObjectID(), // Existing ID of the record to update
	SubjectFieldOnContactUsForm:      false,
	UseSystemEmailForContactUsForm:   true,
	DisplayJavaScriptDisabledWarning: false,
	Log404Errors:                     false,
	BreadcrumbDelimiter:              " / ",
	IgnoreLogWordlist:                []string{"info", "trace"},
	ClearLogOlderThanDays:            60,
	BbcodeEditorOpenLinksInNewWindow: false,
	PopupForTermsOfServiceLinks:      true,
	JqueryMigrateScriptLoggingActive: false,
	UseResponseCompression:           false,
	FaviconAndAppIconsHeadCode:       "<link rel='icon' href='/newfavicon.ico'>",
	EnableHtmlMinification:           true,
	RestartTimeout:                   new(int),
	HeaderCustomHtml:                 "<header>Updated Header</header>",
	FooterCustomHtml:                 "<footer>Updated Footer</footer>",
}

func TestCommonSettingsRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionCommonSettings

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultCommonSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemCommonSettings, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewCommonSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemCommonSettings.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultCommonSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewCommonSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemCommonSettings.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestCommonSettingsRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionCommonSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemCommonSettings).Return(nil, nil).Once()

	repo := repository.NewCommonSettingsRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemCommonSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestCommonSettingsRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionCommonSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemCommonSettings.ID}
	update := bson.M{"$set": mockItemCommonSettings}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewCommonSettingsRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemCommonSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
