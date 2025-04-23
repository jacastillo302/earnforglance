package repository_test

import (
	"context"
	domain "earnforglance/server/domain/catalog"
	repository "earnforglance/server/repository/catalog"
	"earnforglance/server/service/data/mongo/mocks"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type MockSingleResultProductReview struct {
	mock.Mock
}

func (m *MockSingleResultProductReview) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.ProductReview); ok {
		*v.(*domain.ProductReview) = *result
	}
	return args.Error(1)
}

var mockItemProductReview = &domain.ProductReview{
	ID:                      bson.NewObjectID(), // Existing ID of the record to update
	CustomerID:              bson.NewObjectID(),
	ProductID:               bson.NewObjectID(),
	StoreID:                 bson.NewObjectID(),
	IsApproved:              false,
	Title:                   "Updated Review Title",
	ReviewText:              "Updated review text with more details.",
	ReplyText:               "Thank you for your feedback!",
	CustomerNotifiedOfReply: true,
	Rating:                  4,
	HelpfulYesTotal:         15,
	HelpfulNoTotal:          3,
	CreatedOnUtc:            time.Now().AddDate(0, 0, -7), // Created 7 days ago
}

func TestProductReviewRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionProductReview

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultProductReview{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemProductReview, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewProductReviewRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemProductReview.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultProductReview{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewProductReviewRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemProductReview.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestProductReviewRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionProductReview

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemProductReview).Return(nil, nil).Once()

	repo := repository.NewProductReviewRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemProductReview)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestProductReviewRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionProductReview

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemProductReview.ID}
	update := bson.M{"$set": mockItemProductReview}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewProductReviewRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemProductReview)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
