package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/catalog"
	"earnforglance/server/mongo/mocks"
	repository "earnforglance/server/repository/catalog"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultCategoryTemplate struct {
	mock.Mock
}

func (m *MockSingleResultCategoryTemplate) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.CategoryTemplate); ok {
		*v.(*domain.CategoryTemplate) = *result
	}
	return args.Error(1)
}

func TestCategoryTemplateRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionCategoryTemplate

	mockItem := domain.CategoryTemplate{ID: primitive.ObjectID{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, Name: "", ViewPath: "", DisplayOrder: 0}

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultCategoryTemplate{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItem, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewCategoryTemplateRepository(databaseHelper, collectionName)

		result, err := ur.FetchByID(context.Background(), mockItem.ID.Hex())

		assert.NoError(t, err)
		assert.Equal(t, mockItem, result)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultCategoryTemplate{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewCategoryTemplateRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItem.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestCategoryTemplateRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionCategoryTemplate

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)

	mockCategoryTemplate := &domain.CategoryTemplate{
		ID:           primitive.NewObjectID(), // Existing ID of the record to update
		Name:         "Updated Category Template",
		ViewPath:     "/Views/Category/Updated.cshtml",
		DisplayOrder: 2,
	}

	collectionHelper.On("InsertOne", mock.Anything, mockCategoryTemplate).Return(nil, nil).Once()

	repo := repository.NewCategoryTemplateRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockCategoryTemplate)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestCategoryTemplateRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionCategoryTemplate

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)

	mockCategoryTemplate := &domain.CategoryTemplate{
		ID:           primitive.NewObjectID(), // Existing ID of the record to update
		Name:         "Updated Category Template",
		ViewPath:     "/Views/Category/Updated.cshtml",
		DisplayOrder: 2,
	}

	filter := bson.M{"_id": mockCategoryTemplate.ID}
	update := bson.M{"$set": mockCategoryTemplate}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewCategoryTemplateRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockCategoryTemplate)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
