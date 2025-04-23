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

	"go.mongodb.org/mongo-driver/v2/bson"
)

type MockSingleResultReviewType struct {
	mock.Mock
}

func (m *MockSingleResultReviewType) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.ReviewType); ok {
		*v.(*domain.ReviewType) = *result
	}
	return args.Error(1)
}

var mockItemReviewType = &domain.ReviewType{
	ID:                    bson.NewObjectID(), // Existing ID of the record to update
	Name:                  "Durability",
	Description:           "Review the durability of the product.",
	DisplayOrder:          2,
	VisibleToAllCustomers: false,
	IsRequired:            false,
}

func TestReviewTypeRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionReviewType

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultReviewType{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemReviewType, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewReviewTypeRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemReviewType.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultReviewType{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewReviewTypeRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemReviewType.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestReviewTypeRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionReviewType

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemReviewType).Return(nil, nil).Once()

	repo := repository.NewReviewTypeRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemReviewType)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestReviewTypeRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionReviewType

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemReviewType.ID}
	update := bson.M{"$set": mockItemReviewType}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewReviewTypeRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemReviewType)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
