package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/catalog"
	repository "earnforglance/server/repository/catalog"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultProductTemplate struct {
	mock.Mock
}

func (m *MockSingleResultProductTemplate) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.ProductTemplate); ok {
		*v.(*domain.ProductTemplate) = *result
	}
	return args.Error(1)
}

var mockItemProductTemplate = &domain.ProductTemplate{
	ID:                  primitive.NewObjectID(), // Existing ID of the record to update
	Name:                "Updated Template",
	ViewPath:            "/Views/Product/Updated.cshtml",
	DisplayOrder:        2,
	IgnoredProductTypes: "Service",
}

func TestProductTemplateRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionProductTemplate

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultProductTemplate{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemProductTemplate, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewProductTemplateRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemProductTemplate.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultProductTemplate{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewProductTemplateRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemProductTemplate.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestProductTemplateRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionProductTemplate

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemProductTemplate).Return(nil, nil).Once()

	repo := repository.NewProductTemplateRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemProductTemplate)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestProductTemplateRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionProductTemplate

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemProductTemplate.ID}
	update := bson.M{"$set": mockItemProductTemplate}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewProductTemplateRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemProductTemplate)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
