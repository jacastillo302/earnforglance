package repository_test

import (
	"context"
	"errors"
	"testing"
	"time"

	domain "earnforglance/server/domain/shipping"
	repository "earnforglance/server/repository/shipping"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type MockSingleResultShipment struct {
	mock.Mock
}

func (m *MockSingleResultShipment) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.Shipment); ok {
		*v.(*domain.Shipment) = *result
	}
	return args.Error(1)
}

var mockItemShipment = &domain.Shipment{
	ID:                    bson.NewObjectID(), // Existing ID of the record to update
	OrderID:               bson.NewObjectID(),
	TrackingNumber:        "UPDATEDTRACK67890",
	TotalWeight:           new(float64),
	ShippedDateUtc:        new(time.Time),
	DeliveryDateUtc:       new(time.Time),
	ReadyForPickupDateUtc: new(time.Time),
	AdminComment:          "Shipment details updated.",
	CreatedOnUtc:          time.Now().AddDate(0, 0, -7), // Created 7 days ago
}

func TestShipmentRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionShipment

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultShipment{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemShipment, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewShipmentRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemShipment.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultShipment{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewShipmentRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemShipment.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestShipmentRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionShipment

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemShipment).Return(nil, nil).Once()

	repo := repository.NewShipmentRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemShipment)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestShipmentRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionShipment

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemShipment.ID}
	update := bson.M{"$set": mockItemShipment}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewShipmentRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemShipment)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
