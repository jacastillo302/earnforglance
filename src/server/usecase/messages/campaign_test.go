package usecase

import (
	"context"
	domain "earnforglance/server/domain/messages"
	mocks "earnforglance/server/domain/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCampaignUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.CampaignRepository)
	timeout := time.Duration(10)
	usecase := NewCampaignUsecase(mockRepo, timeout)

	campaignID := primitive.NewObjectID().Hex()

	updatedCampaign := domain.Campaign{
		ID:                    primitive.NewObjectID(), // Existing ID of the record to update
		Name:                  "Updated Holiday Sale",
		Subject:               "Updated Discounts for the Holidays!",
		Body:                  "Enjoy up to 60% off on selected items. Offer extended!",
		StoreID:               primitive.NewObjectID(),
		CustomerRoleID:        primitive.NewObjectID(),
		CreatedOnUtc:          time.Now().AddDate(0, 0, -7), // Created 7 days ago
		DontSendBeforeDateUtc: new(time.Time),
	}
	*updatedCampaign.DontSendBeforeDateUtc = time.Now().AddDate(0, 0, 1) // Don't send before tomorrow

	mockRepo.On("FetchByID", mock.Anything, campaignID).Return(updatedCampaign, nil)

	result, err := usecase.FetchByID(context.Background(), campaignID)

	assert.NoError(t, err)
	assert.Equal(t, updatedCampaign, result)
	mockRepo.AssertExpectations(t)
}

func TestCampaignUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.CampaignRepository)
	timeout := time.Duration(10)
	usecase := NewCampaignUsecase(mockRepo, timeout)

	newCampaign := &domain.Campaign{
		Name:                  "Holiday Sale",
		Subject:               "Exclusive Holiday Discounts!",
		Body:                  "Enjoy up to 50% off on selected items. Limited time offer!",
		StoreID:               primitive.NewObjectID(),
		CustomerRoleID:        primitive.NewObjectID(),
		CreatedOnUtc:          time.Now(),
		DontSendBeforeDateUtc: nil,
	}

	mockRepo.On("Create", mock.Anything, newCampaign).Return(nil)

	err := usecase.Create(context.Background(), newCampaign)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCampaignUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.CampaignRepository)
	timeout := time.Duration(10)
	usecase := NewCampaignUsecase(mockRepo, timeout)

	updatedCampaign := &domain.Campaign{
		ID:                    primitive.NewObjectID(), // Existing ID of the record to update
		Name:                  "Updated Holiday Sale",
		Subject:               "Updated Discounts for the Holidays!",
		Body:                  "Enjoy up to 60% off on selected items. Offer extended!",
		StoreID:               primitive.NewObjectID(),
		CustomerRoleID:        primitive.NewObjectID(),
		CreatedOnUtc:          time.Now().AddDate(0, 0, -7), // Created 7 days ago
		DontSendBeforeDateUtc: new(time.Time),
	}
	*updatedCampaign.DontSendBeforeDateUtc = time.Now().AddDate(0, 0, 1) // Don't send before tomorrow

	mockRepo.On("Update", mock.Anything, updatedCampaign).Return(nil)

	err := usecase.Update(context.Background(), updatedCampaign)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCampaignUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.CampaignRepository)
	timeout := time.Duration(10)
	usecase := NewCampaignUsecase(mockRepo, timeout)

	campaignID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, campaignID).Return(nil)

	err := usecase.Delete(context.Background(), campaignID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCampaignUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.CampaignRepository)
	timeout := time.Duration(10)
	usecase := NewCampaignUsecase(mockRepo, timeout)

	fetchedCampaigns := []domain.Campaign{
		{
			ID:                    primitive.NewObjectID(),
			Name:                  "Holiday Sale",
			Subject:               "Exclusive Holiday Discounts!",
			Body:                  "Enjoy up to 50% off on selected items. Limited time offer!",
			StoreID:               primitive.NewObjectID(),
			CustomerRoleID:        primitive.NewObjectID(),
			CreatedOnUtc:          time.Now().AddDate(0, 0, -10), // Created 10 days ago
			DontSendBeforeDateUtc: nil,
		},
		{
			ID:                    primitive.NewObjectID(),
			Name:                  "Spring Sale",
			Subject:               "Fresh Deals for Spring!",
			Body:                  "Get ready for spring with our exclusive discounts.",
			StoreID:               primitive.NewObjectID(),
			CustomerRoleID:        primitive.NewObjectID(),
			CreatedOnUtc:          time.Now().AddDate(0, 0, -5), // Created 5 days ago
			DontSendBeforeDateUtc: new(time.Time),
		},
	}
	*fetchedCampaigns[1].DontSendBeforeDateUtc = time.Now().AddDate(0, 0, 2) // Don't send before 2 days from now

	mockRepo.On("Fetch", mock.Anything).Return(fetchedCampaigns, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedCampaigns, result)
	mockRepo.AssertExpectations(t)
}
