package usecase_test

import (
	"context"
	domain "earnforglance/server/domain/catalog"
	mocks "earnforglance/server/domain/mocks"
	test "earnforglance/server/usecase/catalog"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestBackInStockSubscriptionUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.BackInStockSubscriptionRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewBackInStockSubscriptionUsecase(mockRepo, timeout)

	subscriptionID := bson.NewObjectID().Hex()

	expectedSubscription := domain.BackInStockSubscription{
		ID:           bson.NewObjectID(), // Existing ID of the record to update
		StoreID:      bson.NewObjectID(),
		ProductID:    bson.NewObjectID(),
		CustomerID:   bson.NewObjectID(),
		CreatedOnUtc: time.Now(),
	}

	mockRepo.On("FetchByID", mock.Anything, subscriptionID).Return(expectedSubscription, nil)

	result, err := usecase.FetchByID(context.Background(), subscriptionID)

	assert.NoError(t, err)
	assert.Equal(t, expectedSubscription, result)
	mockRepo.AssertExpectations(t)
}

func TestBackInStockSubscriptionUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.BackInStockSubscriptionRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewBackInStockSubscriptionUsecase(mockRepo, timeout)

	newSubscription := &domain.BackInStockSubscription{
		ID:           bson.NewObjectID(), // Existing ID of the record to update
		StoreID:      bson.NewObjectID(),
		ProductID:    bson.NewObjectID(),
		CustomerID:   bson.NewObjectID(),
		CreatedOnUtc: time.Now(),
	}

	mockRepo.On("Create", mock.Anything, newSubscription).Return(nil)

	err := usecase.Create(context.Background(), newSubscription)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestBackInStockSubscriptionUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.BackInStockSubscriptionRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewBackInStockSubscriptionUsecase(mockRepo, timeout)

	updatedSubscription := &domain.BackInStockSubscription{
		ID:           bson.NewObjectID(), // Existing ID of the record to update
		StoreID:      bson.NewObjectID(),
		ProductID:    bson.NewObjectID(),
		CustomerID:   bson.NewObjectID(),
		CreatedOnUtc: time.Now(),
	}

	mockRepo.On("Update", mock.Anything, updatedSubscription).Return(nil)

	err := usecase.Update(context.Background(), updatedSubscription)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestBackInStockSubscriptionUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.BackInStockSubscriptionRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewBackInStockSubscriptionUsecase(mockRepo, timeout)

	subscriptionID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, subscriptionID).Return(nil)

	err := usecase.Delete(context.Background(), subscriptionID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestBackInStockSubscriptionUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.BackInStockSubscriptionRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewBackInStockSubscriptionUsecase(mockRepo, timeout)

	expectedSubscriptions := []domain.BackInStockSubscription{
		{
			ID:           bson.NewObjectID(),
			StoreID:      bson.NewObjectID(),
			ProductID:    bson.NewObjectID(),
			CustomerID:   bson.NewObjectID(),
			CreatedOnUtc: time.Now().AddDate(0, 0, -10), // 10 days ago
		},
		{
			ID:           bson.NewObjectID(),
			StoreID:      bson.NewObjectID(),
			ProductID:    bson.NewObjectID(),
			CustomerID:   bson.NewObjectID(),
			CreatedOnUtc: time.Now().AddDate(0, 0, -5), // 5 days ago
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(expectedSubscriptions, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, expectedSubscriptions, result)
	mockRepo.AssertExpectations(t)
}
