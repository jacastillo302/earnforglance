package usecase_test

import (
	"context"
	domain "earnforglance/server/domain/blogs"
	mocks "earnforglance/server/domain/mocks"
	test "earnforglance/server/usecase/blogs"

	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestBlogCommentUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.BlogCommentRepository)
	timeout := time.Duration(10)
	usecase := test.NewBlogCommentUsecase(mockRepo, timeout)

	blogCommentID := primitive.NewObjectID().Hex()

	expectedBlogComment := domain.BlogComment{
		CustomerID:   primitive.NewObjectID(),
		CommentText:  "This is a new blog comment.",
		IsApproved:   false,
		StoreID:      primitive.NewObjectID(),
		BlogPostID:   primitive.NewObjectID(),
		CreatedOnUtc: time.Now(),
	}

	mockRepo.On("FetchByID", mock.Anything, blogCommentID).Return(expectedBlogComment, nil)

	result, err := usecase.FetchByID(context.Background(), blogCommentID)

	assert.NoError(t, err)
	assert.Equal(t, expectedBlogComment, result)
	mockRepo.AssertExpectations(t)
}

func TestBlogCommentUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.BlogCommentRepository)
	timeout := time.Duration(10)
	usecase := test.NewBlogCommentUsecase(mockRepo, timeout)

	newBlogComment := &domain.BlogComment{
		CustomerID:   primitive.NewObjectID(),
		CommentText:  "This is a new blog comment.",
		IsApproved:   false,
		StoreID:      primitive.NewObjectID(),
		BlogPostID:   primitive.NewObjectID(),
		CreatedOnUtc: time.Now(),
	}

	mockRepo.On("Create", mock.Anything, newBlogComment).Return(nil)

	err := usecase.Create(context.Background(), newBlogComment)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestBlogCommentUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.BlogCommentRepository)
	timeout := time.Duration(10)
	usecase := test.NewBlogCommentUsecase(mockRepo, timeout)

	updatedBlogComment := &domain.BlogComment{
		ID:           primitive.NewObjectID(), // Existing ID of the record to update
		CustomerID:   primitive.NewObjectID(),
		CommentText:  "This is an updated blog comment.",
		IsApproved:   true,
		StoreID:      primitive.NewObjectID(),
		BlogPostID:   primitive.NewObjectID(),
		CreatedOnUtc: time.Now(),
	}

	mockRepo.On("Update", mock.Anything, updatedBlogComment).Return(nil)

	err := usecase.Update(context.Background(), updatedBlogComment)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestBlogCommentUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.BlogCommentRepository)
	timeout := time.Duration(10)
	usecase := test.NewBlogCommentUsecase(mockRepo, timeout)

	blogCommentID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, blogCommentID).Return(nil)

	err := usecase.Delete(context.Background(), blogCommentID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestBlogCommentUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.BlogCommentRepository)
	timeout := time.Duration(10)
	usecase := test.NewBlogCommentUsecase(mockRepo, timeout)

	expectedBlogComments := []domain.BlogComment{
		{
			ID:           primitive.NewObjectID(),
			CustomerID:   primitive.NewObjectID(),
			CommentText:  "This is the first blog comment.",
			IsApproved:   true,
			StoreID:      primitive.NewObjectID(),
			BlogPostID:   primitive.NewObjectID(),
			CreatedOnUtc: time.Now().AddDate(0, 0, -10), // 10 days ago
		},
		{
			ID:           primitive.NewObjectID(),
			CustomerID:   primitive.NewObjectID(),
			CommentText:  "This is the second blog comment.",
			IsApproved:   false,
			StoreID:      primitive.NewObjectID(),
			BlogPostID:   primitive.NewObjectID(),
			CreatedOnUtc: time.Now().AddDate(0, 0, -5), // 5 days ago
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(expectedBlogComments, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, expectedBlogComments, result)
	mockRepo.AssertExpectations(t)
}
