package usecase_test

import (
	"context"
	domain "earnforglance/server/domain/messages"
	mocks "earnforglance/server/domain/mocks"
	test "earnforglance/server/usecase/messages"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestNewsLetterSubscriptionUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.NewsLetterSubscriptionRepository)
	timeout := time.Duration(10)
	usecase := test.NewNewsLetterSubscriptionUsecase(mockRepo, timeout)

	newsLetterSubscriptionID := bson.NewObjectID().Hex()

	updatedNewsLetterSubscription := domain.NewsLetterSubscription{
		ID:           bson.NewObjectID(), // Existing ID of the record to update
		Guid:         uuid.New().String(),
		Email:        "updated_subscriber@example.com",
		Active:       false,
		StoreID:      bson.NewObjectID(),
		CreatedOnUtc: time.Now().AddDate(0, 0, -7), // Created 7 days ago
		LanguageID:   bson.NewObjectID(),
	}

	mockRepo.On("FetchByID", mock.Anything, newsLetterSubscriptionID).Return(updatedNewsLetterSubscription, nil)

	result, err := usecase.FetchByID(context.Background(), newsLetterSubscriptionID)

	assert.NoError(t, err)
	assert.Equal(t, updatedNewsLetterSubscription, result)
	mockRepo.AssertExpectations(t)
}

func TestNewsLetterSubscriptionUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.NewsLetterSubscriptionRepository)
	timeout := time.Duration(10)
	usecase := test.NewNewsLetterSubscriptionUsecase(mockRepo, timeout)

	newNewsLetterSubscription := &domain.NewsLetterSubscription{
		Guid:         uuid.New().String(),
		Email:        "subscriber@example.com",
		Active:       true,
		StoreID:      bson.NewObjectID(),
		CreatedOnUtc: time.Now(),
		LanguageID:   bson.NewObjectID(),
	}

	mockRepo.On("Create", mock.Anything, newNewsLetterSubscription).Return(nil)

	err := usecase.Create(context.Background(), newNewsLetterSubscription)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestNewsLetterSubscriptionUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.NewsLetterSubscriptionRepository)
	timeout := time.Duration(10)
	usecase := test.NewNewsLetterSubscriptionUsecase(mockRepo, timeout)

	updatedNewsLetterSubscription := &domain.NewsLetterSubscription{
		ID:           bson.NewObjectID(), // Existing ID of the record to update
		Guid:         uuid.New().String(),
		Email:        "updated_subscriber@example.com",
		Active:       false,
		StoreID:      bson.NewObjectID(),
		CreatedOnUtc: time.Now().AddDate(0, 0, -7), // Created 7 days ago
		LanguageID:   bson.NewObjectID(),
	}

	mockRepo.On("Update", mock.Anything, updatedNewsLetterSubscription).Return(nil)

	err := usecase.Update(context.Background(), updatedNewsLetterSubscription)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestNewsLetterSubscriptionUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.NewsLetterSubscriptionRepository)
	timeout := time.Duration(10)
	usecase := test.NewNewsLetterSubscriptionUsecase(mockRepo, timeout)

	newsLetterSubscriptionID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, newsLetterSubscriptionID).Return(nil)

	err := usecase.Delete(context.Background(), newsLetterSubscriptionID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestNewsLetterSubscriptionUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.NewsLetterSubscriptionRepository)
	timeout := time.Duration(10)
	usecase := test.NewNewsLetterSubscriptionUsecase(mockRepo, timeout)

	fetchedNewsLetterSubscriptions := []domain.NewsLetterSubscription{
		{
			ID:           bson.NewObjectID(),
			Guid:         uuid.New().String(),
			Email:        "subscriber1@example.com",
			Active:       true,
			StoreID:      bson.NewObjectID(),
			CreatedOnUtc: time.Now().AddDate(0, 0, -10), // Created 10 days ago
			LanguageID:   bson.NewObjectID(),
		},
		{
			ID:           bson.NewObjectID(),
			Guid:         uuid.New().String(),
			Email:        "subscriber2@example.com",
			Active:       false,
			StoreID:      bson.NewObjectID(),
			CreatedOnUtc: time.Now().AddDate(0, 0, -5), // Created 5 days ago
			LanguageID:   bson.NewObjectID(),
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedNewsLetterSubscriptions, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedNewsLetterSubscriptions, result)
	mockRepo.AssertExpectations(t)
}
