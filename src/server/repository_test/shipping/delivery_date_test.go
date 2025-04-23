package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/shipping"
	repository "earnforglance/server/repository/shipping"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type MockSingleResultDeliveryDate struct {
	mock.Mock
}

func (m *MockSingleResultDeliveryDate) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.DeliveryDate); ok {
		*v.(*domain.DeliveryDate) = *result
	}
	return args.Error(1)
}

var mockItemDeliveryDate = &domain.DeliveryDate{
	ID:           bson.NewObjectID(), // Existing ID of the record to update
	Name:         "Standard Delivery",
	DisplayOrder: 2,
}

func TestDeliveryDateRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionDeliveryDate

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultDeliveryDate{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemDeliveryDate, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewDeliveryDateRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemDeliveryDate.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultDeliveryDate{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewDeliveryDateRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemDeliveryDate.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestDeliveryDateRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionDeliveryDate

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemDeliveryDate).Return(nil, nil).Once()

	repo := repository.NewDeliveryDateRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemDeliveryDate)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestDeliveryDateRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionDeliveryDate

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemDeliveryDate.ID}
	update := bson.M{"$set": mockItemDeliveryDate}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewDeliveryDateRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemDeliveryDate)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
