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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestTierPriceUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.TierPriceRepository)
	timeout := time.Duration(10)
	usecase := test.NewTierPriceUsecase(mockRepo, timeout)

	tierPriceID := primitive.NewObjectID().Hex()

	updatedTierPrice := domain.TierPrice{
		ID:               primitive.NewObjectID(), // Existing ID of the record to update
		ProductID:        primitive.NewObjectID(),
		StoreID:          primitive.NewObjectID(),
		CustomerRoleID:   nil,
		Quantity:         20,
		Price:            39.99,
		StartDateTimeUtc: new(time.Time),
		EndDateTimeUtc:   new(time.Time),
	}

	mockRepo.On("FetchByID", mock.Anything, tierPriceID).Return(updatedTierPrice, nil)

	result, err := usecase.FetchByID(context.Background(), tierPriceID)

	assert.NoError(t, err)
	assert.Equal(t, updatedTierPrice, result)
	mockRepo.AssertExpectations(t)
}

func TestTierPriceUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.TierPriceRepository)
	timeout := time.Duration(10)
	usecase := test.NewTierPriceUsecase(mockRepo, timeout)

	newTierPrice := &domain.TierPrice{
		ID:               primitive.NewObjectID(), // Existing ID of the record to update
		ProductID:        primitive.NewObjectID(),
		StoreID:          primitive.NewObjectID(),
		CustomerRoleID:   nil,
		Quantity:         20,
		Price:            39.99,
		StartDateTimeUtc: new(time.Time),
		EndDateTimeUtc:   new(time.Time),
	}

	mockRepo.On("Create", mock.Anything, newTierPrice).Return(nil)

	err := usecase.Create(context.Background(), newTierPrice)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestTierPriceUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.TierPriceRepository)
	timeout := time.Duration(10)
	usecase := test.NewTierPriceUsecase(mockRepo, timeout)

	updatedTierPrice := &domain.TierPrice{
		ID:               primitive.NewObjectID(), // Existing ID of the record to update
		ProductID:        primitive.NewObjectID(),
		StoreID:          primitive.NewObjectID(),
		CustomerRoleID:   nil,
		Quantity:         20,
		Price:            39.99,
		StartDateTimeUtc: new(time.Time),
		EndDateTimeUtc:   new(time.Time),
	}

	*updatedTierPrice.CustomerRoleID = primitive.NewObjectID()
	*updatedTierPrice.StartDateTimeUtc = time.Now().AddDate(0, 0, -7) // 7 days ago
	*updatedTierPrice.EndDateTimeUtc = time.Now().AddDate(0, 0, 7)    // 7 days from now

	mockRepo.On("Update", mock.Anything, updatedTierPrice).Return(nil)

	err := usecase.Update(context.Background(), updatedTierPrice)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestTierPriceUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.TierPriceRepository)
	timeout := time.Duration(10)
	usecase := test.NewTierPriceUsecase(mockRepo, timeout)

	tierPriceID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, tierPriceID).Return(nil)

	err := usecase.Delete(context.Background(), tierPriceID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestTierPriceUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.TierPriceRepository)
	timeout := time.Duration(10)
	usecase := test.NewTierPriceUsecase(mockRepo, timeout)

	fetchedTierPrices := []domain.TierPrice{
		{
			ID:               primitive.NewObjectID(),
			ProductID:        primitive.NewObjectID(),
			StoreID:          primitive.NewObjectID(),
			CustomerRoleID:   nil,
			Quantity:         10,
			Price:            49.99,
			StartDateTimeUtc: nil,
			EndDateTimeUtc:   nil,
		},
		{
			ID:               primitive.NewObjectID(),
			ProductID:        primitive.NewObjectID(),
			StoreID:          primitive.NewObjectID(),
			CustomerRoleID:   nil,
			Quantity:         20,
			Price:            39.99,
			StartDateTimeUtc: new(time.Time),
			EndDateTimeUtc:   new(time.Time),
		},
	}
	*fetchedTierPrices[1].CustomerRoleID = primitive.NewObjectID()
	*fetchedTierPrices[1].StartDateTimeUtc = time.Now().AddDate(0, 0, -7) // 7 days ago
	*fetchedTierPrices[1].EndDateTimeUtc = time.Now().AddDate(0, 0, 7)    // 7 days from now

	mockRepo.On("Fetch", mock.Anything).Return(fetchedTierPrices, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedTierPrices, result)
	mockRepo.AssertExpectations(t)
}
