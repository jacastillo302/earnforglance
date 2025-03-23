package repository_test

import (
	"context"
	domain "earnforglance/server/domain/blogs"
	"earnforglance/server/mongo/mocks"
	repository "earnforglance/server/repository/blogs"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultBlogComment struct {
	mock.Mock
}

func (m *MockSingleResultBlogComment) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.BlogComment); ok {
		*v.(*domain.BlogComment) = *result
	}
	return args.Error(1)
}

func TestBlogCommentRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionBlogComment

	mockItem := domain.BlogComment{ID: primitive.ObjectID{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, CustomerID: primitive.ObjectID{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, CommentText: "", IsApproved: false, StoreID: primitive.ObjectID{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, BlogPostID: primitive.ObjectID{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, CreatedOnUtc: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultBlogComment{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItem, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewBlogCommentRepository(databaseHelper, collectionName)

		result, err := ur.FetchByID(context.Background(), mockItem.ID.Hex())

		assert.NoError(t, err)
		assert.Equal(t, mockItem, result)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultBlogComment{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewBlogCommentRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItem.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestBlogCommentRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionBlogComment

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)

	mockBlogComment := &domain.BlogComment{
		CustomerID:   primitive.NewObjectID(),
		CommentText:  "This is a new blog comment.",
		IsApproved:   false,
		StoreID:      primitive.NewObjectID(),
		BlogPostID:   primitive.NewObjectID(),
		CreatedOnUtc: time.Now(),
	}
	collectionHelper.On("InsertOne", mock.Anything, mockBlogComment).Return(nil, nil).Once()

	repo := repository.NewBlogCommentRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockBlogComment)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestBlogCommentRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionBlogComment

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)

	mockBlogComment := &domain.BlogComment{
		ID:           primitive.NewObjectID(),
		CustomerID:   primitive.NewObjectID(),
		CommentText:  "This is an updated blog comment.",
		IsApproved:   true,
		StoreID:      primitive.NewObjectID(),
		BlogPostID:   primitive.NewObjectID(),
		CreatedOnUtc: time.Now(),
	}

	filter := bson.M{"_id": mockBlogComment.ID}
	update := bson.M{"$set": mockBlogComment}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewBlogCommentRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockBlogComment)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
