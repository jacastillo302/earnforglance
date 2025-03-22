package usecase

import (
	"context"
	"testing"
	"time"

	domain "earnforglance/server/domain/customers"
	mocks "earnforglance/server/domain/mocks"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestRewardPointsHistoryUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.RewardPointsHistoryRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewRewardPointsHistoryUsecase(mockRepo, timeout)

	rewardPointsHistoryID := primitive.NewObjectID().Hex()

	updatedRewardPointsHistory := domain.RewardPointsHistory{
		ID:            primitive.NewObjectID(), // Existing ID of the record to update
		CustomerID:    primitive.NewObjectID(),
		StoreID:       primitive.NewObjectID(),
		Points:        -50,
		PointsBalance: new(int),
		UsedAmount:    25.0,
		Message:       "Reward points used for order",
		CreatedOnUtc:  time.Now().AddDate(0, 0, -7), // Created 7 days ago
		EndDateUtc:    new(time.Time),
		ValidPoints:   new(int),
		UsedWithOrder: new(uuid.UUID),
	}

	mockRepo.On("FetchByID", mock.Anything, rewardPointsHistoryID).Return(updatedRewardPointsHistory, nil)

	result, err := usecase.FetchByID(context.Background(), rewardPointsHistoryID)

	assert.NoError(t, err)
	assert.Equal(t, updatedRewardPointsHistory, result)
	mockRepo.AssertExpectations(t)
}

func TestRewardPointsHistoryUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.RewardPointsHistoryRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewRewardPointsHistoryUsecase(mockRepo, timeout)

	newRewardPointsHistory := &domain.RewardPointsHistory{
		CustomerID:    primitive.NewObjectID(),
		StoreID:       primitive.NewObjectID(),
		Points:        100,
		PointsBalance: new(int),
		UsedAmount:    0.0,
		Message:       "Reward points added for purchase",
		CreatedOnUtc:  time.Now(),
		EndDateUtc:    nil,
		ValidPoints:   new(int),
		UsedWithOrder: nil,
	}
	*newRewardPointsHistory.PointsBalance = 100
	*newRewardPointsHistory.ValidPoints = 100

	mockRepo.On("Create", mock.Anything, newRewardPointsHistory).Return(nil)

	err := usecase.Create(context.Background(), newRewardPointsHistory)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestRewardPointsHistoryUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.RewardPointsHistoryRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewRewardPointsHistoryUsecase(mockRepo, timeout)

	updatedRewardPointsHistory := &domain.RewardPointsHistory{
		ID:            primitive.NewObjectID(), // Existing ID of the record to update
		CustomerID:    primitive.NewObjectID(),
		StoreID:       primitive.NewObjectID(),
		Points:        -50,
		PointsBalance: new(int),
		UsedAmount:    25.0,
		Message:       "Reward points used for order",
		CreatedOnUtc:  time.Now().AddDate(0, 0, -7), // Created 7 days ago
		EndDateUtc:    new(time.Time),
		ValidPoints:   new(int),
		UsedWithOrder: new(uuid.UUID),
	}
	*updatedRewardPointsHistory.PointsBalance = 50
	*updatedRewardPointsHistory.ValidPoints = 50
	*updatedRewardPointsHistory.EndDateUtc = time.Now().AddDate(0, 0, 30) // Expires in 30 days
	*updatedRewardPointsHistory.UsedWithOrder = uuid.New()

	mockRepo.On("Update", mock.Anything, updatedRewardPointsHistory).Return(nil)

	err := usecase.Update(context.Background(), updatedRewardPointsHistory)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestRewardPointsHistoryUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.RewardPointsHistoryRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewRewardPointsHistoryUsecase(mockRepo, timeout)

	rewardPointsHistoryID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, rewardPointsHistoryID).Return(nil)

	err := usecase.Delete(context.Background(), rewardPointsHistoryID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestRewardPointsHistoryUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.RewardPointsHistoryRepository)
	timeout := time.Duration(10) * time.Second
	usecase := NewRewardPointsHistoryUsecase(mockRepo, timeout)

	fetchedRewardPointsHistories := []domain.RewardPointsHistory{
		{
			ID:            primitive.NewObjectID(),
			CustomerID:    primitive.NewObjectID(),
			StoreID:       primitive.NewObjectID(),
			Points:        100,
			PointsBalance: new(int),
			UsedAmount:    0.0,
			Message:       "Reward points added for purchase",
			CreatedOnUtc:  time.Now().AddDate(0, 0, -10), // Created 10 days ago
			EndDateUtc:    nil,
			ValidPoints:   new(int),
			UsedWithOrder: nil,
		},
		{
			ID:            primitive.NewObjectID(),
			CustomerID:    primitive.NewObjectID(),
			StoreID:       primitive.NewObjectID(),
			Points:        -50,
			PointsBalance: new(int),
			UsedAmount:    25.0,
			Message:       "Reward points used for order",
			CreatedOnUtc:  time.Now().AddDate(0, 0, -7), // Created 7 days ago
			EndDateUtc:    new(time.Time),
			ValidPoints:   new(int),
			UsedWithOrder: new(uuid.UUID),
		},
	}
	*fetchedRewardPointsHistories[0].PointsBalance = 100
	*fetchedRewardPointsHistories[0].ValidPoints = 100
	*fetchedRewardPointsHistories[1].PointsBalance = 50
	*fetchedRewardPointsHistories[1].ValidPoints = 50
	*fetchedRewardPointsHistories[1].EndDateUtc = time.Now().AddDate(0, 0, 30) // Expires in 30 days
	*fetchedRewardPointsHistories[1].UsedWithOrder = uuid.New()

	mockRepo.On("Fetch", mock.Anything).Return(fetchedRewardPointsHistories, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedRewardPointsHistories, result)
	mockRepo.AssertExpectations(t)
}
