package repository_test

import (
	"context"
	"errors"
	"testing"
	"time"

	domain "earnforglance/server/domain/blogs"
	repository "earnforglance/server/repository/blogs"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultBlogPost struct {
	mock.Mock
}

func (m *MockSingleResultBlogPost) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.BlogPost); ok {
		*v.(*domain.BlogPost) = *result
	}
	return args.Error(1)
}

func TestBlogPostRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionBlogPost

	mockItem := domain.BlogPost{
		ID:               primitive.ObjectID{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
		IncludeInSitemap: false,
		Title:            "",
		Body:             "",
		BodyOverview:     "",
		AllowComments:    false,
		Tags:             "",
		StartDateUtc:     nil,
		EndDateUtc:       nil,
		MetaKeywords:     "",
		MetaDescription:  "",
		MetaTitle:        "",
		LimitedToStores:  false,
		CreatedOnUtc:     time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
	}

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultBlogPost{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItem, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewBlogPostRepository(databaseHelper, collectionName)

		result, err := ur.FetchByID(context.Background(), mockItem.ID.Hex())

		assert.NoError(t, err)
		assert.Equal(t, mockItem, result)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultBlogPost{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewBlogPostRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItem.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestBlogPostRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionBlogPost

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)

	mockBlogPost := &domain.BlogPost{
		ID:               primitive.NewObjectID(), // Existing ID of the record to update
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

	collectionHelper.On("InsertOne", mock.Anything, mockBlogPost).Return(nil, nil).Once()

	repo := repository.NewBlogPostRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockBlogPost)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestBlogPostRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionBlogPost

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)

	mockBlogPost := &domain.BlogPost{
		ID:               primitive.NewObjectID(), // Existing ID of the record to update
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

	filter := bson.M{"_id": mockBlogPost.ID}
	update := bson.M{"$set": mockBlogPost}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewBlogPostRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockBlogPost)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
