package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/news"
	repository "earnforglance/server/repository/news"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type MockSingleResultNewsSettings struct {
	mock.Mock
}

func (m *MockSingleResultNewsSettings) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.NewsSettings); ok {
		*v.(*domain.NewsSettings) = *result
	}
	return args.Error(1)
}

var mockItemNewsSettings = &domain.NewsSettings{
	ID:                                     bson.NewObjectID(), // Existing ID of the record to update
	Enabled:                                false,
	AllowNotRegisteredUsersToLeaveComments: true,
	NotifyAboutNewNewsComments:             false,
	ShowNewsOnMainPage:                     false,
	MainPageNewsCount:                      3,
	NewsArchivePageSize:                    15,
	ShowHeaderRssUrl:                       false,
	NewsCommentsMustBeApproved:             false,
	ShowNewsCommentsPerStore:               true,
}

func TestNewsSettingsRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionNewsSettings

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultNewsSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemNewsSettings, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewNewsSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemNewsSettings.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultNewsSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewNewsSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemNewsSettings.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestNewsSettingsRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionNewsSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemNewsSettings).Return(nil, nil).Once()

	repo := repository.NewNewsSettingsRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemNewsSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestNewsSettingsRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionNewsSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemNewsSettings.ID}
	update := bson.M{"$set": mockItemNewsSettings}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewNewsSettingsRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemNewsSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
