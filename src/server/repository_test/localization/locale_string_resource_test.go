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

type MockSingleResultLocaleStringResource struct {
	mock.Mock
}

func (m *MockSingleResultLocaleStringResource) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.LocaleStringResource); ok {
		*v.(*domain.LocaleStringResource) = *result
	}
	return args.Error(1)
}

var mockItemLocaleStringResource = &domain.LocaleStringResource{
	ID:            primitive.NewObjectID(), // Existing ID of the record to update
	LanguageID:    primitive.NewObjectID(),
	ResourceName:  "WelcomeMessage",
	ResourceValue: "Welcome to our updated platform!",
}

func TestLocaleStringResourceRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionLocaleStringResource

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultLocaleStringResource{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemLocaleStringResource, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewLocaleStringResourceRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemLocaleStringResource.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultLocaleStringResource{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewLocaleStringResourceRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemLocaleStringResource.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestLocaleStringResourceRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionLocaleStringResource

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemLocaleStringResource).Return(nil, nil).Once()

	repo := repository.NewLocaleStringResourceRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemLocaleStringResource)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestLocaleStringResourceRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionLocaleStringResource

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemLocaleStringResource.ID}
	update := bson.M{"$set": mockItemLocaleStringResource}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewLocaleStringResourceRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemLocaleStringResource)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
