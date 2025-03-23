package repository_test

import (
	"context"
	"errors"
	"testing"
	"time"

	domain "earnforglance/server/domain/news"
	"earnforglance/server/mongo/mocks"
	repository "earnforglance/server/repository/news"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultNewsComment struct {
	mock.Mock
}

func (m *MockSingleResultNewsComment) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.NewsComment); ok {
		*v.(*domain.NewsComment) = *result
	}
	return args.Error(1)
}

var mockItemNewsComment = &domain.NewsComment{
	ID:           primitive.NewObjectID(), // Existing ID of the record to update
	CommentTitle: "Updated Comment Title",
	CommentText:  "This is an updated comment text.",
	NewsItemID:   primitive.NewObjectID(),
	CustomerID:   primitive.NewObjectID(),
	IsApproved:   false,
	StoreID:      primitive.NewObjectID(),
	CreatedOnUtc: time.Now().AddDate(0, 0, -7), // Created 7 days ago
}

func TestNewsCommentRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionNewsComment

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultNewsComment{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemNewsComment, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewNewsCommentRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemNewsComment.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultNewsComment{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewNewsCommentRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemNewsComment.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestNewsCommentRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionNewsComment

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemNewsComment).Return(nil, nil).Once()

	repo := repository.NewNewsCommentRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemNewsComment)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestNewsCommentRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionNewsComment

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemNewsComment.ID}
	update := bson.M{"$set": mockItemNewsComment}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewNewsCommentRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemNewsComment)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
