package repository_test

import (
	"context"
	"errors"
	"testing"
	"time"

	domain "earnforglance/server/domain/discounts"
	repository "earnforglance/server/repository/discounts"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type MockSingleResultDiscount struct {
	mock.Mock
}

func (m *MockSingleResultDiscount) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.Discount); ok {
		*v.(*domain.Discount) = *result
	}
	return args.Error(1)
}

var mockItemDiscount = &domain.Discount{
	ID:                        bson.NewObjectID(), // Existing ID of the record to update
	Name:                      "Cyber Monday Sale",
	AdminComment:              "Limited to electronics",
	DiscountTypeID:            2,
	UsePercentage:             false,
	DiscountPercentage:        0.0,
	DiscountAmount:            30.0,
	MaximumDiscountAmount:     nil,
	StartDateUtc:              new(time.Time),
	EndDateUtc:                new(time.Time),
	RequiresCouponCode:        false,
	CouponCode:                "",
	IsCumulative:              true,
	DiscountLimitationID:      1,
	LimitationTimes:           3,
	MaximumDiscountedQuantity: nil,
	AppliedToSubCategories:    false,
	IsActive:                  false,
	VendorID:                  nil,
}

func TestDiscountRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionDiscount

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultDiscount{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemDiscount, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewDiscountRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemDiscount.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultDiscount{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewDiscountRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemDiscount.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestDiscountRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionDiscount

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemDiscount).Return(nil, nil).Once()

	repo := repository.NewDiscountRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemDiscount)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestDiscountRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionDiscount

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemDiscount.ID}
	update := bson.M{"$set": mockItemDiscount}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewDiscountRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemDiscount)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
