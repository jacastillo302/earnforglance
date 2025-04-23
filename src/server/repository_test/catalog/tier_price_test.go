package repository_test

import (
	"context"
	"errors"
	"testing"
	"time"

	domain "earnforglance/server/domain/catalog"
	repository "earnforglance/server/repository/catalog"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type MockSingleResultTierPrice struct {
	mock.Mock
}

func (m *MockSingleResultTierPrice) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.TierPrice); ok {
		*v.(*domain.TierPrice) = *result
	}
	return args.Error(1)
}

var mockItemTierPrice = &domain.TierPrice{
	ID:               bson.NewObjectID(), // Existing ID of the record to update
	ProductID:        bson.NewObjectID(),
	StoreID:          bson.NewObjectID(),
	CustomerRoleID:   nil,
	Quantity:         20,
	Price:            39.99,
	StartDateTimeUtc: new(time.Time),
	EndDateTimeUtc:   new(time.Time),
}

func TestTierPriceRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionTierPrice

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultTierPrice{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemTierPrice, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewTierPriceRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemTierPrice.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultTierPrice{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewTierPriceRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemTierPrice.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestTierPriceRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionTierPrice

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemTierPrice).Return(nil, nil).Once()

	repo := repository.NewTierPriceRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemTierPrice)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestTierPriceRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionTierPrice

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemTierPrice.ID}
	update := bson.M{"$set": mockItemTierPrice}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewTierPriceRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemTierPrice)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
