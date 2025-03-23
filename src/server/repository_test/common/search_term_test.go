package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/common"
	"earnforglance/server/mongo/mocks"
	repository "earnforglance/server/repository/common"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultSearchTerm struct {
	mock.Mock
}

func (m *MockSingleResultSearchTerm) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.SearchTerm); ok {
		*v.(*domain.SearchTerm) = *result
	}
	return args.Error(1)
}

var mockItemSearchTerm = &domain.SearchTerm{
	ID:      primitive.NewObjectID(), // Existing ID of the record to update
	Keyword: "smartphone",
	StoreID: primitive.NewObjectID(),
	Count:   200,
}

func TestSearchTermRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionSearchTerm

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultSearchTerm{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemSearchTerm, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewSearchTermRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemSearchTerm.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultSearchTerm{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewSearchTermRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemSearchTerm.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestSearchTermRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionSearchTerm

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemSearchTerm).Return(nil, nil).Once()

	repo := repository.NewSearchTermRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemSearchTerm)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestSearchTermRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionSearchTerm

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemSearchTerm.ID}
	update := bson.M{"$set": mockItemSearchTerm}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewSearchTermRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemSearchTerm)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
