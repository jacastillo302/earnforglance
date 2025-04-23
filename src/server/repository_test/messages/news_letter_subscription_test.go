package repository_test

import (
	"context"
	"errors"
	"testing"
	"time"

	domain "earnforglance/server/domain/messages"
	repository "earnforglance/server/repository/messages"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type MockSingleResultNewsLetterSubscription struct {
	mock.Mock
}

func (m *MockSingleResultNewsLetterSubscription) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.NewsLetterSubscription); ok {
		*v.(*domain.NewsLetterSubscription) = *result
	}
	return args.Error(1)
}

var mockItemNewsLetterSubscription = &domain.NewsLetterSubscription{
	ID:                         bson.NewObjectID(), // Existing ID of the record to update
	NewsLetterSubscriptionGuid: uuid.New(),
	Email:                      "updated_subscriber@example.com",
	Active:                     false,
	StoreID:                    bson.NewObjectID(),
	CreatedOnUtc:               time.Now().AddDate(0, 0, -7), // Created 7 days ago
	LanguageID:                 bson.NewObjectID(),
}

func TestNewsLetterSubscriptionRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionNewsLetterSubscription

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultNewsLetterSubscription{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemNewsLetterSubscription, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewNewsLetterSubscriptionRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemNewsLetterSubscription.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultNewsLetterSubscription{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewNewsLetterSubscriptionRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemNewsLetterSubscription.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestNewsLetterSubscriptionRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionNewsLetterSubscription

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemNewsLetterSubscription).Return(nil, nil).Once()

	repo := repository.NewNewsLetterSubscriptionRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemNewsLetterSubscription)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestNewsLetterSubscriptionRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionNewsLetterSubscription

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemNewsLetterSubscription.ID}
	update := bson.M{"$set": mockItemNewsLetterSubscription}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewNewsLetterSubscriptionRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemNewsLetterSubscription)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
