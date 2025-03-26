package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/shipping"
	"earnforglance/server/mongo/mocks"
	repository "earnforglance/server/repository/shipping"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultPickupPoint struct {
	mock.Mock
}

func (m *MockSingleResultPickupPoint) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.PickupPoint); ok {
		*v.(*domain.PickupPoint) = *result
	}
	return args.Error(1)
}

var mockItemPickupPoint = &domain.PickupPoint{
	ID:                 primitive.NewObjectID(), // Existing ID of the record to update
	Name:               "Downtown Pickup",
	Description:        "Pickup point located downtown.",
	ProviderSystemName: "UPS",
	Address:            "456 Downtown Ave",
	City:               "Los Angeles",
	County:             "Los Angeles County",
	StateAbbreviation:  "CA",
	CountryCode:        "US",
	ZipPostalCode:      "90001",
	Latitude:           new(float64),
	Longitude:          new(float64),
	PickupFee:          7.50,
	OpeningHours:       "10:00 AM - 6:00 PM",
	DisplayOrder:       2,
	TransitDays:        new(int),
}

func TestPickupPointRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionPickupPoint

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultPickupPoint{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemPickupPoint, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewPickupPointRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemPickupPoint.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultPickupPoint{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewPickupPointRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemPickupPoint.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestPickupPointRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionPickupPoint

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemPickupPoint).Return(nil, nil).Once()

	repo := repository.NewPickupPointRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemPickupPoint)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestPickupPointRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionPickupPoint

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemPickupPoint.ID}
	update := bson.M{"$set": mockItemPickupPoint}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewPickupPointRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemPickupPoint)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
