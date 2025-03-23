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

type MockSingleResultProductReviewReviewTypeMapping struct {
	mock.Mock
}

func (m *MockSingleResultProductReviewReviewTypeMapping) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.ProductReviewReviewTypeMapping); ok {
		*v.(*domain.ProductReviewReviewTypeMapping) = *result
	}
	return args.Error(1)
}

var mockItemProductReviewReviewTypeMapping = &domain.ProductReviewReviewTypeMapping{
	ID:              primitive.NewObjectID(), // Existing ID of the record to update
	ProductReviewID: primitive.NewObjectID(),
	ReviewTypeID:    primitive.NewObjectID(),
	Rating:          4,
}

func TestProductReviewReviewTypeMappingRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionProductReviewReviewTypeMapping

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultProductReviewReviewTypeMapping{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemProductReviewReviewTypeMapping, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewProductReviewReviewTypeMappingRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemProductReviewReviewTypeMapping.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultProductReviewReviewTypeMapping{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewProductReviewReviewTypeMappingRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemProductReviewReviewTypeMapping.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestProductReviewReviewTypeMappingRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionProductReviewReviewTypeMapping

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemProductReviewReviewTypeMapping).Return(nil, nil).Once()

	repo := repository.NewProductReviewReviewTypeMappingRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemProductReviewReviewTypeMapping)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestProductReviewReviewTypeMappingRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionProductReviewReviewTypeMapping

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemProductReviewReviewTypeMapping.ID}
	update := bson.M{"$set": mockItemProductReviewReviewTypeMapping}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewProductReviewReviewTypeMappingRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemProductReviewReviewTypeMapping)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
