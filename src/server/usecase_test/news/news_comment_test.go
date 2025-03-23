package usecase_test

import (
	"context"
	mocks "earnforglance/server/domain/mocks"
	domain "earnforglance/server/domain/news"
	test "earnforglance/server/usecase/news"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestNewsCommentUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.NewsCommentRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewNewsCommentUsecase(mockRepo, timeout)

	newsCommentID := primitive.NewObjectID().Hex()

	updatedNewsComment := domain.NewsComment{
		ID:           primitive.NewObjectID(), // Existing ID of the record to update
		CommentTitle: "Updated Comment Title",
		CommentText:  "This is an updated comment text.",
		NewsItemID:   primitive.NewObjectID(),
		CustomerID:   primitive.NewObjectID(),
		IsApproved:   false,
		StoreID:      primitive.NewObjectID(),
		CreatedOnUtc: time.Now().AddDate(0, 0, -7), // Created 7 days ago
	}

	mockRepo.On("FetchByID", mock.Anything, newsCommentID).Return(updatedNewsComment, nil)

	result, err := usecase.FetchByID(context.Background(), newsCommentID)

	assert.NoError(t, err)
	assert.Equal(t, updatedNewsComment, result)
	mockRepo.AssertExpectations(t)
}

func TestNewsCommentUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.NewsCommentRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewNewsCommentUsecase(mockRepo, timeout)

	newNewsComment := &domain.NewsComment{
		CommentTitle: "Great News!",
		CommentText:  "This is an amazing update. Keep up the good work!",
		NewsItemID:   primitive.NewObjectID(),
		CustomerID:   primitive.NewObjectID(),
		IsApproved:   true,
		StoreID:      primitive.NewObjectID(),
		CreatedOnUtc: time.Now(),
	}

	mockRepo.On("Create", mock.Anything, newNewsComment).Return(nil)

	err := usecase.Create(context.Background(), newNewsComment)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestNewsCommentUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.NewsCommentRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewNewsCommentUsecase(mockRepo, timeout)

	updatedNewsComment := &domain.NewsComment{
		ID:           primitive.NewObjectID(), // Existing ID of the record to update
		CommentTitle: "Updated Comment Title",
		CommentText:  "This is an updated comment text.",
		NewsItemID:   primitive.NewObjectID(),
		CustomerID:   primitive.NewObjectID(),
		IsApproved:   false,
		StoreID:      primitive.NewObjectID(),
		CreatedOnUtc: time.Now().AddDate(0, 0, -7), // Created 7 days ago
	}

	mockRepo.On("Update", mock.Anything, updatedNewsComment).Return(nil)

	err := usecase.Update(context.Background(), updatedNewsComment)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestNewsCommentUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.NewsCommentRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewNewsCommentUsecase(mockRepo, timeout)

	newsCommentID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, newsCommentID).Return(nil)

	err := usecase.Delete(context.Background(), newsCommentID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestNewsCommentUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.NewsCommentRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewNewsCommentUsecase(mockRepo, timeout)

	fetchedNewsComments := []domain.NewsComment{
		{
			ID:           primitive.NewObjectID(),
			CommentTitle: "Great News!",
			CommentText:  "This is an amazing update. Keep up the good work!",
			NewsItemID:   primitive.NewObjectID(),
			CustomerID:   primitive.NewObjectID(),
			IsApproved:   true,
			StoreID:      primitive.NewObjectID(),
			CreatedOnUtc: time.Now().AddDate(0, 0, -10), // Created 10 days ago
		},
		{
			ID:           primitive.NewObjectID(),
			CommentTitle: "Needs Improvement",
			CommentText:  "I think this could be better.",
			NewsItemID:   primitive.NewObjectID(),
			CustomerID:   primitive.NewObjectID(),
			IsApproved:   false,
			StoreID:      primitive.NewObjectID(),
			CreatedOnUtc: time.Now().AddDate(0, 0, -5), // Created 5 days ago
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedNewsComments, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedNewsComments, result)
	mockRepo.AssertExpectations(t)
}
