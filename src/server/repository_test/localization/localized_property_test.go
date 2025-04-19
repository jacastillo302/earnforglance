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

type MockSingleResultLocalizedProperty struct {
	mock.Mock
}

func (m *MockSingleResultLocalizedProperty) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.LocalizedProperty); ok {
		*v.(*domain.LocalizedProperty) = *result
	}
	return args.Error(1)
}

var mockItemLocalizedProperty = &domain.LocalizedProperty{
	ID:                 primitive.NewObjectID(), // Existing ID of the record to update
	PermissionRecordID: primitive.NewObjectID(),
	LanguageID:         primitive.NewObjectID(),
	LocaleKeyGroup:     "Category",
	LocaleKey:          "Description",
	LocaleValue:        "Electronics and Gadgets",
}

func TestLocalizedPropertyRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionLocalizedProperty

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultLocalizedProperty{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemLocalizedProperty, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewLocalizedPropertyRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemLocalizedProperty.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultLocalizedProperty{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewLocalizedPropertyRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemLocalizedProperty.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestLocalizedPropertyRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionLocalizedProperty

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemLocalizedProperty).Return(nil, nil).Once()

	repo := repository.NewLocalizedPropertyRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemLocalizedProperty)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestLocalizedPropertyRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionLocalizedProperty

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemLocalizedProperty.ID}
	update := bson.M{"$set": mockItemLocalizedProperty}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewLocalizedPropertyRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemLocalizedProperty)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
