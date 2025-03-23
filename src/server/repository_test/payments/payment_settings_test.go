package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/payments"
	"earnforglance/server/mongo/mocks"
	repository "earnforglance/server/repository/payments"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultPaymentSettings struct {
	mock.Mock
}

func (m *MockSingleResultPaymentSettings) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.PaymentSettings); ok {
		*v.(*domain.PaymentSettings) = *result
	}
	return args.Error(1)
}

var mockItemPaymentSettings = &domain.PaymentSettings{
	ID:                                    primitive.NewObjectID(), // Existing ID of the record to update
	ActivePaymentMethodSystemNames:        []string{"Square", "Authorize.Net"},
	AllowRePostingPayments:                false,
	BypassPaymentMethodSelectionIfOnlyOne: false,
	ShowPaymentMethodDescriptions:         false,
	SkipPaymentInfoStepForRedirectionPaymentMethods: true,
	CancelRecurringPaymentsAfterFailedPayment:       false,
	RegenerateOrderGuidInterval:                     60,
}

func TestPaymentSettingsRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionPaymentSettings

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultPaymentSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemPaymentSettings, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewPaymentSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemPaymentSettings.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultPaymentSettings{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewPaymentSettingsRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemPaymentSettings.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestPaymentSettingsRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionPaymentSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemPaymentSettings).Return(nil, nil).Once()

	repo := repository.NewPaymentSettingsRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemPaymentSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestPaymentSettingsRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionPaymentSettings

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemPaymentSettings.ID}
	update := bson.M{"$set": mockItemPaymentSettings}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewPaymentSettingsRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemPaymentSettings)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
