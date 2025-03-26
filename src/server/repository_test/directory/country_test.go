package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/directory"
	repository "earnforglance/server/repository/directory"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockSingleResultCountry struct {
	mock.Mock
}

func (m *MockSingleResultCountry) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.Country); ok {
		*v.(*domain.Country) = *result
	}
	return args.Error(1)
}

var mockItemCountry = &domain.Country{
	ID:                 primitive.NewObjectID(), // Existing ID of the record to update
	Name:               "Canada",
	AllowsBilling:      true,
	AllowsShipping:     true,
	TwoLetterIsoCode:   "CA",
	ThreeLetterIsoCode: "CAN",
	NumericIsoCode:     124,
	SubjectToVat:       true,
	Published:          false,
	DisplayOrder:       2,
	LimitedToStores:    true,
}

func TestCountryRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionCountry

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultCountry{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemCountry, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewCountryRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemCountry.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultCountry{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewCountryRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemCountry.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestCountryRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionCountry

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemCountry).Return(nil, nil).Once()

	repo := repository.NewCountryRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemCountry)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestCountryRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionCountry

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemCountry.ID}
	update := bson.M{"$set": mockItemCountry}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewCountryRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemCountry)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
