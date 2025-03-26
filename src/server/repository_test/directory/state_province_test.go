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

type MockSingleResultStateProvince struct {
	mock.Mock
}

func (m *MockSingleResultStateProvince) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.StateProvince); ok {
		*v.(*domain.StateProvince) = *result
	}
	return args.Error(1)
}

var mockItemStateProvince = &domain.StateProvince{
	ID:           primitive.NewObjectID(), // Existing ID of the record to update
	CountryID:    2,
	Name:         "Ontario",
	Abbreviation: "ON",
	Published:    false,
	DisplayOrder: 2,
}

func TestStateProvinceRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionStateProvince

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultStateProvince{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemStateProvince, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewStateProvinceRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemStateProvince.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultStateProvince{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewStateProvinceRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemStateProvince.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestStateProvinceRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionStateProvince

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemStateProvince).Return(nil, nil).Once()

	repo := repository.NewStateProvinceRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemStateProvince)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestStateProvinceRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionStateProvince

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemStateProvince.ID}
	update := bson.M{"$set": mockItemStateProvince}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewStateProvinceRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemStateProvince)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
