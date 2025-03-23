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

type MockSingleResultRecurringPayment struct {
	mock.Mock
}

func (m *MockSingleResultRecurringPayment) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.RecurringPayment); ok {
		*v.(*domain.RecurringPayment) = *result
	}
	return args.Error(1)
}

var mockItemRecurringPayment = &domain.RecurringPayment{
	ID:                primitive.NewObjectID(), // Existing ID of the record to update
	CycleLength:       15,
	CyclePeriodID:     2,
	TotalCycles:       6,
	StartDateUtc:      time.Now().AddDate(0, 0, -30), // Started 30 days ago
	IsActive:          false,
	LastPaymentFailed: true,
	Deleted:           true,
	InitialOrderID:    1002,
	CreatedOnUtc:      time.Now().AddDate(0, 0, -60), // Created 60 days ago
	CyclePeriod:       1,
}

func TestRecurringPaymentRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionRecurringPayment

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultRecurringPayment{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemRecurringPayment, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewRecurringPaymentRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemRecurringPayment.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultRecurringPayment{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewRecurringPaymentRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemRecurringPayment.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestRecurringPaymentRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionRecurringPayment

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemRecurringPayment).Return(nil, nil).Once()

	repo := repository.NewRecurringPaymentRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemRecurringPayment)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestRecurringPaymentRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionRecurringPayment

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemRecurringPayment.ID}
	update := bson.M{"$set": mockItemRecurringPayment}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewRecurringPaymentRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemRecurringPayment)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
