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

func TestBlogPostUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.BlogPostRepository)
	timeout := time.Duration(10)
	usecase := test.NewBlogPostUsecase(mockRepo, timeout)

	blogID := primitive.NewObjectID().Hex()

	expectedBlogPost := domain.BlogPost{
		LanguageID:       primitive.NewObjectID(),
		IncludeInSitemap: true,
		Title:            "New Blog Post",
		Body:             "This is the body of the new blog post.",
		BodyOverview:     "This is an overview of the new blog post.",
		AllowComments:    true,
		Tags:             "tag1,tag2",
		StartDateUtc:     nil, // Optional, can be nil
		EndDateUtc:       nil, // Optional, can be nil
		MetaKeywords:     "blog,post,new",
		MetaDescription:  "This is a meta description for the new blog post.",
		MetaTitle:        "New Blog Post Meta Title",
		LimitedToStores:  false,
		CreatedOnUtc:     time.Now(),
	}

	mockRepo.On("FetchByID", mock.Anything, blogID).Return(expectedBlogPost, nil)

	result, err := usecase.FetchByID(context.Background(), blogID)

	assert.NoError(t, err)
	assert.Equal(t, expectedBlogPost, result)
	mockRepo.AssertExpectations(t)
}

func TestBlogPostUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.BlogPostRepository)
	timeout := time.Duration(10)
	usecase := test.NewBlogPostUsecase(mockRepo, timeout)

	newBlogPost := &domain.BlogPost{
		LanguageID:       primitive.NewObjectID(),
		IncludeInSitemap: true,
		Title:            "New Blog Post",
		Body:             "This is the body of the new blog post.",
		BodyOverview:     "This is an overview of the new blog post.",
		AllowComments:    true,
		Tags:             "tag1,tag2",
		StartDateUtc:     nil, // Optional, can be nil
		EndDateUtc:       nil, // Optional, can be nil
		MetaKeywords:     "blog,post,new",
		MetaDescription:  "This is a meta description for the new blog post.",
		MetaTitle:        "New Blog Post Meta Title",
		LimitedToStores:  false,
		CreatedOnUtc:     time.Now(),
	}

	mockRepo.On("Create", mock.Anything, newBlogPost).Return(nil)

	err := usecase.Create(context.Background(), newBlogPost)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestBlogPostUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.BlogPostRepository)
	timeout := time.Duration(10)
	usecase := test.NewBlogPostUsecase(mockRepo, timeout)

	updatedBlogPost := &domain.BlogPost{
		ID:               primitive.NewObjectID(), // Existing ID of the record to update
		LanguageID:       primitive.NewObjectID(),
		IncludeInSitemap: false,
		Title:            "Updated Blog Post",
		Body:             "This is the updated body of the blog post.",
		BodyOverview:     "This is an updated overview of the blog post.",
		AllowComments:    false,
		Tags:             "tag3,tag4",
		StartDateUtc:     nil, // Optional, can be nil
		EndDateUtc:       nil, // Optional, can be nil
		MetaKeywords:     "blog,post,updated",
		MetaDescription:  "This is an updated meta description for the blog post.",
		MetaTitle:        "Updated Blog Post Meta Title",
		LimitedToStores:  true,
		CreatedOnUtc:     time.Now(),
	}

	mockRepo.On("Update", mock.Anything, updatedBlogPost).Return(nil)

	err := usecase.Update(context.Background(), updatedBlogPost)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestBlogPostUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.BlogPostRepository)
	timeout := time.Duration(10)
	usecase := test.NewBlogPostUsecase(mockRepo, timeout)

	blogID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, blogID).Return(nil)

	err := usecase.Delete(context.Background(), blogID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestBlogPostUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.BlogPostRepository)
	timeout := time.Duration(10)
	usecase := test.NewBlogPostUsecase(mockRepo, timeout)

	fetchedBlogPosts := []domain.BlogPost{
		{
			ID:               primitive.NewObjectID(),
			LanguageID:       primitive.NewObjectID(),
			IncludeInSitemap: true,
			Title:            "First Blog Post",
			Body:             "This is the body of the first blog post.",
			BodyOverview:     "This is an overview of the first blog post.",
			AllowComments:    true,
			Tags:             "tag1,tag2",
			StartDateUtc:     nil,
			EndDateUtc:       nil,
			MetaKeywords:     "blog,post,first",
			MetaDescription:  "This is a meta description for the first blog post.",
			MetaTitle:        "First Blog Post Meta Title",
			LimitedToStores:  false,
			CreatedOnUtc:     time.Now().AddDate(0, 0, -10), // 10 days ago
		},
		{
			ID:               primitive.NewObjectID(),
			LanguageID:       primitive.NewObjectID(),
			IncludeInSitemap: false,
			Title:            "Second Blog Post",
			Body:             "This is the body of the second blog post.",
			BodyOverview:     "This is an overview of the second blog post.",
			AllowComments:    false,
			Tags:             "tag3,tag4",
			StartDateUtc:     nil,
			EndDateUtc:       nil,
			MetaKeywords:     "blog,post,second",
			MetaDescription:  "This is a meta description for the second blog post.",
			MetaTitle:        "Second Blog Post Meta Title",
			LimitedToStores:  true,
			CreatedOnUtc:     time.Now().AddDate(0, 0, -5), // 5 days ago
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedBlogPosts, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedBlogPosts, result)
	mockRepo.AssertExpectations(t)
}
