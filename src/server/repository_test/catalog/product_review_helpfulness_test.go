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

type MockSingleResultProductReviewHelpfulness struct {
	mock.Mock
}

func (m *MockSingleResultProductReviewHelpfulness) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.ProductReviewHelpfulness); ok {
		*v.(*domain.ProductReviewHelpfulness) = *result
	}
	return args.Error(1)
}

var mockItemProductReviewHelpfulness = &domain.ProductReviewHelpfulness{
	ID:              primitive.NewObjectID(), // Existing ID of the record to update
	ProductReviewID: primitive.NewObjectID(),
	WasHelpful:      false,
	CustomerID:      primitive.NewObjectID(),
}

func TestProductReviewHelpfulnessRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionProductReviewHelpfulness

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultProductReviewHelpfulness{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemProductReviewHelpfulness, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewProductReviewHelpfulnessRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemProductReviewHelpfulness.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultProductReviewHelpfulness{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewProductReviewHelpfulnessRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemProductReviewHelpfulness.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestProductReviewHelpfulnessRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionProductReviewHelpfulness

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemProductReviewHelpfulness).Return(nil, nil).Once()

	repo := repository.NewProductReviewHelpfulnessRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemProductReviewHelpfulness)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestProductReviewHelpfulnessRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionProductReviewHelpfulness

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemProductReviewHelpfulness.ID}
	update := bson.M{"$set": mockItemProductReviewHelpfulness}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewProductReviewHelpfulnessRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemProductReviewHelpfulness)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
