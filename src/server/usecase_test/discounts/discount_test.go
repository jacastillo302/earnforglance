package usecase_test

import (
	"context"
	domain "earnforglance/server/domain/discounts"
	mocks "earnforglance/server/domain/mocks"
	test "earnforglance/server/usecase/discounts"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestDiscountUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.DiscountRepository)
	timeout := time.Duration(10)
	usecase := test.NewDiscountUsecase(mockRepo, timeout)

	discountID := primitive.NewObjectID().Hex()

	updatedDiscount := domain.Discount{
		ID:                        primitive.NewObjectID(), // Existing ID of the record to update
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

	mockRepo.On("FetchByID", mock.Anything, discountID).Return(updatedDiscount, nil)

	result, err := usecase.FetchByID(context.Background(), discountID)

	assert.NoError(t, err)
	assert.Equal(t, updatedDiscount, result)
	mockRepo.AssertExpectations(t)
}

func TestDiscountUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.DiscountRepository)
	timeout := time.Duration(10)
	usecase := test.NewDiscountUsecase(mockRepo, timeout)

	newDiscount := &domain.Discount{
		Name:                      "Black Friday Sale",
		AdminComment:              "Applies to all products",
		DiscountTypeID:            1,
		UsePercentage:             true,
		DiscountPercentage:        20.0,
		DiscountAmount:            0.0,
		MaximumDiscountAmount:     new(float64),
		StartDateUtc:              new(time.Time),
		EndDateUtc:                new(time.Time),
		RequiresCouponCode:        true,
		CouponCode:                "BLACKFRIDAY",
		IsCumulative:              false,
		DiscountLimitationID:      2,
		LimitationTimes:           1,
		MaximumDiscountedQuantity: new(int),
		AppliedToSubCategories:    true,
		IsActive:                  true,
		VendorID:                  nil,
	}
	*newDiscount.MaximumDiscountAmount = 50.0
	*newDiscount.StartDateUtc = time.Now()
	*newDiscount.EndDateUtc = time.Now().AddDate(0, 0, 7) // Ends in 7 days
	*newDiscount.MaximumDiscountedQuantity = 5

	mockRepo.On("Create", mock.Anything, newDiscount).Return(nil)

	err := usecase.Create(context.Background(), newDiscount)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDiscountUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.DiscountRepository)
	timeout := time.Duration(10)
	usecase := test.NewDiscountUsecase(mockRepo, timeout)

	updatedDiscount := &domain.Discount{
		ID:                        primitive.NewObjectID(), // Existing ID of the record to update
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
		DiscountLimitationID:      3,
		LimitationTimes:           3,
		MaximumDiscountedQuantity: nil,
		AppliedToSubCategories:    false,
		IsActive:                  false,
		VendorID:                  nil,
	}
	*updatedDiscount.StartDateUtc = time.Now().AddDate(0, 0, -7) // Started 7 days ago
	*updatedDiscount.EndDateUtc = time.Now().AddDate(0, 0, 1)    // Ends in 1 day

	mockRepo.On("Update", mock.Anything, updatedDiscount).Return(nil)

	err := usecase.Update(context.Background(), updatedDiscount)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDiscountUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.DiscountRepository)
	timeout := time.Duration(10)
	usecase := test.NewDiscountUsecase(mockRepo, timeout)

	discountID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, discountID).Return(nil)

	err := usecase.Delete(context.Background(), discountID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDiscountUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.DiscountRepository)
	timeout := time.Duration(10)
	usecase := test.NewDiscountUsecase(mockRepo, timeout)

	fetchedDiscounts := []domain.Discount{
		{
			ID:                        primitive.NewObjectID(),
			Name:                      "Black Friday Sale",
			AdminComment:              "Applies to all products",
			DiscountTypeID:            1,
			UsePercentage:             true,
			DiscountPercentage:        20.0,
			DiscountAmount:            0.0,
			MaximumDiscountAmount:     new(float64),
			StartDateUtc:              new(time.Time),
			EndDateUtc:                new(time.Time),
			RequiresCouponCode:        true,
			CouponCode:                "BLACKFRIDAY",
			IsCumulative:              false,
			DiscountLimitationID:      2,
			LimitationTimes:           1,
			MaximumDiscountedQuantity: new(int),
			AppliedToSubCategories:    true,
			IsActive:                  true,
			VendorID:                  nil,
		},
		{
			ID:                        primitive.NewObjectID(),
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
		},
	}
	*fetchedDiscounts[0].MaximumDiscountAmount = 50.0
	*fetchedDiscounts[0].StartDateUtc = time.Now().AddDate(0, 0, -10) // Started 10 days ago
	*fetchedDiscounts[0].EndDateUtc = time.Now().AddDate(0, 0, -3)    // Ended 3 days ago
	*fetchedDiscounts[0].MaximumDiscountedQuantity = 5
	*fetchedDiscounts[1].StartDateUtc = time.Now().AddDate(0, 0, -7) // Started 7 days ago
	*fetchedDiscounts[1].EndDateUtc = time.Now().AddDate(0, 0, 1)    // Ends in 1 day

	mockRepo.On("Fetch", mock.Anything).Return(fetchedDiscounts, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedDiscounts, result)
	mockRepo.AssertExpectations(t)
}
