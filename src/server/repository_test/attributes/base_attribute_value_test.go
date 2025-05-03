package repository_test

import (
	"context"
	"errors"
	"testing"

	domain "earnforglance/server/domain/attributes"
	repository "earnforglance/server/repository/attributes"
	"earnforglance/server/service/data/mongo/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type MockSingleResultBaseAttributeValue struct {
	mock.Mock
}

func (m *MockSingleResultBaseAttributeValue) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.BaseAttributeValue); ok {
		*v.(*domain.BaseAttributeValue) = *result
	}
	return args.Error(1)
}

func TestBaseAttributeValueRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionBaseAttributeValue

	mockItem := domain.BaseAttributeValue{
		ID:              bson.ObjectID{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
		Name:            "",
		IsPreSelected:   false,
		DisplayOrder:    0,
		BaseAttributeID: bson.ObjectID{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
	}

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultBaseAttributeValue{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItem, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewBaseAttributeValueRepository(databaseHelper, collectionName)

		result, err := ur.FetchByID(context.Background(), mockItem.ID.Hex())

		assert.NoError(t, err)
		assert.Equal(t, mockItem, result)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultBaseAttributeValue{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewBaseAttributeValueRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItem.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestBaseAttributeValueRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionBaseAttributeValue

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)

	mockBaseAttributeValue := &domain.BaseAttributeValue{
		ID:              bson.NewObjectID(), // Existing ID of the record to update
		Name:            "Size",
		IsPreSelected:   true,
		DisplayOrder:    2,
		BaseAttributeID: bson.NewObjectID(), // Reference to the related attribute
	}

	collectionHelper.On("InsertOne", mock.Anything, mockBaseAttributeValue).Return(nil, nil).Once()

	repo := repository.NewBaseAttributeValueRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockBaseAttributeValue)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestBaseAttributeValueRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionBaseAttributeValue

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)

	mockBaseAttributeValue := &domain.BaseAttributeValue{
		ID:              bson.NewObjectID(), // Existing ID of the record to update
		Name:            "Size",
		IsPreSelected:   true,
		DisplayOrder:    2,
		BaseAttributeID: bson.NewObjectID(), // Reference to the related attribute
	}

	filter := bson.M{"_id": mockBaseAttributeValue.ID}
	update := bson.M{"$set": mockBaseAttributeValue}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewBaseAttributeValueRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockBaseAttributeValue)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
