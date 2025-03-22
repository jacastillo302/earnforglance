package usecase

import (
	"context"
	mocks "earnforglance/server/domain/mocks"
	domain "earnforglance/server/domain/orders"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGiftCardUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.GiftCardRepository)
	timeout := time.Duration(10)
	usecase := NewGiftCardUsecase(mockRepo, timeout)

	giftCardID := primitive.NewObjectID().Hex()

	updatedGiftCard := domain.GiftCard{
		ID:                       primitive.NewObjectID(), // Existing ID of the record to update
		PurchasedWithOrderItemID: new(primitive.ObjectID),
		GiftCardTypeID:           2,
		Amount:                   150.00,
		IsGiftCardActivated:      false,
		GiftCardCouponCode:       "GIFT150",
		RecipientName:            "Alice Johnson",
		RecipientEmail:           "alice.johnson@example.com",
		SenderName:               "Bob Brown",
		SenderEmail:              "bob.brown@example.com",
		Message:                  "Congratulations!",
		IsRecipientNotified:      true,
		CreatedOnUtc:             time.Now().AddDate(0, 0, -7), // Created 7 days ago
		GiftCardType:             2,
	}

	mockRepo.On("FetchByID", mock.Anything, giftCardID).Return(updatedGiftCard, nil)

	result, err := usecase.FetchByID(context.Background(), giftCardID)

	assert.NoError(t, err)
	assert.Equal(t, updatedGiftCard, result)
	mockRepo.AssertExpectations(t)
}

func TestGiftCardUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.GiftCardRepository)
	timeout := time.Duration(10)
	usecase := NewGiftCardUsecase(mockRepo, timeout)

	newGiftCard := &domain.GiftCard{
		PurchasedWithOrderItemID: nil,
		GiftCardTypeID:           1,
		Amount:                   100.00,
		IsGiftCardActivated:      true,
		GiftCardCouponCode:       "GIFT100",
		RecipientName:            "John Doe",
		RecipientEmail:           "john.doe@example.com",
		SenderName:               "Jane Smith",
		SenderEmail:              "jane.smith@example.com",
		Message:                  "Happy Birthday!",
		IsRecipientNotified:      false,
		CreatedOnUtc:             time.Now(),
		GiftCardType:             1,
	}

	mockRepo.On("Create", mock.Anything, newGiftCard).Return(nil)

	err := usecase.Create(context.Background(), newGiftCard)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGiftCardUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.GiftCardRepository)
	timeout := time.Duration(10)
	usecase := NewGiftCardUsecase(mockRepo, timeout)

	updatedGiftCard := &domain.GiftCard{
		ID:                       primitive.NewObjectID(), // Existing ID of the record to update
		PurchasedWithOrderItemID: new(primitive.ObjectID),
		GiftCardTypeID:           2,
		Amount:                   150.00,
		IsGiftCardActivated:      false,
		GiftCardCouponCode:       "GIFT150",
		RecipientName:            "Alice Johnson",
		RecipientEmail:           "alice.johnson@example.com",
		SenderName:               "Bob Brown",
		SenderEmail:              "bob.brown@example.com",
		Message:                  "Congratulations!",
		IsRecipientNotified:      true,
		CreatedOnUtc:             time.Now().AddDate(0, 0, -7), // Created 7 days ago
		GiftCardType:             2,
	}
	*updatedGiftCard.PurchasedWithOrderItemID = primitive.NewObjectID()

	mockRepo.On("Update", mock.Anything, updatedGiftCard).Return(nil)

	err := usecase.Update(context.Background(), updatedGiftCard)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGiftCardUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.GiftCardRepository)
	timeout := time.Duration(10)
	usecase := NewGiftCardUsecase(mockRepo, timeout)

	giftCardID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, giftCardID).Return(nil)

	err := usecase.Delete(context.Background(), giftCardID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGiftCardUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.GiftCardRepository)
	timeout := time.Duration(10)
	usecase := NewGiftCardUsecase(mockRepo, timeout)

	fetchedGiftCards := []domain.GiftCard{
		{
			ID:                       primitive.NewObjectID(),
			PurchasedWithOrderItemID: nil,
			GiftCardTypeID:           1,
			Amount:                   100.00,
			IsGiftCardActivated:      true,
			GiftCardCouponCode:       "GIFT100",
			RecipientName:            "John Doe",
			RecipientEmail:           "john.doe@example.com",
			SenderName:               "Jane Smith",
			SenderEmail:              "jane.smith@example.com",
			Message:                  "Happy Birthday!",
			IsRecipientNotified:      false,
			CreatedOnUtc:             time.Now().AddDate(0, 0, -10), // Created 10 days ago
			GiftCardType:             1,
		},
		{
			ID:                       primitive.NewObjectID(),
			PurchasedWithOrderItemID: new(primitive.ObjectID),
			GiftCardTypeID:           2,
			Amount:                   150.00,
			IsGiftCardActivated:      false,
			GiftCardCouponCode:       "GIFT150",
			RecipientName:            "Alice Johnson",
			RecipientEmail:           "alice.johnson@example.com",
			SenderName:               "Bob Brown",
			SenderEmail:              "bob.brown@example.com",
			Message:                  "Congratulations!",
			IsRecipientNotified:      true,
			CreatedOnUtc:             time.Now().AddDate(0, 0, -5), // Created 5 days ago
			GiftCardType:             2,
		},
	}
	*fetchedGiftCards[1].PurchasedWithOrderItemID = primitive.NewObjectID()

	mockRepo.On("Fetch", mock.Anything).Return(fetchedGiftCards, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedGiftCards, result)
	mockRepo.AssertExpectations(t)
}
