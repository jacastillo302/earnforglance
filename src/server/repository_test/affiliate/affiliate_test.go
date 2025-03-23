package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/affiliate"
	"earnforglance/server/mongo/mocks"
	repository "earnforglance/server/repository/affiliate"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResult struct {
	mock.Mock
}

func (m *MockSingleResult) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.Affiliate); ok {
		*v.(*domain.Affiliate) = *result
	}
	return args.Error(1)
}

func TestFetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionAffiliate

	mockItem := domain.Affiliate{
		ID:              primitive.ObjectID{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
		AdminComment:    "",
		FriendlyUrlName: "",
		Deleted:         false,
		Active:          false,
	}

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResult{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItem, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewAffiliateRepository(databaseHelper, collectionName)

		result, err := ur.FetchByID(context.Background(), mockItem.ID.Hex())

		assert.NoError(t, err)
		assert.Equal(t, mockItem, result)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResult{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewAffiliateRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItem.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionAffiliate

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)

	mockAffiliate := &domain.Affiliate{
		ID:              primitive.NewObjectID(),
		AdminComment:    "Test Comment",
		FriendlyUrlName: "test-url",
		Deleted:         false,
		Active:          true,
	}

	collectionHelper.On("InsertOne", mock.Anything, mockAffiliate).Return(nil, nil).Once()

	repo := repository.NewAffiliateRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockAffiliate)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestUpdate(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionAffiliate

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)

	mockAffiliate := &domain.Affiliate{
		ID:              primitive.NewObjectID(),
		AdminComment:    "Updated Comment",
		FriendlyUrlName: "updated-url",
		Deleted:         false,
		Active:          true,
	}

	filter := bson.M{"_id": mockAffiliate.ID}
	update := bson.M{"$set": mockAffiliate}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewAffiliateRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockAffiliate)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
