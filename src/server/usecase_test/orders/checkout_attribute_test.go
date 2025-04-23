package usecase_test

import (
	"context"
	mocks "earnforglance/server/domain/mocks"
	domain "earnforglance/server/domain/orders"
	test "earnforglance/server/usecase/orders"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestCheckoutAttributeUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.CheckoutAttributeRepository)
	timeout := time.Duration(10)
	usecase := test.NewCheckoutAttributeUsecase(mockRepo, timeout)

	checkoutAttributeID := bson.NewObjectID().Hex()

	updatedCheckoutAttribute := domain.CheckoutAttribute{
		ID:                              bson.NewObjectID(), // Existing ID of the record to update
		TextPrompt:                      "Update your custom message",
		ShippableProductRequired:        false,
		IsTaxExempt:                     true,
		TaxCategoryID:                   bson.NewObjectID(),
		LimitedToStores:                 true,
		ValidationMinLength:             new(int),
		ValidationMaxLength:             new(int),
		ValidationFileAllowedExtensions: ".pdf,.docx",
		ValidationFileMaximumSize:       new(int),
		DefaultValue:                    "Updated Default Message",
		ConditionAttributeXml:           "<UpdatedConditions></UpdatedConditions>",
	}

	mockRepo.On("FetchByID", mock.Anything, checkoutAttributeID).Return(updatedCheckoutAttribute, nil)

	result, err := usecase.FetchByID(context.Background(), checkoutAttributeID)

	assert.NoError(t, err)
	assert.Equal(t, updatedCheckoutAttribute, result)
	mockRepo.AssertExpectations(t)
}

func TestCheckoutAttributeUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.CheckoutAttributeRepository)
	timeout := time.Duration(10)
	usecase := test.NewCheckoutAttributeUsecase(mockRepo, timeout)

	newCheckoutAttribute := &domain.CheckoutAttribute{
		TextPrompt:                      "Enter your custom message",
		ShippableProductRequired:        true,
		IsTaxExempt:                     false,
		TaxCategoryID:                   bson.NewObjectID(),
		LimitedToStores:                 false,
		ValidationMinLength:             new(int),
		ValidationMaxLength:             new(int),
		ValidationFileAllowedExtensions: ".jpg,.png",
		ValidationFileMaximumSize:       new(int),
		DefaultValue:                    "Default Message",
		ConditionAttributeXml:           "<Conditions></Conditions>",
	}
	*newCheckoutAttribute.ValidationMinLength = 5
	*newCheckoutAttribute.ValidationMaxLength = 100
	*newCheckoutAttribute.ValidationFileMaximumSize = 2048

	mockRepo.On("Create", mock.Anything, newCheckoutAttribute).Return(nil)

	err := usecase.Create(context.Background(), newCheckoutAttribute)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCheckoutAttributeUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.CheckoutAttributeRepository)
	timeout := time.Duration(10)
	usecase := test.NewCheckoutAttributeUsecase(mockRepo, timeout)

	updatedCheckoutAttribute := &domain.CheckoutAttribute{
		ID:                              bson.NewObjectID(), // Existing ID of the record to update
		TextPrompt:                      "Update your custom message",
		ShippableProductRequired:        false,
		IsTaxExempt:                     true,
		TaxCategoryID:                   bson.NewObjectID(),
		LimitedToStores:                 true,
		ValidationMinLength:             new(int),
		ValidationMaxLength:             new(int),
		ValidationFileAllowedExtensions: ".pdf,.docx",
		ValidationFileMaximumSize:       new(int),
		DefaultValue:                    "Updated Default Message",
		ConditionAttributeXml:           "<UpdatedConditions></UpdatedConditions>",
	}
	*updatedCheckoutAttribute.ValidationMinLength = 10
	*updatedCheckoutAttribute.ValidationMaxLength = 200
	*updatedCheckoutAttribute.ValidationFileMaximumSize = 4096

	mockRepo.On("Update", mock.Anything, updatedCheckoutAttribute).Return(nil)

	err := usecase.Update(context.Background(), updatedCheckoutAttribute)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCheckoutAttributeUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.CheckoutAttributeRepository)
	timeout := time.Duration(10)
	usecase := test.NewCheckoutAttributeUsecase(mockRepo, timeout)

	checkoutAttributeID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, checkoutAttributeID).Return(nil)

	err := usecase.Delete(context.Background(), checkoutAttributeID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCheckoutAttributeUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.CheckoutAttributeRepository)
	timeout := time.Duration(10)
	usecase := test.NewCheckoutAttributeUsecase(mockRepo, timeout)

	fetchedCheckoutAttributes := []domain.CheckoutAttribute{
		{
			ID:                              bson.NewObjectID(),
			TextPrompt:                      "Enter your custom message",
			ShippableProductRequired:        true,
			IsTaxExempt:                     false,
			TaxCategoryID:                   bson.NewObjectID(),
			LimitedToStores:                 false,
			ValidationMinLength:             new(int),
			ValidationMaxLength:             new(int),
			ValidationFileAllowedExtensions: ".jpg,.png",
			ValidationFileMaximumSize:       new(int),
			DefaultValue:                    "Default Message",
			ConditionAttributeXml:           "<Conditions></Conditions>",
		},
		{
			ID:                              bson.NewObjectID(),
			TextPrompt:                      "Update your custom message",
			ShippableProductRequired:        false,
			IsTaxExempt:                     true,
			TaxCategoryID:                   bson.NewObjectID(),
			LimitedToStores:                 true,
			ValidationMinLength:             new(int),
			ValidationMaxLength:             new(int),
			ValidationFileAllowedExtensions: ".pdf,.docx",
			ValidationFileMaximumSize:       new(int),
			DefaultValue:                    "Updated Default Message",
			ConditionAttributeXml:           "<UpdatedConditions></UpdatedConditions>",
		},
	}
	*fetchedCheckoutAttributes[0].ValidationMinLength = 5
	*fetchedCheckoutAttributes[0].ValidationMaxLength = 100
	*fetchedCheckoutAttributes[0].ValidationFileMaximumSize = 2048
	*fetchedCheckoutAttributes[1].ValidationMinLength = 10
	*fetchedCheckoutAttributes[1].ValidationMaxLength = 200
	*fetchedCheckoutAttributes[1].ValidationFileMaximumSize = 4096

	mockRepo.On("Fetch", mock.Anything).Return(fetchedCheckoutAttributes, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedCheckoutAttributes, result)
	mockRepo.AssertExpectations(t)
}
