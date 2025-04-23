package repository_test

import (
	"context"
	"errors"
	"testing"
	"time"

	domain "earnforglance/server/domain/forums"
	repository "earnforglance/server/repository/forums"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type MockSingleResultForumGroup struct {
	mock.Mock
}

func (m *MockSingleResultForumGroup) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.ForumGroup); ok {
		*v.(*domain.ForumGroup) = *result
	}
	return args.Error(1)
}

var mockItemForumGroup = &domain.ForumGroup{
	ID:           bson.NewObjectID(), // Existing ID of the record to update
	Name:         "Announcements",
	DisplayOrder: 2,
	CreatedOnUtc: time.Now().AddDate(0, 0, -30), // Created 30 days ago
	UpdatedOnUtc: time.Now(),
}

func TestForumGroupRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionForumGroup

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultForumGroup{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemForumGroup, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewForumGroupRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemForumGroup.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultForumGroup{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewForumGroupRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemForumGroup.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestForumGroupRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionForumGroup

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemForumGroup).Return(nil, nil).Once()

	repo := repository.NewForumGroupRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemForumGroup)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestForumGroupRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionForumGroup

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemForumGroup.ID}
	update := bson.M{"$set": mockItemForumGroup}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewForumGroupRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemForumGroup)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
