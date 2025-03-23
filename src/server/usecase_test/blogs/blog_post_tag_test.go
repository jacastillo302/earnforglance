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

func TestBlogPostTagUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.BlogPostTagRepository)
	timeout := time.Duration(10)
	usecase := test.NewBlogPostTagUsecase(mockRepo, timeout)

	blogposttagID := primitive.NewObjectID().Hex()

	newBlogPostTag := domain.BlogPostTag{
		ID:            primitive.NewObjectID(),
		Name:          "Tech Updates",
		BlogPostCount: 10,
	}

	mockRepo.On("FetchByID", mock.Anything, blogposttagID).Return(newBlogPostTag, nil)

	result, err := usecase.FetchByID(context.Background(), blogposttagID)

	assert.NoError(t, err)
	assert.Equal(t, newBlogPostTag, result)
	mockRepo.AssertExpectations(t)
}

func TestBlogPostTagUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.BlogPostTagRepository)
	timeout := time.Duration(10)
	usecase := test.NewBlogPostTagUsecase(mockRepo, timeout)

	newBlogPostTag := &domain.BlogPostTag{
		Name:          "Technology",
		BlogPostCount: 0, // Initially, no blog posts are associated with this tag
	}

	mockRepo.On("Create", mock.Anything, newBlogPostTag).Return(nil)

	err := usecase.Create(context.Background(), newBlogPostTag)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestBlogPostTagUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.BlogPostTagRepository)
	timeout := time.Duration(10)
	usecase := test.NewBlogPostTagUsecase(mockRepo, timeout)

	updatedBlogPostTag := &domain.BlogPostTag{
		ID:            primitive.NewObjectID(), // Existing ID of the record to update
		Name:          "Tech Updates",
		BlogPostCount: 10, // Updated count of blog posts associated with this tag
	}

	mockRepo.On("Update", mock.Anything, updatedBlogPostTag).Return(nil)

	err := usecase.Update(context.Background(), updatedBlogPostTag)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestBlogPostTagUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.BlogPostTagRepository)
	timeout := time.Duration(10)
	usecase := test.NewBlogPostTagUsecase(mockRepo, timeout)

	blogposttagID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, blogposttagID).Return(nil)

	err := usecase.Delete(context.Background(), blogposttagID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestBlogPostTagUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.BlogPostTagRepository)
	timeout := time.Duration(10)
	usecase := test.NewBlogPostTagUsecase(mockRepo, timeout)

	expectedBlogPostTags := []domain.BlogPostTag{
		{
			ID:            primitive.NewObjectID(),
			Name:          "Technology",
			BlogPostCount: 5,
		},
		{
			ID:            primitive.NewObjectID(),
			Name:          "Lifestyle",
			BlogPostCount: 8,
		},
		{
			ID:            primitive.NewObjectID(),
			Name:          "Travel",
			BlogPostCount: 3,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(expectedBlogPostTags, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, expectedBlogPostTags, result)
	mockRepo.AssertExpectations(t)
}
