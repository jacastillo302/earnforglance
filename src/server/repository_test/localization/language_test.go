package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/localization"
	repository "earnforglance/server/repository/localization"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultLanguage struct {
	mock.Mock
}

func (m *MockSingleResultLanguage) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.Language); ok {
		*v.(*domain.Language) = *result
	}
	return args.Error(1)
}

var mockItemLanguage = &domain.Language{
	ID:                primitive.NewObjectID(), // Existing ID of the record to update
	Name:              "Spanish",
	LanguageCulture:   "es-ES",
	UniqueSeoCode:     "es",
	FlagImageFileName: "es.png",
	Rtl:               false,
	LimitedToStores:   true,
	DefaultCurrencyID: primitive.NewObjectID(),
	Published:         false,
	DisplayOrder:      2,
}

func TestLanguageRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionLanguage

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultLanguage{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemLanguage, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewLanguageRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemLanguage.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultLanguage{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewLanguageRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemLanguage.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestLanguageRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionLanguage

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemLanguage).Return(nil, nil).Once()

	repo := repository.NewLanguageRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemLanguage)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestLanguageRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionLanguage

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemLanguage.ID}
	update := bson.M{"$set": mockItemLanguage}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewLanguageRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemLanguage)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
