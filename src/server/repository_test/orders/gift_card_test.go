package repository_test

import (
	"context"
	"errors"
	"testing"
	"time"

	domain "earnforglance/server/domain/orders"
	"earnforglance/server/mongo/mocks"
	repository "earnforglance/server/repository/orders"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultGiftCard struct {
	mock.Mock
}

func (m *MockSingleResultGiftCard) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.GiftCard); ok {
		*v.(*domain.GiftCard) = *result
	}
	return args.Error(1)
}

var mockItemGiftCard = &domain.GiftCard{
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

func TestGiftCardRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionGiftCard

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultGiftCard{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemGiftCard, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewGiftCardRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemGiftCard.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultGiftCard{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewGiftCardRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemGiftCard.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestGiftCardRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionGiftCard

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemGiftCard).Return(nil, nil).Once()

	repo := repository.NewGiftCardRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemGiftCard)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestGiftCardRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionGiftCard

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemGiftCard.ID}
	update := bson.M{"$set": mockItemGiftCard}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewGiftCardRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemGiftCard)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
