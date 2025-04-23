package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/tax"
	repository "earnforglance/server/repository/tax"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type MockSingleResultTaxCategory struct {
	mock.Mock
}

func (m *MockSingleResultTaxCategory) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.TaxCategory); ok {
		*v.(*domain.TaxCategory) = *result
	}
	return args.Error(1)
}

var mockItemTaxCategory = &domain.TaxCategory{
	ID:           bson.NewObjectID(), // Existing ID of the record to update
	Name:         "Reduced Rate",
	DisplayOrder: 2,
}

func TestTaxCategoryRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionTaxCategory

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultTaxCategory{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemTaxCategory, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewTaxCategoryRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemTaxCategory.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultTaxCategory{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewTaxCategoryRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemTaxCategory.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestTaxCategoryRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionTaxCategory

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemTaxCategory).Return(nil, nil).Once()

	repo := repository.NewTaxCategoryRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemTaxCategory)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestTaxCategoryRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionTaxCategory

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemTaxCategory.ID}
	update := bson.M{"$set": mockItemTaxCategory}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewTaxCategoryRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemTaxCategory)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
