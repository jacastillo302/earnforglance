package usecase_test

import (
	"context"
	mocks "earnforglance/server/domain/mocks"
	domain "earnforglance/server/domain/payments"
	test "earnforglance/server/usecase/payments"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestPaymentSettingsUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.PaymentSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewPaymentSettingsUsecase(mockRepo, timeout)

	paymentSettingsID := bson.NewObjectID().Hex()

	updatedPaymentSettings := domain.PaymentSettings{
		ID:                                    bson.NewObjectID(), // Existing ID of the record to update
		ActivePaymentMethodSystemNames:        []string{"Square", "Authorize.Net"},
		AllowRePostingPayments:                false,
		BypassPaymentMethodSelectionIfOnlyOne: false,
		ShowPaymentMethodDescriptions:         false,
		SkipPaymentInfoStepForRedirectionPaymentMethods: true,
		CancelRecurringPaymentsAfterFailedPayment:       false,
		RegenerateOrderGuidInterval:                     60,
	}

	mockRepo.On("FetchByID", mock.Anything, paymentSettingsID).Return(updatedPaymentSettings, nil)

	result, err := usecase.FetchByID(context.Background(), paymentSettingsID)

	assert.NoError(t, err)
	assert.Equal(t, updatedPaymentSettings, result)
	mockRepo.AssertExpectations(t)
}

func TestPaymentSettingsUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.PaymentSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewPaymentSettingsUsecase(mockRepo, timeout)

	newPaymentSettings := &domain.PaymentSettings{
		ActivePaymentMethodSystemNames:                  []string{"PayPal", "Stripe"},
		AllowRePostingPayments:                          true,
		BypassPaymentMethodSelectionIfOnlyOne:           true,
		ShowPaymentMethodDescriptions:                   true,
		SkipPaymentInfoStepForRedirectionPaymentMethods: false,
		CancelRecurringPaymentsAfterFailedPayment:       true,
		RegenerateOrderGuidInterval:                     30,
	}

	mockRepo.On("Create", mock.Anything, newPaymentSettings).Return(nil)

	err := usecase.Create(context.Background(), newPaymentSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestPaymentSettingsUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.PaymentSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewPaymentSettingsUsecase(mockRepo, timeout)

	updatedPaymentSettings := &domain.PaymentSettings{
		ID:                                    bson.NewObjectID(), // Existing ID of the record to update
		ActivePaymentMethodSystemNames:        []string{"Square", "Authorize.Net"},
		AllowRePostingPayments:                false,
		BypassPaymentMethodSelectionIfOnlyOne: false,
		ShowPaymentMethodDescriptions:         false,
		SkipPaymentInfoStepForRedirectionPaymentMethods: true,
		CancelRecurringPaymentsAfterFailedPayment:       false,
		RegenerateOrderGuidInterval:                     60,
	}

	mockRepo.On("Update", mock.Anything, updatedPaymentSettings).Return(nil)

	err := usecase.Update(context.Background(), updatedPaymentSettings)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestPaymentSettingsUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.PaymentSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewPaymentSettingsUsecase(mockRepo, timeout)

	paymentSettingsID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, paymentSettingsID).Return(nil)

	err := usecase.Delete(context.Background(), paymentSettingsID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestPaymentSettingsUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.PaymentSettingsRepository)
	timeout := time.Duration(10)
	usecase := test.NewPaymentSettingsUsecase(mockRepo, timeout)

	fetchedPaymentSettings := []domain.PaymentSettings{
		{
			ID:                                    bson.NewObjectID(),
			ActivePaymentMethodSystemNames:        []string{"PayPal", "Stripe"},
			AllowRePostingPayments:                true,
			BypassPaymentMethodSelectionIfOnlyOne: true,
			ShowPaymentMethodDescriptions:         true,
			SkipPaymentInfoStepForRedirectionPaymentMethods: false,
			CancelRecurringPaymentsAfterFailedPayment:       true,
			RegenerateOrderGuidInterval:                     30,
		},
		{
			ID:                                    bson.NewObjectID(),
			ActivePaymentMethodSystemNames:        []string{"Square", "Authorize.Net"},
			AllowRePostingPayments:                false,
			BypassPaymentMethodSelectionIfOnlyOne: false,
			ShowPaymentMethodDescriptions:         false,
			SkipPaymentInfoStepForRedirectionPaymentMethods: true,
			CancelRecurringPaymentsAfterFailedPayment:       false,
			RegenerateOrderGuidInterval:                     60,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedPaymentSettings, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedPaymentSettings, result)
	mockRepo.AssertExpectations(t)
}
