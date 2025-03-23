package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/blogs"
	"earnforglance/server/mongo/mocks"
	repository "earnforglance/server/repository/blogs"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultBlogPostTag struct {
	mock.Mock
}

func (m *MockSingleResultBlogPostTag) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.BlogPostTag); ok {
		*v.(*domain.BlogPostTag) = *result
	}
	return args.Error(1)
}

func TestBlogPostTagRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionBlogPostTag

	mockItem := domain.BlogPostTag{
		ID:            primitive.ObjectID{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
		Name:          "",
		BlogPostCount: 0,
	}

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultBlogPostTag{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItem, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewBlogPostTagRepository(databaseHelper, collectionName)

		result, err := ur.FetchByID(context.Background(), mockItem.ID.Hex())

		assert.NoError(t, err)
		assert.Equal(t, mockItem, result)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultBlogPostTag{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewBlogPostTagRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItem.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestBlogPostTagRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionBlogPostTag

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)

	mockBlogPostTag := &domain.BlogPostTag{
		ID:            primitive.NewObjectID(), // Existing ID of the record to update
		Name:          "Tech Updates",
		BlogPostCount: 10, // Updated count of blog posts associated with this tag
	}

	collectionHelper.On("InsertOne", mock.Anything, mockBlogPostTag).Return(nil, nil).Once()

	repo := repository.NewBlogPostTagRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockBlogPostTag)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestBlogPostTagRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionBlogPostTag

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)

	mockBlogPostTag := &domain.BlogPostTag{
		Name:          "Technology",
		BlogPostCount: 0, // Initially, no blog posts are associated with this tag
	}

	filter := bson.M{"_id": mockBlogPostTag.ID}
	update := bson.M{"$set": mockBlogPostTag}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewBlogPostTagRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockBlogPostTag)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
