package usecase

import (
	"context"
	mocks "earnforglance/server/domain/mocks"
	domain "earnforglance/server/domain/shipping"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestShippingOptionUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.ShippingOptionRepository)
	timeout := time.Duration(10)
	usecase := NewShippingOptionUsecase(mockRepo, timeout)

	shippingOptionID := primitive.NewObjectID().Hex()

	updatedShippingOption := domain.ShippingOption{
		ID:                                      primitive.NewObjectID(), // Existing ID of the record to update
		ShippingRateComputationMethodSystemName: "ExpressRate",
		Rate:                                    20.00,
		Name:                                    "Express Shipping",
		Description:                             "Delivery within 1-2 business days.",
		TransitDays:                             new(int),
		IsPickupInStore:                         true,
		DisplayOrder:                            new(int),
	}

	mockRepo.On("FetchByID", mock.Anything, shippingOptionID).Return(updatedShippingOption, nil)

	result, err := usecase.FetchByID(context.Background(), shippingOptionID)

	assert.NoError(t, err)
	assert.Equal(t, updatedShippingOption, result)
	mockRepo.AssertExpectations(t)
}

func TestShippingOptionUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.ShippingOptionRepository)
	timeout := time.Duration(10)
	usecase := NewShippingOptionUsecase(mockRepo, timeout)
	newShippingOption := &domain.ShippingOption{
		ShippingRateComputationMethodSystemName: "FlatRate",
		Rate:                                    10.00,
		Name:                                    "Standard Shipping",
		Description:                             "Delivery within 5-7 business days.",
		TransitDays:                             new(int),
		IsPickupInStore:                         false,
		DisplayOrder:                            new(int),
	}
	*newShippingOption.TransitDays = 5
	*newShippingOption.DisplayOrder = 1

	mockRepo.On("Create", mock.Anything, newShippingOption).Return(nil)

	err := usecase.Create(context.Background(), newShippingOption)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestShippingOptionUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.ShippingOptionRepository)
	timeout := time.Duration(10)
	usecase := NewShippingOptionUsecase(mockRepo, timeout)

	updatedShippingOption := &domain.ShippingOption{
		ID:                                      primitive.NewObjectID(), // Existing ID of the record to update
		ShippingRateComputationMethodSystemName: "ExpressRate",
		Rate:                                    20.00,
		Name:                                    "Express Shipping",
		Description:                             "Delivery within 1-2 business days.",
		TransitDays:                             new(int),
		IsPickupInStore:                         true,
		DisplayOrder:                            new(int),
	}
	*updatedShippingOption.TransitDays = 2
	*updatedShippingOption.DisplayOrder = 2

	mockRepo.On("Update", mock.Anything, updatedShippingOption).Return(nil)

	err := usecase.Update(context.Background(), updatedShippingOption)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestShippingOptionUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.ShippingOptionRepository)
	timeout := time.Duration(10)
	usecase := NewShippingOptionUsecase(mockRepo, timeout)

	shippingOptionID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, shippingOptionID).Return(nil)

	err := usecase.Delete(context.Background(), shippingOptionID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestShippingOptionUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.ShippingOptionRepository)
	timeout := time.Duration(10)
	usecase := NewShippingOptionUsecase(mockRepo, timeout)

	fetchedShippingOptions := []domain.ShippingOption{
		{
			ID:                                      primitive.NewObjectID(),
			ShippingRateComputationMethodSystemName: "FlatRate",
			Rate:                                    10.00,
			Name:                                    "Standard Shipping",
			Description:                             "Delivery within 5-7 business days.",
			TransitDays:                             new(int),
			IsPickupInStore:                         false,
			DisplayOrder:                            new(int),
		},
		{
			ID:                                      primitive.NewObjectID(),
			ShippingRateComputationMethodSystemName: "ExpressRate",
			Rate:                                    20.00,
			Name:                                    "Express Shipping",
			Description:                             "Delivery within 1-2 business days.",
			TransitDays:                             new(int),
			IsPickupInStore:                         true,
			DisplayOrder:                            new(int),
		},
	}
	*fetchedShippingOptions[0].TransitDays = 5
	*fetchedShippingOptions[0].DisplayOrder = 1
	*fetchedShippingOptions[1].TransitDays = 2
	*fetchedShippingOptions[1].DisplayOrder = 2

	mockRepo.On("Fetch", mock.Anything).Return(fetchedShippingOptions, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedShippingOptions, result)
	mockRepo.AssertExpectations(t)
}
