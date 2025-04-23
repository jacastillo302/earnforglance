package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/catalog"
	repository "earnforglance/server/repository/catalog"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type MockSingleResultSpecificationAttributeOption struct {
	mock.Mock
}

func (m *MockSingleResultSpecificationAttributeOption) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.SpecificationAttributeOption); ok {
		*v.(*domain.SpecificationAttributeOption) = *result
	}
	return args.Error(1)
}

var mockItemSpecificationAttributeOption = &domain.SpecificationAttributeOption{
	ID:                       bson.NewObjectID(), // Existing ID of the record to update
	SpecificationAttributeID: bson.NewObjectID(),
	Name:                     "Size",
	ColorSquaresRgb:          "",
	DisplayOrder:             2,
}

func TestSpecificationAttributeOptionRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionSpecificationAttributeOption

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultSpecificationAttributeOption{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItemSpecificationAttributeOption, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewSpecificationAttributeOptionRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemSpecificationAttributeOption.ID.Hex())

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultSpecificationAttributeOption{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewSpecificationAttributeOptionRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItemSpecificationAttributeOption.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestSpecificationAttributeOptionRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionSpecificationAttributeOption

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	collectionHelper.On("InsertOne", mock.Anything, mockItemSpecificationAttributeOption).Return(nil, nil).Once()

	repo := repository.NewSpecificationAttributeOptionRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockItemSpecificationAttributeOption)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestSpecificationAttributeOptionRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionSpecificationAttributeOption

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)
	filter := bson.M{"_id": mockItemSpecificationAttributeOption.ID}
	update := bson.M{"$set": mockItemSpecificationAttributeOption}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewSpecificationAttributeOptionRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockItemSpecificationAttributeOption)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
